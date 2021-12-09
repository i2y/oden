package widget

import (
	"fmt"
)

type TextAreaWidget struct {
	Base
	model *TextAreaModel
}

func TextArea(placeholder string) *TextAreaWidget {
	return TextAreaWithModel(NewTextAreaModel(placeholder))
}

func TextAreaWithModel(m *TextAreaModel) *TextAreaWidget {
	i := &TextAreaWidget{
		Base:  NewBase(),
		model: m,
	}
	m.AddListener(i)
	i.Base.SetWidget(i)
	return i
}

func (i *TextAreaWidget) View() string {
	return fmt.Sprintf(
		`<sl-textarea id="%s" style="%s %s" placeholder="%s" size="medium" resize="none"></sl-button>
		 <style>sl-textarea#%s::part(base) {%s; %s}</style>`,
		i.IDStr(),
		i.SizeStyle(),
		i.OtherStyle(),
		i.model.placeholder,

		i.IDStr(),
		"--sl-textarea-height-medium: 100%",
		i.TextStyle(),
	)
}

type TextAreaModel struct {
	Model
	placeholder string
}

func NewTextAreaModel(placeholder string) *TextAreaModel {
	return &TextAreaModel{
		Model:       NewModel(),
		placeholder: placeholder,
	}
}

func (im *TextAreaModel) Placeholder() string {
	return im.placeholder
}

func (im *TextAreaModel) SetPlaceholder(placeholder string) {
	im.placeholder = placeholder
}
