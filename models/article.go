package models

/*
 *
 *  SiGG-Proof-Einstein-Rosen-Bridge-Theory
 *
 */

type Article struct {
	BaseModel

	UserID   uint   `gorm:"index" json:"-"`
	Title    string `gorm:"type:text" form:"title" json:"title" binding:"required"`
	Abstract string `gorm:"type:text" json:"abstract"`
	Content  string `gorm:"type:longtext" form:"content" json:"content" binding:"required"`
	Language string `gorm:"column:lang;size:64" json:"language"`

	ContentId  string `gorm:"type:varchar(128);unique_index" json:"content_id"`
	ContentDNA string `gorm:"type:varchar(128);unique_index" json:"content_dna"`
}

func (article *Article) DetectLanguage() string {
	return ""
}
