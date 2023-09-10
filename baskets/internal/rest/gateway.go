package rest

import (
	"context"
	"github.com/Sraik25/event-driven-architecture/baskets/basketspb"
	"github.com/go-chi/chi/v5"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RegisterGateway(ctx context.Context, mux *chi.Mux, grpcAddr string) error {
	const apiRoot = "/api/baskets"

	gateway := runtime.NewServeMux()
	err := basketspb.RegisterBasketServiceHandlerFromEndpoint(ctx, gateway, grpcAddr, []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	})
	if err != nil {
		return err
	}

	// Mount the GRPC gateway
	mux.Mount(apiRoot, gateway)
	return nil
}
