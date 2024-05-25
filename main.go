package main

import (
	"log"

	"github.com/mukeshmahato17/godistributedfilesystem/p2p"
)

func OnPeer(peer p2p.Peer) error {
	peer.Close()
	// fmt.Println("doing some logic with the peer outside the TCPTransport")
	return nil
}

func main() {
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
		OnPeer:        OnPeer,
	}

	tr := p2p.NewTCPTransport(tcpOpts)

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}
