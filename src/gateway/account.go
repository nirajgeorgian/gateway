package gateway

import (
	"context"

	"go.opencensus.io/trace"
	api "github.com/nirajgeorgian/gateway/src/account/api"
	accountmodel "github.com/nirajgeorgian/gateway/src/account/models"
)

func (c *GatewayServer) CreateAccount(ctx context.Context, account accountmodel.Account) (*accountmodel.Account, error) {
	_, span := trace.StartSpan(ctx, "gateway.http.gateway.CreateAccount")
	defer span.End()

	r, err := api.NewAccountServiceClient(c.AccountClient).CreateAccount(
		ctx,
		&api.CreateAccountReq{Account: &account},
	)
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeUnknown, Message: err.Error()})
		return nil, err
	}
	defer span.End()

	span.Annotate([]trace.Attribute{
		trace.StringAttribute("fetch", "CreateAccount"),
	}, "fetch CreateAccount from client")

	return r.Account, nil
}

func (c *GatewayServer) UpdateAccount(ctx context.Context, account accountmodel.Account) (*api.UpdateAccountRes, error) {
	_, span := trace.StartSpan(ctx, "gateway.http.gateway.UpdateAccount")
	defer span.End()

	r, err := api.NewAccountServiceClient(c.AccountClient).UpdateAccount(
		ctx,
		&api.UpdateAccountReq{Account: &account},
	)
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeUnknown, Message: err.Error()})
		return nil, err
	}

	span.Annotate([]trace.Attribute{
		trace.StringAttribute("fetch", "UpdateAccount"),
	}, "fetch UpdateAccount from client")

	return r, nil
}

func (c *GatewayServer) ReadAccount(ctx context.Context, account_id string) (*api.ReadAccountRes, error) {
	_, span := trace.StartSpan(ctx, "gateway.http.gateway.ReadAccount")
	defer span.End()

	r, err := api.NewAccountServiceClient(c.AccountClient).ReadAccount(
		ctx,
		&api.ReadAccountReq{AccountId: account_id},
	)
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeUnknown, Message: err.Error()})
		return nil, err
	}

	span.Annotate([]trace.Attribute{
		trace.StringAttribute("fetch", "ReadAccount"),
	}, "fetch ReadAccount from client")

	return r, nil
}

func (c *GatewayServer) ValidateUsername(ctx context.Context, username string) (*api.ValidateUsernameRes, error) {
	_, span := trace.StartSpan(ctx, "gateway.http.gateway.ValidateUsername")
	defer span.End()

	r, err := api.NewAccountServiceClient(c.AccountClient).ValidateUsername(
		ctx,
		&api.ValidateUsernameReq{Username: username},
	)
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeUnknown, Message: err.Error()})
		return nil, err
	}

	span.Annotate([]trace.Attribute{
		trace.StringAttribute("fetch", "ValidateUsername"),
	}, "fetch ValidateUsername from client")

	return r, nil
}

func (c *GatewayServer) ValidateEmail(ctx context.Context, email string) (*api.ValidateEmailRes, error) {
	_, span := trace.StartSpan(ctx, "gateway.http.gateway.ValidateEmail")
	defer span.End()

	r, err := api.NewAccountServiceClient(c.AccountClient).ValidateEmail(
		ctx,
		&api.ValidateEmailReq{Email: email},
	)
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeUnknown, Message: err.Error()})
		return nil, err
	}

	span.Annotate([]trace.Attribute{
		trace.StringAttribute("fetch", "ValidateEmail"),
	}, "fetch ValidateEmail from client")

	return r, nil
}

func (c *GatewayServer) Auth(ctx context.Context, account accountmodel.Account) (*api.AuthRes, error) {
	_, span := trace.StartSpan(ctx, "gateway.http.gateway.Auth")
	defer span.End()

	r, err := api.NewAccountServiceClient(c.AccountClient).Auth(
		ctx,
		&api.AuthReq{Account: &account},
	)
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeUnknown, Message: err.Error()})
		return nil, err
	}

	span.Annotate([]trace.Attribute{
		trace.StringAttribute("fetch", "Auth"),
	}, "fetch Auth from client")

	return r, nil
}
