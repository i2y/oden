package widget

import core "github.com/i2y/oden/core"

type Component struct {
	Base
	widgetTree Widget
	builder    func() Widget
}

func NewComponent(builder func() Widget) *Component {
	return &Component{
		Base: NewBase(),
		builder: builder,
	}
}

func (c *Component) Attach(a *core.App) {
	c.Base.Attach(a)
	c.build().Attach(a)
}

func (c *Component) Detach() {
	c.Base.Detach()
}

func (c *Component) View() string {
	return c.build().View()
}

func (c *Component) build() Widget {
	if c.widgetTree == nil {
		c.widgetTree = c.builder()
	}
	return c.widgetTree
}
