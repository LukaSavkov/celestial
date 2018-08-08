package main

import (
	"fmt"
	"github.com/c12s/celestial/model/config"
	"github.com/c12s/celestial/storage/etcd"
	"log"
)

func main() {
	// Load configurations
	conf, err := config.ConfigFile()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(conf)

	//Load database
	db, err := etcd.New(conf.GetClientConfig())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(db)

	//Start server

	// c := client.NewClient(conf.GetClientConfig())
	// defer c.Close()

	// for i := 0; i < 10; i++ {
	// 	for n := range c.GetClusterNodes("novisad", "grbavica") {
	// 		fmt.Println(n)
	// 	}
	// }

	// a := api.NewApi(conf)
	// a.Run()

}
