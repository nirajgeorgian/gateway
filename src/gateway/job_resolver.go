package gateway

import (
  "context"
  "time"

	"go.opencensus.io/trace"
	models "github.com/nirajgeorgian/gateway/src/models"
)

type jobResolver struct {
  server *GatewayServer
}

func (r *jobResolver) JobType(ctx context.Context, in *models.Job) (int, error) {
	_, span := trace.StartSpan(ctx, "gateway.http.gateway.JobType")
	defer span.End()

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
  defer cancel()

	span.Annotate([]trace.Attribute{
		trace.StringAttribute("fetch", "JobType"),
	}, "fetch JobType from client")

	return int(in.JobType), nil
}

func (r *jobResolver) JobStatus(ctx context.Context, in *models.Job) (int, error) {
	_, span := trace.StartSpan(ctx, "gateway.http.gateway.JobStatus")
	defer span.End()

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
  defer cancel()

	span.Annotate([]trace.Attribute{
		trace.StringAttribute("fetch", "JobStatus"),
	}, "fetch JobStatus from client")

	return int(in.JobStatus), nil
}
