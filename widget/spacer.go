package widget

import (
	"fmt"
)

type SpacerWidget struct {
	Base
}

func Spacer() *SpacerWidget {
	s := &SpacerWidget{
		Base: NewBase(),
	}
	s.Base.SetWidget(s)
	return s
}

func (s *SpacerWidget) View() string {
	return fmt.Sprintf(
		`<div id="%s" style="%s"></div>`,
		s.ID(),
		s.SizeStyle(),
	)
}
