package widget

import (
	"fmt"
	"html"
)

type ButtonWidget struct {
	Base
	model *ButtonModel
	style *ButtonStyle
}

func Button(label string) *ButtonWidget {
	return ButtonWithModel(NewButtonModel(label))
}

func ButtonWithModel(m *ButtonModel) *ButtonWidget {
	b := &ButtonWidget{
		Base:  NewBase(),
		model: m,
		style: &ButtonStyle{
			kind: DefaultKind,
		},
	}
	m.AddListener(b)
	b.Base.SetWidget(b)
	return b
}

func (b *ButtonWidget) View() string {
	return fmt.Sprintf(
		`<sl-button class="btn" id="%s" %s style="%s" size="medium">%s</sl-button>
		 <style> sl-button#%s::part(base) {%s; %s}</style>`,
		b.IDStr(),
		b.style,
		b.SizeStyle(),
		html.EscapeString(b.model.label),
		b.IDStr(),
		"--sl-input-height-medium: 100%",
		b.TextStyle(),
	)
}

func (b *ButtonWidget) Label() string {
	return b.model.label
}

func (b *ButtonWidget) SetLabel(label string) *ButtonWidget {
	b.model.SetLabel(label)
	return b
}

func (b *ButtonWidget) Disable() *ButtonWidget {
	b.model.Disable()
	return b
}

func (b *ButtonWidget) Enable() *ButtonWidget {
	b.model.Enable()
	return b
}

func (b *ButtonWidget) SetStyle(style *ButtonStyle) *ButtonWidget {
	b.style = style
	return b
}

func (b *ButtonWidget) Kind(k Kind) *ButtonWidget {
	b.style.kind = k
	return b
}

func (b *ButtonWidget) Shape(s Shape) *ButtonWidget {
	b.style.shape = s
	return b
}

type ButtonModel struct {
	Model
	label    string
	disabled bool
	loading  bool
}

func NewButtonModel(label string) *ButtonModel {
	return &ButtonModel{
		Model:    NewModel(),
		label:    label,
		disabled: false,
		loading:  false,
	}
}

func (bm *ButtonModel) Label() string {
	return bm.label
}

func (bm *ButtonModel) SetLabel(label string) {
	bm.label = label
	bm.Notify()
}

func (bm *ButtonModel) Disable() {
	bm.disabled = true
	bm.Notify()
}

func (bm *ButtonModel) Enable() {
	bm.disabled = false
	bm.Notify()
}

func (bm *ButtonModel) Loading(flag bool) {
	bm.loading = flag
	bm.Notify()
}

type Kind int

const (
	DefaultKind Kind = iota
	Primary
	Success
	Neutral
	Warning
	Dangerous
)

func (k Kind) String() string {
	switch k {
	case DefaultKind:
		return "default"
	case Primary:
		return "primary"
	case Success:
		return "success"
	case Neutral:
		return "neutral"
	case Warning:
		return "warning"
	case Dangerous:
		return "danger"
	default:
		return "default"
	}
}

type Shape int

const (
	DefaultShape Shape = iota
	Outline
	Pill
	Circle
)

func (s Shape) String() string {
	switch s {
	case DefaultShape:
		return ""
	case Outline:
		return "outline"
	case Pill:
		return "pill"
	case Circle:
		return "circle"
	}
	return ""
}

type ButtonStyle struct {
	kind  Kind
	shape Shape
}

func (bs *ButtonStyle) String() string {
	return fmt.Sprintf(`type = "%s" %s`, bs.kind, bs.shape)
}
