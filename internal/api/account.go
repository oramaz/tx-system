package api

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofiber/fiber/v2"
	db "github.com/oramaz/tx-system/internal/db/sqlc"
)

type createAccountRequest struct {
	Owner string `json:"owner" `
}

func (r *createAccountRequest) validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Owner, validation.Required),
	)
}

func (s *Server) createAccount(c *fiber.Ctx) error {
	req := new(createAccountRequest)

	// Parse request body to struct
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse(err))
	}

	// Validate the struct
	if err := req.validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse(err))
	}

	arg := db.CreateAccountParams{
		Owner:   req.Owner,
		Balance: 0,
	}

	account, err := s.store.CreateAccount(c.Context(), arg)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(account)
}
