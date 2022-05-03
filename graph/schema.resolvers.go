package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/xyzuba/docs-server/graph/generated"
	"github.com/xyzuba/docs-server/graph/model"
)

func (r *mutationResolver) CreateSnippet(ctx context.Context, input *model.SnippetInput) (*model.Snippet, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Snippet(ctx context.Context, id string) (*model.Snippet, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Snippets(ctx context.Context) ([]*model.Snippet, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
