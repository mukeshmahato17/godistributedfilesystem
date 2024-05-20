package p2p

import "net"

// Message holds any arbitrary that is being sent over the
// each transport between two nodes in the network.
type RPC struct {
	From    net.Addr
	payload []byte
}
