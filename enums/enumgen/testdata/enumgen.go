// Code generated by "enumgen.test -test.testlogfile=/var/folders/x1/r8shprmj7j71zbw3qvgl9dqc0000gq/T/go-build675384445/b647/testlog.txt -test.paniconexit0 -test.timeout=10m0s -test.v=true"; DO NOT EDIT.

package testdata

import (
	"database/sql/driver"
	"io"
	"strconv"

	"cogentcore.org/core/enums"
)

var _FruitsValues = []Fruits{0, 1, 2, 3, 4, 5, 6}

// FruitsN is the highest valid value for type Fruits, plus one.
const FruitsN Fruits = 7

var _FruitsValueMap = map[string]Fruits{`Apple`: 0, `apple`: 0, `Orange`: 1, `orange`: 1, `Peach`: 2, `peach`: 2, `Strawberry`: 3, `strawberry`: 3, `Blackberry`: 4, `blackberry`: 4, `Blueberry`: 5, `blueberry`: 5, `Apricot`: 6, `apricot`: 6}

var _FruitsDescMap = map[Fruits]string{0: ``, 1: ``, 2: ``, 3: ``, 4: ``, 5: ``, 6: ``}

var _FruitsMap = map[Fruits]string{0: `Apple`, 1: `Orange`, 2: `Peach`, 3: `Strawberry`, 4: `Blackberry`, 5: `Blueberry`, 6: `Apricot`}

// String returns the string representation of this Fruits value.
func (i Fruits) String() string { return enums.String(i, _FruitsMap) }

// SetString sets the Fruits value from its string representation,
// and returns an error if the string is invalid.
func (i *Fruits) SetString(s string) error {
	return enums.SetStringLower(i, s, _FruitsValueMap, "Fruits")
}

// Int64 returns the Fruits value as an int64.
func (i Fruits) Int64() int64 { return int64(i) }

// SetInt64 sets the Fruits value from an int64.
func (i *Fruits) SetInt64(in int64) { *i = Fruits(in) }

// Desc returns the description of the Fruits value.
func (i Fruits) Desc() string { return enums.Desc(i, _FruitsDescMap) }

// FruitsValues returns all possible values for the type Fruits.
func FruitsValues() []Fruits { return _FruitsValues }

// Values returns all possible values for the type Fruits.
func (i Fruits) Values() []enums.Enum { return enums.Values(_FruitsValues) }

// IsValid returns whether the value is a valid option for type Fruits.
func (i Fruits) IsValid() bool { _, ok := _FruitsMap[i]; return ok }

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Fruits) MarshalText() ([]byte, error) { return []byte(i.String()), nil }

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Fruits) UnmarshalText(text []byte) error { return enums.UnmarshalText(i, text, "Fruits") }

var _FoodsValues = []Foods{7, 8, 9, 10}

// FoodsN is the highest valid value for type Foods, plus one.
const FoodsN Foods = 11

var _FoodsValueMap = map[string]Foods{`Bread`: 7, `Lettuce`: 8, `Cheese`: 9, `Meat`: 10}

var _FoodsDescMap = map[Foods]string{7: ``, 8: ``, 9: ``, 10: ``}

var _FoodsMap = map[Foods]string{7: `Bread`, 8: `Lettuce`, 9: `Cheese`, 10: `Meat`}

// String returns the string representation of this Foods value.
func (i Foods) String() string { return enums.StringExtended[Foods, Fruits](i, _FoodsMap) }

// SetString sets the Foods value from its string representation,
// and returns an error if the string is invalid.
func (i *Foods) SetString(s string) error {
	return enums.SetStringExtended(i, (*Fruits)(i), s, _FoodsValueMap)
}

// Int64 returns the Foods value as an int64.
func (i Foods) Int64() int64 { return int64(i) }

// SetInt64 sets the Foods value from an int64.
func (i *Foods) SetInt64(in int64) { *i = Foods(in) }

// Desc returns the description of the Foods value.
func (i Foods) Desc() string { return enums.DescExtended[Foods, Fruits](i, _FoodsDescMap) }

// FoodsValues returns all possible values for the type Foods.
func FoodsValues() []Foods { return enums.ValuesGlobalExtended(_FoodsValues, FruitsValues()) }

// Values returns all possible values for the type Foods.
func (i Foods) Values() []enums.Enum { return enums.ValuesExtended(_FoodsValues, FruitsValues()) }

// IsValid returns whether the value is a valid option for type Foods.
func (i Foods) IsValid() bool { _, ok := _FoodsMap[i]; return ok || Fruits(i).IsValid() }

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Foods) MarshalText() ([]byte, error) { return []byte(i.String()), nil }

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Foods) UnmarshalText(text []byte) error { return enums.UnmarshalText(i, text, "Foods") }

var _DaysValues = []Days{-11, -9, -7, -5, -3, -1, 1}

// DaysN is the highest valid value for type Days, plus one.
const DaysN Days = 2

var _DaysValueMap = map[string]Days{`DAY_SATURDAY`: -11, `DAY_FRIDAY`: -9, `DAY_THURSDAY`: -7, `DAY_WEDNESDAY`: -5, `DAY_TUESDAY`: -3, `DAY_MONDAY`: -1, `DAY_SUNDAY`: 1}

var _DaysDescMap = map[Days]string{-11: `Saturday is the seventh day of the week`, -9: `Friday is the sixth day of the week`, -7: `Thursday is the fifth day of the week`, -5: `Wednesday is the fourth day of the week`, -3: `Tuesday is the third day of the week`, -1: `Monday is the second day of the week`, 1: `Sunday is the first day of the week`}

var _DaysMap = map[Days]string{-11: `DAY_SATURDAY`, -9: `DAY_FRIDAY`, -7: `DAY_THURSDAY`, -5: `DAY_WEDNESDAY`, -3: `DAY_TUESDAY`, -1: `DAY_MONDAY`, 1: `DAY_SUNDAY`}

// String returns the string representation of this Days value.
func (i Days) String() string { return enums.String(i, _DaysMap) }

// SetString sets the Days value from its string representation,
// and returns an error if the string is invalid.
func (i *Days) SetString(s string) error { return enums.SetString(i, s, _DaysValueMap, "Days") }

// Int64 returns the Days value as an int64.
func (i Days) Int64() int64 { return int64(i) }

// SetInt64 sets the Days value from an int64.
func (i *Days) SetInt64(in int64) { *i = Days(in) }

// Desc returns the description of the Days value.
func (i Days) Desc() string { return enums.Desc(i, _DaysDescMap) }

// DaysValues returns all possible values for the type Days.
func DaysValues() []Days { return _DaysValues }

// Values returns all possible values for the type Days.
func (i Days) Values() []enums.Enum { return enums.Values(_DaysValues) }

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Days) MarshalText() ([]byte, error) { return []byte(i.String()), nil }

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Days) UnmarshalText(text []byte) error { return enums.UnmarshalText(i, text, "Days") }

// MarshalGQL implements the [graphql.Marshaler] interface.
func (i Days) MarshalGQL(w io.Writer) { w.Write([]byte(strconv.Quote(i.String()))) }

// UnmarshalGQL implements the [graphql.Unmarshaler] interface.
func (i *Days) UnmarshalGQL(value any) error { return enums.Scan(i, value, "Days") }

var _StatesValues = []States{1, 3, 5, 7, 9, 11, 13}

// StatesN is the highest valid value for type States, plus one.
const StatesN States = 14

var _StatesValueMap = map[string]States{`enabled`: 1, `not-enabled`: 3, `focused`: 5, `vered`: 7, `currently-being-pressed-by-user`: 9, `actively-focused`: 11, `selected`: 13}

var _StatesDescMap = map[States]string{1: `Enabled indicates the widget is enabled`, 3: `Disabled indicates the widget is disabled`, 5: `Focused indicates the widget has keyboard focus`, 7: `Hovered indicates the widget is being hovered over`, 9: `Active indicates the widget is being interacted with`, 11: `ActivelyFocused indicates the widget has active keyboard focus`, 13: `Selected indicates the widget is selected`}

var _StatesMap = map[States]string{1: `enabled`, 3: `not-enabled`, 5: `focused`, 7: `vered`, 9: `currently-being-pressed-by-user`, 11: `actively-focused`, 13: `selected`}

// String returns the string representation of this States value.
func (i States) String() string { return enums.BitFlagString(i, _StatesValues) }

// BitIndexString returns the string representation of this States value
// if it is a bit index value (typically an enum constant), and
// not an actual bit flag value.
func (i States) BitIndexString() string { return enums.String(i, _StatesMap) }

// SetString sets the States value from its string representation,
// and returns an error if the string is invalid.
func (i *States) SetString(s string) error { *i = 0; return i.SetStringOr(s) }

// SetStringOr sets the States value from its string representation
// while preserving any bit flags already set, and returns an
// error if the string is invalid.
func (i *States) SetStringOr(s string) error {
	return enums.SetStringOr(i, s, _StatesValueMap, "States")
}

// Int64 returns the States value as an int64.
func (i States) Int64() int64 { return int64(i) }

// SetInt64 sets the States value from an int64.
func (i *States) SetInt64(in int64) { *i = States(in) }

// Desc returns the description of the States value.
func (i States) Desc() string { return enums.Desc(i, _StatesDescMap) }

// StatesValues returns all possible values for the type States.
func StatesValues() []States { return _StatesValues }

// Values returns all possible values for the type States.
func (i States) Values() []enums.Enum { return enums.Values(_StatesValues) }

// HasFlag returns whether these bit flags have the given bit flag set.
func (i *States) HasFlag(f enums.BitFlag) bool { return enums.HasFlag((*int64)(i), f) }

// SetFlag sets the value of the given flags in these flags to the given value.
func (i *States) SetFlag(on bool, f ...enums.BitFlag) { enums.SetFlag((*int64)(i), on, f...) }

// Value implements the [driver.Valuer] interface.
func (i States) Value() (driver.Value, error) { return i.String(), nil }

// Scan implements the [sql.Scanner] interface.
func (i *States) Scan(value any) error { return enums.Scan(i, value, "States") }

var _LanguagesValues = []Languages{6, 10, 14, 18, 22, 26, 30, 34, 38, 42, 46, 50, 54}

// LanguagesN is the highest valid value for type Languages, plus one.
const LanguagesN Languages = 55

var _LanguagesValueMap = map[string]Languages{`Go`: 6, `Python`: 10, `JavaScript`: 14, `Dart`: 18, `Rust`: 22, `Ruby`: 26, `C`: 30, `CPP`: 34, `ObjectiveC`: 38, `Java`: 42, `TypeScript`: 46, `Kotlin`: 50, `Swift`: 54}

var _LanguagesDescMap = map[Languages]string{6: `Go is the best programming language`, 10: ``, 14: `JavaScript is the worst programming language`, 18: ``, 22: ``, 26: ``, 30: ``, 34: ``, 38: ``, 42: ``, 46: ``, 50: ``, 54: ``}

var _LanguagesMap = map[Languages]string{6: `Go`, 10: `Python`, 14: `JavaScript`, 18: `Dart`, 22: `Rust`, 26: `Ruby`, 30: `C`, 34: `CPP`, 38: `ObjectiveC`, 42: `Java`, 46: `TypeScript`, 50: `Kotlin`, 54: `Swift`}

// String returns the string representation of this Languages value.
func (i Languages) String() string { return enums.BitFlagString(i, _LanguagesValues) }

// BitIndexString returns the string representation of this Languages value
// if it is a bit index value (typically an enum constant), and
// not an actual bit flag value.
func (i Languages) BitIndexString() string { return enums.String(i, _LanguagesMap) }

// SetString sets the Languages value from its string representation,
// and returns an error if the string is invalid.
func (i *Languages) SetString(s string) error { *i = 0; return i.SetStringOr(s) }

// SetStringOr sets the Languages value from its string representation
// while preserving any bit flags already set, and returns an
// error if the string is invalid.
func (i *Languages) SetStringOr(s string) error {
	return enums.SetStringOr(i, s, _LanguagesValueMap, "Languages")
}

// Int64 returns the Languages value as an int64.
func (i Languages) Int64() int64 { return int64(i) }

// SetInt64 sets the Languages value from an int64.
func (i *Languages) SetInt64(in int64) { *i = Languages(in) }

// Desc returns the description of the Languages value.
func (i Languages) Desc() string { return enums.Desc(i, _LanguagesDescMap) }

// LanguagesValues returns all possible values for the type Languages.
func LanguagesValues() []Languages { return _LanguagesValues }

// Values returns all possible values for the type Languages.
func (i Languages) Values() []enums.Enum { return enums.Values(_LanguagesValues) }

// HasFlag returns whether these bit flags have the given bit flag set.
func (i *Languages) HasFlag(f enums.BitFlag) bool { return enums.HasFlag((*int64)(i), f) }

// SetFlag sets the value of the given flags in these flags to the given value.
func (i *Languages) SetFlag(on bool, f ...enums.BitFlag) { enums.SetFlag((*int64)(i), on, f...) }

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Languages) MarshalText() ([]byte, error) { return []byte(i.String()), nil }

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Languages) UnmarshalText(text []byte) error {
	return enums.UnmarshalText(i, text, "Languages")
}

var _MoreLanguagesValues = []MoreLanguages{55}

// MoreLanguagesN is the highest valid value for type MoreLanguages, plus one.
const MoreLanguagesN MoreLanguages = 56

var _MoreLanguagesValueMap = map[string]MoreLanguages{`Perl`: 55}

var _MoreLanguagesDescMap = map[MoreLanguages]string{55: ``}

var _MoreLanguagesMap = map[MoreLanguages]string{55: `Perl`}

// String returns the string representation of this MoreLanguages value.
func (i MoreLanguages) String() string {
	return enums.BitFlagStringExtended(i, _MoreLanguagesValues, LanguagesValues())
}

// BitIndexString returns the string representation of this MoreLanguages value
// if it is a bit index value (typically an enum constant), and
// not an actual bit flag value.
func (i MoreLanguages) BitIndexString() string {
	return enums.BitIndexStringExtended[MoreLanguages, Languages](i, _MoreLanguagesMap)
}

// SetString sets the MoreLanguages value from its string representation,
// and returns an error if the string is invalid.
func (i *MoreLanguages) SetString(s string) error { *i = 0; return i.SetStringOr(s) }

// SetStringOr sets the MoreLanguages value from its string representation
// while preserving any bit flags already set, and returns an
// error if the string is invalid.
func (i *MoreLanguages) SetStringOr(s string) error {
	return enums.SetStringOrExtended(i, (*Languages)(i), s, _MoreLanguagesValueMap)
}

// Int64 returns the MoreLanguages value as an int64.
func (i MoreLanguages) Int64() int64 { return int64(i) }

// SetInt64 sets the MoreLanguages value from an int64.
func (i *MoreLanguages) SetInt64(in int64) { *i = MoreLanguages(in) }

// Desc returns the description of the MoreLanguages value.
func (i MoreLanguages) Desc() string {
	return enums.DescExtended[MoreLanguages, Languages](i, _MoreLanguagesDescMap)
}

// MoreLanguagesValues returns all possible values for the type MoreLanguages.
func MoreLanguagesValues() []MoreLanguages {
	return enums.ValuesGlobalExtended(_MoreLanguagesValues, LanguagesValues())
}

// Values returns all possible values for the type MoreLanguages.
func (i MoreLanguages) Values() []enums.Enum {
	return enums.ValuesExtended(_MoreLanguagesValues, LanguagesValues())
}

// HasFlag returns whether these bit flags have the given bit flag set.
func (i *MoreLanguages) HasFlag(f enums.BitFlag) bool { return enums.HasFlag((*int64)(i), f) }

// SetFlag sets the value of the given flags in these flags to the given value.
func (i *MoreLanguages) SetFlag(on bool, f ...enums.BitFlag) { enums.SetFlag((*int64)(i), on, f...) }

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i MoreLanguages) MarshalText() ([]byte, error) { return []byte(i.String()), nil }

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *MoreLanguages) UnmarshalText(text []byte) error {
	return enums.UnmarshalText(i, text, "MoreLanguages")
}
