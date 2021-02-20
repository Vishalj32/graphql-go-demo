package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"graphql-go-demo/graph/generated"
	"graphql-go-demo/graph/model"
)

func (r *mutationResolver) CreateBook(ctx context.Context, title string, author string) (*model.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreatAuthor(ctx context.Context, firstName string, lastName string) (*model.Author, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) BookByID(ctx context.Context, id *string) (*model.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AllBooks(ctx context.Context) ([]*model.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

/* func (r *queryResolver) AllBooks(ctx context.Context) ([]*model.Book, error) {
	var books []*model.Book
	book := model.Book{
		Title: "Dummy Book!!",
		ID:    "0",
		Author: []*model.Author{
			{
				ID:        "1",
				FirstName: "Demo",
				LastName:  "Name",
			},
		},
	}

	books = append(books, &book)
	return books, nil

} */

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
