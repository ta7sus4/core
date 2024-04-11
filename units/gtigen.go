// Code generated by "core generate"; DO NOT EDIT.

package units

import (
	"cogentcore.org/core/types"
)

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/units.Value", IDName: "value", Doc: "Value and units, and converted value into raw pixels (dots in DPI)", Directives: []types.Directive{{Tool: "gti", Directive: "add"}}, Fields: []types.Field{{Name: "Value", Doc: "Value is the value in terms of the specified unit"}, {Name: "Unit", Doc: "Unit is the unit used for the value"}, {Name: "Dots", Doc: "Dots is the computed value in raw pixels (dots in DPI)"}, {Name: "Custom", Doc: "Custom is a custom function that returns the dots of the value.\nIf it is non-nil, it overrides all other fields.\nOtherwise, the standard ToDots with the other fields is used."}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/units.XY", IDName: "xy", Doc: "XY represents unit Value for X and Y dimensions", Directives: []types.Directive{{Tool: "gti", Directive: "add"}}, Fields: []types.Field{{Name: "X", Doc: "X is the horizontal axis value"}, {Name: "Y", Doc: "Y is the vertical axis value"}}})
