package gateway

import (
  "context"
  "errors"
  "time"
  "log"

  models "github.com/nirajgeorgian/gateway/src/models"

  account "github.com/nirajgeorgian/account/src/model"
  job "github.com/nirajgeorgian/job/src/model"
)

var (
  ErrInvalidParameter = errors.New("invalid parameter")
)

type mutationResolver struct {
  server *GatewayServer
}

func (r *mutationResolver)  Dummy(ctx context.Context) (*string, error)  {
  ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
  defer cancel()

  msg := "dodo duck"

  return &msg, nil
}

func (r *mutationResolver) CreateJob(ctx context.Context, in models.CreateJobReq) (*job.Job, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
  defer cancel()

	// sallary min and max for range
	MinSallary := &job.Sallary{
		Value: uint64(in.MinSallary.Value),
		Currency: in.MinSallary.Currency,
	}
	MaxSallary := &job.Sallary{
		Value: uint64(in.MaxSallary.Value),
		Currency: in.MaxSallary.Currency,
	}

	job := &job.Job{
		JobName: in.JobName,
		JobDescription: in.JobDescription,
		JobCategory: in.JobCategory,
		Location: *in.Location,
		JobTag: in.JobTag,
		SkillsRequired: in.SkillsRequired,
		JobType: job.Job_DEFAULT,
		JobStatus: job.Job_ACTIVE,
		MinSallary: MinSallary,
		MaxSallary: MaxSallary,
	}

	job, err := r.server.JobClient.CreateJob(ctx, *job)
  if err != nil {
    log.Fatalf("could not greet: %v", err)
  }

	return job, nil
}

func (r *mutationResolver) CreateAccount(ctx context.Context, in models.CreateAccountReq) (*account.Account, error) {
  ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
  defer cancel()

  acc := &account.Account{
    Email: in.Email,
    Username: in.Username,
    Description: in.Description,
    PasswordHash: in.PasswordHash,
  }
  acc, err := r.server.AccountClient.CreateAccount(ctx, *acc)
  if err != nil {
    log.Fatalf("could not greet: %v", err)
  }

  return acc, nil
}

func (r *mutationResolver) Auth(ctx context.Context, in models.AuthReq) (*models.AuthRes, error) {
  ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
  defer cancel()

  acc := &account.Account{
    Email: *in.Email,
    PasswordHash: *in.PasswordHash,
  }

  token, err := r.server.AccountClient.Auth(ctx, *acc)
  if err != nil {
    log.Fatalf("could not greet: %v", err)
  }

  return &models.AuthRes{
    Token: &token.Token,
    Valid: &token.Valid,
  }, nil
}
