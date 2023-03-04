package models

/*
 *
 *  SiGG-Proof-Einstein-Rosen-Bridge-Theory
 *
 */

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"

	"github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
)

type URLContentCommentVote struct {
	BaseModel
	UniqueID            string `json:"id" gorm:"type:varchar(128);unique_index"`
	UserID              uint   `json:"-"`
	URLContentCommentID uint   `json:"-"`
	Like                bool   `json:"like"`

	User User `gorm:"save_associations:false" json:"user"`
}

func (comment *URLContentCommentVote) SetUniqueID() error {

	if comment.UserID == 0 || comment.URLContentCommentID == 0 {
		return errors.New("UserID Or URLContentCommentID Zero")
	}

	h := sha1.New()
	io.WriteString(h, fmt.Sprintf("%d%d", comment.UserID, comment.URLContentCommentID))
	comment.UniqueID = fmt.Sprintf("%x", h.Sum(nil))

	return nil

}

func (common *URLContentCommentVote) CheckVoteExists(dbi *gorm.DB, uniqueID string) (bool, error) {
	vote := &URLContentCommentVote{}
	err := dbi.Where("unique_id = ?", uniqueID).First(vote).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	logrus.Errorf("%v+", vote)

	return vote.ID != 0, nil
}
