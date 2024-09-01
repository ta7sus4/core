// Copyright (c) 2023, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"image"
	"strings"

	"cogentcore.org/core/base/errors"
	"cogentcore.org/core/colors"
	"cogentcore.org/core/icons"
	"cogentcore.org/core/svg"
	"cogentcore.org/core/system"
)

var (
	// TheApp is the current [App]; only one is ever in effect.
	TheApp = &App{App: system.TheApp}

	// AppAbout is the about information for the current app.
	// It is set by a linker flag in the core command line tool.
	AppAbout string

	// AppIcon is the svg icon for the current app.
	// It is set by a linker flag in the core command line tool.
	// It defaults to [icons.CogentCore] otherwise.
	AppIcon string = string(icons.CogentCore)
)

// App represents a Cogent Core app. It extends [system.App] to provide both system-level
// and high-level data and functions to do with the currently running application. The
// single instance of it is [TheApp], which embeds [system.TheApp].
type App struct { //types:add -setters
	system.App `set:"-"`

	// SceneInit is a function called on every newly created [Scene].
	// This can be used to set global configuration and styling for all
	// widgets in conjunction with [Scene.WidgetInit].
	SceneInit func(sc *Scene) `edit:"-"`
}

// appIconImagesCache is a cached version of [appIconImages].
var appIconImagesCache []image.Image

// appIconImages returns a slice of images of sizes 16x16, 32x32, and 48x48
// rendered from [AppIcon]. It returns nil if [AppIcon] is "" or if there is
// an error. It automatically logs any errors. It caches the result for future
// calls.
func appIconImages() []image.Image {
	if appIconImagesCache != nil {
		return appIconImagesCache
	}
	if AppIcon == "" {
		return nil
	}

	res := make([]image.Image, 3)

	sv := svg.NewSVG(16, 16)
	sv.Color = colors.Uniform(colors.FromRGB(66, 133, 244)) // Google Blue (#4285f4)
	err := sv.ReadXML(strings.NewReader(AppIcon))
	if errors.Log(err) != nil {
		return nil
	}

	sv.Render()
	res[0] = sv.Pixels

	sv.Resize(image.Pt(32, 32))
	sv.Render()
	res[1] = sv.Pixels

	sv.Resize(image.Pt(48, 48))
	sv.Render()
	res[2] = sv.Pixels
	appIconImagesCache = res
	return res
}

// makeAppBar configures a new top app bar in the given parent.
func makeAppBar(bar *Frame) {
	tb := NewToolbar(bar)
	if len(tb.Scene.AppBars) > 0 {
		tb.Makers.Normal = append(tb.Makers.Normal, tb.Scene.AppBars...)
	}
}
