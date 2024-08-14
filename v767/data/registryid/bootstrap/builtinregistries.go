package bootstrap

import (
	"github.com/maxsupermanhd/go-vmc/v767/data/registryid"
	"github.com/maxsupermanhd/go-vmc/v767/level/block"
	"github.com/maxsupermanhd/go-vmc/v767/registry"
)

func RegisterBlocks(reg *registry.Registry[block.Block]) {
	reg.Clear()
	for i, key := range registryid.Block {
		id, val := reg.Put(key, block.FromID[key])
		if int32(i) != id || val == nil || *val == nil {
			panic("register blocks failed")
		}
	}
}
