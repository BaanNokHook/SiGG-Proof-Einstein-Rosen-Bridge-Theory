package models_test

/*
 *
 *  SiGG-Proof-Einstein-Rosen-Bridge-Theory
 *
 */

import (
	"crypto/sha1"
	"fmt"
	"io"
	"testing"

	"github.com/magiconair/properties/assert"

	"github.com/primasio/wormhole/models"
)

func TestSetUniqueID(t *testing.T) {
	history := models.IntegrationHistory{}

	err := history.SetUniqueID()
	if err == nil {
		t.Errorf("err should be nil, but got error: %s", err)
	}

	history.Data = "Hello, there"
	history.SetUniqueID()

	h := sha1.New()
	io.WriteString(h, history.Data)
	uid := fmt.Sprintf("%x", h.Sum(nil))
	assert.Equal(t, uid, history.UniqueID)
}
