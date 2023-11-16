package main

import (
	"os"

	"github.com/soraritt/kawaii-shop-tutorial/config"
	"github.com/soraritt/kawaii-shop-tutorial/modules/servers"
	"github.com/soraritt/kawaii-shop-tutorial/pkg/databases"
)

func envPath() string {

	if len(os.Args) == 1 {
		return ".env"
	} else {
		// fmt.Println("os.Args :", os.Args)
		return os.Args[1]
	}
}

func main() {
	config := config.LoadConfig(envPath())
	// fmt.Println(config.App().Url())
	// fmt.Println(config.Db())
	// fmt.Println(config.Jwt())
	db := databases.DbConnect(config.Db())
	defer db.Close()

	servers.NewServer(config, db).Start()

}
