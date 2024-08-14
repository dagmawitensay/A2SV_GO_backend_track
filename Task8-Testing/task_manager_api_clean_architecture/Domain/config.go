package domain

type Config struct {
	Database DatabaseConfig
	TimeZone string
	SecretKey string
}


type DatabaseConfig struct {
	DBURI string
	DbName string
	Username string
	Password string
}