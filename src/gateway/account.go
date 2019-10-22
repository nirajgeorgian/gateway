package gateway

import (
	"context"

	"github.com/nirajgeorgian/gateway/src/api"
	"github.com/nirajgeorgian/gateway/src/models"
)

func (c *GatewayServer) CreateAccount(ctx context.Context, account models.Account) (*models.Account, error) {
	r, err := api.NewAccountServiceClient(c.AccountClient).CreateAccount(
		ctx,
		&api.CreateAccountReq{Account: &account},
	)
	if err != nil {
		return nil, err
	}
	return r.Account, nil
}

func (c *GatewayServer) UpdateAccount(ctx context.Context, account models.Account) (*api.UpdateAccountRes, error) {
	r, err := api.NewAccountServiceClient(c.AccountClient).UpdateAccount(
		ctx,
		&api.UpdateAccountReq{Account: &account},
	)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (c *GatewayServer) ReadAccount(ctx context.Context, account_id string) (*api.ReadAccountRes, error) {
	r, err := api.NewAccountServiceClient(c.AccountClient).ReadAccount(
		ctx,
		&api.ReadAccountReq{AccountId: account_id},
	)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (c *GatewayServer) ValidateUsername(ctx context.Context, username string) (*api.ValidateUsernameRes, error) {
	r, err := api.NewAccountServiceClient(c.AccountClient).ValidateUsername(
		ctx,
		&api.ValidateUsernameReq{Username: username},
	)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (c *GatewayServer) ValidateEmail(ctx context.Context, email string) (*api.ValidateEmailRes, error) {
	r, err := api.NewAccountServiceClient(c.AccountClient).ValidateEmail(
		ctx,
		&api.ValidateEmailReq{Email: email},
	)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (c *GatewayServer) Auth(ctx context.Context, account models.Account) (*api.AuthRes, error) {
	r, err := api.NewAccountServiceClient(c.AccountClient).Auth(
		ctx,
		&api.AuthReq{Account: &account},
	)
	if err != nil {
		return nil, err
	}
	return r, nil
}
