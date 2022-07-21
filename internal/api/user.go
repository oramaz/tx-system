package api

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
	db "github.com/oramaz/tx-system/internal/db/sqlc"
	"github.com/oramaz/tx-system/internal/util"
)

type createUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *createUserRequest) validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Username, validation.Required, is.Alphanumeric),
		validation.Field(&r.Password, validation.Required, validation.Length(6, 0)),
	)
}

type createUserResponse struct {
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

func (s *Server) createUser(c *fiber.Ctx) error {
	var req createUserRequest

	// Parse request body to struct
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse(err))
	}

	// Validate the struct
	if err := req.validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse(err))
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse(err))
	}

	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
	}

	user, err := s.store.CreateUser(c.Context(), arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return c.Status(fiber.StatusForbidden).JSON(errorResponse(err))
			}
		}
		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse(err))
	}

	resp := createUserResponse{
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}
