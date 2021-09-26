package main

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/bgw"
)

func main() {
	fmt.Println("bgwars")
	calcpath := bgw.CalculatePath(bgw.NewHexagon(1, 1), bgw.NewHexagon(7, 3))
	fmt.Println(calcpath)
}
