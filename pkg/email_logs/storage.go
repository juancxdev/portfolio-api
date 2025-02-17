package email_logs

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
	Oracle     = "oci8"
)

type ServicesEmailLogsRepository interface {
	create(m *EmailLog) error
}

func FactoryStorage(graphqlURL string) ServicesEmailLogsRepository {
	return newGraphQLRepository(graphqlURL)
}
