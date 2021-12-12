package widget

import (
	"fmt"
	"html"
)

type ButtonWidget struct {
	Base
	model  *ButtonModel
	option *ButtonOption
}

func Button(label string, options ...func(*ButtonOption)) *ButtonWidget {
	return ButtonWithModel(NewButtonModel(label), options...)
}

func ButtonWithModel(m *ButtonModel, options ...func(*ButtonOption)) *ButtonWidget {
	o := &ButtonOption{
		kind:  DefaultKind,
		shape: DefaultShape,
	}
	for _, option := range options {
		option(o)
	}

	b := &ButtonWidget{
		Base:   NewBase(),
		model:  m,
		option: o,
	}
	m.AddListener(b)
	b.Base.SetWidget(b)
	return b
}

func (b *ButtonWidget) View() string {
	return fmt.Sprintf(
		`<sl-button class="btn" id="%s" %s style="%s %s" size="medium">%s</sl-button>
		 <style> sl-button#%s::part(base) {%s; %s}</style>`,
		b.ID(),
		b.option,
		b.SizeStyle(),
		b.OtherStyle(),
		html.EscapeString(b.model.label),
		b.ID(),
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

type kind int

const (
	DefaultKind kind = iota
	Primary
	Success
	Neutral
	Warning
	Dangerous
)

func (k kind) String() string {
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

type shape int

const (
	DefaultShape shape = iota
	Outline
	Pill
	Circle
)

func (s shape) String() string {
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

type ButtonOption struct {
	kind  kind
	shape shape
}

func Type(k kind) func(*ButtonOption) {
	return func(option *ButtonOption) {
		option.kind = k
	}
}

func Shape(s shape) func(*ButtonOption) {
	return func(option *ButtonOption) {
		option.shape = s
	}
}

func (o *ButtonOption) String() string {
	return fmt.Sprintf(`type="%s" %s`, o.kind, o.shape)
}
