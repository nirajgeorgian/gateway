package gateway

import (
  "context"

  "github.com/spf13/viper"
  "github.com/99designs/gqlgen/graphql"

  accountapi "github.com/nirajgeorgian/account/src/api"
)

type GatewayServer struct {
  AccountSvcAddr string
  AccountClient *accountapi.Client
}

func NewGraphQLServer(ctx context.Context) (*GatewayServer, error) {
  // create an empty server
  svc := new(GatewayServer)

  svc.AccountSvcAddr = viper.GetString("accounturi")
  accountClient, err := accountapi.NewClient(svc.AccountSvcAddr)
  if err != nil {
		return nil, err
	}
  svc.AccountClient = accountClient

  return svc, nil
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

func (s *GatewayServer) ToExecutableSchema() graphql.ExecutableSchema {
  return NewExecutableSchema(Config{
    Resolvers: s,
  })
}
