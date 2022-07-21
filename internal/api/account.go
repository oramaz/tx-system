package api

import (
	"database/sql"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
	db "github.com/oramaz/tx-system/internal/db/sqlc"
)

type createAccountRequest struct {
	Owner string `json:"owner"`
}

func (r *createAccountRequest) validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Owner, validation.Required, validation.Length(2, 0)),
	)
}

func (s *Server) createAccount(c *fiber.Ctx) error {
	var req createAccountRequest

	// Parse request body to struct
	if err := c.BodyParser(&req); err != nil {
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
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation":
				return c.Status(fiber.StatusForbidden).JSON(errorResponse(err))
			}
		}
		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(account)
}

type getAccountRequest struct {
	ID int64 `json:"id"`
}

func (r *getAccountRequest) validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.ID, validation.Required, validation.Min(1)),
	)
}

func (s *Server) getAccount(c *fiber.Ctx) error {
	var req getAccountRequest

	// Parse request params to struct
	c.ParamsParser(&req)

	// Validate the struct
	if err := req.validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse(err))
	}

	account, err := s.store.GetAccount(c.Context(), req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(errorResponse(err))
		}

		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse(err))

	}

	return c.Status(fiber.StatusOK).JSON(account)
}
