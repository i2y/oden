package widget

import (
	"fmt"
)

type InputWidget struct {
	Base
	model *InputModel
}

func Input(inputType InputType, placeholder string) *InputWidget {
	return InputWithModel(NewInputModel(inputType, placeholder))
}

func InputWithModel(m *InputModel) *InputWidget {
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
		`<div style="%s">
		   <sl-input id="%s" style="%s" type="%s" placeholder="%s" size="medium" clearable></sl-button>
		 </div>
		 <style>sl-input#%s::part(base) {%s; %s}</style>`,
		i.SizeStyle(),
		i.ID(),
		i.OtherStyle(),
		i.model.inputType,
		i.model.placeholder,

		i.ID(),
		"--sl-input-height-medium: 100%",
		i.TextStyle(),
	)
}

type InputModel struct {
	Model
	inputType   InputType
	placeholder string
}

func NewInputModel(t InputType, placeholder string) *InputModel {
	return &InputModel{
		Model:       NewModel(),
		inputType:   t,
		placeholder: placeholder,
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
