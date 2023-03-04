package models

/*
 *
 *  SiGG-Proof-Einstein-Rosen-Bridge-Theory
 *
 */

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/primasio/wormhole/util"
)

type URLContentComment struct {
	BaseModel

	UniqueID     string `gorm:"type:varchar(128);unique_index" json:"id"`
	UserID       uint   `json:"-"`
	URLContentId uint   `json:"-"`
	Content      string `gorm:"type:longtext" json:"content"`

	CommentUpVotes   uint `json:"comment_up_votes" gorm:"type:INT(11);default:0"`
	CommentDownVotes uint `json:"comment_down_votes" gorm:"type:INT(11);default:0"`

	User User `gorm:"save_associations:false" json:"user"`

	IsDeleted bool `gorm:"default:false" json:"is_deleted"`
}

func (comment *URLContentComment) SetUniqueID(db *gorm.DB) error {
	var counter = 0

	for {
		counter = counter + 1
		uid := util.RandStringUppercase(8)

		check := &URLContentComment{UniqueID: uid}

		db.Where(&check).First(&check)

		if check.ID == 0 {
			comment.UniqueID = uid
			return nil
		}

		if counter >= 5 {
			// This is unlikely to happen
			// Must be error from other parts
			return errors.New("too many iterations while generating new session key")
		}
	}
}

func (comment *URLContentComment) IncrementVote(like bool) {
	if like {
		comment.CommentUpVotes++
	} else {
		comment.CommentDownVotes++
	}
}

func (comment *URLContentComment) SwitchVote(like bool) {
	if like {
		comment.CommentUpVotes++
		comment.CommentDownVotes--
	} else {
		comment.CommentUpVotes--
		comment.CommentDownVotes++
	}
}

func (comment *URLContentComment) CancelVote(like bool) {
	if like {
		comment.CommentUpVotes--
	} else {
		comment.CommentDownVotes--
	}
}
