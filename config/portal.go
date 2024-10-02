package config

import (
	"errors"
	"time"

	"github.com/spf13/viper"
)

type PortalCORSConfiguration struct {
	Enabled bool `mapstructure:"enabled"`
}

type PortalConfiguration struct {
	Enabled         bool                    `mapstructure:"enabled"`
	CORS            PortalCORSConfiguration `mapstructure:"cors"`
	TokenSecret     string                  `mapstructure:"tokenSecret"`
	TokenExpiration time.Duration           `mapstructure:"tokenExpiration"`
}

// Validate validates the configuration.
func (c PortalConfiguration) Validate() error {
	if !c.Enabled {
		return nil
	}

	var errs []error

	if c.TokenSecret == "" {
		errs = append(errs, errors.New("token secret is required"))
	}

	if c.TokenExpiration.Seconds() == 0 {
		errs = append(errs, errors.New("token duration is required"))
	}

	return errors.Join(errs...)
}

// ConfigurePortal configures some defaults in the Viper instance.
func ConfigurePortal(v *viper.Viper) {
	v.SetDefault("portal.enabled", false)
	v.SetDefault("portal.cors.enabled", true)
	v.SetDefault("portal.tokenSecret", "")
	v.SetDefault("portal.tokenExpiration", "1h")
}
