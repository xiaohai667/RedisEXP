package main

import (
	"RedisExp/pkg/help"
	"fmt"
)

func main() {
	logo := `
██████╗ ███████╗██████╗ ██╗███████╗    ███████╗██╗  ██╗██████╗ 
██╔══██╗██╔════╝██╔══██╗██║██╔════╝    ██╔════╝╚██╗██╔╝██╔══██╗
██████╔╝█████╗  ██║  ██║██║███████╗    █████╗   ╚███╔╝ ██████╔╝
██╔══██╗██╔══╝  ██║  ██║██║╚════██║    ██╔══╝   ██╔██╗ ██╔═══╝ 
██║  ██║███████╗██████╔╝██║███████║    ███████╗██╔╝ ██╗██║     
╚═╝  ╚═╝╚══════╝╚═════╝ ╚═╝╚══════╝    ╚══════╝╚═╝  ╚═╝╚═╝
`
	fmt.Println(logo)
	help.Help()
}
