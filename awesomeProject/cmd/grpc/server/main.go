package main

import (
	"awesomeProject/proto"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
)

func New(db_ *sql.DB) *server {
	return &server{
		db: db_,
	}
}

type server struct {
	proto.UnimplementedAccountManagerServer
	db *sql.DB
}

func (s *server) CreateAccount(ctx context.Context, req *proto.CreateAccountRequest) (*proto.CreateAccountReply, error) {
	if len(req.Name) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "empty name")
	}

	row, err := s.db.QueryContext(ctx, "SELECT name FROM accounts WHERE name=$1", req.Name)
	if err == nil {
		return nil, status.Errorf(codes.AlreadyExists, "account already exists")
	}
	defer func() {
		_ = row.Close()
	}()

	result, err := s.db.ExecContext(ctx, "INSERT INTO accounts(name, amount) VALUES($1, $2)", req.Name, req.Amount)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create account: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to make id of account: %v", err)
	}
	repl := proto.CreateAccountReply{
		AccountId: id,
	}
	return &repl, nil
}

func (s *server) ChangeAmountAccount(ctx context.Context, req *proto.ChangeAmountAccountRequest) (*proto.ChangeAmountAccountReply, error) {
	if len(req.Name) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "empty name")
	}

	row, err := s.db.QueryContext(ctx, "UPDATE accounts SET amount=$1 WHERE name=$2", req.NewAmount, req.Name)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to query account: %v", err)
	}
	defer func() {
		_ = row.Close()
	}()
	return nil, nil
}

func (s *server) DeleteAccount(ctx context.Context, req *proto.DeleteAccountRequest) (*proto.DeleteAccountReply, error) {
	if len(req.Name) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "empty name")
	}

	row := s.db.QueryRowContext(ctx, "SELECT name FROM accounts WHERE name=$1", req.Name)

	var name string
	err := row.Scan(&name)
	if err == sql.ErrNoRows {
		return nil, status.Errorf(codes.NotFound, "account does not exist")
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to query account: %v", err)
	}

	_, err = s.db.ExecContext(ctx, "DELETE FROM accounts WHERE name=$1", req.Name)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete account: %v", err)
	}

	return nil, nil
}

func (s *server) ChangeNameAccount(ctx context.Context, req *proto.ChangeNameAccountRequest) (*proto.ChangeNameAccountReply, error) {
	if len(req.NewName) == 0 || len(req.NewName) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "empty name")
	}

	row, err := s.db.QueryContext(ctx, "UPDATE accounts SET name=$1 WHERE name=$2", req.NewName, req.PrevName)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "cannot change name with such params")
	}
	defer func() {
		_ = row.Close()
	}()
	return nil, nil
}

func (s *server) GetAccount(ctx context.Context, req *proto.GetAccountRequest) (*proto.GetAccountReply, error) {
	if len(req.Name) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "empty name")
	}

	row := s.db.QueryRowContext(ctx, "SELECT name, amount FROM accounts WHERE name=$1", req.Name)

	var name string
	var amount int32
	err := row.Scan(&name, &amount)
	if err == sql.ErrNoRows {
		return nil, status.Errorf(codes.NotFound, "account does not exist")
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to query account: %v", err)
	}

	repl := &proto.GetAccountReply{
		Name:   name,
		Amount: amount,
	}
	return repl, nil
}

func main() {
	connectionString := "host=0.0.0.0 port=5432 dbname=postgres user=postgres password=mysecretpassword"

	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		fmt.Errorf("failed to open database: %v", err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		fmt.Errorf("failed to ping database: %v", err)
	}

	ctx := context.Background()
	//
	//res, err := db.ExecContext(ctx, "INSERT INTO accounts(name, amount) VALUES($1, $2)", "bob", 10)
	//if err != nil {
	//	panic(err)
	//}
	//
	//id, err := res.RowsAffected()
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(id)
	//
	rows, err := db.QueryContext(ctx, "SELECT name, amount FROM accounts")
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = rows.Close()
	}()

	type Account struct {
		Name   string
		Amount int
	}

	accounts := make([]Account, 0)

	for rows.Next() {
		var account Account

		if err := rows.Scan(&account.Name, &account.Amount); err != nil {
			panic(err)
		}

		accounts = append(accounts, account)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

	fmt.Println(accounts)

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 4567))

	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	proto.RegisterAccountManagerServer(s, New(db))
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}
