package postgressDb

import (
	"fmt"

	"github.com/aeriech/social/internal/env"
)

func GetDbConfig() DbConfig {
	return DbConfig{
		host:     env.GetString("DB_HOST", "localhost"),
		user:     env.GetString("DB_USER", "user"),
		password: env.GetString("DB_PASSWORD", "password"),
		dbName:   env.GetString("DB_NAME", "social"),
		port:     env.GetString("DB_PORT", "5432"),
		sslMode:  env.GetString("DB_SSL_MODE", "disable"),
		timeZone: env.GetString("DB_TIMEZONE", "Asia/Shanghai"),
	}
}

func GetDns() string {
	db := GetDbConfig()

	dnsFormat := "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s"

	return fmt.Sprintf(dnsFormat, db.host, db.user, db.password, db.dbName, db.port, db.sslMode, db.timeZone)
}
