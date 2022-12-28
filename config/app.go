package config

type Config struct {
	HttpAddress      string `mapstructure:"HTTP_ADDRESS"`
	AllowOrigins     string `mapstructure:"ALLOW_ORIGINS"`
	PostgresHost     string `mapstructure:"POSTGRES_HOST"`
	PostgresPort     uint16 `mapstructure:"POSTGRES_PORT"`
	PostgresDatabase string `mapstructure:"POSTGRES_DATABASE"`
	PostgresMaxConns int32  `mapstructure:"POSTGRES_MAX_CONNS"`
	PostgresUsername string `mapstructure:"POSTGRES_USERNAME"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
}
