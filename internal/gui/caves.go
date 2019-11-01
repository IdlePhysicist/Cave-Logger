package gui

import (
	"time"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"

	"github.com/idlephysicist/cave-logger/internal/model"
)

type caves struct {
	*tview.Table
	caves chan *model.Cave
	filterWord string
}

func newCaves(g *Gui) *caves {
	caves := &caves{
		Table: tview.NewTable().SetSelectable(true, false).Select(0,0).SetFixed(1,1),
	}

	caves.SetTitle(` Caves `).SetTitleAlign(tview.AlignLeft)
	caves.SetBorder(true)
	caves.setEntries(g)
	caves.setKeybinding(g)
	return caves
}

func (c *caves) name() string {
	return `caves`
}

func (c *caves) setKeybinding(g *Gui) {
	c.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		g.setGlobalKeybinding(event)

		return event
	})
}

func (c *caves) setEntries(g *Gui) {}

func (c *caves) updateEntries(g *Gui) {}

func (c *caves) entries(g *Gui) {}

func (c *caves) focus(g *Gui) {
	c.SetSelectable(true, false)
	g.app.SetFocus(c)
}

func (c *caves) unfocus() {
	c.SetSelectable(false, false)
}

func (c *caves) setFilterWord(word string) {
	c.filterWord = word
}

func (c *caves) monitoringCaves(g *Gui) {
	ticker := time.NewTicker(5 * time.Second)

LOOP:
	for {
		select {
		case <-ticker.C:
			c.updateEntries(g)
		case <-g.state.stopChans["caves"]:
			ticker.Stop()
			break LOOP
		}
	}
}