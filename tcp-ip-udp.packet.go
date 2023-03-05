package base

/*
 *
 *  SiGG-Proof-Einstein-Rosen-Bridge-Theory
 *
 */

import "net"

type UdpPacket struct {
	RemoteAddr *net.UDPAddr
	Data       []byte
}

func NewUdpPacket(remoteAddr *net.UDPAddr, data []byte) *UdpPacket {
	return &UdpPacket{
		RemoteAddr: remoteAddr,
		Data:       data,
	}
}
