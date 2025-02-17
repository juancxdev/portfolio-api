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
func (s *graphQLRepository) create(m *EmailLog) (*EmailLogRecord, error) {
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

	variables := map[string]interface{}{
		"request": map[string]interface{}{
			"email_to":      m.EmailTo,
			"email_from":    m.EmailFrom,
			"subject":       m.Subject,
			"status":        m.Status,
			"content":       m.Content,
			"template_name": m.TemplateName,
			"template_data": m.TemplateData,
			"metadata":      m.Metadata,
		},
	}

	var response GraphQLResponse
	err := s.graphqlClient.Execute(query, variables, &response)
	if err != nil {
		return nil, fmt.Errorf("create email log error: %w", err)
	}

	// Verificar si se creó al menos un registro
	if len(response.Data.Collection.Records) == 0 {
		return nil, fmt.Errorf("no email log record was created")
	}

	// Obtener el primer registro creado
	record := response.Data.Collection.Records[0]

	// Crear la respuesta
	return &EmailLogRecord{
		ID:      record.ID,
		EmailTo: record.EmailTo,
		Status:  record.Status,
		SentAt:  record.SentAt,
	}, nil
}
