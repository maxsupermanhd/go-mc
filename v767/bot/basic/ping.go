package basic

import (
	"github.com/Tnze/go-vmc/v767/data/packetid"
	pk "github.com/Tnze/go-vmc/v767/net/packet"
)

func (p *Player) handlePingPacket(packet pk.Packet) error {
	var pingID pk.Int
	if err := packet.Scan(&pingID); err != nil {
		return Error{err}
	}

	// Response
	err := p.c.Conn.WritePacket(pk.Packet{
		ID:   int32(packetid.ServerboundPong),
		Data: packet.Data,
	})
	if err != nil {
		return Error{err}
	}
	return nil
}
