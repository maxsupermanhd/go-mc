package command

import (
	"github.com/maxsupermanhd/go-vmc/v764/data/packetid"
	pk "github.com/maxsupermanhd/go-vmc/v764/net/packet"
)

type Client interface {
	SendPacket(p pk.Packet)
}

// ClientJoin implement server.Component for Graph
func (g *Graph) ClientJoin(client Client) {
	client.SendPacket(pk.Marshal(
		packetid.ClientboundCommands, g,
	))
}
