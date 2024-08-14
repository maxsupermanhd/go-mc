package component

import (
	pk "github.com/Tnze/go-vmc/v767/net/packet"
)

var _ DataComponent = (*OminousBottleAmplifier)(nil)

type OminousBottleAmplifier struct {
	pk.VarInt
}

// ID implements DataComponent.
func (OminousBottleAmplifier) ID() string {
	return "minecraft:ominous_bottle_amplifier"
}
