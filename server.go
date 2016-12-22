package osin

import (
	"time"
)

// Server is an OAuth2 implementation
type Server struct {
	Config             *ServerConfig
	Storage            Storage // deprecated in favor of StorageWithContext
	StorageWithContext StorageWithContext
	AuthorizeTokenGen  AuthorizeTokenGen
	AccessTokenGen     AccessTokenGen
	Now                func() time.Time
}

// NewServer creates a new server instance
func NewServer(config *ServerConfig, storage Storage) *Server {
	s := NewServerWithContext(config, &oldStorageWithContext{storage})
	s.Storage = storage
	return s
}

// NewServerWithContext creates a new server instance that has a context-aware storage.
func NewServerWithContext(config *ServerConfig, storage StorageWithContext) *Server {
	return &Server{
		Config:             config,
		StorageWithContext: storage,
		AuthorizeTokenGen:  &AuthorizeTokenGenDefault{},
		AccessTokenGen:     &AccessTokenGenDefault{},
		Now:                time.Now,
	}
}

// NewResponse creates a new response for the server
func (s *Server) NewResponse() *Response {
	r := NewResponseWithContext(s.GetStorage())
	r.ErrorStatusCode = s.Config.ErrorStatusCode
	return r
}

// GetStorage is a getter for the StorageWithContext associated with the Server.
func (s *Server) GetStorage() StorageWithContext {
	if s.StorageWithContext == nil {
		return &oldStorageWithContext{Storage: s.Storage}
	}
	return s.StorageWithContext
}
