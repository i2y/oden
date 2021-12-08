package widget

import (
	"fmt"
)

type InputWidget struct {
	Base
	model *InputModel
}

func Input(m *InputModel) *InputWidget {
	i := &InputWidget{
		Base:  NewBase(),
		model: m,
	}
	m.AddListener(i)
	i.Base.SetWidget(i)
	return i
}

func (i *InputWidget) View() string {
	return fmt.Sprintf(
		`<sl-input id="%s" style="%s" type="%s" placeholder="%s" size="%s" clearable></sl-button>
		 <style>sl-input#%s::part(base) {%s; %s}</style>`,
		i.IDStr(),
		i.SizeStyle(),
		i.model.inputType,
		i.model.placeholder,
		i.model.size,

		i.IDStr(),
		"--sl-input-height-medium: 100%", // TODO
		i.TextStyle(),
	)
}

type InputModel struct {
	Model
	inputType   InputType
	placeholder string
	size        InputSize
}

func NewInputModel(t InputType, placeholder string, size InputSize) *InputModel {
	return &InputModel{
		Model:       NewModel(),
		inputType:   t,
		placeholder: placeholder,
		size:        size,
	}
}

func (im *InputModel) InputType() InputType {
	return im.inputType
}

func (im *InputModel) SetInputType(inputType InputType) {
	im.inputType = inputType
}

func (im *InputModel) Placeholder() string {
	return im.placeholder
}

func (im *InputModel) SetPlaceholder(placeholder string) {
	im.placeholder = placeholder
}

type InputType int

const (
	TextInputType InputType = iota
	DateInputType
	EmailInputType
	NumberInputType
	PasswordInputType
	SearchInputType
	TelInputType
	URLInputType
)

func (i InputType) String() string {
	switch i {
	case TextInputType:
		return "text"
	case DateInputType:
		return "date"
	case EmailInputType:
		return "email"
	case NumberInputType:
		return "number"
	case PasswordInputType:
		return "password"
	case SearchInputType:
		return "search"
	case TelInputType:
		return "tel"
	}
	return "text"
}

type InputSize int

const (
	MediumInput InputSize = iota
	LargeInput
	SmallInput
)

func (i InputSize) String() string {
	switch i {
	case MediumInput:
		return "medium"
	case LargeInput:
		return "large"
	case SmallInput:
		return "small"
	}
	return "text"
}
