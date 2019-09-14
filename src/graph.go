package gateway

import (
  "context"
  "time"

  "github.com/pkg/errors"
  "google.golang.org/grpc"
  "github.com/spf13/viper"
  "github.com/99designs/gqlgen/graphql"
)

type GatewayServer struct {
  JobSvcAddr string
  JobClient *grpc.ClientConn

  AccountSvcAddr string
  AccountClient *grpc.ClientConn
}

func NewGraphQLServer(ctx context.Context) (*GatewayServer, error) {
  // create an empty server
  svc := new(GatewayServer)

  svc.AccountSvcAddr = viper.GetString("accounturi")
  svc.JobSvcAddr = viper.GetString("joburi")

  mustConnGRPC(ctx, &svc.AccountClient, svc.AccountSvcAddr)
  mustConnGRPC(ctx, &svc.JobClient, svc.JobSvcAddr)

  return svc, nil
}

func mustConnGRPC(ctx context.Context, conn **grpc.ClientConn, addr string) {
	var err error
	*conn, err = grpc.DialContext(ctx, addr,
		grpc.WithInsecure(),
		grpc.WithTimeout(time.Second*3),
  )
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

// func (s *GatewayServer) Account() AccountResolver {
//   return &accountResolver{
//     server: s,
//   }
// }

func (s *GatewayServer) ToExecutableSchema() graphql.ExecutableSchema {
  return NewExecutableSchema(Config{
    Resolvers: s,
  })
}
