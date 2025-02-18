package email_logs

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"portfolio-api/services/graphql"
)

// sqlServer estructura de conexión a la BD de mssql
type graphQLRepository struct {
	graphqlClient graphql_service.GraphQLClient
}

func newGraphQLRepository(graphqlURL string) *graphQLRepository {
	return &graphQLRepository{
		graphqlClient: graphql_service.NewService(graphqlURL),
	}
}

// Create registra en la BD
func (s *graphQLRepository) create(variables []*graphql_service.Variable, headers []*graphql_service.Header) error {
	query := `
        mutation CreateEmailLog($request: email_logInsertInput!) {
  insertIntoemail_logCollection(
    objects: [$request]
  ) {
    records {
      id
      email_to
      status
      sent_at
    }
  }
}
    `

	var response DataGraph
	err := s.graphqlClient.Execute(query, variables, headers, &response)
	if err != nil {
		return fmt.Errorf("create email log error: %w", err)
	}

	if len(response.Data.InsertIntoemailLogCollection.Records) == 0 {
		return fmt.Errorf("no email log record was created")
	}

	return nil
}
