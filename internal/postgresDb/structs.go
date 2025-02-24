package postgresDb

type DbConfig struct {
	host     string
	user     string
	password string
	dbName   string
	port     string
	sslMode  string
	timeZone string
}
