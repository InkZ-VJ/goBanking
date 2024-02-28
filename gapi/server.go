package gapi

import (
	"fmt"

	db "github.com/VatJittiprasert/goBanking/db/sqlc"
	"github.com/VatJittiprasert/goBanking/pb"
	"github.com/VatJittiprasert/goBanking/token"
	"github.com/VatJittiprasert/goBanking/utils"
	"github.com/VatJittiprasert/goBanking/worker"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	pb.UnimplementedSimpleBankServer
	config          utils.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

// NewServer creates a new gRPC server.
func NewServer(config utils.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil
}
