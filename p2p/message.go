package p2p

// Message holds any arbitrary that is being sent over the
// each transport between two nodes in the network.
type RPC struct {
	From    string
	Payload []byte
}
