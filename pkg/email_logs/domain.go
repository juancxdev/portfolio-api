package email_logs

import (
	graphql_service "portfolio-api/services/graphql"
)

type DataGraph struct {
	Data struct {
		InsertIntoemailLogCollection struct {
			Records []struct {
				ID      string `json:"id"`
				EmailTo string `json:"email_to"`
				Status  string `json:"status"`
				SentAt  string `json:"sent_at"`
			} `json:"records"`
		} `json:"insertIntoemail_logCollection"`
	} `json:"data"`
}

func NewEmailLog(idRegister, emailTo, emailFrom, subject, content, status string) []*graphql_service.Variable {
	return []*graphql_service.Variable{
		{
			Key: "request",
			Value: map[string]interface{}{
				"id_register": idRegister,
				"email_to":    emailTo,
				"email_from":  emailFrom,
				"subject":     subject,
				"status":      status,
				"content":     content,
			},
		},
	}
}

func NewHeadersEmailLog(apiKey string) []*graphql_service.Header {
	return []*graphql_service.Header{
		{Key: "Cache-Control", Value: "no-cache"},
		{Key: "Content-Type", Value: "application/json"},
		{Key: "apiKey", Value: apiKey},
	}
}
