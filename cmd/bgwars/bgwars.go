package main

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/hexagon"
	"github.com/CmdSoda/boardgamewars/internal/path"
)

func main() {
	fmt.Println("bgwars")
	calcpath := path.CalculatePath(hexagon.NewHexagon(1, 1), hexagon.NewHexagon(7, 3))
	fmt.Println(calcpath)
}
