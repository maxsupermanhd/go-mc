package server

import (
	"github.com/maxsupermanhd/go-vmc/v767/chat"
	"github.com/maxsupermanhd/go-vmc/v767/data/packetid"
	"github.com/maxsupermanhd/go-vmc/v767/net"
	pk "github.com/maxsupermanhd/go-vmc/v767/net/packet"
	"github.com/maxsupermanhd/go-vmc/v767/registry"
)

type ConfigHandler interface {
	AcceptConfig(conn *net.Conn) error
}

type Configurations struct {
	Registries registry.Registries
}

func (c *Configurations) AcceptConfig(conn *net.Conn) error {
	err := conn.WritePacket(pk.Marshal(
		packetid.ClientboundConfigRegistryData,
		pk.NBT(c.Registries),
	))
	if err != nil {
		return err
	}
	err = conn.WritePacket(pk.Marshal(
		packetid.ClientboundConfigFinishConfiguration,
	))
	return err
}

type ConfigFailErr struct {
	reason chat.Message
}

func (c ConfigFailErr) Error() string {
	return "config error: " + c.reason.ClearString()
}
