// Copyright (c) 2023, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"strings"

	"cogentcore.org/core/styles"
	"cogentcore.org/core/tree"
)

// BarFuncs are functions for creating control bars,
// attached to different sides of a Scene (e.g., TopAppBar at Top,
// NavBar at Bottom, etc).  Functions are called in forward order
// so first added are called first.
type BarFuncs []func(parent Widget)

// Add adds the given function for configuring a control bar
func (bf *BarFuncs) Add(fun func(parent Widget)) *BarFuncs {
	*bf = append(*bf, fun)
	return bf
}

// Call calls all the functions for configuring given widget
func (bf *BarFuncs) Call(parent Widget) {
	for _, fun := range *bf {
		fun(parent)
	}
}

// IsEmpty returns true if there are no functions added
func (bf *BarFuncs) IsEmpty() bool {
	return len(*bf) == 0
}

// Inherit adds other bar funcs in front of any existing
func (bf *BarFuncs) Inherit(obf BarFuncs) {
	if len(obf) == 0 {
		return
	}
	nbf := make(BarFuncs, len(obf), len(obf)+len(*bf))
	copy(nbf, obf)
	nbf = append(nbf, *bf...)
	*bf = nbf
}

// AddAppBar adds an AppBar Config function for an element within the scene
func (sc *Scene) AddAppBar(fun func(cf *Config)) {
	sc.AppBars.Add(fun)
}

// AddAppChooser adds an AppChooser Config function for an element within the scene
func (sc *Scene) AddAppChooser(fun func(ch *Chooser)) {
	sc.AppChoosers.Add(fun)
}

// ConfigSceneBars configures the side control bars, for main scenes
func (sc *Scene) ConfigSceneBars() {
	// at last possible moment, add app-specific app bar config
	if TheApp.AppBarConfig != nil && sc.Stage.Type.IsMain() && (sc.Stage.NewWindow || sc.Stage.FullWindow) {
		if sc.Bars.Top.IsEmpty() {
			sc.Bars.Top.Add(TheApp.AppBarConfig) // put in the top by default
		}
	}
	if !sc.Bars.Top.IsEmpty() {
		head := NewLayout(sc).Style(func(s *styles.Style) {
			s.Align.Items = styles.Center
		})
		sc.Bars.Top.Call(head)
	}
	if !sc.Bars.Left.IsEmpty() || !sc.Bars.Right.IsEmpty() {
		mid := NewLayout(sc)
		if !sc.Bars.Left.IsEmpty() {
			left := NewLayout(mid).Style(func(s *styles.Style) {
				s.Direction = styles.Column
				s.Align.Items = styles.Center
			})
			sc.Bars.Left.Call(left)
		}
		if sc.Body != nil {
			mid.AddChild(sc.Body)
		}
		if !sc.Bars.Right.IsEmpty() {
			right := NewLayout(mid).Style(func(s *styles.Style) {
				s.Direction = styles.Column
				s.Align.Items = styles.Center
			})
			sc.Bars.Right.Call(right)
		}
	} else {
		if sc.Body != nil {
			sc.AddChild(sc.Body)
		}
	}
	if !sc.Bars.Bottom.IsEmpty() {
		foot := NewLayout(sc).Style(func(s *styles.Style) {
			s.Justify.Content = styles.End
			s.Align.Items = styles.Center
		})
		sc.Bars.Bottom.Call(foot)
	}
}

// GetBar returns Bar layout widget at given side, nil if not there.
func (sc *Scene) GetBar(side styles.SideIndexes) *Layout {
	nm := strings.ToLower(side.String()) + "-bar"
	bar := sc.ChildByName(nm)
	if bar != nil {
		return bar.(*Layout)
	}
	return nil
}

// GetTopAppBar returns the TopAppBar Toolbar if it exists, nil otherwise.
func (sc *Scene) GetTopAppBar() *Toolbar {
	tb := sc.GetBar(styles.Top)
	if tb == nil {
		return nil
	}
	tab := tb.ChildByType(ToolbarType, tree.NoEmbeds)
	if tab != nil {
		return tab.(*Toolbar)
	}
	return nil
}

// RecycleToolbar constructs or returns a Toolbar in given parent Widget
func RecycleToolbar(parent Widget) *Toolbar {
	tb := parent.ChildByType(ToolbarType, tree.NoEmbeds)
	if tb != nil {
		return tb.(*Toolbar)
	}
	return NewToolbar(parent)
}

// InheritBarsWidget inherits Bar functions based on a source widget
// (e.g., Context of dialog)
func (sc *Scene) InheritBarsWidget(wi Widget) {
	if wi == nil || wi.This() == nil {
		return
	}
	wb := wi.AsWidget()
	if wb.Scene == nil {
		return
	}
	sc.InheritBars(wb.Scene)
}

// InheritBars inherits Bars functions from given other scene
// for each side that the other scene marks as inherited.
func (sc *Scene) InheritBars(osc *Scene) {
	if osc == nil {
		return
	}
	if osc.BarsInherit.Top || sc.BarsInherit.Top {
		sc.Bars.Top.Inherit(osc.Bars.Top)
		sc.BarsInherit.Top = true
	}
	if osc.BarsInherit.Bottom || sc.BarsInherit.Bottom {
		sc.Bars.Bottom.Inherit(osc.Bars.Bottom)
		sc.BarsInherit.Bottom = true
	}
	if osc.BarsInherit.Left || sc.BarsInherit.Left {
		sc.Bars.Left.Inherit(osc.Bars.Left)
		sc.BarsInherit.Left = true
	}
	if osc.BarsInherit.Right || sc.BarsInherit.Right {
		sc.Bars.Right.Inherit(osc.Bars.Right)
		sc.BarsInherit.Right = true
	}
}

//////////////////////////////////////////////////////////////
// 	Scene wrappers

// AddTopBar adds the given function for configuring a control bar
// at the top of the window
func (bd *Body) AddTopBar(fun func(parent Widget)) {
	bd.Scene.Bars.Top.Add(fun)
}

// AddLeftBar adds the given function for configuring a control bar
// on the left of the window
func (bd *Body) AddLeftBar(fun func(parent Widget)) {
	bd.Scene.Bars.Left.Add(fun)
}

// AddRightBar adds the given function for configuring a control bar
// on the right of the window
func (bd *Body) AddRightBar(fun func(parent Widget)) {
	bd.Scene.Bars.Right.Add(fun)
}

// AddBottomBar adds the given function for configuring a control bar
// at the bottom of the window
func (bd *Body) AddBottomBar(fun func(parent Widget)) {
	bd.Scene.Bars.Bottom.Add(fun)
}

// AddAppBar adds an AppBar function for an element within the scene
func (bd *Body) AddAppBar(fun func(c *Config)) {
	bd.Scene.AddAppBar(fun)
}

// AddAppChooser adds a ConfigAppChooser function for an element within the scene
func (bd *Body) AddAppChooser(fun func(ch *Chooser)) {
	bd.Scene.AddAppChooser(fun)
}

// GetTopAppBar returns the TopAppBar Toolbar if it exists, nil otherwise.
func (bd *Body) GetTopAppBar() *Toolbar {
	return bd.Scene.GetTopAppBar()
}
