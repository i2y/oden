package widget

import (
	"fmt"
)

type RowLayout struct {
	Layout
}

func Row(children ...Widget) *RowLayout {
	r := &RowLayout{
		Layout: NewLayout(children...),
	}
	r.Base.SetWidget(r)
	return r
}

func (r *RowLayout) View() string {
	r.layout()
	return fmt.Sprintf(
		`<div id="%s" style="%s %s">%s</div>`,
		r.IDStr(),
		r.style(),
		r.SizeStyle(),
		r.Layout.View(),
	)
}

func (r *RowLayout) style() string {
	return fmt.Sprintf("display: flex; flex-direction: row;")
}

func (r *RowLayout) layout() {
	for _, w := range r.children {
		switch w.SizePolicy() {
		case Expanding:
			w.SetSizeStyle("flex: 1 1 0; width: 100%; height: 100%;")
		case Fixed:
			w.SetSizeStyle(fmt.Sprintf("flex: 0 0 %dpx; height: %dpx;", w.Width(), w.Height()))
		case FixedWidth:
			w.SetSizeStyle(fmt.Sprintf("flex: 0 0 %dpx; height: 100%%; width: %dpx", w.Width(), w.Width()))
		case FixedRatioWidth:
			w.SetSizeStyle(fmt.Sprintf("flex: 0 0 %d%%; height: 100%%; width: %dpx", w.Width(), w.Width()))
		case FixedHeight:
			w.SetSizeStyle(fmt.Sprintf("flex: 1 1 0; width: 100%%; height: %dpx;", w.Height()))
		case FixedRatioHeight:
			w.SetSizeStyle(fmt.Sprintf("flex: 1 1 0; width: 100%%; height: %d%%;", w.Height()))
		}
	}
}
