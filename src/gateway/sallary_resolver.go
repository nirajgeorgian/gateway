package gateway

import (
  "context"
  "time"

	models "github.com/nirajgeorgian/gateway/src/models"
)

type sallaryResolver struct {
  server *GatewayServer
}

func (r *sallaryResolver) Value(ctx context.Context, in *models.Sallary) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
  defer cancel()

	return int(in.Value), nil
}
