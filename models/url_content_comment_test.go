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

func TestIncrementVote(t *testing.T) {
	comment := models.URLContentComment{CommentUpVotes: 30, CommentDownVotes: 20}

	like := true
	comment.IncrementVote(like)
	assert.Equal(t, comment.CommentUpVotes, uint(31))
	assert.Equal(t, comment.CommentDownVotes, uint(20))

	like = false
	comment.IncrementVote(like)
	assert.Equal(t, comment.CommentUpVotes, uint(31))
	assert.Equal(t, comment.CommentDownVotes, uint(21))
}

func TestSwitchVote(t *testing.T) {
	comment := models.URLContentComment{CommentUpVotes: 30, CommentDownVotes: 20}

	like := true
	comment.SwitchVote(like)
	assert.Equal(t, comment.CommentUpVotes, uint(31))
	assert.Equal(t, comment.CommentDownVotes, uint(19))

	like = false
	comment.SwitchVote(like)
	assert.Equal(t, comment.CommentUpVotes, uint(30))
	assert.Equal(t, comment.CommentDownVotes, uint(20))
}

func TestCancelVote(t *testing.T) {
	comment := models.URLContentComment{CommentUpVotes: 30, CommentDownVotes: 20}

	like := true
	comment.CancelVote(like)
	assert.Equal(t, comment.CommentUpVotes, uint(29))
	assert.Equal(t, comment.CommentDownVotes, uint(20))

	like = false
	comment.CancelVote(like)
	assert.Equal(t, comment.CommentUpVotes, uint(29))
	assert.Equal(t, comment.CommentDownVotes, uint(19))
}
