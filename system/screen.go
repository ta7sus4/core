// Copyright (c) 2018, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// based on golang.org/x/exp/shiny:
// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package system

import (
	"image"
	"math"
)

var (
	// LogicalDPIScale is the default scaling factor for Logical DPI
	// as a multiplier on Physical DPI.
	// Smaller numbers produce smaller font sizes etc.
	LogicalDPIScale = float32(1)

	// LogicalDPIScales are per-screen name versions of LogicalDPIScale
	// these can be set from preferences (as in gi/prefs) on a per-screen
	// basis.
	LogicalDPIScales map[string]float32
)

// note: fields obtained from QScreen in Qt

// Screen contains data about each physical and / or logical screen
type Screen struct {
	// ScreenNumber is the index of this screen in the list of screens
	// maintained under Screen.
	ScreenNumber int

	// Geometry contains the geometry of the screen in window manager
	// size units, which may not be same as raw pixels (dots)
	Geometry image.Rectangle

	// DevicePixelRatio is a factor that scales the screen's
	// "natural" pixel coordinates into actual device pixels.
	// On OS-X, it is backingScaleFactor = 2.0 on "retina"
	DevicePixelRatio float32

	// PixelSize is the number of actual pixels in the screen
	// (raw display dots), computed as Size * DevicePixelRatio
	PixelSize image.Point

	// PhysicalSize is the actual physical size of the screen, in mm.
	PhysicalSize image.Point

	// LogicalDPI is the logical dots per inch of the screen,
	// which is used for all rendering.
	// It is: transient zoom factor * screen-specific multiplier * PhysicalDPI
	LogicalDPI float32

	// PhysicalDPI is the physical dots per inch of the screen,
	// for generating true-to-physical-size output.
	// It is computed as 25.4 * (PixelSize.X / PhysicalSize.X)
	// where 25.4 is the number of mm per inch.
	PhysicalDPI float32

	// Color depth of the screen, in bits.
	Depth int

	// Refresh rate in Hz
	RefreshRate float32

	// todo: not using these yet
	// AvailableGeometry        image.Rectangle
	// VirtualGeometry          image.Rectangle
	// AvailableVirtualGeometry image.Rectangle

	Orientation        ScreenOrientation `table:"-"`
	NativeOrientation  ScreenOrientation `table:"-"`
	PrimaryOrientation ScreenOrientation `table:"-"`

	Name         string
	Manufacturer string `table:"-"`
	Model        string `table:"-"`
	SerialNumber string `table:"-"`
}

// ScreenOrientation is the orientation of the device screen.
type ScreenOrientation int32 //enums:enum

const (
	// OrientationUnknown means device orientation cannot be determined.
	//
	// Equivalent on Android to Configuration.ORIENTATION_UNKNOWN
	// and on iOS to:
	//	UIDeviceOrientationUnknown
	//	UIDeviceOrientationFaceUp
	//	UIDeviceOrientationFaceDown
	OrientationUnknown ScreenOrientation = iota

	// Portrait is a device oriented so it is tall and thin.
	//
	// Equivalent on Android to Configuration.ORIENTATION_PORTRAIT
	// and on iOS to:
	//	UIDeviceOrientationPortrait
	//	UIDeviceOrientationPortraitUpsideDown
	Portrait

	// Landscape is a device oriented so it is short and wide.
	//
	// Equivalent on Android to Configuration.ORIENTATION_LANDSCAPE
	// and on iOS to:
	//	UIDeviceOrientationLandscapeLeft
	//	UIDeviceOrientationLandscapeRight
	Landscape
)

// LogicalFromPhysicalDPI computes the logical DPI used in actual screen scaling
// based on the given logical DPI scale factor (logScale), and also makes it a
// multiple of 6 to make normal font sizes look best.
func LogicalFromPhysicalDPI(logScale, pdpi float32) float32 {
	idpi := int(math.Round(float64(pdpi * logScale)))
	mdpi := idpi / 6
	mdpi *= 6
	return float32(mdpi)
}

// SetLogicalDPIScale sets the LogicalDPIScale factor for given screen name
func SetLogicalDPIScale(scrnName string, dpiScale float32) {
	if LogicalDPIScales == nil {
		LogicalDPIScales = make(map[string]float32)
	}
	LogicalDPIScales[scrnName] = dpiScale
}

// UpdateLogicalDPI updates the logical DPI of the screen
// based on ZoomFactor and LogicalDPIScale (per screen if exists)
func (sc *Screen) UpdateLogicalDPI() {
	dpisc := LogicalDPIScale
	if LogicalDPIScales != nil {
		if dsc, has := LogicalDPIScales[sc.Name]; has {
			dpisc = dsc
		}
	}
	sc.LogicalDPI = LogicalFromPhysicalDPI(dpisc, sc.PhysicalDPI)
	// fmt.Println("log dpi:", sc.LogicalDPI, "zoom:", ZoomFactor, "dpisc:", dpisc)
}

// UpdatePhysicalDPI updates the value of [Screen.PhysicalDPI] based on
// [Screen.PixelSize] and [Screen.PhysicalSize]
func (sc *Screen) UpdatePhysicalDPI() {
	sc.PhysicalDPI = 25.4 * (float32(sc.PixelSize.X) / float32(sc.PhysicalSize.X))
}

// WindowSizeToPixels returns window manager size units
// (where DevicePixelRatio is ignored) converted to pixel units --
// i.e., multiply by DevicePixelRatio
func (sc *Screen) WindowSizeToPixels(sz image.Point) image.Point {
	var psz image.Point
	psz.X = int(float32(sz.X) * sc.DevicePixelRatio)
	psz.Y = int(float32(sz.Y) * sc.DevicePixelRatio)
	return psz
}

// WindowSizeFromPixels returns window manager size units
// (where DevicePixelRatio is ignored) converted from pixel units --
// i.e., divide by DevicePixelRatio
func (sc *Screen) WindowSizeFromPixels(sz image.Point) image.Point {
	var wsz image.Point
	wsz.X = int(float32(sz.X) / sc.DevicePixelRatio)
	wsz.Y = int(float32(sz.Y) / sc.DevicePixelRatio)
	return wsz
}

// ConstrainWindowGeometry constrains window geometry to fit in the screen.
// Size is in pixel units.
func (sc *Screen) ConstrainWindowGeometry(pos, sz image.Point) (cpos, csz image.Point) {
	scSize := sc.Geometry.Size() // in window coords size
	if TheApp.Platform() == Windows {
		// these are windows-specific special numbers for minimized windows
		// can be sent here for WinGeom saved geom.
		if pos.X == -32000 {
			pos.X = 0
		}
		if pos.Y == -32000 {
			pos.Y = 50
		}
	}
	cpos, csz = ConstrainWindowGeometry(pos, sc.WindowSizeFromPixels(sz), scSize)
	csz = sc.WindowSizeToPixels(csz)
	return
}

// ConstrainWindowGeometry constrains the size and position of a window within
// given screen size, preserving the size to the extent possible.
// size is in window manager coordinates.
func ConstrainWindowGeometry(pos, sz, scSize image.Point) (cpos, csz image.Point) {
	csz = sz
	cpos = pos
	if csz.X > scSize.X {
		csz.X = scSize.X
	}
	if csz.Y > scSize.Y {
		csz.Y = scSize.Y
	}
	// don't go off the edge
	if cpos.X+csz.X > scSize.X {
		cpos.X = scSize.X - csz.X
	}
	if cpos.Y+csz.Y > scSize.Y {
		cpos.Y = scSize.Y - csz.Y
	}
	if cpos.X < 0 {
		cpos.X = 0
	}
	if cpos.Y < 0 {
		cpos.Y = 0
	}
	return
}

// InitScreenLogicalDPIFunc is a function that can be set to initialize the
// screen LogicalDPI values based on user preferences etc.  Called just before
// first window is opened.
var InitScreenLogicalDPIFunc func()
