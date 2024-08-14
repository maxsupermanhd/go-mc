package world

import "github.com/maxsupermanhd/go-vmc/v767/level"

type EventsListener struct {
	LoadChunk   func(pos level.ChunkPos) error
	UnloadChunk func(pos level.ChunkPos) error
}