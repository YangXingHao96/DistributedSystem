package main

import (
	"github.com/YangXingHao96/DistributedSystem/package/server/database"
	"github.com/YangXingHao96/DistributedSystem/package/server/handler"
)

func main() {
	db := database.Init()
	handler.HandleUDPRequestAtLeastOnce(db)
}
