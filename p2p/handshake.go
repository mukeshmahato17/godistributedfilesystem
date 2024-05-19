package p2p

import "errors"

// ErrInvalidHandshake is returned if the handshake between
// the local and the remote node couldnot be established.
var ErrInvalidHandshake = errors.New("invalid handshake")

// HandshakeFunc... ?
type HandshakeFunc func(any) error

func NOPHandshakeFunc(any) error { return nil }
