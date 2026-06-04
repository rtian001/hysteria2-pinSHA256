package main

import (
	"context"
	"crypto/sha256"
	"crypto/tls"
	"fmt"
	"os"
	"time"

	"github.com/quic-go/quic-go"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "用法: %s host:port\n", os.Args[0])
		os.Exit(1)
	}
	addr := os.Args[1]


	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"h3"},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	conn, err := quic.DialAddr(ctx, addr, tlsConf, &quic.Config{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "连接失败: %v\n", err)
		os.Exit(1)
	}
	defer conn.CloseWithError(0, "")

	state := conn.ConnectionState().TLS
	certs := state.PeerCertificates
	if len(certs) == 0 {
		fmt.Fprintln(os.Stderr, "未收到证书")
		os.Exit(1)
	}

	leaf := certs[0]
	certDER := leaf.Raw
	hash := sha256.Sum256(certDER)
	pinHex := fmt.Sprintf("%x", hash)
	fmt.Printf(pinHex)

}
