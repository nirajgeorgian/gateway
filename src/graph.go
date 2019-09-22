package gateway

import (
  "context"

  "github.com/spf13/viper"
  "github.com/99designs/gqlgen/graphql"

  accountapi "github.com/nirajgeorgian/account/src/api"
	jobapi "github.com/nirajgeorgian/job/src/api"
)

type GatewayServer struct {
  AccountSvcAddr string
  AccountClient *accountapi.Client

	JobSvcAddr string
	JobClient *jobapi.Client
}

func NewGraphQLServer(ctx context.Context) (*GatewayServer, error) {
  // create an empty server
  svc := new(GatewayServer)

	// account uri to copnnect to account service
  svc.AccountSvcAddr = viper.GetString("accounturi")
  accountClient, err := accountapi.NewClient(svc.AccountSvcAddr)
  if err != nil {
		return nil, err
	}
  svc.AccountClient = accountClient

	// job uri to connect to job service
	svc.JobSvcAddr = viper.GetString("joburi")
	jobClient, err := jobapi.NewClient(svc.JobSvcAddr)
	if err != nil {
		return nil, err
	}
	svc.JobClient = jobClient

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
