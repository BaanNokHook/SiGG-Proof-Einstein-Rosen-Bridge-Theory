package migrations

/*
 *
 *  SiGG-Proof-Einstein-Rosen-Bridge-Theory
 *
 */

import (
	"flag"
	"github.com/magiconair/properties/assert"
	"github.com/primasio/wormhole/config"
	"github.com/primasio/wormhole/db"
	"github.com/primasio/wormhole/db/migrations"
	"log"
	"os"
	"testing"
)

func TestMigrate(t *testing.T) {

	// Manually initialize test environment

	environment := flag.String("e", "test", "")

	flag.Parse()

	// Init Config
	path := "../../config/"
	err := config.Init(*environment, &path)
	assert.Equal(t, err, nil)

	// Init Database
	if err := db.Init(); err != nil {
		log.Println("Database:", err)
		os.Exit(1)
	}

	err = migrations.Migrate()
	assert.Equal(t, err, nil)
}
