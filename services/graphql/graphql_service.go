package graphql_service

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
)

type Service struct {
	client *graphql.Client
}

type GraphQLClient interface {
	Execute(query string, variables map[string]interface{}, result interface{}) error
}

func NewService(urlGraphQL string) GraphQLClient {
	return &Service{
		client: graphql.NewClient(urlGraphQL),
	}
}

func (s Service) Execute(query string, variables map[string]interface{}, result interface{}) error {
	ctx := context.Background()

	req := graphql.NewRequest(query)

	req.Header.Set("Cache-Control", "no-cache")

	for key, value := range variables {
		req.Var(key, value)
	}

	if err := s.client.Run(ctx, req, result); err != nil {
		return fmt.Errorf("graphql execute error: %w", err)
	}

	return nil
}
