package helper

import (
	"fmt"
	"portfolio-api/internal/env"
)

func String(c *env.Configuration) string {
	return fmt.Sprintf(
		"Configuration:\n"+
			"  Port: %d\n"+
			"  Log Path: %s\n"+
			"  SMTP API Key: %s\n",
		c.Port,
		c.LogConfig.Path,
		maskApiKey(c.SmtpConfig.ApiKey),
	)
}

func maskApiKey(key string) string {
	if len(key) <= 4 {
		return "****"
	}
	return key[:4] + "****" + key[len(key)-4:]
}
