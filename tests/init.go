package tests

/*
 *
 *  SiGG-Proof-Einstein-Rosen-Bridge-Theory
 *
 */

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/primasio/wormhole/cache"
	"github.com/primasio/wormhole/config"
	"github.com/primasio/wormhole/db"
	"github.com/primasio/wormhole/db/migrations"
)

func InitTestEnv(configPath string) {
	environment := flag.String("e", "test", "")

	flag.Parse()

	// Init Config
	config.Init(*environment, &configPath)

	// Init Database
	if err := db.Init(); err != nil {
		log.Println("Database:", err)
		os.Exit(1)
	}

	// Init Cache
	if err := cache.InitCache(); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	if err := migrations.Migrate(); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	rand.Seed(time.Now().UnixNano())
}
