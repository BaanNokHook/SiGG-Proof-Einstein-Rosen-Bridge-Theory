package models_test

/*
 *
 *  SiGG-Proof-Einstein-Rosen-Bridge-Theory
 *
 */

import (
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/primasio/wormhole/models"
)

func TestIncrementCommentVote(t *testing.T) {
	user := models.User{CommentUpVotes: 30, CommentDownVotes: 20}

	like := true
	user.IncrementCommentVote(like)
	assert.Equal(t, user.CommentUpVotes, uint(31))
	assert.Equal(t, user.CommentDownVotes, uint(20))

	like = false
	user.IncrementCommentVote(like)
	assert.Equal(t, user.CommentUpVotes, uint(31))
	assert.Equal(t, user.CommentDownVotes, uint(21))
}

func TestSwitchCommentVote(t *testing.T) {
	user := models.User{CommentUpVotes: 30, CommentDownVotes: 20}

	like := true
	user.SwitchCommentVote(like)
	assert.Equal(t, user.CommentUpVotes, uint(31))
	assert.Equal(t, user.CommentDownVotes, uint(19))

	like = false
	user.SwitchCommentVote(like)
	assert.Equal(t, user.CommentUpVotes, uint(30))
	assert.Equal(t, user.CommentDownVotes, uint(20))
}

func TestCancelCommentVote(t *testing.T) {
	user := models.User{CommentUpVotes: 30, CommentDownVotes: 20}

	like := true
	user.CancelCommentVote(like)
	assert.Equal(t, user.CommentUpVotes, uint(29))
	assert.Equal(t, user.CommentDownVotes, uint(20))

	like = false
	user.CancelCommentVote(like)
	assert.Equal(t, user.CommentUpVotes, uint(29))
	assert.Equal(t, user.CommentDownVotes, uint(19))
}

func TestIncrementIntegration(t *testing.T) {
	user := models.User{CommentUpVotes: 30, CommentDownVotes: 20, Integration: 55}

	user.IncrementIntegration(5)
	assert.Equal(t, user.Integration, int64(60))

	user.IncrementIntegration(-5)
	assert.Equal(t, user.Integration, int64(55))
}
