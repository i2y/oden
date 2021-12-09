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
	t := &TextAreaWidget{
		Base:  NewBase(),
		model: m,
	}
	m.AddListener(t)
	t.Base.SetWidget(t)
	return t
}

func (t *TextAreaWidget) View() string {
	return fmt.Sprintf(
		`<sl-textarea id="%s" style="%s %s" placeholder="%s" size="medium" resize="none"></sl-textarea>
		 <style>sl-textarea#%s::part(base) {%s; %s}</style>`,
		t.IDStr(),
		t.SizeStyle(),
		t.OtherStyle(),
		t.model.placeholder,

		t.IDStr(),
		"--sl-textarea-height-medium: 100%",
		t.TextStyle(),
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

func (sm *TextAreaModel) Placeholder() string {
	return sm.placeholder
}

func (sm *TextAreaModel) SetPlaceholder(placeholder string) {
	sm.placeholder = placeholder
}
