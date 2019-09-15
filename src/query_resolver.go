package gateway

import (
  "context"
  "time"
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
