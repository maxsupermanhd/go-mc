package server

import (
	_ "embed"

	"github.com/maxsupermanhd/go-vmc/v764/net"
	"github.com/maxsupermanhd/go-vmc/v764/yggdrasil/user"

	"github.com/google/uuid"
)

type GamePlay interface {
	// AcceptPlayer handle everything after "LoginSuccess" is sent.
	//
	// Note: the connection will be closed after this function returned.
	// You don't need to close the connection, but to keep not returning while the player is playing.
	AcceptPlayer(name string, id uuid.UUID, profilePubKey *user.PublicKey, properties []user.Property, protocol int32, conn *net.Conn)
}