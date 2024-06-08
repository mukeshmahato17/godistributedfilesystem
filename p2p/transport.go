package p2p

// Peer is an interface that represents the remote nodes.
type Peer interface {
	Close() error
}

// Transport is anything that handles the communication
// between the nodes in the network. This can be of the
// form (TCP, UDP, web sockets, ... )
type Transport interface {
	Dial(string) error
	ListenAndAccept() error
	Consume() <-chan RPC
	Close() error
}
