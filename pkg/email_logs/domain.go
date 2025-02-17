package email_logs

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type GraphQLResponse struct {
	Data DataEmailLogResponse `json:"data"`
}

type DataEmailLogResponse struct {
	Collection InsertIntoemailLogCollection `json:"insertIntoemail_logCollection"`
}

type InsertIntoemailLogCollection struct {
	Records []EmailLogRecord `json:"records"`
}

type EmailLogRecord struct {
	ID      string    `json:"id"`
	EmailTo string    `json:"email_to"`
	Status  string    `json:"status"`
	SentAt  time.Time `json:"sent_at"`
}

type EmailLog struct {
	EmailTo      string      `json:"email_to" valid:"email,required"`
	EmailFrom    string      `json:"email_from" valid:"email,required"`
	Subject      string      `json:"subject" valid:"required"`
	Status       string      `json:"status" valid:"required"`
	Content      string      `json:"content" valid:"required"`
	TemplateName string      `json:"template_name" valid:"required"`
	TemplateData interface{} `json:"template_data"`
	Metadata     interface{} `json:"metadata"`
}

func NewEmailLog(emailTo, emailFrom, subject, content, status, templateName string, templateData, metadata interface{}) *EmailLog {
	return &EmailLog{
		EmailTo:      emailTo,
		EmailFrom:    emailFrom,
		Subject:      subject,
		Status:       status,
		Content:      content,
		TemplateName: templateName,
		TemplateData: templateData,
		Metadata:     metadata,
	}
}

func (m *EmailLog) Valid() (bool, error) {
	return govalidator.ValidateStruct(m)
}
