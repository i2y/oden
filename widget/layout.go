package widget

import core "github.com/i2y/oden/core"

type Layout struct {
	Base
	children []Widget
}

func NewLayout(children ...Widget) Layout {
	return Layout{
		Base:     NewBase(),
		children: children,
	}
}

func (l *Layout) Attach(a *core.App) {
	l.Base.Attach(a)

	for _, c := range l.children {
		c.Attach(a)
	}
}

func (l *Layout) Detach() {
	l.Base.Detach()

	for _, c := range l.children {
		c.Detach()
	}
}

func (l *Layout) Children() []Widget {
	return l.children
}

func (l *Layout) Add(w Widget) {
	l.children = append(l.children, w)
	l.Update()
}

func (l *Layout) Remove(w Widget) {
	new := make([]Widget, len(l.children))
	i := 0
	for _, c := range l.children {
		if c != w {
			new[i] = c
			i++
		}
	}
	l.children = new[:i]
	l.Update()
}

func (l *Layout) View() string {
	var ret string
	for _, w := range l.children {
		ret += w.View()
	}
	return ret
}
