package config

import "os"

type DBSqliteConfig struct {
	URL  string
	Name string
}

func LoadDBSqliteConfig() DBSqliteConfig {
	return DBSqliteConfig{
		URL:  os.Getenv("SQLITE_DB_URL"),
		Name: os.Getenv("SQLITE_DB_NAME"),
	}
}
