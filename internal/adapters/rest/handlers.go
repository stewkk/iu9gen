package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/stewkk/iu9gen/internal/app"
)

// Make sure we conform to ServerInterface interface
var _ ServerInterface = (*Server)(nil)

type Server struct {
	App app.App
}

// GenerateReport implements ServerInterface
func (*Server) GenerateReport(ctx echo.Context) error {
	panic("unimplemented")
}

// Ping implements ServerInterface
func (s *Server) Ping(ctx echo.Context) error {
	return ctx.NoContent(http.StatusOK)
}

// 4 spaces
const identation = "    "
