package gateway

import (
  "context"
  "time"

  job "github.com/nirajgeorgian/job/src/model"
)

type jobResolver struct {
  server *GatewayServer
}

func (r *jobResolver) JobType(ctx context.Context, in *job.Job) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
  defer cancel()

	return int(in.JobType), nil
}

func (r *jobResolver) JobStatus(ctx context.Context, in *job.Job) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
  defer cancel()

	return int(in.JobStatus), nil
}
