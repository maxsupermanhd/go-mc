package component

import pk "github.com/maxsupermanhd/go-vmc/v767/net/packet"

var _ DataComponent = (*Damage)(nil)

type Damage struct {
	pk.VarInt
}

// ID implements DataComponent.
func (Damage) ID() string {
	return "minecraft:damage"
}
