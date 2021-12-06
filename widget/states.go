package widget

import (
	"fmt"
	"strconv"
)

type StringEventPublisher interface {
	EventPublisher
	fmt.Stringer
	SetString(value string)
}

type StrStateModel struct {
	Model
	value string
}

func StrState(value string) *StrStateModel {
	return &StrStateModel{
		Model: NewModel(),
		value: value,
	}
}

func (s *StrStateModel) String() string {
	return s.value
}

func (s *StrStateModel) SetValue(value string) {
	s.value = value
	s.Notify()
}

func (s *StrStateModel) SetString(value string) {
	s.value = value
	s.Notify()
}


type IntStateModel struct {
	Model
	value int
}

func IntState(value int) *IntStateModel {
	return &IntStateModel{
		Model: NewModel(),
		value: value,
	}
}

func (i *IntStateModel) SetValue(value int) {
	i.value = value
	i.Notify()
}

func (i *IntStateModel) Value() int {
	return i.value
}

func (i *IntStateModel) String() string {
	return strconv.Itoa(i.value)
}

func (i *IntStateModel) SetString(value string) {
	v, _ := strconv.Atoi(value)
	i.value = v
	i.Notify()
}

func (i *IntStateModel) Increment() {
	i.value++
	i.Notify()
}

func (i *IntStateModel) Decrement() {
	i.value--
	i.Notify()
}

func (i *IntStateModel) Add(n int) {
	i.value = i.value + n
	i.Notify()
}

func (i *IntStateModel) Sub(n int) {
	i.value = i.value - n
	i.Notify()
}

func (i *IntStateModel) Mul(n int) {
	i.value = i.value * n
	i.Notify()
}

func (i *IntStateModel) Div(n int) {
	i.value = i.value / n
	i.Notify()
}
