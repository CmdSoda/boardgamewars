package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"log"
)

func main() {
	fmt.Println("boardgamewars playground")

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

//goland:noinspection GoUnhandledErrorResult
func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("hello", maxX/2-7, maxY/2, maxX/2+7, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "Hello world!")
	}
	return nil
}

//goland:noinspection GoUnusedParameter
func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
