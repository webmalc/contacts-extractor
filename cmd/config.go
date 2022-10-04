package cmd

import "github.com/spf13/viper"

// Config is the cmd configuration struct.
type Config struct {
	Sources  []string
	Contacts []string
}

// setDefaults sets the default values.
func setDefaults() {
	viper.SetDefault("sources", []string{"yeahdesk", "email"})
	viper.SetDefault("contacts", []string{"email", "phone", "vk"})
}

// NewConfig returns the configuration object.
func NewConfig() *Config {
	setDefaults()
	config := &Config{
		Sources:  viper.GetStringSlice("sources"),
		Contacts: viper.GetStringSlice("contacts"),
	}

	return config
}
