// Copyright (c) 2018, Randall C. O'Reilly. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
	Package Gi (GoGi) provides a complete Graphical Interface based on GoKi Tree Node structs

	2D and 3D scenegraphs supported, each rendering to respective Viewport2D or 3D
	which in turn can be integrated within the other type of scenegraph.
	Within 2D scenegraph, the following are supported
		* SVG-based rendering nodes for basic shapes, paths, curves, arcs etc, with SVG / CSS properties
		* Widget nodes for GUI actions (Buttons, etc), with support for full SVG-based rendering of styles, using Qt-based naming and functionality, including TreeView, TableView
		* Layouts for placing widgets, based on QtQuick model
		* Primary geometry is managed in terms of inherited position offsets from top-left,
		  in a widget-like manner, but svg-based transforms also supported.
*/
package gi

import (
	// "fmt"
	"github.com/rcoreilly/goki/ki"
	// "gopkg.in/go-playground/colors.v1"
	"image"
	// "log"
	// "reflect"
	// "strconv"
	// "strings"
)

////////////////////////////////////////////////////////////////////////////////////////
// 2D  Nodes

/*
Base struct node for 2D rendering tree -- renders to a bitmap using Paint / Viewport rendering functions

Rendering is done in 3 separate passes:
	1. Style2D: In a MeFirst downward pass, all properties are cached out in an inherited manner, and incorporating any css styles, into either the Paint (SVG) or Style (Widget) object for each Node.
	2. Layout2D: In a DepthFirst downward pass, layout is updated for each node, with Layout parent nodes arranging layout-aware child nodes according to their properties.  Text2D nodes are layout aware, but basic SVG nodes are not -- they must be incorporated into widget parents to obtain layout (e.g., Icon widget).
	3. Render2D: Final MeFirst rendering pass -- also individual nodes can optionally re-render directly depending on their type, without requiring a full re-render. Layout geom is incorporated and WinBBox bounding box computed at this stage.
*/
type Node2DBase struct {
	NodeBase
	Style    Style       `desc:"styling settings for this item -- set in SetStyle2D during an initialization step, and when the structure changes"`
	Paint    Paint       `json:"-",desc:"full paint information for this node"`
	Viewport *Viewport2D `json:"-",desc:"our viewport -- set in InitNode2D (Base typically) and used thereafter"`
	Layout   LayoutData  `desc:"all the layout information for this item"`
}

// must register all new types so type names can be looked up by name -- e.g., for json
var KiT_Node2DBase = ki.KiTypes.AddType(&Node2DBase{})

// primary interface for all Node2D nodes
type Node2D interface {
	// get a generic Node2DBase for our node -- not otherwise possible -- don't want to interface everything that a base node can do as that would be highly redundant
	GiNode2D() *Node2DBase
	// if this is a Viewport2D-derived node, get it as a Viewport2D, else return nil
	GiViewport2D() *Viewport2D
	// initialize a node -- setup connections etc -- before this call, InitNodeBase is called to set basic inits including setting Viewport and connecting node signal to parent vp -- must be robust to being called repeatedly
	InitNode2D()
	// In a MeFirst downward pass, all properties are cached out in an inherited manner, and incorporating any css styles, into either the Paint or Style object for each Node, depending on the type of node (SVG does Paint, Widget does Style)
	Style2D()
	// Layout2D: In a DepthFirst downward pass, layout is updated for each node, with Layout parent nodes arranging layout-aware child nodes according to their properties.  Text2D nodes are layout aware, but basic SVG nodes are not -- they must be incorporated into widget parents to obtain layout (e.g., Icon widget).  WinBBox bounding box is computed at this stage.
	Layout2D(iter int)
	// get the bounding box of this node relative to its parent viewport -- used in computing WinBBox, must be called during Render
	Node2DBBox() image.Rectangle
	// Render2D: Final MeFirst rendering pass -- individual nodes can optionally re-render directly depending on their type, without requiring a full re-render.
	Render2D()
	// Can this node re-render itself directly using cached data?  only for nodes that paint an opaque background first (e.g., widgets) -- optimizes local redraw when possible -- always true for sub-viewports
	CanReRender2D() bool
	// function called on node when it gets or loses focus -- focus flag has current state too
	FocusChanged2D(gotFocus bool)
}

////////////////////////////////////////////////////////////////////////////////////////
// 2D basic infrastructure code

func (g *Node2DBase) ParentNode2D() *Node2DBase {
	if g.Parent == nil {
		return nil
	}
	gii, ok := g.Parent.(Node2D)
	if ok {
		return gii.GiNode2D()
	}
	return nil
}

// handles basic node initialization -- InitNode2D can then do special things
func (g *Node2DBase) InitNode2DBase() {
	g.Viewport = g.FindViewportParent()
	if g.Viewport != nil { // default for most cases -- delete connection of not
		g.NodeSig.Connect(g.Viewport.This, SignalViewport2D)
	}
	g.Style.Defaults()
	g.Paint.Defaults()
	g.Layout.Defaults() // doesn't overwrite
}

// style the Paint values directly from node properties -- for SVG-style nodes
func (g *Node2DBase) Style2DSVG() {
	gii, ok := g.This.(Node2D)
	if g.Viewport == nil { // robust
		// fmt.Printf("in style, initializing node %v\n", g.PathUnique())
		g.InitNode2DBase()
		if ok {
			gii.InitNode2D()
		}
	}
	if g.Viewport == nil {
		return
	}
	pg := g.CopyParentPaint() // svg always inherits all paint settings from parent
	g.Paint.SetStyle(&pg.Paint, &PaintDefault, g.KiProperties())
	g.Paint.SetUnitContext(&g.Viewport.Render, 0) // svn only has to set units here once
	g.Layout.Reset()                              // start with a fresh layout

}

// style the Style values from node properties -- for Widget-style nodes
func (g *Node2DBase) Style2DWidget() {
	gii, ok := g.This.(Node2D)
	if g.Viewport == nil { // robust
		// fmt.Printf("in style, initializing node %v\n", g.PathUnique())
		g.InitNode2DBase()
		if ok {
			gii.InitNode2D()
		}
	}
	if g.Viewport == nil {
		return
	}
	pg := g.ParentNode2D()
	if pg != nil {
		g.Style.SetStyle(&pg.Style, &StyleDefault, g.KiProperties())
	} else {
		g.Style.SetStyle(nil, &StyleDefault, g.KiProperties())
	}
	g.Style.SetUnitContext(&g.Viewport.Render, 0) // todo: test for use of el-relative
	g.Layout.SetFromStyle(&g.Style.Layout)        // also does reset
}

// find parent viewport -- uses GiViewport2D() method on Node2D interface
func (g *Node2DBase) FindViewportParent() *Viewport2D {
	var parVp *Viewport2D
	g.FunUpParent(0, g.This, func(k ki.Ki, level int, d interface{}) bool {
		gii, ok := k.(Node2D)
		if !ok {
			return false // don't keep going up
		}
		vp := gii.GiViewport2D()
		if vp != nil {
			parVp = vp
			return false // done
		}
		return true
	})
	return parVp
}

// copy our paint from our parents -- called during Style for SVG
func (g *Node2DBase) CopyParentPaint() *Node2DBase {
	pg := g.ParentNode2D()
	if pg != nil {
		g.Paint = pg.Paint
	}
	return pg
}

// get our bbox from Layout allocation
func (g *Node2DBase) WinBBoxFromAlloc() image.Rectangle {
	tp := g.Paint.TransformPoint(g.Layout.AllocPos.X, g.Layout.AllocPos.Y)
	ts := g.Paint.TransformPoint(g.Layout.AllocSize.X, g.Layout.AllocSize.Y)
	return image.Rect(int(tp.X), int(tp.Y), int(tp.X+ts.X), int(tp.Y+ts.Y))
}

// set our window-level BBox from vp and our bbox
func (g *Node2DBase) SetWinBBox(bb image.Rectangle) {
	if g.Viewport != nil {
		g.WinBBox = bb.Add(image.Point{g.Viewport.WinBBox.Min.X, g.Viewport.WinBBox.Min.Y})
	} else {
		g.WinBBox = bb
	}
}

// add the position of our parent to our layout position, to be called after getting geom from layout
func (g *Node2DBase) AddParentPos() {
	pg := g.ParentNode2D()
	if pg != nil {
		g.Layout.AllocPos = g.Layout.AllocPos.Add(pg.Layout.AllocPos)
	}
}

// for widgets: if a layout positioned us, then use that, otherwise use our
// user-specified pos, size this should be called in layout2d for iter > 0
func (g *Node2DBase) GeomFromLayout() {
	g.AddParentPos()
	g.Layout.UsePos(&g.Style.Layout)
	g.Layout.UseSize(&g.Style.Layout)
}

func (g *Node2DBase) DefaultGeom() {
	g.SetWinBBox(g.WinBBoxFromAlloc())
}
