package widget

import (
	"fmt"
	"reflect"
)

/* I wrote some parts that use "reflect" referring to go-funk's code */

func ForEach(items interface{}, builder interface{}) []Widget {
	var widgets []Widget

	if !isIterable(items) {
		panic("First parameter must be an iterable")
	}

	if !isFunction(builder, 1, 1) {
		panic("Second argument must be a function")
	}

	funcValue := reflect.ValueOf(builder)

	funcType := funcValue.Type()
	if funcType.Out(0).String() != "widget.Widget" {
		panic("Return argument should be a Widget")
	}

	itemValue := reflect.ValueOf(items)
	itemType := itemValue.Type()

	if funcType.In(0) != itemType.Elem() {
		panic(fmt.Sprintf("First parameter must be %s", itemType))
	}

	for i := 0; i < itemValue.Len(); i++ {
		elem := itemValue.Index(i)
		widget := funcValue.Call([]reflect.Value{elem})[0].Interface().(Widget)
		widgets = append(widgets, widget)
	}

	return widgets
}

/* これらのreflectを使ったコードはgo-funkのコードを参考にしました */
func isIterable(in interface{}) bool {
	if in == nil {
		return false
	}
	arrType := reflect.TypeOf(in)

	kind := arrType.Kind()

	return kind == reflect.Array || kind == reflect.Slice || kind == reflect.Map
}

func isFunction(in interface{}, num ...int) bool {
	funcType := reflect.TypeOf(in)

	result := funcType != nil && funcType.Kind() == reflect.Func

	if len(num) >= 1 {
		result = result && funcType.NumIn() == num[0]
	}

	if len(num) == 2 {
		result = result && funcType.NumOut() == num[1]
	}

	return result
}
