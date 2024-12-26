package main

import (
	"dwnldr/pkg/config"
	"dwnldr/pkg/handlers"
	"fmt"
)

func main() {
    config.LoadConfig("config/config.toml")
    handlers.HandleURL("https://www.youtube.com/watch?v=example")
    fmt.Println("Download started!")
}
