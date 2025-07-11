package gostlib

import (
	"log"
	"net"
	"net/url"

	"github.com/ginuerzh/gost"
)

var (
	server *gost.Server
)

// getTransporter returns the appropriate transporter based on transport type
func getTransporter(transport string) gost.Transporter {
	switch transport {
	case "ws":
		return gost.WSTransporter(nil)
	case "wss":
		return gost.WSSTransporter(nil)
	case "ssh":
		return gost.SSHTunnelTransporter()
	default:
		// Default to websocket if unknown transport
		return gost.WSTransporter(nil)
	}
}

// StartTunnel sets up and starts the gost tunnel with specified transport and address.
func StartTunnel(transport, addr, username, password string) error {
	// Create chain with configurable transport and address
	node := gost.Node{
		Protocol:  "socks5",
		Transport: transport,
		Addr:      addr,
		Client: &gost.Client{
			Connector:   gost.SOCKS5Connector(nil),
			Transporter: getTransporter(transport),
		},
	}
	if username != "" {
		node.User = url.UserPassword(username, password)
	}

	chain := gost.NewChain(node)

	// Create SOCKS5 listener
	ln, err := net.Listen("tcp", "0.0.0.0:1080")
	if err != nil {
		return err
	}

	// Create SOCKS5 handler with chain
	h := gost.SOCKS5Handler(
		gost.ChainHandlerOption(chain),
	)

	// Create and start server
	server = &gost.Server{
		Listener: ln,
		Handler:  h,
	}

	go func() {
		if err := server.Run(); err != nil {
			log.Printf("Server error: %v", err)
		}
	}()

	return nil
}

// StopTunnel stops the tunnel.
func StopTunnel() error {
	if server != nil {
		return server.Close()
	}
	return nil
}
