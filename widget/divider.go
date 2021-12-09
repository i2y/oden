package widget

import (
	"fmt"
)

type DividerWidget struct {
	Base
}

func Divider() *DividerWidget {
	d := &DividerWidget{
		Base: NewBase(),
	}
	d.Base.SetWidget(d)
	return d
}

func (d *DividerWidget) View() string {
	return fmt.Sprintf(
		`<sl-divider id="%s" style="%s; height: 32px"></sl-divider>`, // TODO height
		d.IDStr(),
		d.SizeStyle(),
	)
}
