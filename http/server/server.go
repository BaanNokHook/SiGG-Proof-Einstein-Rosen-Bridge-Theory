package server

/*
 *
 *  SiGG-Proof-Einstein-Rosen-Bridge-Theory
 *
 */

import "github.com/primasio/wormhole/config"

func Init() {
	c := config.GetConfig()
	r := NewRouter()
	r.Run(c.GetString("http.server.host") + ":" + c.GetString("http.server.port"))
}
