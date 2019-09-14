package gateway

import (
  "context"
  "errors"
  "time"
  _ "log"

  models "github.com/nirajgeorgian/gateway/src/models"

  account "github.com/nirajgeorgian/account/src/model"
  _ "github.com/nirajgeorgian/account/src/api"

  job "github.com/nirajgeorgian/job/src/model"
  _ "github.com/nirajgeorgian/job/src/api"
)

var (
  ErrInvalidParameter = errors.New("invalid parameter")
)

type mutationResolver struct {
  server *GatewayServer
}

func (r *mutationResolver)	Dummy(ctx context.Context) (*string, error)  {
  ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
  defer cancel()

  msg := "dodo duck"

  return &msg, nil
}

func (r *mutationResolver) CreateAccount(ctx context.Context, in models.CreateAccountReq) (*account.Account, error) {
  // ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
  // defer cancel()
  //
  // acc := &account.Account{Email: *in.Email}
  // _, err := accountapi.NewAccountServiceClient(s.).CreateAccount(ctx, &accountapi.CreateAccountReq{Account: acc})
  // if err != nil {
  //   log.Fatalf("could not greet: %v", err)
  // }

  return &account.Account{
    Email:   *in.Email,
  }, nil
}

func (r *mutationResolver) CreateJob(ctx context.Context, in models.CreateJobReq) (*job.Job, error) {
  // ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
  // defer cancel()
  //
  // acc := &job.Job{JobName: *in.JobName}
  // _, err = jobapi.NewJobServiceClient(r.server.JobClient).CreateJob(ctx, &jobapi.CreateJobReq{Job: acc})
  // if err != nil {
  //   log.Fatalf("could not greet: %v", err)
  // }

  return &job.Job{
    JobName:   *in.JobName,
  }, nil
}
