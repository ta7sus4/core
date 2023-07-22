// Code generated by "stringer -type=VpFlags"; DO NOT EDIT.

package gi

import (
	"errors"
	"strconv"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[VpFlagPopup-24]
	_ = x[VpFlagMenu-25]
	_ = x[VpFlagCompleter-26]
	_ = x[VpFlagCorrector-27]
	_ = x[VpFlagTooltip-28]
	_ = x[VpFlagPopupDestroyAll-29]
	_ = x[VpFlagSVG-30]
	_ = x[VpFlagUpdatingNode-31]
	_ = x[VpFlagNeedsFullRender-32]
	_ = x[VpFlagDoingFullRender-33]
	_ = x[VpFlagPrefSizing-34]
	_ = x[VpFlagsN-35]
}

const _VpFlags_name = "VpFlagPopupVpFlagMenuVpFlagCompleterVpFlagCorrectorVpFlagTooltipVpFlagPopupDestroyAllVpFlagSVGVpFlagUpdatingNodeVpFlagNeedsFullRenderVpFlagDoingFullRenderVpFlagPrefSizingVpFlagsN"

var _VpFlags_index = [...]uint8{0, 11, 21, 36, 51, 64, 85, 94, 112, 133, 154, 170, 178}

func (i VpFlags) String() string {
	i -= 24
	if i < 0 || i >= VpFlags(len(_VpFlags_index)-1) {
		return "VpFlags(" + strconv.FormatInt(int64(i+24), 10) + ")"
	}
	return _VpFlags_name[_VpFlags_index[i]:_VpFlags_index[i+1]]
}

func StringToVpFlags(s string) (VpFlags, error) {
	for i := 0; i < len(_VpFlags_index)-1; i++ {
		if s == _VpFlags_name[_VpFlags_index[i]:_VpFlags_index[i+1]] {
			return VpFlags(i + 24), nil
		}
	}
	return 0, errors.New("String: " + s + " is not a valid option for type: VpFlags")
}

var _VpFlags_descMap = map[VpFlags]string{
	24: `VpFlagPopup means viewport is a popup (menu or dialog) -- does not obey
parent bounds (otherwise does)
`,
	25: `VpFlagMenu means viewport is serving as a popup menu -- affects how window
processes clicks
`,
	26: `VpFlagCompleter means viewport is serving as a popup menu for code completion --
only applies if the VpFlagMenu is also set
`,
	27: `VpFlagCorrector means viewport is serving as a popup menu for spelling correction --
only applies if the VpFlagMenu is also set
`,
	28: `VpFlagTooltip means viewport is serving as a tooltip
`,
	29: `VpFlagPopupDestroyAll means that if this is a popup, then destroy all
the children when it is deleted -- otherwise children below the main
layout under the vp will not be destroyed -- it is up to the caller to
manage those (typically these are reusable assets)
`,
	30: `VpFlagSVG means that this viewport is an SVG viewport -- SVG elements
look for this for re-rendering
`,
	31: `VpFlagUpdatingNode means that this viewport is currently handling the
update of a node, and is under the UpdtMu mutex lock.
This can be checked to see about whether to add another update or not.
`,
	32: `VpFlagNeedsFullRender means that this viewport needs to do a full
render -- this is set during signal processing and will preempt
other lower-level updates etc.
`,
	33: `VpFlagDoingFullRender means that this viewport is currently doing a
full render -- can be used by elements to drive deep rebuild in case
underlying data has changed.
`,
	34: `VpFlagPrefSizing means that this viewport is currently doing a
PrefSize computation to compute the size of the viewport
(for sizing window for example) -- affects layout size computation
only for Over
`,
	35: ``,
}

func (i VpFlags) Desc() string {
	if str, ok := _VpFlags_descMap[i]; ok {
		return str
	}
	return "VpFlags(" + strconv.FormatInt(int64(i), 10) + ")"
}
