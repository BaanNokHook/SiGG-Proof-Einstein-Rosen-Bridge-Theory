package main

/*
 *
 *  SiGG-Proof-Einstein-Rosen-Bridge-Theory
 *
 */

import (
	"flag"
	"math/rand"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/primasio/wormhole/cache"
	"github.com/primasio/wormhole/config"
	"github.com/primasio/wormhole/db"
	"github.com/primasio/wormhole/db/migrations"
	"github.com/primasio/wormhole/http/server"
	"github.com/primasio/wormhole/worker"
)

func main() {

	// Init random seed
	rand.Seed(time.Now().UnixNano())

	migrate := flag.Bool("migrate", false, "whether to run the database migration")

	flag.Parse()

	// Init Environment
	env := os.Getenv("APP_ENV")

	if env == "" {
		env = config.AppEnvDevelopment
	}

	if env == config.AppEnvProduction {
		gin.SetMode(gin.ReleaseMode)
	}

	// Init Config
	configDir := os.Getenv("APP_CONFIG")

	var err error

	if configDir != "" {
		err = config.Init(env, &configDir)
	} else {
		err = config.Init(env, nil)
	}

	if err != nil {
		glog.Error(err)
		os.Exit(1)
	}

	// Init Database
	if err := db.Init(); err != nil {
		glog.Error(err)
		os.Exit(1)
	}

	// Run database migration

	// In a large scale production level deployment
	// we might want to run the migration separately

	if *migrate {
		if err := migrations.Migrate(); err != nil {
			glog.Error(err)
			os.Exit(1)
		}
	}

	// Init Cache
	if err := cache.InitCache(); err != nil {
		glog.Error(err)
		os.Exit(1)
	}

	w := worker.NewRegisterIntegrationWorker()
	go w.Run()

	// Start HTTP server
	server.Init()
}
