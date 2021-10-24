package main

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/hexagon"
)

func main() {
	fmt.Println("bgwars")
	calcpath := hexagon.CalculatePath(hexagon.NewHexagon(1, 1), hexagon.NewHexagon(7, 3))
	fmt.Println(calcpath)
}
