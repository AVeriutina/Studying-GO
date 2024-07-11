package main

import (
	"awesomeProject/cmd"
	"awesomeProject/proto"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type Client struct {
	client_imp proto.AccountManagerClient
}

func (c *Client) do(cmd cmd.Command) error {
	switch cmd.Cmd {
	case "create":
		if err := c.create(cmd); err != nil {
			return fmt.Errorf("create account failed: %w", err)
		}
		return nil
	case "get":
		if err := c.get(cmd); err != nil {
			return fmt.Errorf("get account failed: %w", err)
		}
		return nil
	case "change_name":
		if err := c.change_name(cmd); err != nil {
			return fmt.Errorf("change name of account failed: %w", err)
		}
		return nil
	case "change_amount":
		if err := c.change_amount(cmd); err != nil {
			return fmt.Errorf("change amount of account failed: %w", err)
		}
		return nil
	case "delete":
		if err := c.delete(cmd); err != nil {
			return fmt.Errorf("delete account failed: %w", err)
		}
		return nil
	default:
		return fmt.Errorf("unknown command %s", cmd.Cmd)
	}
}

func (c *Client) create(cmd cmd.Command) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := c.client_imp.CreateAccount(ctx, &proto.CreateAccountRequest{Name: cmd.Name, Amount: int32(cmd.Amount)})
	return err
}

func (c *Client) change_name(cmd cmd.Command) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := c.client_imp.ChangeNameAccount(ctx, &proto.ChangeNameAccountRequest{PrevName: cmd.Name, NewName: cmd.NewName})
	return err
}

func (c *Client) get(cmd cmd.Command) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := c.client_imp.GetAccount(ctx, &proto.GetAccountRequest{Name: cmd.Name})
	if err != nil {
		return err
	}
	fmt.Printf("response account name: %s and amount: %d", resp.Name, resp.Amount)
	return err
}

func (c *Client) change_amount(cmd cmd.Command) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := c.client_imp.ChangeAmountAccount(ctx, &proto.ChangeAmountAccountRequest{Name: cmd.Name, NewAmount: int32(cmd.Amount)})
	return err
}

func (c *Client) delete(cmd cmd.Command) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := c.client_imp.DeleteAccount(ctx, &proto.DeleteAccountRequest{Name: cmd.Name})
	return err
}

func main() {
	portVal := flag.Int("port", 8080, "server port")
	hostVal := flag.String("host", "0.0.0.0", "server host")
	cmdVal := flag.String("cmd", "", "command to execute")
	nameVal := flag.String("name", "", "name of account")
	amountVal := flag.Int("amount", 0, "amount of account")
	newnameVal := flag.String("new_name", "", "new name of account")

	flag.Parse()

	command := cmd.Command{
		Port:    *portVal,
		Host:    *hostVal,
		Cmd:     *cmdVal,
		Name:    *nameVal,
		Amount:  *amountVal,
		NewName: *newnameVal,
	}

	conn, err := grpc.NewClient("0.0.0.0:4567", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}

	defer func() {
		_ = conn.Close()
	}()

	client := Client{proto.NewAccountManagerClient(conn)}

	if err := client.do(command); err != nil {
		panic(err)
	}

}
