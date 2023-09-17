// Code generated by "goki generate"; DO NOT EDIT.

package config

import (
	"errors"
	"strconv"
	"strings"

	"goki.dev/enums"
)

var _TypesValues = []Types{0, 1}

// TypesN is the highest valid value
// for type Types, plus one.
const TypesN Types = 2

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _TypesNoOp() {
	var x [1]struct{}
	_ = x[TypeApp-(0)]
	_ = x[TypeLibrary-(1)]
}

var _TypesNameToValueMap = map[string]Types{
	`TypeApp`:     0,
	`typeapp`:     0,
	`TypeLibrary`: 1,
	`typelibrary`: 1,
}

var _TypesDescMap = map[Types]string{
	0: `TypeApp is an executable app`,
	1: `TypeLibrary is an importable library`,
}

var _TypesMap = map[Types]string{
	0: `TypeApp`,
	1: `TypeLibrary`,
}

// String returns the string representation
// of this Types value.
func (i Types) String() string {
	if str, ok := _TypesMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the Types value from its
// string representation, and returns an
// error if the string is invalid.
func (i *Types) SetString(s string) error {
	if val, ok := _TypesNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	if val, ok := _TypesNameToValueMap[strings.ToLower(s)]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type Types")
}

// Int64 returns the Types value as an int64.
func (i Types) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the Types value from an int64.
func (i *Types) SetInt64(in int64) {
	*i = Types(in)
}

// Desc returns the description of the Types value.
func (i Types) Desc() string {
	if str, ok := _TypesDescMap[i]; ok {
		return str
	}
	return i.String()
}

// TypesValues returns all possible values
// for the type Types.
func TypesValues() []Types {
	return _TypesValues
}

// Values returns all possible values
// for the type Types.
func (i Types) Values() []enums.Enum {
	res := make([]enums.Enum, len(_TypesValues))
	for i, d := range _TypesValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type Types.
func (i Types) IsValid() bool {
	_, ok := _TypesMap[i]
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Types) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Types) UnmarshalText(text []byte) error {
	return i.SetString(string(text))
}
