// Code generated by "core generate -add-types"; DO NOT EDIT.

package parse

import (
	"cogentcore.org/core/enums"
)

var _LangFlagsValues = []LanguageFlags{0, 1, 2, 3}

// LangFlagsN is the highest valid value for type LangFlags, plus one.
const LangFlagsN LanguageFlags = 4

var _LangFlagsValueMap = map[string]LanguageFlags{`NoFlags`: 0, `IndentSpace`: 1, `IndentTab`: 2, `ReAutoIndent`: 3}

var _LangFlagsDescMap = map[LanguageFlags]string{0: `NoFlags = nothing special`, 1: `IndentSpace means that spaces must be used for this language`, 2: `IndentTab means that tabs must be used for this language`, 3: `ReAutoIndent causes current line to be re-indented during AutoIndent for Enter (newline) -- this should only be set for strongly indented languages where the previous + current line can tell you exactly what indent the current line should be at.`}

var _LangFlagsMap = map[LanguageFlags]string{0: `NoFlags`, 1: `IndentSpace`, 2: `IndentTab`, 3: `ReAutoIndent`}

// String returns the string representation of this LangFlags value.
func (i LanguageFlags) String() string { return enums.String(i, _LangFlagsMap) }

// SetString sets the LangFlags value from its string representation,
// and returns an error if the string is invalid.
func (i *LanguageFlags) SetString(s string) error {
	return enums.SetString(i, s, _LangFlagsValueMap, "LangFlags")
}

// Int64 returns the LangFlags value as an int64.
func (i LanguageFlags) Int64() int64 { return int64(i) }

// SetInt64 sets the LangFlags value from an int64.
func (i *LanguageFlags) SetInt64(in int64) { *i = LanguageFlags(in) }

// Desc returns the description of the LangFlags value.
func (i LanguageFlags) Desc() string { return enums.Desc(i, _LangFlagsDescMap) }

// LangFlagsValues returns all possible values for the type LangFlags.
func LangFlagsValues() []LanguageFlags { return _LangFlagsValues }

// Values returns all possible values for the type LangFlags.
func (i LanguageFlags) Values() []enums.Enum { return enums.Values(_LangFlagsValues) }

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i LanguageFlags) MarshalText() ([]byte, error) { return []byte(i.String()), nil }

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *LanguageFlags) UnmarshalText(text []byte) error {
	return enums.UnmarshalText(i, text, "LangFlags")
}
