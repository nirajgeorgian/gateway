package gateway

import (
  "context"
  "time"
	"log"

	"go.opencensus.io/trace"
  "github.com/nirajgeorgian/gateway/src/models"
)

type queryResolver struct {
  server *GatewayServer
}

func (r *queryResolver) Dummy(ctx context.Context) (*string, error)  {
	_, span := trace.StartSpan(ctx, "gateway.http.gateway.Dummy")
	defer span.End()

  ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
  defer cancel()

  msg := "dodo duck"

	span.Annotate([]trace.Attribute{
		trace.StringAttribute("fetch", "Dummy"),
	}, "fetch Dummy from client")

  return &msg, nil
}

func (r *queryResolver) ReadAccount(ctx context.Context, in models.ReadAccountReq) (*models.Account, error) {
	_, span := trace.StartSpan(ctx, "gateway.http.gateway.ReadAccount")
	defer span.End()

  ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
  defer cancel()

  acc, err := r.server.ReadAccount(ctx, in.AccountID)
  if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeUnknown, Message: err.Error()})
    log.Fatalf("could not greet: %v", err)
  }

	span.Annotate([]trace.Attribute{
		trace.StringAttribute("fetch", "ReadAccount"),
	}, "fetch ReadAccount from client")

  return acc.Account, nil
}

func (r *queryResolver) ValidateUsername(ctx context.Context, in models.ValidateUsernameReq) (*models.ValidationResponse, error) {
	_, span := trace.StartSpan(ctx, "gateway.http.gateway.ValidateUsername")
	defer span.End()

  ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
  defer cancel()

  acc, err := r.server.ValidateUsername(ctx, in.Username)
  if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeUnknown, Message: err.Error()})
    log.Fatalf("could not greet: %v", err)
  }

	span.Annotate([]trace.Attribute{
		trace.StringAttribute("fetch", "ValidateUsername"),
	}, "fetch ValidateUsername from client")

  return &models.ValidationResponse{Success: &acc.Success}, nil
}

func (r *queryResolver) ValidateEmail(ctx context.Context, in models.ValidateEmailReq) (*models.ValidationResponse, error) {
	_, span := trace.StartSpan(ctx, "gateway.http.gateway.ValidateEmail")
	defer span.End()

  ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
  defer cancel()

  acc, err := r.server.ValidateEmail(ctx, in.Email)
  if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeUnknown, Message: err.Error()})
    log.Fatalf("could not greet: %v", err)
  }

	span.Annotate([]trace.Attribute{
		trace.StringAttribute("fetch", "ValidateEmail"),
	}, "fetch ValidateEmail from client")

  return &models.ValidationResponse{Success: &acc.Success}, nil
}
