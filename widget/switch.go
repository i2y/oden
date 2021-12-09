package widget

import (
	"fmt"
)

type SwitchWidget struct {
	Base
	model *BoolStateModel
	label string
}

func Switch(checked bool, label string) *SwitchWidget {
	return SwitchWithModel(BoolState(checked), label)
}

func SwitchWithModel(b *BoolStateModel, label string) *SwitchWidget {
	s := &SwitchWidget{
		Base:  NewBase(),
		model: b,
		label: label,
	}
	b.AddListener(s)
	s.Base.SetWidget(s)
	return s
}

func (s *SwitchWidget) View() string {
	return fmt.Sprintf(
		`<sl-switch id="%s" style="%s %s" checked>%s</sl-switch>
		 <style>sl-switch#%s::part(base) {%s}</style>`,
		s.IDStr(),
		s.SizeStyle(),
		s.OtherStyle(),
		s.label,

		s.IDStr(),
		s.TextStyle(),
	)
}

type SwitchModel struct {
	Model
	checked bool
}

func NewSwitchModel(checked bool) *SwitchModel {
	return &SwitchModel{
		Model:   NewModel(),
		checked: checked,
	}
}

func (sm *SwitchModel) Checked() bool {
	return sm.checked
}

func (sm *SwitchModel) SetChecked(checked bool) {
	sm.checked = checked
}
