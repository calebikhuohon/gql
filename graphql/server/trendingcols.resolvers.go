package server

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"quicknode/graphql/pkg/middleware"
	"quicknode/graphql/server/generated"
	"quicknode/graphql/server/model"
)

// GetTrendingCollection is the resolver for the getTrendingCollection field.
func (r *queryResolver) GetTrendingCollection(ctx context.Context, orderBy string, orderDirection string) (*model.TrendingCollections, error) {
	if !middleware.Auth() {
		return nil, gqlerror.Errorf("user not authenticated")
	}

	trendingCols, err := r.Client.GetTrendingCollection()
	if err != nil {
		return nil, gqlerror.Errorf("failed to get trending collections")
	}
	result := &model.TrendingCollections{
		Edges: make([]*model.Edges, 0),
	}

	for _, edge := range trendingCols.Edges {
		result.Edges = append(result.Edges, &model.Edges{
			Node: &model.Node{
				Address: edge.Node.Address,
				Name:    edge.Node.Name,
				Stats: &model.Stats{
					TotalSales: edge.Node.Stats.TotalSales,
					Average:    edge.Node.Stats.Average,
					Ceiling:    edge.Node.Stats.Ceiling,
					Floor:      edge.Node.Stats.Floor,
					Volume:     edge.Node.Stats.Volume,
				},
				Symbol: edge.Node.Symbol,
			}})
	}

	return result, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
