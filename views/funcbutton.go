// Copyright (c) 2023, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package views

import (
	"log/slog"
	"reflect"
	"strings"
	"unicode"

	"cogentcore.org/core/base/reflectx"
	"cogentcore.org/core/base/strcase"
	"cogentcore.org/core/core"
	"cogentcore.org/core/cursors"
	"cogentcore.org/core/events"
	"cogentcore.org/core/icons"
	"cogentcore.org/core/keymap"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/styles/abilities"
	"cogentcore.org/core/tree"
	"cogentcore.org/core/types"
)

// CallFunc calls the given function in the context of the given widget,
// popping up a dialog to prompt for any arguments and show the return
// values of the function. It is a helper function that uses [NewSoloFuncButton]
// under the hood.
func CallFunc(ctx core.Widget, fun any) {
	NewSoloFuncButton(ctx, fun).CallFunc()
}

// NewSoloFuncButton returns a standalone FuncButton with the given context
// for popping up any dialog elements.
func NewSoloFuncButton(ctx core.Widget, fun any) *FuncButton {
	fb := NewFuncButton(core.NewWidgetBase(), fun)
	fb.SetContext(ctx)
	return fb
}

// FuncButton is a button that is set up to call a function when it
// is pressed, using a dialog to prompt the user for any arguments.
// Also, it automatically sets various properties of the button like
// the name, text, tooltip, and icon based on the properties of the
// function, using [reflect] and [types]. The function must be registered
// with [types] to get documentation information, but that is not
// required; add a `//types:add` comment directive and run `core generate`
// if you want tooltips. If the function is a method, both the method and
// its receiver type must be added to [types] to get documentation.
type FuncButton struct { //core:no-new
	core.Button

	// Func is the [types.Func] associated with this button.
	// This function can also be a method, but it must be
	// converted to a [types.Func] first. It should typically
	// be set using [FuncButton.SetFunc].
	Func *types.Func `set:"-"`

	// ReflectFunc is the [reflect.Value] of the function or
	// method associated with this button. It should typically
	// bet set using [FuncButton.SetFunc].
	ReflectFunc reflect.Value `set:"-"`

	// Args are the [FuncArg] objects associated with the
	// arguments of the function. They are automatically set in
	// [SetFunc], but they can be customized to configure
	// default values and other options.
	Args []FuncArg `set:"-"`

	// Returns are the [FuncArg] objects associated with the
	// return values of the function. They are automatically
	// set in [SetFunc], but they can be customized to configure
	// options. The [FuncArg.Value]s are not set until the
	// function is called, and are thus not typically applicable
	// to access.
	Returns []FuncArg `set:"-"`

	// Confirm is whether to prompt the user for confirmation
	// before calling the function.
	Confirm bool

	// ShowReturn is whether to display the return values of
	// the function (and a success message if there are none).
	// The way that the return values are shown is determined
	// by ShowReturnAsDialog. Non-nil error return values will
	// always be shown, even if ShowReturn is set to false.
	ShowReturn bool

	// ShowReturnAsDialog, if and only if ShowReturn is true,
	// indicates to show the return values of the function in
	// a dialog, instead of in a snackbar, as they are by default.
	// If there are multiple return values from the function, or if
	// one of them is a complex type (pointer, struct, slice,
	// array, map), then ShowReturnAsDialog will
	// automatically be set to true.
	ShowReturnAsDialog bool

	// NewWindow makes the ReturnDialog a NewWindow dialog
	// (if supported by platform).
	NewWindow bool

	// WarnUnadded is whether to log warnings when a function that
	// has not been added to [types] is used. It is on by default and
	// must be set before [FuncButton.SetFunc] is called for it to
	// have any effect.
	WarnUnadded bool `default:"true"`

	// Context is used for opening Dialogs if non-nil.
	Context core.Widget

	// AfterFunc is an optional function called after the func button
	// function is executed
	AfterFunc func()
}

// FuncArg represents one argument or return value of a function
// in the context of a [FuncButton].
type FuncArg struct { //types:add -setters

	// Name is the name of the argument or return value.
	Name string

	// Tag contains any tags associated with the argument or return value,
	// which can be added programmatically to customize [core.Value] behavior.
	Tag reflect.StructTag

	// Value is the actual value of the function argument or return value.
	// It can be modified when creating a [FuncButton] to set a default value.
	Value any
}

// NewFuncButton adds a new [FuncButton] with the given function
// to the given parent. TODO(config): remove this and use autogenerated version.
func NewFuncButton(parent tree.Node, fun any) *FuncButton {
	if parent == nil {
		return tree.New[*FuncButton]().SetFunc(fun)
	} else {
		return parent.NewChild(FuncButtonType).(*FuncButton).SetFunc(fun)
	}
}

func (fb *FuncButton) OnInit() {
	fb.Button.OnInit()
	fb.WarnUnadded = true
	fb.Style(func(s *styles.Style) {
		if fb.IsReadOnly() {
			s.SetAbilities(false, abilities.Hoverable, abilities.Clickable, abilities.Activatable)
			s.Cursor = cursors.None
		}
	})
	fb.OnClick(func(e events.Event) {
		if !fb.IsReadOnly() {
			fb.CallFunc()
		}
	})
}

// SetText sets the [FuncButton.Text] and updates the tooltip to correspond to the new name.
func (fb *FuncButton) SetText(v string) *FuncButton {
	ptext := fb.Text
	fb.Text = v
	if fb.Func != nil && fb.Text != ptext && ptext != "" {
		fb.Func.Doc = types.FormatDoc(fb.Func.Doc, ptext, fb.Text)
		fb.SetTooltip(fb.Func.Doc)
	}
	return fb
}

// SetFunc sets the function associated with the FuncButton to the
// given function or method value. For documentation information for
// the function to be obtained, it must be added to [types]. SetFunc is
// automatically called by [NewFuncButton].
func (fb *FuncButton) SetFunc(fun any) *FuncButton {
	fnm := types.FuncName(fun)
	// the "-fm" suffix indicates that it is a method
	if strings.HasSuffix(fnm, "-fm") {
		fnm = strings.TrimSuffix(fnm, "-fm")
		// the last dot separates the function name
		li := strings.LastIndex(fnm, ".")
		metnm := fnm[li+1:]
		typnm := fnm[:li]
		// get rid of any parentheses and pointer receivers
		// that may surround the type name
		typnm = strings.ReplaceAll(typnm, "(*", "")
		typnm = strings.TrimSuffix(typnm, ")")
		gtyp := types.TypeByName(typnm)
		var met *types.Method
		if gtyp == nil {
			if fb.WarnUnadded {
				slog.Warn("views.FuncButton.SetFunc called with a method whose receiver type has not been added to types", "function", fnm)
			}
			met = &types.Method{Name: metnm}
		} else {
			for _, m := range gtyp.Methods {
				if m.Name == metnm {
					met = &m
					break
				}
			}
			if met == nil {
				if fb.WarnUnadded {
					slog.Warn("views.FuncButton.SetFunc called with a method that has not been added to types (even though the receiver type was, you still need to add the method itself)", "function", fnm)
				}
				met = &types.Method{Name: metnm}
			}
		}
		return fb.SetMethodImpl(met, reflect.ValueOf(fun))
	}

	if isAnonymousFunction(fnm) {
		f := &types.Func{Name: fnm, Doc: "Anonymous function " + fnm}
		return fb.SetFuncImpl(f, reflect.ValueOf(fun))
	}

	f := types.FuncByName(fnm)
	if f == nil {
		if fb.WarnUnadded {
			slog.Warn("views.FuncButton.SetFunc called with a function that has not been added to types", "function", fnm)
		}
		f = &types.Func{Name: fnm}
	}
	return fb.SetFuncImpl(f, reflect.ValueOf(fun))
}

func isAnonymousFunction(fnm string) bool {
	// FuncName.funcN indicates that a function was defined anonymously
	return len(fnm) > 0 && unicode.IsDigit(rune(fnm[len(fnm)-1])) && strings.Contains(fnm, ".func")
}

// SetFuncImpl is the underlying implementation of [FuncButton.SetFunc].
// It should typically not be used by end-user code.
func (fb *FuncButton) SetFuncImpl(gfun *types.Func, rfun reflect.Value) *FuncButton {
	fb.Func = gfun
	fb.ReflectFunc = rfun
	fb.SetArgs()
	fb.SetReturns()
	snm := fb.Func.Name
	// get name without package
	li := strings.LastIndex(snm, ".")
	isAnonymous := isAnonymousFunction(snm)
	if isAnonymous {
		snm = strings.TrimRightFunc(snm, func(r rune) bool {
			return unicode.IsDigit(r) || r == '.'
		})
		snm = strings.TrimSuffix(snm, ".func")
		// we cut at the second to last period (we want to keep the
		// receiver / package for anonymous functions)
		li = strings.LastIndex(snm[:strings.LastIndex(snm, ".")], ".")
	}
	if li >= 0 {
		snm = snm[li+1:] // must also get rid of "."
		// if we are a global function, we may have gone too far with the second to last period,
		// so we go after the last slash if there still is one
		if strings.Contains(snm, "/") {
			snm = snm[strings.LastIndex(snm, "/")+1:]
		}
	}
	snm = strings.Map(func(r rune) rune {
		if r == '(' || r == ')' || r == '*' {
			return -1
		}
		return r
	}, snm)
	txt := strcase.ToSentence(snm)
	fb.SetText(txt)
	// doc formatting interferes with anonymous functions
	if !isAnonymous {
		fb.Func.Doc = types.FormatDoc(fb.Func.Doc, snm, txt)
	}
	fb.SetTooltip(fb.Func.Doc)
	// we default to the icon with the same name as
	// the function, if it exists
	ic := icons.Icon(strcase.ToSnake(snm))
	if ic.IsValid() {
		fb.SetIcon(ic)
	}
	return fb
}

func (fb *FuncButton) GoodContext() core.Widget {
	ctx := fb.Context
	if fb.Context == nil {
		if fb.This() == nil {
			return nil
		}
		ctx = fb.This().(core.Widget)
	}
	return ctx
}

func (fb *FuncButton) CallFuncShowReturns() {
	if fb.AfterFunc != nil {
		defer fb.AfterFunc()
	}
	if len(fb.Args) == 0 {
		rets := fb.ReflectFunc.Call(nil)
		fb.ShowReturnsDialog(rets)
		return
	}
	rargs := make([]reflect.Value, len(fb.Args))
	for i, arg := range fb.Args {
		rargs[i] = reflect.ValueOf(arg.Value)
	}
	rets := fb.ReflectFunc.Call(rargs)
	fb.ShowReturnsDialog(rets)
}

// ConfirmDialog runs the confirm dialog
func (fb *FuncButton) ConfirmDialog() {
	ctx := fb.GoodContext()
	d := core.NewBody().AddTitle(fb.Text + "?").AddText("Are you sure you want to run " + fb.Text + "? " + fb.Tooltip)
	d.AddBottomBar(func(parent core.Widget) {
		d.AddCancel(parent)
		d.AddOK(parent).SetText(fb.Text).OnClick(func(e events.Event) {
			fb.CallFuncShowReturns()
		})
	})
	d.RunDialog(ctx)
}

// CallFunc calls the function or method associated with this button,
// prompting the user for any arguments.
func (fb *FuncButton) CallFunc() {
	ctx := fb.GoodContext()
	if len(fb.Args) == 0 {
		if !fb.Confirm {
			fb.CallFuncShowReturns()
			return
		}
		fb.ConfirmDialog()
		return
	}
	// TODO(config)
	// // go directly to the dialog if there is one
	// if len(fb.Args) == 1 && OpenDialog(fb.Args[0], ctx, func() {
	// 	fb.CallFuncShowReturns()
	// }, func() {
	// 	makeTmpWidget(fb.Args[0])
	// }) {
	// 	return
	// }
	d := core.NewBody().AddTitle(fb.Text).AddText(fb.Tooltip)
	str := FuncArgsToStruct(fb.Args)
	NewStructView(d).SetStruct(str.Addr().Interface())
	d.AddBottomBar(func(parent core.Widget) {
		d.AddCancel(parent)
		d.AddOK(parent).SetText(fb.Text).OnClick(func(e events.Event) {
			d.Close() // note: the other Close event happens too late!
			fb.CallFuncShowReturns()
		})
	})
	d.RunDialog(ctx)
}

// FuncArgsToStruct converts a slice of [FuncArg] objects
// to a new non-pointer struct [reflect.Value].
func FuncArgsToStruct(args []FuncArg) reflect.Value {
	fields := make([]reflect.StructField, len(args))
	for i, arg := range args {
		fields[i] = reflect.StructField{
			Name: strcase.ToCamel(arg.Name),
			Type: reflect.TypeOf(arg.Value),
			Tag:  arg.Tag,
		}
	}
	typ := reflect.StructOf(fields)
	value := reflect.New(typ).Elem()
	for i, arg := range args {
		value.Field(i).Set(reflect.ValueOf(arg.Value))
	}
	return value
}

// SetMethodImpl is the underlying implementation of [FuncButton.SetFunc] for methods.
// It should typically not be used by end-user code.
func (fb *FuncButton) SetMethodImpl(gmet *types.Method, rmet reflect.Value) *FuncButton {
	return fb.SetFuncImpl(&types.Func{
		Name:       gmet.Name,
		Doc:        gmet.Doc,
		Directives: gmet.Directives,
		Args:       gmet.Args,
		Returns:    gmet.Returns,
	}, rmet)
}

// ShowReturnsDialog runs a dialog displaying the given function return
// values for the function associated with the function button. It does
// nothing if [FuncButton.ShowReturn] is dialog
func (fb *FuncButton) ShowReturnsDialog(rets []reflect.Value) {
	if !fb.ShowReturn {
		for _, ret := range rets {
			if err, ok := ret.Interface().(error); ok && err != nil {
				core.ErrorSnackbar(fb, err, fb.Text+" failed")
				return
			}
		}
		return
	}
	ctx := fb.GoodContext()
	if ctx == nil {
		return
	}
	for i, ret := range rets {
		fb.Returns[i].Value = ret.Interface()
	}
	main := "Result of " + fb.Text
	if len(rets) == 0 {
		main = fb.Text + " succeeded"
	}
	if !fb.ShowReturnAsDialog {
		txt := main
		if len(fb.Returns) > 0 {
			txt += ": "
			for i, ret := range fb.Returns {
				txt += reflectx.ToString(ret.Value)
				if i < len(fb.Returns)-1 {
					txt += ", "
				}
			}
		}
		core.MessageSnackbar(ctx, txt)
		return
	}
	// TODO(config)
	// // go directly to the dialog if there is one
	// if len(fb.Returns) == 1 && OpenDialog(fb.Returns[0], ctx, nil, func() {
	// 	makeTmpWidget(fb.Returns[0])
	// }) {
	// 	return
	// }
	d := core.NewBody().AddTitle(main).AddText(fb.Tooltip).AddOKOnly()
	str := FuncArgsToStruct(fb.Returns)
	NewStructView(d).SetStruct(str.Addr().Interface()).SetReadOnly(true)
	if fb.NewWindow {
		d.RunWindowDialog(ctx)
	} else {
		d.RunDialog(ctx)
	}
}

// SetArgs sets the appropriate [Value] objects for the
// arguments of the function associated with the function button.
// It is called in [FuncButton.SetFunc] and should typically not
// be called by end-user code.
func (fb *FuncButton) SetArgs() {
	narg := fb.ReflectFunc.Type().NumIn()
	fb.Args = make([]FuncArg, narg)
	for i := range fb.Args {
		typ := fb.ReflectFunc.Type().In(i)

		name := ""
		if fb.Func.Args != nil && len(fb.Func.Args) > i {
			name = fb.Func.Args[i]
		} else {
			name = reflectx.NonPointerType(typ).Name()
		}

		fb.Args[i] = FuncArg{
			Name:  name,
			Value: reflect.New(typ).Elem().Interface(),
		}
	}
}

// SetReturns sets the appropriate [Value] objects for the
// return values of the function associated with the function
// button. It is called in [FuncButton.SetFunc] and should
// typically not be called by end-user code.
func (fb *FuncButton) SetReturns() {
	nret := fb.ReflectFunc.Type().NumOut()
	fb.Returns = make([]FuncArg, nret)
	hasComplex := false
	for i := range fb.Returns {
		typ := fb.ReflectFunc.Type().Out(i)
		if !hasComplex {
			k := typ.Kind()
			if k == reflect.Pointer || k == reflect.Struct || k == reflect.Slice || k == reflect.Array || k == reflect.Map {
				hasComplex = true
			}
		}

		name := ""
		if fb.Func.Returns != nil && len(fb.Func.Returns) > i {
			name = fb.Func.Returns[i]
		} else {
			name = reflectx.NonPointerType(typ).Name()
		}

		fb.Returns[i] = FuncArg{
			Name:  name,
			Value: reflect.New(typ).Elem().Interface(),
		}
	}
	if nret > 1 || hasComplex {
		fb.ShowReturnAsDialog = true
	}
}

// SetKey sets the shortcut of the function button from the given [keymap.Functions]
func (fb *FuncButton) SetKey(kf keymap.Functions) *FuncButton {
	fb.Button.SetKey(kf)
	return fb
}

// makeTmpWidget makes a temporary widget in a temporary parent for the given value.
func makeTmpWidget(v Value) {
	v.SetWidget(core.NewWidgetBase().NewChild(v.WidgetType()).(core.Widget))
}
