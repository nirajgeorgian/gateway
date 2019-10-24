package gateway

import (
  "context"
  "time"

	"go.opencensus.io/trace"
	jobmodels "github.com/nirajgeorgian/gateway/src/job/models"
)

type sallaryResolver struct {
  server *GatewayServer
}

func (r *sallaryResolver) Value(ctx context.Context, in *jobmodels.Sallary) (int, error) {
	_, span := trace.StartSpan(ctx, "gateway.http.gateway.Value")
	defer span.End()

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
  defer cancel()

	span.Annotate([]trace.Attribute{
		trace.StringAttribute("fetch", "Value"),
	}, "fetch Value from client")

	return int(in.Value), nil
}
