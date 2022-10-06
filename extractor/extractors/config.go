package extractors

import (
	"github.com/spf13/viper"
)

// Config is the extractor configuration struct.
type Config struct {
	YeahdeskURL   string
	YeahdeskToken string
	EmailsPath    string
}

// setDefaults sets the default values.
func setDefaults() {
	viper.SetDefault(
		"yeahdesk_url", "https://app.yeahdesk.ru/api/clients/person/read",
	)
	viper.SetDefault("emails_path", "emails/")
}

// NewConfig returns the configuration object.
func NewConfig() *Config {
	setDefaults()
	config := &Config{
		YeahdeskURL:   viper.GetString("yeahdesk_url"),
		YeahdeskToken: viper.GetString("yeahdesk_token"),
		EmailsPath:    viper.GetString("base_dir") + viper.GetString("emails_path"),
	}

	return config
}
