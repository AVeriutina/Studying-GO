package accounts

import (
	"awesomeProject/accounts/dto"
	"awesomeProject/accounts/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"sync"
)

func New() *Handler {
	return &Handler{
		accounts: make(map[string]*models.Account),
		guard:    &sync.RWMutex{},
	}
}

type Handler struct {
	accounts map[string]*models.Account
	guard    *sync.RWMutex
}

func (h *Handler) CreateAccount(c echo.Context) error {
	var request dto.CreateAccountRequest // {"name": "alice", "amount": 50}
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)

		return c.String(http.StatusBadRequest, "invalid request")
	}

	if len(request.Name) == 0 {
		return c.String(http.StatusBadRequest, "empty name")
	}

	h.guard.Lock()

	if _, ok := h.accounts[request.Name]; ok {
		h.guard.Unlock()

		return c.String(http.StatusForbidden, "account already exists")
	}

	h.accounts[request.Name] = &models.Account{
		Name:   request.Name,
		Amount: request.Amount,
	}

	h.guard.Unlock()

	return c.NoContent(http.StatusCreated)
}

func (h *Handler) GetAccount(c echo.Context) error {
	name := c.QueryParams().Get("name")

	h.guard.RLock()

	account, ok := h.accounts[name]

	h.guard.RUnlock()

	if !ok {
		return c.String(http.StatusNotFound, "account not found")
	}

	response := dto.GetAccountResponse{
		Name:   account.Name,
		Amount: account.Amount,
	}

	return c.JSON(http.StatusOK, response)
}

// Удаляет аккаунт
func (h *Handler) DeleteAccount(c echo.Context) error {
	var request dto.DeleteAccountRequest // {"name": "alice"}
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)

		return c.String(http.StatusBadRequest, "invalid request")
	}

	if len(request.Name) == 0 {
		return c.String(http.StatusBadRequest, "empty name")
	}

	if _, ok := h.accounts[request.Name]; !ok {
		return c.String(http.StatusBadRequest, "account doesn't exist")
	}

	h.guard.Lock()

	delete(h.accounts, request.Name)

	h.guard.Unlock()
	return c.NoContent(http.StatusNoContent)
}

// Меняет баланс
func (h *Handler) PatchAccount(c echo.Context) error {
	var request dto.ChangeAccountRequest // {"name": "alice", "amount": 50}
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)

		return c.String(http.StatusBadRequest, "invalid request")
	}

	if len(request.Name) == 0 {
		return c.String(http.StatusBadRequest, "empty name")
	}

	if _, ok := h.accounts[request.Name]; !ok {
		return c.String(http.StatusBadRequest, "account doesn't exist")
	}

	h.guard.Lock()

	h.accounts[request.Name].Amount = request.Amount

	h.guard.Unlock()

	return c.NoContent(http.StatusOK)
}

// Меняет имя
func (h *Handler) ChangeAccount(c echo.Context) error {
	var request dto.PatchAccountRequest // {"name": "alice", "new_name": "alica111"}
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)

		return c.String(http.StatusBadRequest, "invalid request")
	}

	if len(request.NewName) == 0 || len(request.PrevName) == 0 {
		return c.String(http.StatusBadRequest, "empty name")
	}

	if _, ok := h.accounts[request.PrevName]; !ok {
		return c.String(http.StatusBadRequest, "account doesn't exist")
	}
	if _, ok := h.accounts[request.NewName]; ok {
		return c.String(http.StatusBadRequest, "account with such name already exists")
	}

	h.guard.Lock()

	h.accounts[request.NewName] = &models.Account{
		Name:   request.NewName,
		Amount: h.accounts[request.PrevName].Amount,
	}
	delete(h.accounts, request.PrevName)

	h.guard.Unlock()

	return c.NoContent(http.StatusOK)
}
