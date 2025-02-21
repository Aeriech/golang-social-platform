package postgressDb

import (
	"fmt"

	"github.com/aeriech/social/internal/env"
)

func GetDbConfig() DbConfig {
	return DbConfig{
		host:     env.GetString("DB_HOST", ""),
		user:     env.GetString("DB_USER", ""),
		password: env.GetString("DB_PASSWORD", ""),
		dbName:   env.GetString("DB_NAME", ""),
		port:     env.GetString("DB_PORT", ""),
		sslMode:  env.GetString("DB_SSL_MODE", ""),
		timeZone: env.GetString("DB_TIMEZONE", ""),
	}
}

func GetDns() string {
	db := GetDbConfig()

	dnsFormat := "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s"

	return fmt.Sprintf(dnsFormat, db.host, db.user, db.password, db.dbName, db.port, db.sslMode, db.timeZone)
}
