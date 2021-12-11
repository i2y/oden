package widget

import (
	"fmt"
	"html"
)

type TextWidget struct {
	Base
	model StringEventPublisher
	style *TextStyle
}

// Go 1.18がリリースされたらstring | State[T] のようになるだろう..
func Text(s StringEventPublisher) *TextWidget {
	l := &TextWidget{
		Base:  NewBase(),
		model: s,
		style: &TextStyle{
			align:         Center,
			verticalAlign: Middle,
		},
	}
	s.AddListener(l)
	l.Base.SetWidget(l)
	return l
}

func (t *TextWidget) View() string {
	return fmt.Sprintf(
		`<div id="%s" style="%s display: table;"><span class="label" style="%s">%s</span></div>`,
		t.ID(),
		t.SizeStyle(),
		t.TextStyle(),
		html.EscapeString(t.model.String()),
	)
}

func (t *TextWidget) SetLabel(label string) {
	t.model.SetString(label)
}
