package main

import (
	"ServeGin/internal/core"
	"fmt"
)

/*
入口
*/
func main() {
	cfg := core.ReadConfig()
	fmt.Println(cfg.DB)
}
