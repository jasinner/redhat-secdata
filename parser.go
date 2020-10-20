package main

import (
	"context"
	"fmt"
	"os"

	"github.com/quay/claircore/libvuln"
	"github.com/quay/claircore/updater/defaults"
	"github.com/rs/zerolog/log"
)

func getEnvOrDefault(key string, def string) string {
	value := os.Getenv(key)
	if value == "" {
		value = def
	}
	return value
}

func getDBConn(ctx context.Context) string {
	password, exist := os.LookupEnv("DB_PASS")
	if exist != true {
		log.Ctx(ctx).Error().Msg("DB_PASS environment variable not found")
	}
	host := getEnvOrDefault("DB_HOST", "postgres")
	user := getEnvOrDefault("DB_USER", "postgres")
	dbname := getEnvOrDefault("DB_NAME", "claircore")
	return fmt.Sprintf("host=%v port=5432 sslmode=disable user=%v dbname=%v password=%v", host, user, dbname, password)
}

func main() {

	globalLogger := log.With().Timestamp().Logger()
	ctx := globalLogger.WithContext(context.Background())

	connection := getDBConn(ctx)

	opts := libvuln.Opts{
		ConnString: connection,
		Migrations: true,
		//Only get the RHEL data, change to nil for all Clair data
		UpdaterSets: []string{"rhel"},

		DisableBackgroundUpdates: true,
		//TODO filter on RHCOS cpes?
		//UpdaterFilter func(name string) (keep bool)
	}

	//This is necessary to load the updaters
	err := defaults.Error()
	if err != nil {
		globalLogger.Error().Err(err).Msg("error during updater initialization")
	}

	_, err = libvuln.New(ctx, &opts)
	if err != nil {
		globalLogger.Error().Err(err).Msg("Error creating vuln library")
	}
	globalLogger.Info().Msg("Update complete, will exit")
}
