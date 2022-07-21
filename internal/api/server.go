package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	db "github.com/oramaz/tx-system/internal/db/sqlc"
)

type Server struct {
	store  db.Store
	router *fiber.App
}

func (server *Server) Start(address string) error {
	return server.router.Listen(address)
}

func NewServer(st db.Store) *Server {
	s := &Server{store: st}
	r := fiber.New()

	r.Use(logger.New())

	r.Post("/users", s.createUser)
	r.Post("/accounts", s.createAccount)
	r.Get("/accounts/:id", s.getAccount)
	r.Post("/transfers", s.createTransfer)

	s.router = r

	return s
}

func errorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"error": err.Error(),
	}
}
