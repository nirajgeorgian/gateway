package gateway

import (
  "context"
  "time"
	"log"

  models "github.com/nirajgeorgian/gateway/src/models"

  account "github.com/nirajgeorgian/account/src/model"
)

type queryResolver struct {
  server *GatewayServer
}

func (r *queryResolver) Dummy(ctx context.Context) (*string, error)  {
  ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
  defer cancel()

  msg := "dodo duck"

  return &msg, nil
}

func (r *queryResolver) ReadAccount(ctx context.Context, in models.ReadAccountReq) (*account.Account, error) {
  ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
  defer cancel()

  acc, err := r.server.AccountClient.ReadAccount(ctx, in.AccountID)
  if err != nil {
    log.Fatalf("could not greet: %v", err)
  }

  return acc.Account, nil
}

func (r *queryResolver) ValidateUsername(ctx context.Context, in models.ValidateUsernameReq) (*models.ValidationResponse, error) {
  ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
  defer cancel()

  acc, err := r.server.AccountClient.ValidateUsername(ctx, in.Username)
  if err != nil {
    log.Fatalf("could not greet: %v", err)
  }

  return &models.ValidationResponse{Success: &acc.Success}, nil
}

func (r *queryResolver) ValidateEmail(ctx context.Context, in models.ValidateEmailReq) (*models.ValidationResponse, error) {
  ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
  defer cancel()

  acc, err := r.server.AccountClient.ValidateEmail(ctx, in.Email)
  if err != nil {
    log.Fatalf("could not greet: %v", err)
  }

  return &models.ValidationResponse{Success: &acc.Success}, nil
}
