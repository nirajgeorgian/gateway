package gateway

import (
	"context"

	api "github.com/nirajgeorgian/gateway/src/mails/api"
	"github.com/nirajgeorgian/gateway/src/models"
)

func (c *GatewayServer) SendAccountConfirmation(ctx context.Context, in models.AccountConfirmationReq) (*api.ConfirmationRes, error) {
	r, err := api.NewMailsServiceClient(c.MailClient).SendAccountConfirmation(
		ctx,
		&api.AccountConfirmationReq{Username: in.Username, Message: in.Message, ConfirmationCode: in.ConfirmationCode, Email: in.Email},
	)
	if err != nil {
		return nil, err
	}

	return r, nil
}
