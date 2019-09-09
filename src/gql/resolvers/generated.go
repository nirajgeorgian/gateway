//go:generate go run scripts/gqlgen.go -v
package resolvers

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"github.com/nirajgeorgian/gateway/src/gql"
	"github.com/nirajgeorgian/gateway/src/gql/models"

	model "github.com/nirajgeorgian/job/src/model"
	proto "github.com/nirajgeorgian/job/src/api"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() gql.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() gql.QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Subscription() gql.SubscriptionResolver {
	return &subscriptionResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) Dummy(ctx context.Context) (*string, error) {
	message := "Dodo Duck"

	return &message, nil
}
func (r *mutationResolver) CreateJob(ctx context.Context, input models.CreateJobRequest) (*model.Job, error) {
	job := model.Job{
		JobId: *input.JobID,
		JobName: *input.JobName,
	}

	// call the server
	address     := "localhost:3000"

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewJobServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	result, err := c.CreateJob(ctx, &proto.CreateJobRequest{Job: &job})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", result.Job.JobName)

	return &job, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Dummy(ctx context.Context) (*string, error) {
	message := "Dodo Duck"

	return &message, nil
}
func (r *queryResolver) Job(ctx context.Context) (*model.Job, error) {
	panic("not implemented")
}

type subscriptionResolver struct{ *Resolver }

func (r *subscriptionResolver) Dummy(ctx context.Context) (<-chan *string, error) {
	panic("not implemented")
}
