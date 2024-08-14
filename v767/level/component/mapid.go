package component

import pk "github.com/Tnze/go-vmc/v767/net/packet"

var _ DataComponent = (*MapID)(nil)

type MapID struct {
	pk.VarInt
}

// ID implements DataComponent.
func (MapID) ID() string {
	return "minecraft:map_id"
}
