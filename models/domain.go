package models

/*
 *
 *  SiGG-Proof-Einstein-Rosen-Bridge-Theory
 *
 */

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"net/url"

	"github.com/jinzhu/gorm"
	"github.com/primasio/wormhole/db"
)

type Domain struct {
	BaseModel
	UserID  uint   `json:"-"`
	Domain  string `gorm:"type:text" json:"domain"`
	Title   string `gorm:"type:text" json:"title"`
	HashKey string `gorm:"type:varchar(128);unique_index" json:"-"`

	IsActive bool `gorm:"default:false" json:"is_active"`
	Votes    uint `gorm:"default:1" json:"votes"`
}

func ExtractDomainFromURL(urlStr string) (error, string) {
	u, err := url.Parse(urlStr)

	if err != nil {
		return err, ""
	}

	return nil, u.Host
}

func CleanDomain(domain string) string {
	//TODO: Remove trailing slash, remove www, etc
	return domain
}

func GetDomainByDomainName(domain string, dbi *gorm.DB, forUpdate bool) (error, *Domain) {

	if domain == "" {
		return errors.New("domain is empty"), nil
	}

	hashKey := GetDomainHashKey(domain)

	var domainModel Domain

	sql := "SELECT * FROM domains WHERE hash_key = ?"

	if forUpdate && db.GetDbType() != db.SQLITE {
		sql = sql + " FOR UPDATE"
	}

	dbi.Raw(sql, hashKey).Scan(&domainModel)

	if domainModel.ID == 0 {
		return nil, nil
	}

	return nil, &domainModel
}

func GetDomainHashKey(domain string) string {

	sumBytes := sha1.Sum([]byte(domain))

	return hex.EncodeToString(sumBytes[:])
}
