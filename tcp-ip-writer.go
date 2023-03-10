package base

/*
 *
 *  SiGG-Proof-Einstein-Rosen-Bridge-Theory
 *
 */

import "net"

type CachedTCPWriter interface {
	// WriteString send the string to the writing queue
	WriteString(message string)
	// Write send the buffer to the writing queue
	Write(message []byte)
}

type CachedTCPWriter interface {
	WriteTo(remoteAddr *net.UDPAddr, data []byte)
}
