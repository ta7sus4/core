// Copyright (c) 2024, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package metric

import (
	"math"

	"cogentcore.org/core/tensor"
)

// ClosestRow returns the closest fit between probe pattern and patterns in
// a "vocabulary" tensor with outermost row dimension, using given metric
// function, *which must have the Increasing property*, i.e., larger = further.
// returns the row and metric value for that row.
// note: this does _not_ use any existing Indexes for the probe (but does for the vocab).
func ClosestRow(probe, vocab *tensor.Indexed, mfun MetricFunc) (int, float64) {
	rows, _ := vocab.Tensor.RowCellSize()
	mi := -1
	out := tensor.NewFloatScalar(0.0)
	// todo: need a 1d view of both spaces
	mind := math.MaxFloat64
	prview := tensor.NewIndexed(tensor.New1DViewOf(probe.Tensor))
	for ri := range rows {
		sub := tensor.NewIndexed(tensor.New1DViewOf(vocab.SubSpace(ri)))
		mfun(prview, sub, out)
		d := out.Tensor.Float1D(0)
		if d < mind {
			mi = ri
			mind = d
		}
	}
	return mi, mind
}
