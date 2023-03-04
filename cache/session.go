package cache

/*
 *
 *  SiGG-Proof-Einstein-Rosen-Bridge-Theory
 *
 */

import (
	"errors"
	"time"

	"github.com/primasio/wormhole/util"
)

const sessionPrefix = "wormhole_session_"

func NewSessionKey() (error, string) {

	var counter = 0

	for {
		counter = counter + 1
		key := util.RandString(32)
		err, check := SessionGet(key)

		if err != nil {
			return err, ""
		}

		if check == "" {
			return nil, key
		}

		if counter >= 5 {
			// This is unlikely to happen
			// Must be error from other parts
			return errors.New("too many iterations while generating new session key"), ""
		}
	}
}

func SessionSet(token, userId string, expires bool) error {
	store := GetCache()

	duration := time.Hour * 24 * 30

	if expires {
		duration = time.Hour * 2
	}

	return store.Set(sessionPrefix+token, userId, duration)
}

func SessionGet(token string) (err error, userId string) {

	store := GetCache()

	if store == nil {
		return errors.New("cache store is nil"), ""
	}

	var userIdStore string

	if err := store.Get(sessionPrefix+token, &userIdStore); err != nil {
		if err != ErrCacheMiss && err != ErrNotStored {
			return err, ""
		}
	}

	return nil, userIdStore
}
