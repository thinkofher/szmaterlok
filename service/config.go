package service

import (
	"context"
	"log"
	"os"

	env "github.com/joho/godotenv"
)

// Pathts of configuration files.
const (

	// ConfigSystemFile is the path for default system wide
	// config file. It is the first config file loaded
	// by the szmaterlok.
	ConfigSystemFile = "/etc/szmaterlok/config.env"

	// ConfigLocalFile is the path for the default local
	// config file. It is the second config file loaded
	// by the szmaterlok. Any configure variables saved
	// in this file will overwrite config variables from
	// ConfigSystemFile.
	ConfigLocalFile = ".env"
)

// Names of configuration environmental variables.
const (

	// ConfigAddressVarName is env variable for listening address.
	ConfigAddressVarName = "S8K_ADDR"

	// ConfigSessionSecretVarName is env variable for secret session password.
	ConfigSessionSecretVarName = "S8K_SESSION_SECRET"
)

// Default values for configuration variables.
const (

	// ConfigAddressDefaultVal is default value for address
	// configuration variable.
	ConfigAddressDefaultVal = "0.0.0.0:8080"

	// ConfigSessionSecretDefaultVal is default value for session
	// secret variable. Remember to change this value during
	// production deployment of szmaterlok!
	ConfigSessionSecretDefaultVal = "secret_password"
)

// ConfigVariables represents state read from environmental
// variables, which are used for configuration of szmaterlok.
type ConfigVariables struct {
	// Address is combination of IP addres and port
	// which is used for listening to TCP/IP connections.
	Address string

	// SessionSecret is secret password which is used to encrypt
	// and decrypt session state data.
	SessionSecret string
}

// ConfigLoad loads all the config files with environmental variables.
func ConfigLoad(ctx context.Context) error {
	if err := env.Load(ConfigSystemFile); err != nil {
		log.Printf("config: failed to open system config file: %s", err)
	}

	if err := env.Load(ConfigLocalFile); err != nil {
		log.Printf("config: failed to load config file: %s", err)
	}

	return nil
}

// ConfigDefault returns default configuration for szmaterlok.
func ConfigDefault() ConfigVariables {
	return ConfigVariables{
		Address:       ConfigAddressDefaultVal,
		SessionSecret: ConfigSessionSecretDefaultVal,
	}
}

// ConfigRead overwrites fields of given config variables with
// their environmental correspondent values (when they're set).
func ConfigRead(c *ConfigVariables) error {
	if addr := os.Getenv(ConfigAddressVarName); addr != "" {
		c.Address = addr
	}

	if secret := os.Getenv(ConfigSessionSecretVarName); secret != "" {
		c.SessionSecret = secret
	}

	return nil
}
