package component

import (
	"io"

	"github.com/maxsupermanhd/go-vmc/v767/nbt/dynbt"
	pk "github.com/maxsupermanhd/go-vmc/v767/net/packet"
)

var _ DataComponent = (*BucketEntityData)(nil)

type BucketEntityData struct {
	dynbt.Value
}

// ID implements DataComponent.
func (BucketEntityData) ID() string {
	return "minecraft:bucket_entity_data"
}

// ReadFrom implements DataComponent.
func (b *BucketEntityData) ReadFrom(r io.Reader) (n int64, err error) {
	return pk.NBT(&b.Value).ReadFrom(r)
}

// WriteTo implements DataComponent.
func (b *BucketEntityData) WriteTo(w io.Writer) (n int64, err error) {
	return pk.NBT(&b.Value).WriteTo(w)
}
