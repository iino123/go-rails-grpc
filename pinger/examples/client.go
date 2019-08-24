package main

import (
	"context"
	"fmt"
	"os"

	pinger "../lib"
	"google.golang.org/grpc"
)

/* server.goを起動した状態で下記コマンドを実行してgrpcサーバが動いていることを確認するための実装。
$ cd grpc-sample/pinger
$ go run examples/client.go
*/

func main() {
	conn, err := grpc.Dial("localhost:5300", grpc.WithInsecure())
	if err != nil {
		fmt.Fprintf(os.Stderr, "grpc.Dial: %v\n", err)
		return
	}
	defer conn.Close()
	client := pinger.NewPingerClient(conn)
	req := &pinger.Empty{}
	pong, err := client.Ping(context.Background(), req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ping: %v\n", err)
		return
	}
	fmt.Fprintf(os.Stdout, "Pong: %s\n", pong.Text)
}
