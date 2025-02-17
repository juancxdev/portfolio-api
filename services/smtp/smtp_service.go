package smtp

import (
	"github.com/resend/resend-go/v2"
	"portfolio-api/internal/logger"
)

type Service struct {
	client *resend.Client
}

type Server interface {
	Send(params *resend.SendEmailRequest) (string, error)
	Get(id string) (string, error)
}

func NewService(apiKey string) Server {
	return &Service{
		client: resend.NewClient(apiKey),
	}
}

func (r *Service) Send(params *resend.SendEmailRequest) (string, error) {
	sent, err := r.client.Emails.Send(params)
	if err != nil {
		return "", err
	}

	email, err := r.client.Emails.Get(sent.Id)
	if err != nil {
		logger.Error.Printf("Email enviado pero no se pudo verificar: %v", err)
	} else {
		logger.Trace.Printf("Email enviado exitosamente con ID: %s",
			email.Id)
	}

	return sent.Id, nil
}

func (r *Service) Get(id string) (string, error) {
	email, err := r.client.Emails.Get(id)
	if err != nil {
		logger.Error.Printf("Email enviado pero no se pudo verificar: %v", err)
	} else {
		logger.Trace.Printf("Email enviado exitosamente con ID: %s",
			email.Id)
	}

	return email.Id, nil
}
