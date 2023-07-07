// Copyright 2019 The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ios

package mobile

func (app *appImpl) FontPaths() []string {
	return []string{"/System/Library/Fonts"}
}
