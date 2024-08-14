package bot

import (
	"errors"

	"github.com/google/uuid"

	"github.com/maxsupermanhd/go-vmc/v762/data/packetid"
	"github.com/maxsupermanhd/go-vmc/v762/net"
	pk "github.com/maxsupermanhd/go-vmc/v762/net/packet"
	"github.com/maxsupermanhd/go-vmc/v762/net/queue"
)

// Client is used to access Minecraft server
type Client struct {
	Conn *Conn
	Auth Auth

	Name string
	UUID uuid.UUID

	Events      Events
	LoginPlugin map[string]func(data []byte) ([]byte, error)
}

func (c *Client) Close() error {
	return c.Conn.Close()
}

// NewClient init and return a new Client.
//
// A new Client has default name "Steve" and zero UUID.
// It is usable for an offline-mode game.
//
// For online-mode, you need login your Mojang account
// and load your Name, UUID and AccessToken to client.
func NewClient() *Client {
	return &Client{
		Auth:   Auth{Name: "Steve"},
		Events: Events{handlers: make([][]PacketHandler, packetid.ClientboundPacketIDGuard)},
	}
}

// Conn is a concurrently-safe warpper of net.Conn with packet queue.
// Note that not all methods are concurrently-safe.
type Conn struct {
	*net.Conn
	send, recv queue.Queue[pk.Packet]
	rerr       error
}

func warpConn(c *net.Conn, qr, qw queue.Queue[pk.Packet]) *Conn {
	wc := Conn{
		Conn: c,
		send: qw,
		recv: qr,
		rerr: nil,
	}
	go func() {
		for {
			var p pk.Packet
			if err := c.ReadPacket(&p); err != nil {
				wc.rerr = err
				break
			}
			if ok := wc.recv.Push(p); !ok {
				wc.rerr = errors.New("receive queue is full")
				break
			}
		}
		wc.recv.Close()
	}()
	go func() {
		for {
			p, ok := wc.send.Pull()
			if !ok {
				break
			}
			if err := c.WritePacket(p); err != nil {
				break
			}
		}
	}()

	return &wc
}

func (c *Conn) ReadPacket(p *pk.Packet) error {
	packet, ok := c.recv.Pull()
	if !ok {
		return c.rerr
	}
	*p = packet
	return nil
}

func (c *Conn) WritePacket(p pk.Packet) error {
	ok := c.send.Push(p)
	if !ok {
		return errors.New("queue is full")
	}
	return nil
}

func (c *Conn) Close() error {
	c.send.Close()
	return c.Conn.Close()
}

// Position is a 3D vector.
type Position struct {
	X, Y, Z int
}