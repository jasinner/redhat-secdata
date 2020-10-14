package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/quay/claircore/libvuln"
	"github.com/rs/zerolog"
)

func getEnvOrDefault(key string, def string) string {
	value := os.Getenv(key)
	if value == "" {
		value = def
	}
	return value
}

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	fmt.Println("Set Global log level")

	password, exist := os.LookupEnv("DB_PASS")
	if exist != true {
		log.Fatal("DB_PASS not found in environment")
	}
	host := getEnvOrDefault("DB_HOST", "postgres")
	user := getEnvOrDefault("DB_USER", "postgres")
	dbname := getEnvOrDefault("DB_NAME", "claircore")
	connection := fmt.Sprintf("host=%v port=5432 sslmode=disable user=%v dbname=%v password=%v", host, user, dbname, password)
	opts := libvuln.Opts{
		//MaxConnPool int32
		ConnString: connection,
		//UpdateInterval time.Duration
		// Determines if Livuln will manage database migrations
		Migrations: true,

		UpdaterSets: []string{"rhel"},
		// A list of out-of-tree updaters to run.
		//
		// This list will be merged with any defined UpdaterSets.
		//
		// If you desire no updaters to run do not add an updater
		// into this slice.
		//Updaters []driver.Updater
		// A list of out-of-tree matchers you'd like libvuln to
		// use.
		//
		// This list will me merged with the default matchers.
		//Matchers []driver.Matcher

		// If set to true, there will not be a goroutine launched to periodically
		// run updaters.
		//DisableBackgroundUpdates false

		// UpdaterConfigs is a map of functions for configuration of Updaters.
		//UpdaterConfigs map[string]driver.ConfigUnmarshaler

		//UpdaterFilter func(name string) (keep bool)
	}

	ctx := context.TODO()
	libvuln, err := libvuln.New(ctx, &opts)
	fmt.Printf("created lib vuln %v\n", libvuln)
	if err != nil {
		log.Fatal(err)
	}
}
