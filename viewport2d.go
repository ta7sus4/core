// Copyright (c) 2018, Randall C. O'Reilly. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gi

import (
	"fmt"
	"github.com/rcoreilly/goki/ki"
	// "golang.org/x/image/font"
	"image"
	"image/draw"
	"image/png"
	"io"
	"log"
	"os"
	"reflect"
)

// Viewport2D provides an image and a stack of Paint contexts for drawing onto the image
// with a convenience forwarding of the Paint methods operating on the current Paint
type Viewport2D struct {
	Node2DBase
	ViewBox ViewBox2D   `svg:"viewBox",desc:"viewbox within any parent Viewport2D"`
	Render  RenderState `json:"-",desc:"render state for rendering"`
	Pixels  *image.RGBA `json:"-",desc:"pixels that we render into"`
	Backing *image.RGBA `json:"-",desc:"if non-nil, this is what goes behind our image -- copied from our region in parent image -- allows us to re-render cleanly into parent, even with transparency"`
}

// must register all new types so type names can be looked up by name -- e.g., for json
var KiT_Viewport2D = ki.KiTypes.AddType(&Viewport2D{})

// NewViewport2D creates a new image.RGBA with the specified width and height
// and prepares a context for rendering onto that image.
func NewViewport2D(width, height int) *Viewport2D {
	return NewViewport2DForRGBA(image.NewRGBA(image.Rect(0, 0, width, height)))
}

// NewViewport2DForImage copies the specified image into a new image.RGBA
// and prepares a context for rendering onto that image.
func NewViewport2DForImage(im image.Image) *Viewport2D {
	return NewViewport2DForRGBA(imageToRGBA(im))
}

// NewViewport2DForRGBA prepares a context for rendering onto the specified image.
// No copy is made.
func NewViewport2DForRGBA(im *image.RGBA) *Viewport2D {
	vp := &Viewport2D{
		ViewBox: ViewBox2D{Size: image.Point{im.Bounds().Size().X, im.Bounds().Size().Y}},
		Pixels:  im,
	}
	vp.Render.Image = vp.Pixels
	return vp
}

// resize viewport, creating a new image (no point in trying to resize the image -- need to re-render) -- updates ViewBox Size too -- triggers update -- wrap in other UpdateStart/End calls as appropriate
func (vp *Viewport2D) Resize(width, height int) {
	if vp.Pixels.Bounds().Size().X == width && vp.Pixels.Bounds().Size().Y == height {
		return // already good
	}
	vp.UpdateStart()
	vp.Pixels = image.NewRGBA(image.Rect(0, 0, width, height))
	vp.Render.Image = vp.Pixels
	vp.ViewBox.Size = image.Point{width, height}
	vp.UpdateEnd()
	vp.FullRender2DRoot()
}

////////////////////////////////////////////////////////////////////////////////////////
//  Main Rendering code

// draw our image into parents -- called at right place in Render
func (vp *Viewport2D) DrawIntoParent(parVp *Viewport2D) {
	r := vp.ViewBox.Bounds()
	if vp.Backing != nil {
		draw.Draw(parVp.Pixels, r, vp.Backing, image.ZP, draw.Src)
	}
	draw.Draw(parVp.Pixels, r, vp.Pixels, image.ZP, draw.Src)
}

// copy our backing image from parent -- called at right place in Render
func (vp *Viewport2D) CopyBacking(parVp *Viewport2D) {
	r := vp.ViewBox.Bounds()
	if vp.Backing == nil {
		vp.Backing = image.NewRGBA(vp.ViewBox.SizeRect())
	}
	draw.Draw(vp.Backing, r, parVp.Pixels, image.ZP, draw.Src)
}

func (vp *Viewport2D) DrawIntoWindow() {
	wini := vp.FindParentByType(reflect.TypeOf(Window{}))
	if wini != nil {
		win := (wini).(*Window)
		// width, height := win.Win.Size() // todo: update size of our window
		s := win.Win.Screen()
		s.CopyRGBA(vp.Pixels, vp.Pixels.Bounds())
		win.Win.FlushImage()
	}
}

////////////////////////////////////////////////////////////////////////////////////////
// Node2D interface

func (vp *Viewport2D) GiNode2D() *Node2DBase {
	return &vp.Node2DBase
}

func (vp *Viewport2D) GiViewport2D() *Viewport2D {
	return vp
}

func (vp *Viewport2D) InitNode2D() {
	vp.NodeSig.Connect(vp.This, func(vpki, vpa ki.Ki, sig int64, data interface{}) {
		vp, ok := vpki.(*Viewport2D)
		if !ok {
			return
		}
		fmt.Printf("viewport: %v rendering due to signal: %v from node: %v\n", vp.PathUnique(), sig, vpa.PathUnique())
		vp.FullRender2DRoot()
	})
}

func (vp *Viewport2D) Style2D() {
	vp.Style2DWidget()
}

func (vp *Viewport2D) Layout2D(iter int) {
	if iter == 0 {
		vp.Layout.AllocSize.SetFromPoint(vp.ViewBox.Size)
	}
}

func (vp *Viewport2D) Node2DBBox() image.Rectangle {
	return vp.ViewBox.Bounds()
}

func (vp *Viewport2D) Render2D() {
	// todo: we should get layout info and set our size??
	vp.SetWinBBox(vp.Node2DBBox())
	if vp.Viewport != nil {
		vp.CopyBacking(vp.Viewport) // full re-render is when we copy the backing
		vp.DrawIntoParent(vp.Viewport)
	} else { // top-level, try drawing into window
		vp.DrawIntoWindow()
	}
}

func (vp *Viewport2D) CanReRender2D() bool {
	return true // always true for viewports
}

func (g *Viewport2D) FocusChanged2D(gotFocus bool) {
}

////////////////////////////////////////////////////////////////////////////////////////
//  Signal Handling

// each node calls this signal method to notify its parent viewport whenever it changes, causing a re-render
func SignalViewport2D(vpki, node ki.Ki, sig int64, data interface{}) {
	vpgi, ok := vpki.(Node2D)
	if !ok {
		return
	}
	vp := vpgi.GiViewport2D()
	if vp == nil { // should not happen -- should only be called on viewports
		return
	}
	gii, ok := node.(Node2D)
	if !ok { // should not happen..
		return
	}
	fmt.Printf("viewport: %v rendering due to signal: %v from node: %v\n", vp.PathUnique(), sig, node.PathUnique())

	// todo: probably need better ways of telling how much re-rendering is needed
	if sig == ki.SignalChildAdded {
		vp.Init2DRoot()
		vp.Style2DRoot()
		vp.Render2DRoot()
	} else {
		if gii.CanReRender2D() {
			gii.Render2D()
			vp.Render2D() // redraw us
		} else {
			vp.Style2DFromNode(gii.GiNode2D()) // restyle only from affected node downward
			vp.Render2DRoot()                  // need to re-render entirely..
		}
	}
}

////////////////////////////////////////////////////////////////////////////////////////
// Root-level Viewport API -- does all the recursive calls

// initialize scene graph
func (vp *Viewport2D) Init2DRoot() {
	vp.FunDownMeFirst(0, vp, func(k ki.Ki, level int, d interface{}) bool {
		gii, ok := (interface{}(k)).(Node2D)
		if !ok {
			// todo: need to detect the thing that wraps a 3D node inside a 2D, and stop there
			log.Printf("Node %v in Viewport2D does NOT implement Node2D interface -- it should!\n", k.PathUnique())
			return false
		}
		gi := gii.GiNode2D()
		gi.InitNode2DBase()
		gii.InitNode2D()
		return true
	})
}

// full render of the tree
func (vp *Viewport2D) FullRender2DRoot() {
	vp.Init2DRoot()
	vp.Style2DRoot()
	vp.Layout2DRoot()
	vp.RenderOnly2DRoot()
}

// render of the tree -- after it has already been initialized and styled
func (vp *Viewport2D) Render2DRoot() {
	vp.Layout2DRoot()
	vp.RenderOnly2DRoot()
}

// this only needs to be done on a structural update
func (vp *Viewport2D) Style2DFromNode(gi *Node2DBase) {
	gi.FunDownMeFirst(0, vp, func(k ki.Ki, level int, d interface{}) bool {
		gii, ok := k.(Node2D)
		if !ok { // todo: error message already in InitNode2D
			log.Printf("Node %v in Viewport2D does NOT implement Node2D interface -- it should!\n", k.PathUnique())
			return false // going into a different type of thing, bail
		}
		gii.Style2D()
		return true
	})
}

// this only needs to be done on a structural update
func (vp *Viewport2D) Style2DRoot() {
	vp.FunDownMeFirst(0, vp, func(k ki.Ki, level int, d interface{}) bool {
		gii, ok := k.(Node2D)
		if !ok { // error message already in InitNode2D
			log.Printf("Node %v in Viewport2D does NOT implement Node2D interface -- it should!\n", k.PathUnique())
			return false // going into a different type of thing, bail
		}
		gii.Style2D()
		return true
	})
}

func (vp *Viewport2D) Layout2DRoot() {
	// todo: support multiple iterations if necc

	// layout happens in depth-first manner -- requires two functions
	vp.FunDownDepthFirst(0, vp,
		func(k ki.Ki, level int, d interface{}) bool { // this is for testing whether to process node
			gii, ok := k.(Node2D)
			if !ok {
				return false
			}
			gi := gii.GiNode2D()
			if gi.Paint.Off { // off below this
				return false
			}
			return true
		},
		func(k ki.Ki, level int, d interface{}) bool {
			gii, ok := k.(Node2D)
			if !ok {
				return false
			}
			gi := gii.GiNode2D()
			if gi.Paint.Off { // off below this
				return false
			}
			gii.Layout2D(0)
			return true
		})

	// second pass we add the parent positions after layout -- don't want to do that in
	// render b/c then it doesn't work for local re-renders..
	vp.FunDownMeFirst(0, vp, func(k ki.Ki, level int, d interface{}) bool {
		gii, ok := k.(Node2D)
		if !ok {
			return false
		}
		gi := gii.GiNode2D()
		if gi.Paint.Off { // off below this
			return false
		}
		gii.Layout2D(1) // todo: check for multiple iterations needed..
		return true
	})

}

// just do the render only part -- not the full 3-pass version called in Render2DRoot
func (vp *Viewport2D) RenderOnly2DRoot() {
	vp.FunDownMeFirst(0, vp, func(k ki.Ki, level int, d interface{}) bool {
		gii, ok := k.(Node2D)
		if !ok {
			return false
		}
		gi := gii.GiNode2D()
		if gi.Paint.Off { // off below this
			return false
		}
		gii.Render2D()
		return true
	})
}

// SavePNG encodes the image as a PNG and writes it to disk.
func (vp *Viewport2D) SavePNG(path string) error {
	return SavePNG(path, vp.Pixels)
}

// EncodePNG encodes the image as a PNG and writes it to the provided io.Writer.
func (vp *Viewport2D) EncodePNG(w io.Writer) error {
	return png.Encode(w, vp.Pixels)
}

// todo:

// DrawPoint is like DrawCircle but ensures that a circle of the specified
// size is drawn regardless of the current transformation matrix. The position
// is still transformed, but not the shape of the point.
// func (vp *Viewport2D) DrawPoint(x, y, r float64) {
// 	pc := vp.PushNewPaint()
// 	p := pc.TransformPoint(x, y)
// 	pc.Identity()
// 	pc.DrawCircle(p.X, p.Y, r)
// 	vp.PopPaint()
// }

//////////////////////////////////////////////////////////////////////////////////
//  Image utilities

func LoadImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	im, _, err := image.Decode(file)
	return im, err
}

func LoadPNG(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return png.Decode(file)
}

func SavePNG(path string, im image.Image) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return png.Encode(file, im)
}

func imageToRGBA(src image.Image) *image.RGBA {
	dst := image.NewRGBA(src.Bounds())
	draw.Draw(dst, dst.Rect, src, image.ZP, draw.Src)
	return dst
}
