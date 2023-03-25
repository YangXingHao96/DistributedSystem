package main

import (
	"DistributedSystem/package/server/database"
	"DistributedSystem/package/server/handler"
)

func main() {
	db := database.Init()
	handler.HandleUDPRequestAtLeastOnce(db)
}
