package server

import "quicknode/graphql/client"

type Client interface {
	GetTrendingCollection() (client.TrendingCollections, error)
}

type Resolver struct {
	Client Client
}
