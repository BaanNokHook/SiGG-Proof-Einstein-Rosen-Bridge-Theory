package service_test

/*
 *
 *  SiGG-Proof-Einstein-Rosen-Bridge-Theory
 *
 */

import (
	"testing"

	"github.com/magiconair/properties/assert"

	"github.com/primasio/wormhole/service"
)

func TestGetIntegration(t *testing.T) {
	ig := service.GetIntegration()
	ig2 := service.GetIntegration()
	assert.Equal(t, ig, ig2)
}
