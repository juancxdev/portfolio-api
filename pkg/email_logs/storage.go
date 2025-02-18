package email_logs

import graphql_service "portfolio-api/services/graphql"

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
	Oracle     = "oci8"
)

type ServicesEmailLogsRepository interface {
	create(variables []*graphql_service.Variable, headers []*graphql_service.Header) error
}

func FactoryStorage(graphqlURL string) ServicesEmailLogsRepository {
	return newGraphQLRepository(graphqlURL)
}
