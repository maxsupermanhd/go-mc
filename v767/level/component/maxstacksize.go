package component

import pk "github.com/maxsupermanhd/go-vmc/v767/net/packet"

var _ DataComponent = (*MaxStackSize)(nil)

type MaxStackSize struct {
	pk.VarInt
}

// ID implements DataComponent.
func (MaxStackSize) ID() string {
	return "minecraft:max_stack_size"
}
