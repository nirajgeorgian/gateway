package gateway

import (
	"context"

	"github.com/nirajgeorgian/gateway/src/api"
	"github.com/nirajgeorgian/gateway/src/models"
)

func (c *GatewayServer) CreateJob(ctx context.Context, job models.Job) (*models.Job, error) {
	r, err := api.NewJobServiceClient(c.JobClient).CreateJob(
		ctx,
		&api.CreateJobReq{Job: &job},
	)
	if err != nil {
		return nil, err
	}
	return r.Job, nil
}
