package main

import (
	"fmt"

	"maneeshaindrachapa.github.io/go-server-gin/configs"
)

func main() {
	// load env variables just once in here so can be use in any other place
	configs.InitEnvConfigs()
  
  	// print the env variables
	fmt.Printf("Port:%s\n", configs.EnvConfigs.LocalServerPort)
	fmt.Printf("key:%s", configs.EnvConfigs.SecretKey)
}