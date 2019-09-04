package resolvers

import (
	"context"

	"github.com/nirajgeorgian/gateway/src/gql"
	"github.com/nirajgeorgian/gateway/src/gql/models"
	job "github.com/nirajgeorgian/job/src/proto"
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
	panic("not implemented")
}
func (r *mutationResolver) CreateJob(ctx context.Context, input models.CreateJobRequest) (*job.Job, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Dummy(ctx context.Context) (*string, error) {
	panic("not implemented")
}
func (r *queryResolver) Job(ctx context.Context) (*job.Job, error) {
	panic("not implemented")
}

type subscriptionResolver struct{ *Resolver }

func (r *subscriptionResolver) Dummy(ctx context.Context) (<-chan *string, error) {
	panic("not implemented")
}
