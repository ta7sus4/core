// Code generated by "stringer -type=TextFieldSignals"; DO NOT EDIT.

package gi

import (
	"errors"
	"strconv"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TextFieldDone-0]
	_ = x[TextFieldDeFocused-1]
	_ = x[TextFieldSelected-2]
	_ = x[TextFieldCleared-3]
	_ = x[TextFieldInsert-4]
	_ = x[TextFieldBackspace-5]
	_ = x[TextFieldDelete-6]
	_ = x[TextFieldSignalsN-7]
}

const _TextFieldSignals_name = "TextFieldDoneTextFieldDeFocusedTextFieldSelectedTextFieldClearedTextFieldInsertTextFieldBackspaceTextFieldDeleteTextFieldSignalsN"

var _TextFieldSignals_index = [...]uint8{0, 13, 31, 48, 64, 79, 97, 112, 129}

func (i TextFieldSignals) String() string {
	if i < 0 || i >= TextFieldSignals(len(_TextFieldSignals_index)-1) {
		return "TextFieldSignals(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TextFieldSignals_name[_TextFieldSignals_index[i]:_TextFieldSignals_index[i+1]]
}

func (i *TextFieldSignals) FromString(s string) error {
	for j := 0; j < len(_TextFieldSignals_index)-1; j++ {
		if s == _TextFieldSignals_name[_TextFieldSignals_index[j]:_TextFieldSignals_index[j+1]] {
			*i = TextFieldSignals(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: TextFieldSignals")
}

var _TextFieldSignals_descMap = map[TextFieldSignals]string{
	0: `TextFieldDone is main signal -- return or tab was pressed and the edit was
intentionally completed.  data is the text.
`,
	1: `TextFieldDeFocused means that the user has transitioned focus away from
the text field due to interactions elsewhere, and any ongoing changes have been
applied and the editor is no longer active.  data is the text.
If you have a button that performs the same action as pressing enter in a textfield,
then pressing that button will trigger a TextFieldDeFocused event, for any active
edits.  Otherwise, you probably want to respond to both TextFieldDone and
TextFieldDeFocused as "apply" events that trigger actions associated with the field.
`,
	2: `TextFieldSelected means that some text was selected (for Inactive state,
selection is via WidgetSig)
`,
	3: `TextFieldCleared means the clear button was clicked
`,
	4: `TextFieldInsert is emitted when a character is inserted into the textfield
`,
	5: `TextFieldBackspace is emitted when a character before cursor is deleted
`,
	6: `TextFieldDelete is emitted when a character after cursor is deleted
`,
	7: ``,
}

func (i TextFieldSignals) Desc() string {
	if str, ok := _TextFieldSignals_descMap[i]; ok {
		return str
	}
	return "TextFieldSignals(" + strconv.FormatInt(int64(i), 10) + ")"
}
