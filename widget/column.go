package widget

import (
	"fmt"
)

type ColumnLayout struct {
	Layout
}

func Column(children ...Widget) *ColumnLayout {
	c := &ColumnLayout{
		Layout: NewLayout(children...),
	}
	c.Base.SetWidget(c)
	return c
}

func (c *ColumnLayout) View() string {
	c.layout()
	return fmt.Sprintf(
		`<div id="oden-%s" style="%s %s">%s</div>`,
		c.ID(),
		c.style(),
		c.SizeStyle(),
		c.Layout.View(),
	)
}

func (c *ColumnLayout) style() string {
	return fmt.Sprintf("display: flex; flex-direction: column; flex: 1 1 0; width: 100%%; height: 100%%;")
}

func (c *ColumnLayout) layout() {
	for _, w := range c.children {
		switch w.SizePolicy() {
		case Expanding:
			w.SetSizeStyle("flex: 1 1 0; width: 100%; height: 100%;")
		case Fixed:
			w.SetSizeStyle(fmt.Sprintf("flex: 0 0 %dpx; width: %dpx;", w.Height(), w.Width()))
		case FixedWidth:
			w.SetSizeStyle(fmt.Sprintf("flex: 1 1 0; width: %dpx; height: 100%%;", w.Width()))
		case FixedRatioWidth:
			w.SetSizeStyle(fmt.Sprintf("flex: 1 1 0; width: %d%%; height: 100%%;", w.Width()))
		case FixedHeight:
			w.SetSizeStyle(fmt.Sprintf("flex: 0 0 %dpx; width: 100%%;", w.Height()))
		case FixedRatioHeight:
			w.SetSizeStyle(fmt.Sprintf("flex: 0 0 %d%%; width: 100%%;", w.Height()))
		}
	}
}
