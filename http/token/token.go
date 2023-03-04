package token

/*
 *
 *  SiGG-Proof-Einstein-Rosen-Bridge-Theory
 *
 */

import (
	"fmt"

	"github.com/primasio/wormhole/cache"
)

type Token struct {
	Token string `json:"token"`
}

/**
 * Create new token for a given user
 */
func IssueToken(userId uint, expires bool) (error, *Token) {

	err, token := cache.NewSessionKey()

	if err != nil {
		return err, nil
	}

	userIdStr := fmt.Sprint(userId)
	cache.SessionSet(token, userIdStr, expires)

	tokenStruct := &Token{Token: token}

	return nil, tokenStruct
}
