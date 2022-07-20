package api

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofiber/fiber/v2"
	db "github.com/oramaz/tx-system/internal/db/sqlc"
)

type transferRequest struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

func (r *transferRequest) validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.FromAccountID, validation.Required, validation.Min(1)),
		validation.Field(&r.ToAccountID, validation.Required, validation.Min(1)),
		validation.Field(&r.Amount, validation.Required, validation.Min(0)),
	)
}

func (s *Server) createTransfer(c *fiber.Ctx) error {
	var req transferRequest

	// Parse request body to struct
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse(err))
	}

	// Validate the struct
	if err := req.validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse(err))
	}

	arg := db.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	}

	result, err := s.store.TransferTx(c.Context(), arg)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(result)
}
