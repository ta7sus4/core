// Code generated by "stringer -type=Action"; DO NOT EDIT.

package dnd

import (
	"fmt"
	"strconv"
)

const _Action_name = "NoActionDropOnTargetDropFmSourceMoveEnterExitActionN"

var _Action_index = [...]uint8{0, 8, 20, 32, 36, 41, 45, 52}

func (i Action) String() string {
	if i < 0 || i >= Action(len(_Action_index)-1) {
		return "Action(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Action_name[_Action_index[i]:_Action_index[i+1]]
}

func (i *Action) FromString(s string) error {
	for j := 0; j < len(_Action_index)-1; j++ {
		if s == _Action_name[_Action_index[j]:_Action_index[j+1]] {
			*i = Action(j)
			return nil
		}
	}
	return fmt.Errorf("String %v is not a valid option for type Action", s)
}
