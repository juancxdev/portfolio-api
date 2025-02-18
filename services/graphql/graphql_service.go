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
	Execute(query string, variables []*Variable, headers []*Header, result interface{}) error
}

func NewService(urlGraphQL string) GraphQLClient {
	return &Service{
		client: graphql.NewClient(urlGraphQL),
	}
}

func (s Service) Execute(query string, variables []*Variable, headers []*Header, result interface{}) error {
	ctx := context.Background()

	req := graphql.NewRequest(query)

	if headers != nil {
		for _, header := range headers {
			req.Header.Set(header.Key, header.Value)
		}
	}

	if variables != nil {
		for _, variable := range variables {
			req.Var(variable.Key, variable.Value)
		}
	}

	if err := s.client.Run(ctx, req, result); err != nil {
		return fmt.Errorf("graphql execute error: %w", err)
	}

	return nil
}
