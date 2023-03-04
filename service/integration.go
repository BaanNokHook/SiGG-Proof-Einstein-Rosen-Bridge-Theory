package service

/*
 *
 *  SiGG-Proof-Einstein-Rosen-Bridge-Theory
 *
 */

import (
	"sync"

	"github.com/primasio/wormhole/config"
)

var integration *Integration
var integrationOnce sync.Once

type Integration struct{}

func GetIntegration() *Integration {
	integrationOnce.Do(func() {
		integration = &Integration{}
	})

	return integration
}

func (s *Integration) GetURLContentCommentVoteScore(like bool) int64 {
	c := config.GetConfig()

	if like {
		return c.GetInt64("integration.url_content_comment_vote.like")
	}

	return c.GetInt64("integration.url_content_comment_vote.hate")
}

func (s *Integration) GetRegisterScore() int64 {
	return config.GetConfig().GetInt64("integration.register")
}
