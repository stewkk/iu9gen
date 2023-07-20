package html

import "github.com/stewkk/iu9gen/internal/app"

// Make sure we conform to ServerInterface interface
var _ ServerInterface = (*Server)(nil)

type Server struct {
	App app.App
}
