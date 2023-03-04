package models

/*
 *
 *  SiGG-Proof-Einstein-Rosen-Bridge-Theory
 *
 */

const OAuthGoogle = 1

type UserOAuth struct {
	BaseModel
	UserID     uint
	VendorType uint
	VendorID   string
}
