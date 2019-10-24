package gateway

import (
	"context"

	"go.opencensus.io/trace"
	api "github.com/nirajgeorgian/gateway/src/job/api"
	"github.com/nirajgeorgian/gateway/src/job/models"
)

func (c *GatewayServer) CreateJob(ctx context.Context, job models.Job) (*models.Job, error) {
	_, span := trace.StartSpan(ctx, "gateway.http.gateway.CreateJob")
	defer span.End()

	r, err := api.NewJobServiceClient(c.JobClient).CreateJob(
		ctx,
		&api.CreateJobReq{Job: &job},
	)
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeUnknown, Message: err.Error()})
		return nil, err
	}

	span.Annotate([]trace.Attribute{
		trace.StringAttribute("fetch", "CreateJob"),
	}, "fetch CreateJob from client")

	return r.Job, nil
}
