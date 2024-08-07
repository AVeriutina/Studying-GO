package main

import (
	"awesomeProject/accounts/dto"
	"awesomeProject/cmd"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
)

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

	if err := do(command); err != nil {
		panic(err)
	}
}

func do(cmd cmd.Command) error {
	switch cmd.Cmd {
	case "create":
		if err := create(cmd); err != nil {
			return fmt.Errorf("create account failed: %w", err)
		}

		return nil
	case "get":
		if err := get(cmd); err != nil {
			return fmt.Errorf("get account failed: %w", err)
		}

		return nil

	case "change_amount":
		if err := change_amount(cmd); err != nil {
			return fmt.Errorf("change amount of account failed: %w", err)
		}

		return nil

	case "change_name":
		if err := change_name(cmd); err != nil {
			return fmt.Errorf("change name of account failed: %w", err)
		}

		return nil

	case "delete":
		if err := deleted(cmd); err != nil {
			return fmt.Errorf("delete account failed: %w", err)
		}

		return nil

	default:
		return fmt.Errorf("unknown command %s", cmd.Cmd)
	}
}

func deleted(cmd cmd.Command) error {
	if len(cmd.Name) == 0 {
		return fmt.Errorf("name is empty")
	}

	request := dto.DeleteAccountRequest{
		Name: cmd.Name,
	}

	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	req, err := http.NewRequest(
		"DELETE",
		fmt.Sprintf("http://%s:%d/account/delete", cmd.Host, cmd.Port),
		bytes.NewReader(data),
	)
	req.Header.Add("Content-type", "application/json")

	if err != nil {
		return fmt.Errorf("http delete failed: %w", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode == http.StatusNoContent {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	return fmt.Errorf("resp error %s", string(body))
}

func change_name(cmd cmd.Command) error {
	if len(cmd.Name) == 0 || len(cmd.NewName) == 0 {
		return fmt.Errorf("name is empty")
	}

	request := dto.PatchAccountRequest{
		PrevName: cmd.Name,
		NewName:  cmd.NewName,
	}

	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	req, err := http.NewRequest(
		"PATCH",
		fmt.Sprintf("http://%s:%d/account/change_name", cmd.Host, cmd.Port),
		bytes.NewReader(data),
	)
	req.Header.Add("Content-type", "application/json")

	if err != nil {
		return fmt.Errorf("http patch failed: %w", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	return fmt.Errorf("resp error %s", string(body))
}

func change_amount(cmd cmd.Command) error {
	if len(cmd.Name) == 0 {
		return fmt.Errorf("name is empty")
	}

	request := dto.ChangeAccountRequest{
		Name:   cmd.Name,
		Amount: cmd.Amount,
	}

	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	req, err := http.NewRequest(
		"PATCH",
		fmt.Sprintf("http://%s:%d/account/change_amount", cmd.Host, cmd.Port),
		bytes.NewReader(data),
	)
	req.Header.Add("Content-type", "application/json")

	if err != nil {
		return fmt.Errorf("http patch failed: %w", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	return fmt.Errorf("resp error %s", string(body))
}

func get(cmd cmd.Command) error {
	resp, err := http.Get(
		fmt.Sprintf("http://%s:%d/account?name=%s", cmd.Host, cmd.Port, cmd.Name),
	)
	if err != nil {
		return fmt.Errorf("http get failed: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("read body failed: %w", err)
		}

		return fmt.Errorf("resp error %s", string(body))
	}

	var response dto.GetAccountResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return fmt.Errorf("json decode failed: %w", err)
	}

	fmt.Printf("response account name: %s and amount: %d", response.Name, response.Amount)

	return nil
}

func create(cmd cmd.Command) error {
	if len(cmd.Name) == 0 {
		return fmt.Errorf("name is empty")
	}

	request := dto.CreateAccountRequest{
		Name:   cmd.Name,
		Amount: cmd.Amount,
	}

	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	resp, err := http.Post(
		fmt.Sprintf("http://%s:%d/account/create", cmd.Host, cmd.Port),
		"application/json",
		bytes.NewReader(data),
	)
	if err != nil {
		return fmt.Errorf("http post failed: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode == http.StatusCreated {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	return fmt.Errorf("resp error %s", string(body))
}
