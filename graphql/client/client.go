package client

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
)

type TrendingCollections struct {
	Edges []Edges `json:"edges"`
}

type Edges struct {
	Node Node `json:"node"`
}

type Node struct {
	Address string  `json:"address"`
	Name    string  `json:"name"`
	Stats   Stats   `json:"stats"`
	Symbol  *string `json:"symbol"`
}

type Stats struct {
	TotalSales int     `json:"totalSales"`
	Average    float64 `json:"average"`
	Ceiling    float64 `json:"ceiling"`
	Floor      float64 `json:"floor"`
	Volume     float64 `json:"volume"`
}

type Client struct {
	ApiKey string
}

func NewClient(key string) *Client {
	return &Client{ApiKey: key}
}

func (c *Client) GetTrendingCollection() (TrendingCollections, error) {
	client := graphql.NewClient("https://graphql.icy.tools/graphql")

	req := graphql.NewRequest(`query TrendingCollections {
    trendingCollections(orderBy: SALES, orderDirection: DESC) {
      edges {
        node {
          address
          ... on ERC721Contract {
            name
            stats {
              totalSales
              average
              ceiling
              floor
              volume
            }
            symbol
          }
        }
      }
    }
  }`)

	req.Header.Set("Cache-Control", "no-cache")
	fmt.Println(c.ApiKey)
	req.Header.Set("x-api-key", c.ApiKey)

	ctx := context.Background()

	var response struct {
		Data struct {
			TrendingCollections TrendingCollections `json:"trendingCollections"`
		} `json:"data"`
	}
	if err := client.Run(ctx, req, &response.Data); err != nil {
		return TrendingCollections{}, err
	}
	fmt.Println("res: ", response)

	return response.Data.TrendingCollections, nil
}
