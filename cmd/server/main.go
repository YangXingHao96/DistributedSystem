package main

import (
	"DistributedSystem/package/server/handler"
	"DistributedSystem/package/server/service/database"
)

func main() {
	database.Init()
	handler.HandleUDPRequest()
}
