package gateway

import (
  "context"
  "errors"
  "time"
  "log"
	"fmt"

  models "github.com/nirajgeorgian/gateway/src/models"
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

func (r *mutationResolver) CreateJob(ctx context.Context, in models.CreateJobReq) (*models.Job, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
  defer cancel()

	// sallary min and max for range
	MinSallary := &models.Sallary{
		Value: uint64(in.MinSallary.Value),
		Currency: in.MinSallary.Currency,
	}
	MaxSallary := &models.Sallary{
		Value: uint64(in.MaxSallary.Value),
		Currency: in.MaxSallary.Currency,
	}

	newJob := &models.Job{
		JobName: in.JobName,
		JobDescription: in.JobDescription,
		JobCategory: in.JobCategory,
		Location: *in.Location,
		JobTag: in.JobTag,
		SkillsRequired: in.SkillsRequired,
		JobType: models.Job_DEFAULT,
		JobStatus: models.Job_ACTIVE,
		MinSallary: MinSallary,
		MaxSallary: MaxSallary,
	}

	job, err := r.server.CreateJob(ctx, *newJob)
  if err != nil {
    log.Fatalf("could not greet: %v", err)
  }

	return job, nil
}

func (r *mutationResolver) CreateAccount(ctx context.Context, in models.AccountReq) (*models.Account, error) {
  ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
  defer cancel()

  acc := &models.Account{
		Email: *in.Email,
    Username: *in.Username,
    Description: *in.Description,
    PasswordHash: *in.PasswordHash,
  }
  acc, err := r.server.CreateAccount(ctx, *acc)
  if err != nil {
    log.Fatalf("could not greet: %v", err)
  }

  return acc, nil
}

func (r *mutationResolver) Auth(ctx context.Context, in models.AuthReq) (*models.AuthRes, error) {
  ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
  defer cancel()

  acc := &models.Account{
    Email: *in.Email,
    PasswordHash: *in.PasswordHash,
  }

  token, err := r.server.Auth(ctx, *acc)
  if err != nil {
    log.Fatalf("could not greet: %v", err)
  }

  return &models.AuthRes{
    Token: &token.Token,
    Valid: &token.Valid,
  }, nil
}

func (r *mutationResolver) UpdateAccount(ctx context.Context, in models.AccountReq) (*models.UpdatedAccount, error) {
  ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
  defer cancel()

	acc := &models.Account{
		AccountId: *in.AccountID,
    Email: *in.Email,
    Username: *in.Username,
    Description: *in.Description,
    PasswordHash: *in.PasswordHash,
  }

  upacc, err := r.server.UpdateAccount(ctx, *acc)
  if err != nil {
    log.Fatalf("could not greet: %v", err)
  }

  return &models.UpdatedAccount{Account: upacc.Account, Success: &upacc.Success}, nil
}

func (r *mutationResolver) SendAccountConfirmation(ctx context.Context, in models.AccountConfirmationReq) (*models.ConfirmationRes, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
  defer cancel()
	fmt.Println(in.Username)

	confirmationRes, err := r.server.SendAccountConfirmation(ctx, in)
	if err != nil {
    log.Fatalf("could not greet: %v", err)
  }

	return &models.ConfirmationRes{Status: &confirmationRes.Status}, nil
}
