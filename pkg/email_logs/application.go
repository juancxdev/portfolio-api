package email_logs

import "portfolio-api/internal/logger"

type PortsServerEmailLogs interface {
	CreateRegister(idRegister, emailTo, emailFrom, subject, content, status, apiKey string) error
}

type service struct {
	repository ServicesEmailLogsRepository
}

func NewEmailLogsService(repository ServicesEmailLogsRepository) PortsServerEmailLogs {
	return &service{repository: repository}
}

func (s service) CreateRegister(idRegister, emailTo, emailFrom, subject, content, status, apiKey string) error {
	variables := NewEmailLog(idRegister, emailTo, emailFrom, subject, content, status)
	headers := NewHeadersEmailLog(apiKey)

	if err := s.repository.create(variables, headers); err != nil {
		if err.Error() == "rows affected error" {
			return nil
		}
		logger.Error.Println(" - couldn't create log :", err)
		return err
	}
	return nil
}
