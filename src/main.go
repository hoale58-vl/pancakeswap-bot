package main

import (
	"swap_bot/websocket"
)

func main() {
	for {
		websocket.Init().Connect()
	}
}
