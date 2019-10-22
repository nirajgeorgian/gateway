package gateway

import (
	"os"
	"fmt"
	"time"
	"context"

	"google.golang.org/grpc"
	"github.com/pkg/errors"
  "github.com/spf13/viper"
  "github.com/99designs/gqlgen/graphql"
)

type GatewayServer struct {
  AccountSvcAddr string
  AccountClient *grpc.ClientConn

	JobSvcAddr string
	JobClient *grpc.ClientConn
}

func NewGraphQLServer(ctx context.Context) (*GatewayServer, error) {
  svc := new(GatewayServer)

  svc.AccountSvcAddr = viper.GetString("accounturi")
	svc.JobSvcAddr = viper.GetString("joburi")

	mustConnGRPC(ctx, &svc.AccountClient, svc.AccountSvcAddr)
	mustConnGRPC(ctx, &svc.JobClient, svc.JobSvcAddr)

  return svc, nil
}

// mustMapEnv map environment varibale to server configuration
func mustMapEnv(target *string, envKey string) {
	v := os.Getenv(envKey)
	if v == "" {
		panic(fmt.Sprintf("environment variable %q not set", envKey))
	}
	*target = v
}

// mustConnGRPC map gRPC client to appropiate client
func mustConnGRPC(ctx context.Context, conn **grpc.ClientConn, addr string) {
	var err error
	*conn, err = grpc.DialContext(ctx, addr,
		grpc.WithInsecure(),
		grpc.WithTimeout(time.Second*3),)
	if err != nil {
		panic(errors.Wrapf(err, "grpc: failed to connect %s", addr))
	}
}

func (s *GatewayServer) Mutation() MutationResolver {
 return  &mutationResolver{
   server: s,
 }
}

func (s *GatewayServer) Query() QueryResolver {
 return  &queryResolver{
   server: s,
 }
}

func (s *GatewayServer) Job() JobResolver {
  return &jobResolver{
    server: s,
  }
}

func (s *GatewayServer) Sallary() SallaryResolver {
  return &sallaryResolver{
    server: s,
  }
}

func (s *GatewayServer) ToExecutableSchema() graphql.ExecutableSchema {
  return NewExecutableSchema(Config{
    Resolvers: s,
  })
}
