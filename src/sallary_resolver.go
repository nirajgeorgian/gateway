package gateway

import (
  "context"
  "time"

  job "github.com/nirajgeorgian/job/src/model"
)

type sallaryResolver struct {
  server *GatewayServer
}

func (r *sallaryResolver) Value(ctx context.Context, in *job.Sallary) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
  defer cancel()

	return int(in.Value), nil
}
