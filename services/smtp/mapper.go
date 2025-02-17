package smtp

import (
	"github.com/resend/resend-go/v2"
)

func MapperSimpleEmailMessage(to, from, html, subject string, tags []string) *resend.SendEmailRequest {
	var resendTags []resend.Tag
	for _, tag := range tags {
		resendTags = append(resendTags, resend.Tag{
			Name:  "category",
			Value: tag,
		})
	}

	if len(resendTags) == 0 {
		resendTags = make([]resend.Tag, 0)
	}

	return &resend.SendEmailRequest{
		To:      []string{to},
		From:    from,
		Html:    html,
		Subject: subject,
		Tags:    resendTags,
	}
}
