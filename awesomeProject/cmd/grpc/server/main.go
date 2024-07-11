package main

import (
	"awesomeProject/accounts/models"
	"awesomeProject/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
	"sync"
)

func New() *server {
	return &server{
		accounts: make(map[string]*models.Account),
		guard:    &sync.RWMutex{},
	}
}

type server struct {
	proto.UnimplementedAccountManagerServer
	accounts map[string]*models.Account
	guard    *sync.RWMutex
}

func (s *server) CreateAccount(ctx context.Context, req *proto.CreateAccountRequest) (*proto.CreateAccountReply, error) {
	if len(req.Name) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "empty name")
	}
	s.guard.Lock()
	if _, ok := s.accounts[req.Name]; ok {
		s.guard.Unlock()

		return nil, status.Errorf(codes.AlreadyExists, "account already exists")
	}

	s.accounts[req.Name] = &models.Account{
		Name:   req.Name,
		Amount: int(req.Amount),
	}
	s.guard.Unlock()

	return nil, nil
}

func (s *server) ChangeAmountAccount(ctx context.Context, req *proto.ChangeAmountAccountRequest) (*proto.ChangeAmountAccountReply, error) {
	if len(req.Name) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "empty name")
	}

	if _, ok := s.accounts[req.Name]; !ok {
		return nil, status.Errorf(codes.InvalidArgument, "account does not exist")
	}

	s.guard.Lock()

	s.accounts[req.Name].Amount = int(req.NewAmount)

	s.guard.Unlock()

	return nil, nil
}

func (s *server) DeleteAccount(ctx context.Context, req *proto.DeleteAccountRequest) (*proto.DeleteAccountReply, error) {
	if len(req.Name) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "empty name")
	}

	if _, ok := s.accounts[req.Name]; !ok {
		return nil, status.Errorf(codes.InvalidArgument, "account does not exist")
	}

	s.guard.Lock()

	delete(s.accounts, req.Name)

	s.guard.Unlock()
	return nil, nil
}

func (s *server) ChangeNameAccount(ctx context.Context, req *proto.ChangeNameAccountRequest) (*proto.ChangeNameAccountReply, error) {
	if len(req.PrevName) == 0 || len(req.NewName) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "empty name")
	}

	if _, ok := s.accounts[req.PrevName]; !ok {
		return nil, status.Errorf(codes.InvalidArgument, "account does not exist")
	}
	if _, ok := s.accounts[req.NewName]; ok {
		return nil, status.Errorf(codes.AlreadyExists, "account already exists")
	}

	s.guard.Lock()

	s.accounts[req.NewName] = &models.Account{
		Name:   req.NewName,
		Amount: s.accounts[req.PrevName].Amount,
	}
	delete(s.accounts, req.PrevName)

	s.guard.Unlock()

	return nil, nil
}

func (s *server) GetAccount(ctx context.Context, req *proto.GetAccountRequest) (*proto.GetAccountReply, error) {
	if len(req.Name) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "empty name")
	}

	s.guard.RLock()

	account, ok := s.accounts[req.Name]

	s.guard.RUnlock()

	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "account does not exist")
	}

	repl := proto.GetAccountReply{
		Name:   account.Name,
		Amount: int32(account.Amount),
	}
	return &repl, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 4567))

	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	proto.RegisterAccountManagerServer(s, New())
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}
