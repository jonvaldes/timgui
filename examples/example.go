package main

import (
	"fmt"
	"os"
	"github.com/nsf/termbox-go"
	"github.com/jonvaldes/timui"
)

type Data struct {
	tree bool
	ls bool

	selectedDir int
	otherDir string
}

func redraw(state *timui.State, data *Data) {
	termbox.Clear(state.Colors.Default, state.Colors.Default)

	mouseX := fmt.Sprintf("%d", state.MouseCursor.X)
	mouseY := fmt.Sprintf("%d", state.MouseCursor.Y)

	timui.Box(state, 2, 1, "Commands",
		&timui.CheckBox{&data.tree, "tree"},
		&timui.CheckBox{&data.ls, "ls"},
		&timui.Separator{"Mouse coords"},
		&timui.TextEdit{&mouseX},
		&timui.TextEdit{&mouseY},
	)

	timui.Box(state, 19, 1, "Dirs",
		&timui.RadioBox{0, &data.selectedDir, "/"},
		&timui.RadioBox{1, &data.selectedDir, "~"},
		&timui.RadioBox{2, &data.selectedDir, "~/Downloads"},
		&timui.RadioBox{5, &data.selectedDir, "Other:"},
		&timui.TextEdit{&data.otherDir},
	)

	timui.Box(state, 40, 1, "",
		&timui.Button{"Run!", func() {
			termbox.Close()
			os.Exit(0)
		}},
	)

	state.Flush()
	termbox.Flush()
}

func main() {
	termbox.Init()
	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)
	data := Data{}

	state := timui.NewState()
	state.Colors.Selected = termbox.ColorCyan
	state.Colors.Cursor = termbox.ColorCyan | termbox.AttrBold
	redraw(&state, &data)

mainloop:
	for {
		ev := termbox.PollEvent()
		state.HandleEvent(ev)
		if ev.Type == termbox.EventKey && ev.Key == termbox.KeyEsc {
			break mainloop
		}
	repeat:
		redraw(&state, &data)
		if state.NeedsRedraw {
			state.NeedsRedraw = false
			goto repeat
		}
	}

	termbox.Close()
}
