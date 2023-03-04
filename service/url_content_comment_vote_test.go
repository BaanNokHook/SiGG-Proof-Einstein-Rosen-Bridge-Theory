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

func TestGetURLContentCommentVote(t *testing.T) {
	uccVote := service.GetURLContentCommentVote()
	uccVote2 := service.GetURLContentCommentVote()
	assert.Equal(t, uccVote, uccVote2)
}
