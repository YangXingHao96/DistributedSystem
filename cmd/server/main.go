package main

import (
	"flag"
	"fmt"
	"github.com/YangXingHao96/DistributedSystem/package/server/database"
	"github.com/YangXingHao96/DistributedSystem/package/server/handler"
)

func main() {
	deliveryMode := flag.Int("mode", 0, "server request processing mode, default at least once integer 0, at most once is integer 1")
	randomTimeout := flag.Bool("timeout", false, "server request timeout mode, default is no timeout boolean false, timeout is boolean true")
	host := flag.String("host", "localhost", "server listening host")
	port := flag.String("port", "2222", "server listening port")
	flag.Parse()
	fmt.Printf("Deliver mode: %v, timeout: %v, host: %v, port:%v\n", *deliveryMode, *randomTimeout, *host, *port)
	db := database.Init()

	if *deliveryMode == 0 {
		fmt.Println("Server starting up, running at least once delivery")
		handler.HandleUDPRequestAtLeastOnce(db, *randomTimeout, *host, *port)
	} else {
		fmt.Println("Server starting up, running at most once delivery")
		handler.HandleUDPRequestAtMostOnce(db, *randomTimeout, *host, *port)
	}

}
