package gapi

import (
	"fmt"

	db "github.com/dungnguyen/simplebank/db/sqlc"
	"github.com/dungnguyen/simplebank/pb"
	"github.com/dungnguyen/simplebank/token"
	"github.com/dungnguyen/simplebank/util"
)

// Server serves gRPC request for our banking service.
type Server struct {
	pb.UnimplementedSimplebankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer creates a new gRPC server.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
