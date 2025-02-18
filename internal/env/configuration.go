package env

import (
	"flag"
	"fmt"
	"os"
	"sync"
)

var (
	config = &Configuration{}
	once   sync.Once
)

var (
	portApp        *int
	secretApp      *string
	pathLog        *string
	smtpApiKey     *string
	smtpTo         *string
	supaBaseURL    *string
	supaBaseAPIKey *string
)

type Configuration struct {
	Port           int      `json:"port"`
	SecretApp      string   `json:"secret_app"`
	LogConfig      Log      `json:"log_config"`
	SmtpConfig     Smtp     `json:"smtp_config"`
	SupaBaseConfig SupaBase `json:"supabase_config"`
}

type Log struct {
	Path string `json:"path"`
}

type Smtp struct {
	To     string `json:"to"`
	ApiKey string `json:"api_key"`
}

type SupaBase struct {
	SupaBaseURL    string `json:"supabase_url"`
	SupaBaseAPIKey string `json:"supabase_api_key"`
}

func init() {
	// Definir los flags una sola vez durante la inicialización
	portApp = flag.Int("port", 0, "Puerto para exponer servicio (required)")
	secretApp = flag.String("secret", "", "Secreto para encryptar información")
	pathLog = flag.String("log-path", "", "Ruta de logs de aplicación")
	smtpApiKey = flag.String("smtp-api-key", "", "Api key del smtp")
	smtpTo = flag.String("smtp-to", "", "To del smtp")
	supaBaseURL = flag.String("supabase-url", "", "URL de SupaBase")
	supaBaseAPIKey = flag.String("supabase-api-key", "", "Api key de SupaBase")
}

func NewConfiguration() *Configuration {
	once.Do(func() {
		if !flag.Parsed() {
			flag.Parse()
		}

		if err := validateAndAssignConfig(); err != nil {
			fmt.Fprintf(os.Stderr, "Error en la configuración: %v\n", err)
			flag.Usage()
			os.Exit(1)
		}
	})
	return config
}

func validateAndAssignConfig() error {
	// Validar puerto
	if *portApp <= 0 || *portApp > 65535 {
		return fmt.Errorf("el puerto debe estar entre 1 y 65535, valor actual: %d", *portApp)
	}
	config.Port = *portApp

	if *secretApp == "" {
		return fmt.Errorf("secret es requerida")
	}

	config.SecretApp = *secretApp

	// Validar ruta de logs
	if *pathLog == "" {
		return fmt.Errorf("la ruta de logs es requerida")
	}

	config.LogConfig.Path = *pathLog

	// Validar API Key de SMTP
	if *smtpApiKey == "" {
		return fmt.Errorf("la API key de SMTP es requerida")
	}

	if len(*smtpApiKey) < 8 {
		return fmt.Errorf("la API key de SMTP debe tener al menos 8 caracteres")
	}
	config.SmtpConfig.ApiKey = *smtpApiKey

	if *smtpTo == "" {
		return fmt.Errorf("to de SMTP es requerida")
	}

	config.SmtpConfig.To = *smtpTo

	if *supaBaseURL == "" {
		return fmt.Errorf("to de SMTP es requerida")
	}

	config.SupaBaseConfig.SupaBaseURL = *supaBaseURL

	if *supaBaseAPIKey == "" {
		return fmt.Errorf("to de SMTP es requerida")
	}

	config.SupaBaseConfig.SupaBaseAPIKey = *supaBaseAPIKey

	return nil
}
