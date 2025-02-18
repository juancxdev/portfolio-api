package pkg

import "portfolio-api/pkg/email_logs"

type ServerPKG struct {
	SrvEmailLog email_logs.PortsServerEmailLogs
}

func NewServerPKG(graphqlURL string) *ServerPKG {
	return &ServerPKG{
		SrvEmailLog: email_logs.NewEmailLogsService(email_logs.FactoryStorage(graphqlURL)),
	}
}
