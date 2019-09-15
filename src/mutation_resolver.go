package gateway

import (
  "context"
  "errors"
  "time"
  "log"

  models "github.com/nirajgeorgian/gateway/src/models"

  account "github.com/nirajgeorgian/account/src/model"
  _ "github.com/nirajgeorgian/account/src/api"
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

  return &account.Account{
    Email: acc.Email,
    Username: acc.Username,
    Description: acc.Description,
    PasswordHash: acc.PasswordHash,
  }, nil
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
