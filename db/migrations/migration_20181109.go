package migrations

/*
 *
 *  SiGG-Proof-Einstein-Rosen-Bridge-Theory
 *
 */

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func Migration20181109() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		{
			ID: "201811091130",
			Migrate: func(tx *gorm.DB) error {

				type BaseModel struct {
					ID        uint `gorm:"primary_key" json:"-"`
					CreatedAt uint `json:"created_at"`
					UpdatedAt uint `json:"updated_at"`
				}

				type User struct {
					BaseModel
					UniqueID  string `json:"id" gorm:"type:varchar(128);unique_index"`
					Username  string `json:"-" gorm:"type:varchar(128);index"`
					Password  string `json:"-"`
					Salt      string `json:"-"`
					Nickname  string `json:"nickname"`
					AvatarURL string `json:"avatar_url"`
					Balance   string `json:"balance"`
				}

				if err := tx.AutoMigrate(&User{}).Error; err != nil {
					return err
				}

				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				return nil
			},
		},
	}
}
