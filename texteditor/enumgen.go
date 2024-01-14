// Code generated by "goki generate"; DO NOT EDIT.

package texteditor

import (
	"errors"
	"log"
	"strconv"
	"strings"
	"sync/atomic"

	"goki.dev/enums"
	"goki.dev/gi"
)

var _BufSignalsValues = []BufSignals{0, 1, 2, 3, 4, 5, 6}

// BufSignalsN is the highest valid value
// for type BufSignals, plus one.
const BufSignalsN BufSignals = 7

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _BufSignalsNoOp() {
	var x [1]struct{}
	_ = x[BufDone-(0)]
	_ = x[BufNew-(1)]
	_ = x[BufMods-(2)]
	_ = x[BufInsert-(3)]
	_ = x[BufDelete-(4)]
	_ = x[BufMarkUpdt-(5)]
	_ = x[BufClosed-(6)]
}

var _BufSignalsNameToValueMap = map[string]BufSignals{
	`BufDone`:     0,
	`bufdone`:     0,
	`BufNew`:      1,
	`bufnew`:      1,
	`BufMods`:     2,
	`bufmods`:     2,
	`BufInsert`:   3,
	`bufinsert`:   3,
	`BufDelete`:   4,
	`bufdelete`:   4,
	`BufMarkUpdt`: 5,
	`bufmarkupdt`: 5,
	`BufClosed`:   6,
	`bufclosed`:   6,
}

var _BufSignalsDescMap = map[BufSignals]string{
	0: `BufDone means that editing was completed and applied to Txt field -- data is Txt bytes`,
	1: `BufNew signals that entirely new text is present. All views should do full layout update.`,
	2: `BufMods signals that potentially diffuse modifications have been made. Views should do a Layout and Render.`,
	3: `BufInsert signals that some text was inserted. data is textbuf.Edit describing change. The Buf always reflects the current state *after* the edit.`,
	4: `BufDelete signals that some text was deleted. data is textbuf.Edit describing change. The Buf always reflects the current state *after* the edit.`,
	5: `BufMarkUpdt signals that the Markup text has been updated This signal is typically sent from a separate goroutine, so should be used with a mutex`,
	6: `BufClosed signals that the textbuf was closed.`,
}

var _BufSignalsMap = map[BufSignals]string{
	0: `BufDone`,
	1: `BufNew`,
	2: `BufMods`,
	3: `BufInsert`,
	4: `BufDelete`,
	5: `BufMarkUpdt`,
	6: `BufClosed`,
}

// String returns the string representation
// of this BufSignals value.
func (i BufSignals) String() string {
	if str, ok := _BufSignalsMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the BufSignals value from its
// string representation, and returns an
// error if the string is invalid.
func (i *BufSignals) SetString(s string) error {
	if val, ok := _BufSignalsNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	if val, ok := _BufSignalsNameToValueMap[strings.ToLower(s)]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type BufSignals")
}

// Int64 returns the BufSignals value as an int64.
func (i BufSignals) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the BufSignals value from an int64.
func (i *BufSignals) SetInt64(in int64) {
	*i = BufSignals(in)
}

// Desc returns the description of the BufSignals value.
func (i BufSignals) Desc() string {
	if str, ok := _BufSignalsDescMap[i]; ok {
		return str
	}
	return i.String()
}

// BufSignalsValues returns all possible values
// for the type BufSignals.
func BufSignalsValues() []BufSignals {
	return _BufSignalsValues
}

// Values returns all possible values
// for the type BufSignals.
func (i BufSignals) Values() []enums.Enum {
	res := make([]enums.Enum, len(_BufSignalsValues))
	for i, d := range _BufSignalsValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type BufSignals.
func (i BufSignals) IsValid() bool {
	_, ok := _BufSignalsMap[i]
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i BufSignals) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *BufSignals) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println("BufSignals.UnmarshalText:", err)
	}
	return nil
}

var _BufFlagsValues = []BufFlags{8, 9, 10, 11, 12}

// BufFlagsN is the highest valid value
// for type BufFlags, plus one.
const BufFlagsN BufFlags = 13

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _BufFlagsNoOp() {
	var x [1]struct{}
	_ = x[BufAutoSaving-(8)]
	_ = x[BufMarkingUp-(9)]
	_ = x[BufChanged-(10)]
	_ = x[BufNotSaved-(11)]
	_ = x[BufFileModOk-(12)]
}

var _BufFlagsNameToValueMap = map[string]BufFlags{
	`AutoSaving`: 8,
	`autosaving`: 8,
	`MarkingUp`:  9,
	`markingup`:  9,
	`Changed`:    10,
	`changed`:    10,
	`NotSaved`:   11,
	`notsaved`:   11,
	`FileModOk`:  12,
	`filemodok`:  12,
}

var _BufFlagsDescMap = map[BufFlags]string{
	8:  `BufAutoSaving is used in atomically safe way to protect autosaving`,
	9:  `BufMarkingUp indicates current markup operation in progress -- don&#39;t redo`,
	10: `BufChanged indicates if the text has been changed (edited) relative to the original, since last EditDone`,
	11: `BufNotSaved indicates if the text has been changed (edited) relative to the original, since last Save`,
	12: `BufFileModOk have already asked about fact that file has changed since being opened, user is ok`,
}

var _BufFlagsMap = map[BufFlags]string{
	8:  `AutoSaving`,
	9:  `MarkingUp`,
	10: `Changed`,
	11: `NotSaved`,
	12: `FileModOk`,
}

// String returns the string representation
// of this BufFlags value.
func (i BufFlags) String() string {
	str := ""
	for _, ie := range gi.WidgetFlagsValues() {
		if i.HasFlag(ie) {
			ies := ie.BitIndexString()
			if str == "" {
				str = ies
			} else {
				str += "|" + ies
			}
		}
	}
	for _, ie := range _BufFlagsValues {
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
// representation of this BufFlags value
// if it is a bit index value
// (typically an enum constant), and
// not an actual bit flag value.
func (i BufFlags) BitIndexString() string {
	if str, ok := _BufFlagsMap[i]; ok {
		return str
	}
	return gi.WidgetFlags(i).BitIndexString()
}

// SetString sets the BufFlags value from its
// string representation, and returns an
// error if the string is invalid.
func (i *BufFlags) SetString(s string) error {
	*i = 0
	return i.SetStringOr(s)
}

// SetStringOr sets the BufFlags value from its
// string representation while preserving any
// bit flags already set, and returns an
// error if the string is invalid.
func (i *BufFlags) SetStringOr(s string) error {
	flgs := strings.Split(s, "|")
	for _, flg := range flgs {
		if val, ok := _BufFlagsNameToValueMap[flg]; ok {
			i.SetFlag(true, &val)
		} else if val, ok := _BufFlagsNameToValueMap[strings.ToLower(flg)]; ok {
			i.SetFlag(true, &val)
		} else if flg == "" {
			continue
		} else {
			err := (*gi.WidgetFlags)(i).SetStringOr(flg)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Int64 returns the BufFlags value as an int64.
func (i BufFlags) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the BufFlags value from an int64.
func (i *BufFlags) SetInt64(in int64) {
	*i = BufFlags(in)
}

// Desc returns the description of the BufFlags value.
func (i BufFlags) Desc() string {
	if str, ok := _BufFlagsDescMap[i]; ok {
		return str
	}
	return gi.WidgetFlags(i).Desc()
}

// BufFlagsValues returns all possible values
// for the type BufFlags.
func BufFlagsValues() []BufFlags {
	es := gi.WidgetFlagsValues()
	res := make([]BufFlags, len(es))
	for i, e := range es {
		res[i] = BufFlags(e)
	}
	res = append(res, _BufFlagsValues...)
	return res
}

// Values returns all possible values
// for the type BufFlags.
func (i BufFlags) Values() []enums.Enum {
	es := gi.WidgetFlagsValues()
	les := len(es)
	res := make([]enums.Enum, les+len(_BufFlagsValues))
	for i, d := range es {
		res[i] = d
	}
	for i, d := range _BufFlagsValues {
		res[i+les] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type BufFlags.
func (i BufFlags) IsValid() bool {
	_, ok := _BufFlagsMap[i]
	if !ok {
		return gi.WidgetFlags(i).IsValid()
	}
	return ok
}

// HasFlag returns whether these
// bit flags have the given bit flag set.
func (i BufFlags) HasFlag(f enums.BitFlag) bool {
	return atomic.LoadInt64((*int64)(&i))&(1<<uint32(f.Int64())) != 0
}

// SetFlag sets the value of the given
// flags in these flags to the given value.
func (i *BufFlags) SetFlag(on bool, f ...enums.BitFlag) {
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
func (i BufFlags) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *BufFlags) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println("BufFlags.UnmarshalText:", err)
	}
	return nil
}

var _EditorFlagsValues = []EditorFlags{8, 9, 10, 11, 12}

// EditorFlagsN is the highest valid value
// for type EditorFlags, plus one.
const EditorFlagsN EditorFlags = 13

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _EditorFlagsNoOp() {
	var x [1]struct{}
	_ = x[EditorHasLineNos-(8)]
	_ = x[EditorNeedsLayout-(9)]
	_ = x[EditorLastWasTabAI-(10)]
	_ = x[EditorLastWasUndo-(11)]
	_ = x[EditorTargetSet-(12)]
}

var _EditorFlagsNameToValueMap = map[string]EditorFlags{
	`EditorHasLineNos`:   8,
	`editorhaslinenos`:   8,
	`EditorNeedsLayout`:  9,
	`editorneedslayout`:  9,
	`EditorLastWasTabAI`: 10,
	`editorlastwastabai`: 10,
	`EditorLastWasUndo`:  11,
	`editorlastwasundo`:  11,
	`EditorTargetSet`:    12,
	`editortargetset`:    12,
}

var _EditorFlagsDescMap = map[EditorFlags]string{
	8:  `EditorHasLineNos indicates that this editor has line numbers (per Buf option)`,
	9:  `EditorNeedsLayout is set by SetNeedsLayout: Editor does significant internal layout in LayoutAllLines, and its layout is simply based on what it gets allocated, so it does not affect the rest of the Scene.`,
	10: `EditorLastWasTabAI indicates that last key was a Tab auto-indent`,
	11: `EditorLastWasUndo indicates that last key was an undo`,
	12: `EditorTargetSet indicates that the CursorTarget is set`,
}

var _EditorFlagsMap = map[EditorFlags]string{
	8:  `EditorHasLineNos`,
	9:  `EditorNeedsLayout`,
	10: `EditorLastWasTabAI`,
	11: `EditorLastWasUndo`,
	12: `EditorTargetSet`,
}

// String returns the string representation
// of this EditorFlags value.
func (i EditorFlags) String() string {
	str := ""
	for _, ie := range gi.WidgetFlagsValues() {
		if i.HasFlag(ie) {
			ies := ie.BitIndexString()
			if str == "" {
				str = ies
			} else {
				str += "|" + ies
			}
		}
	}
	for _, ie := range _EditorFlagsValues {
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
// representation of this EditorFlags value
// if it is a bit index value
// (typically an enum constant), and
// not an actual bit flag value.
func (i EditorFlags) BitIndexString() string {
	if str, ok := _EditorFlagsMap[i]; ok {
		return str
	}
	return gi.WidgetFlags(i).BitIndexString()
}

// SetString sets the EditorFlags value from its
// string representation, and returns an
// error if the string is invalid.
func (i *EditorFlags) SetString(s string) error {
	*i = 0
	return i.SetStringOr(s)
}

// SetStringOr sets the EditorFlags value from its
// string representation while preserving any
// bit flags already set, and returns an
// error if the string is invalid.
func (i *EditorFlags) SetStringOr(s string) error {
	flgs := strings.Split(s, "|")
	for _, flg := range flgs {
		if val, ok := _EditorFlagsNameToValueMap[flg]; ok {
			i.SetFlag(true, &val)
		} else if val, ok := _EditorFlagsNameToValueMap[strings.ToLower(flg)]; ok {
			i.SetFlag(true, &val)
		} else if flg == "" {
			continue
		} else {
			err := (*gi.WidgetFlags)(i).SetStringOr(flg)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Int64 returns the EditorFlags value as an int64.
func (i EditorFlags) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the EditorFlags value from an int64.
func (i *EditorFlags) SetInt64(in int64) {
	*i = EditorFlags(in)
}

// Desc returns the description of the EditorFlags value.
func (i EditorFlags) Desc() string {
	if str, ok := _EditorFlagsDescMap[i]; ok {
		return str
	}
	return gi.WidgetFlags(i).Desc()
}

// EditorFlagsValues returns all possible values
// for the type EditorFlags.
func EditorFlagsValues() []EditorFlags {
	es := gi.WidgetFlagsValues()
	res := make([]EditorFlags, len(es))
	for i, e := range es {
		res[i] = EditorFlags(e)
	}
	res = append(res, _EditorFlagsValues...)
	return res
}

// Values returns all possible values
// for the type EditorFlags.
func (i EditorFlags) Values() []enums.Enum {
	es := gi.WidgetFlagsValues()
	les := len(es)
	res := make([]enums.Enum, les+len(_EditorFlagsValues))
	for i, d := range es {
		res[i] = d
	}
	for i, d := range _EditorFlagsValues {
		res[i+les] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type EditorFlags.
func (i EditorFlags) IsValid() bool {
	_, ok := _EditorFlagsMap[i]
	if !ok {
		return gi.WidgetFlags(i).IsValid()
	}
	return ok
}

// HasFlag returns whether these
// bit flags have the given bit flag set.
func (i EditorFlags) HasFlag(f enums.BitFlag) bool {
	return atomic.LoadInt64((*int64)(&i))&(1<<uint32(f.Int64())) != 0
}

// SetFlag sets the value of the given
// flags in these flags to the given value.
func (i *EditorFlags) SetFlag(on bool, f ...enums.BitFlag) {
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
func (i EditorFlags) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *EditorFlags) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println("EditorFlags.UnmarshalText:", err)
	}
	return nil
}
