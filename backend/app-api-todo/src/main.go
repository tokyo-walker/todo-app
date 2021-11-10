package main
// package main じゃないとRunできない

import (
	"./api"

)

func main() {

	api.Setting{}.Connect()
}