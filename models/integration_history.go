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
)

type IntegrationHistory struct {
	BaseModel

	UserID      uint   `json:"-"`
	UniqueID    string `json:"id" gorm:"type:varchar(128);unique_index"`
	Integration int64  `json:"integration"`
	Description string `json:"description"`
	Data        string `json:"-"`

	User User `gorm:"save_associations:false" json:"user"`
}

func (m *IntegrationHistory) SetUniqueID() error {
	if m.Data == "" {
		return errors.New("Integration History Data Required")
	}

	h := sha1.New()
	io.WriteString(h, m.Data)
	m.UniqueID = fmt.Sprintf("%x", h.Sum(nil))

	return nil
}
