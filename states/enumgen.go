// Code generated by "enumgen"; DO NOT EDIT.

package states

import (
	"errors"
	"strconv"
	"strings"
	"sync/atomic"

	"goki.dev/enums"
)

var _AbilitiesValues = []Abilities{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

// AbilitiesN is the highest valid value
// for type Abilities, plus one.
const AbilitiesN Abilities = 12

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _AbilitiesNoOp() {
	var x [1]struct{}
	_ = x[Editable-(0)]
	_ = x[Selectable-(1)]
	_ = x[Activatable-(2)]
	_ = x[Draggable-(3)]
	_ = x[Droppable-(4)]
	_ = x[Slideable-(5)]
	_ = x[Scrollable-(6)]
	_ = x[Focusable-(7)]
	_ = x[FocusWithinable-(8)]
	_ = x[Checkable-(9)]
	_ = x[Hoverable-(10)]
	_ = x[LongHoverable-(11)]
}

var _AbilitiesNameToValueMap = map[string]Abilities{
	`Editable`:        0,
	`editable`:        0,
	`Selectable`:      1,
	`selectable`:      1,
	`Activatable`:     2,
	`activatable`:     2,
	`Draggable`:       3,
	`draggable`:       3,
	`Droppable`:       4,
	`droppable`:       4,
	`Slideable`:       5,
	`slideable`:       5,
	`Scrollable`:      6,
	`scrollable`:      6,
	`Focusable`:       7,
	`focusable`:       7,
	`FocusWithinable`: 8,
	`focuswithinable`: 8,
	`Checkable`:       9,
	`checkable`:       9,
	`Hoverable`:       10,
	`hoverable`:       10,
	`LongHoverable`:   11,
	`longhoverable`:   11,
}

var _AbilitiesDescMap = map[Abilities]string{
	0:  `Editable means it can switch between ReadOnly and not`,
	1:  `Selectable means it can be Selected`,
	2:  `Activatable means it can be made Active`,
	3:  `Draggable means it can be Dragged`,
	4:  `Droppable means it can receive Drop events (not specific to current Drag item, just generally)`,
	5:  `Slideable means it has a slider element that can be dragged to change value. Cannot be both Draggable and Slideable.`,
	6:  `Scrollable means it can be Scrolled`,
	7:  `Focusable means it can be Focused`,
	8:  `FocusWithinable means it can be FocusedWithin`,
	9:  `Checkable means it can be Checked`,
	10: `Hoverable means it can be Hovered`,
	11: `LongHoverable means it can be LongHovered`,
}

var _AbilitiesMap = map[Abilities]string{
	0:  `Editable`,
	1:  `Selectable`,
	2:  `Activatable`,
	3:  `Draggable`,
	4:  `Droppable`,
	5:  `Slideable`,
	6:  `Scrollable`,
	7:  `Focusable`,
	8:  `FocusWithinable`,
	9:  `Checkable`,
	10: `Hoverable`,
	11: `LongHoverable`,
}

// String returns the string representation
// of this Abilities value.
func (i Abilities) String() string {
	str := ""
	for _, ie := range _AbilitiesValues {
		if i.HasFlag(ie) {
			ies := ie.BitIndexString()
			if str == "" {
				str = ies
			} else {
				str += "|" + ies
			}
		}
	}
	return str
}

// BitIndexString returns the string
// representation of this Abilities value
// if it is a bit index value
// (typically an enum constant), and
// not an actual bit flag value.
func (i Abilities) BitIndexString() string {
	if str, ok := _AbilitiesMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the Abilities value from its
// string representation, and returns an
// error if the string is invalid.
func (i *Abilities) SetString(s string) error {
	*i = 0
	return i.SetStringOr(s)
}

// SetStringOr sets the Abilities value from its
// string representation while preserving any
// bit flags already set, and returns an
// error if the string is invalid.
func (i *Abilities) SetStringOr(s string) error {
	flgs := strings.Split(s, "|")
	for _, flg := range flgs {
		if val, ok := _AbilitiesNameToValueMap[flg]; ok {
			i.SetFlag(true, &val)
		} else if val, ok := _AbilitiesNameToValueMap[strings.ToLower(flg)]; ok {
			i.SetFlag(true, &val)
		} else {
			return errors.New(flg + " is not a valid value for type Abilities")
		}
	}
	return nil
}

// Int64 returns the Abilities value as an int64.
func (i Abilities) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the Abilities value from an int64.
func (i *Abilities) SetInt64(in int64) {
	*i = Abilities(in)
}

// Desc returns the description of the Abilities value.
func (i Abilities) Desc() string {
	if str, ok := _AbilitiesDescMap[i]; ok {
		return str
	}
	return i.String()
}

// AbilitiesValues returns all possible values
// for the type Abilities.
func AbilitiesValues() []Abilities {
	return _AbilitiesValues
}

// Values returns all possible values
// for the type Abilities.
func (i Abilities) Values() []enums.Enum {
	res := make([]enums.Enum, len(_AbilitiesValues))
	for i, d := range _AbilitiesValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type Abilities.
func (i Abilities) IsValid() bool {
	_, ok := _AbilitiesMap[i]
	return ok
}

// HasFlag returns whether these
// bit flags have the given bit flag set.
func (i Abilities) HasFlag(f enums.BitFlag) bool {
	return atomic.LoadInt64((*int64)(&i))&(1<<uint32(f.Int64())) != 0
}

// SetFlag sets the value of the given
// flags in these flags to the given value.
func (i *Abilities) SetFlag(on bool, f ...enums.BitFlag) {
	var mask int64
	for _, v := range f {
		mask |= 1 << v.Int64()
	}
	in := int64(*i)
	if on {
		in |= mask
		atomic.StoreInt64((*int64)(i), in)
	} else {
		in &^= mask
		atomic.StoreInt64((*int64)(i), in)
	}
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Abilities) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Abilities) UnmarshalText(text []byte) error {
	return i.SetString(string(text))
}

var _StatesValues = []States{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}

// StatesN is the highest valid value
// for type States, plus one.
const StatesN States = 20

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _StatesNoOp() {
	var x [1]struct{}
	_ = x[Disabled-(0)]
	_ = x[ReadOnly-(1)]
	_ = x[Selected-(2)]
	_ = x[Active-(3)]
	_ = x[Dragged-(4)]
	_ = x[Sliding-(5)]
	_ = x[Scrolled-(6)]
	_ = x[Focused-(7)]
	_ = x[FocusedWithin-(8)]
	_ = x[Checked-(9)]
	_ = x[Hovered-(10)]
	_ = x[LongHovered-(11)]
	_ = x[DragHovered-(12)]
	_ = x[DropOK-(13)]
	_ = x[Invalid-(14)]
	_ = x[Required-(15)]
	_ = x[Blank-(16)]
	_ = x[Link-(17)]
	_ = x[Visited-(18)]
	_ = x[AnyLink-(19)]
}

var _StatesNameToValueMap = map[string]States{
	`Disabled`:      0,
	`disabled`:      0,
	`ReadOnly`:      1,
	`readonly`:      1,
	`Selected`:      2,
	`selected`:      2,
	`Active`:        3,
	`active`:        3,
	`Dragged`:       4,
	`dragged`:       4,
	`Sliding`:       5,
	`sliding`:       5,
	`Scrolled`:      6,
	`scrolled`:      6,
	`Focused`:       7,
	`focused`:       7,
	`FocusedWithin`: 8,
	`focusedwithin`: 8,
	`Checked`:       9,
	`checked`:       9,
	`Hovered`:       10,
	`hovered`:       10,
	`LongHovered`:   11,
	`longhovered`:   11,
	`DragHovered`:   12,
	`draghovered`:   12,
	`DropOK`:        13,
	`dropok`:        13,
	`Invalid`:       14,
	`invalid`:       14,
	`Required`:      15,
	`required`:      15,
	`Blank`:         16,
	`blank`:         16,
	`Link`:          17,
	`link`:          17,
	`Visited`:       18,
	`visited`:       18,
	`AnyLink`:       19,
	`anylink`:       19,
}

var _StatesDescMap = map[States]string{
	0:  `Disabled elements cannot be interacted with or selected, but do display.`,
	1:  `ReadOnly elements cannot be changed, but can be selected.`,
	2:  `Selected elements have been marked for clipboard or other such actions.`,
	3:  `Active elements are currently being interacted with, usually involving a mouse button being pressed in the element. A text field will be active while being clicked on, and this can also result in a Focused state. If further movement happens, an element can also end up being Dragged or Sliding.`,
	4:  `Dragged means this element is currently being dragged by the mouse (i.e., a MouseDown event followed by MouseMove), as part of a drag-n-drop sequence.`,
	5:  `Sliding means this element is currently being manipulated via mouse to change the slider state, which will continue until the mouse is released, even if it goes off the element. It should also still be Active.`,
	6:  `Scrolled means this element is currently being scrolled.`,
	7:  `Focused elements receive keyboard input.`,
	8:  `FocusedWithin elements have a Focused element within them, including self.`,
	9:  `Checked is for check boxes or radio buttons or other similar state.`,
	10: `Hovered indicates that a mouse pointer has entered the space over an element, but it is not Active (nor DragHovered).`,
	11: `LongHovered indicates a Hover that persists without significant movement for a minimum period of time (e.g., 500 msec), which typically triggers a tooltip popup.`,
	12: `DragHovered indicates that a mouse pointer has entered the space over an element, during a drag-n-drop sequence. This makes it a candidate for a potential drop target. See DropOK for state in relation to that.`,
	13: `DropOK indicates that a DragHovered element is OK to receive a Drop from the current Dragged item, subject also to the Droppable ability.`,
	14: `Invalid indicates that the element has invalid input and needs to be corrected by the user`,
	15: `Required indicates that the element must be set by the user`,
	16: `Blank indicates that the element has yet to be set by user`,
	17: `Link indicates a URL link that has not been visited yet`,
	18: `Visited indicates a URL link that has been visited`,
	19: `AnyLink is either Link or Visited`,
}

var _StatesMap = map[States]string{
	0:  `Disabled`,
	1:  `ReadOnly`,
	2:  `Selected`,
	3:  `Active`,
	4:  `Dragged`,
	5:  `Sliding`,
	6:  `Scrolled`,
	7:  `Focused`,
	8:  `FocusedWithin`,
	9:  `Checked`,
	10: `Hovered`,
	11: `LongHovered`,
	12: `DragHovered`,
	13: `DropOK`,
	14: `Invalid`,
	15: `Required`,
	16: `Blank`,
	17: `Link`,
	18: `Visited`,
	19: `AnyLink`,
}

// String returns the string representation
// of this States value.
func (i States) String() string {
	str := ""
	for _, ie := range _StatesValues {
		if i.HasFlag(ie) {
			ies := ie.BitIndexString()
			if str == "" {
				str = ies
			} else {
				str += "|" + ies
			}
		}
	}
	return str
}

// BitIndexString returns the string
// representation of this States value
// if it is a bit index value
// (typically an enum constant), and
// not an actual bit flag value.
func (i States) BitIndexString() string {
	if str, ok := _StatesMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the States value from its
// string representation, and returns an
// error if the string is invalid.
func (i *States) SetString(s string) error {
	*i = 0
	return i.SetStringOr(s)
}

// SetStringOr sets the States value from its
// string representation while preserving any
// bit flags already set, and returns an
// error if the string is invalid.
func (i *States) SetStringOr(s string) error {
	flgs := strings.Split(s, "|")
	for _, flg := range flgs {
		if val, ok := _StatesNameToValueMap[flg]; ok {
			i.SetFlag(true, &val)
		} else if val, ok := _StatesNameToValueMap[strings.ToLower(flg)]; ok {
			i.SetFlag(true, &val)
		} else {
			return errors.New(flg + " is not a valid value for type States")
		}
	}
	return nil
}

// Int64 returns the States value as an int64.
func (i States) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the States value from an int64.
func (i *States) SetInt64(in int64) {
	*i = States(in)
}

// Desc returns the description of the States value.
func (i States) Desc() string {
	if str, ok := _StatesDescMap[i]; ok {
		return str
	}
	return i.String()
}

// StatesValues returns all possible values
// for the type States.
func StatesValues() []States {
	return _StatesValues
}

// Values returns all possible values
// for the type States.
func (i States) Values() []enums.Enum {
	res := make([]enums.Enum, len(_StatesValues))
	for i, d := range _StatesValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type States.
func (i States) IsValid() bool {
	_, ok := _StatesMap[i]
	return ok
}

// HasFlag returns whether these
// bit flags have the given bit flag set.
func (i States) HasFlag(f enums.BitFlag) bool {
	return atomic.LoadInt64((*int64)(&i))&(1<<uint32(f.Int64())) != 0
}

// SetFlag sets the value of the given
// flags in these flags to the given value.
func (i *States) SetFlag(on bool, f ...enums.BitFlag) {
	var mask int64
	for _, v := range f {
		mask |= 1 << v.Int64()
	}
	in := int64(*i)
	if on {
		in |= mask
		atomic.StoreInt64((*int64)(i), in)
	} else {
		in &^= mask
		atomic.StoreInt64((*int64)(i), in)
	}
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i States) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *States) UnmarshalText(text []byte) error {
	return i.SetString(string(text))
}
