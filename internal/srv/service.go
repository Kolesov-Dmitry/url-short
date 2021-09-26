package srv

import (
	"github.com/gorilla/mux"
)

// Service is an HTTP endpoint abstraction
type Service interface {
	// Register binds service routes with router
	Register(router *mux.Router)
}
