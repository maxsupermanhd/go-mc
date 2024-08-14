package component

import pk "github.com/maxsupermanhd/go-vmc/v767/net/packet"

var _ DataComponent = (*RepairCost)(nil)

type RepairCost struct {
	pk.VarInt
}

// ID implements DataComponent.
func (RepairCost) ID() string {
	return "minecraft:repair_cost"
}
