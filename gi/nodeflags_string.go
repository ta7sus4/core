// Code generated by "stringer -type=NodeFlags"; DO NOT EDIT.

package gi

import (
	"errors"
	"strconv"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NoLayout-11]
	_ = x[EventsConnected-12]
	_ = x[CanFocus-13]
	_ = x[HasFocus-14]
	_ = x[FullReRender-15]
	_ = x[ReRenderAnchor-16]
	_ = x[Invisible-17]
	_ = x[Inactive-18]
	_ = x[Selected-19]
	_ = x[MouseHasEntered-20]
	_ = x[DNDHasEntered-21]
	_ = x[NodeDragging-22]
	_ = x[InstaDrag-23]
	_ = x[NodeFlagsN-24]
	_ = x[TextFieldFocusActive-24]
}

const _NodeFlags_name = "NoLayoutEventsConnectedCanFocusHasFocusFullReRenderReRenderAnchorInvisibleInactiveSelectedMouseHasEnteredDNDHasEnteredNodeDraggingInstaDragNodeFlagsN"

var _NodeFlags_index = [...]uint8{0, 8, 23, 31, 39, 51, 65, 74, 82, 90, 105, 118, 130, 139, 149}

func (i NodeFlags) String() string {
	i -= 11
	if i < 0 || i >= NodeFlags(len(_NodeFlags_index)-1) {
		return "NodeFlags(" + strconv.FormatInt(int64(i+11), 10) + ")"
	}
	return _NodeFlags_name[_NodeFlags_index[i]:_NodeFlags_index[i+1]]
}

func StringToNodeFlags(s string) (NodeFlags, error) {
	for i := 0; i < len(_NodeFlags_index)-1; i++ {
		if s == _NodeFlags_name[_NodeFlags_index[i]:_NodeFlags_index[i+1]] {
			return NodeFlags(i + 11), nil
		}
	}
	return 0, errors.New("String: " + s + " is not a valid option for type: NodeFlags")
}

var _NodeFlags_descMap = map[NodeFlags]string{
	11: `NoLayout means that this node does not participate in the layout
process (Size, Layout, Move) -- set by e.g., SVG nodes
`,
	12: `EventsConnected: this node has been connected to receive events from
the window -- to optimize event processing, connections are typically
only established for visible nodes during render, and disconnected when
not visible
`,
	13: `CanFocus: can this node accept focus to receive keyboard input events
-- set by default for typical nodes that do so, but can be overridden,
including by the style 'can-focus' property
`,
	14: `HasFocus: does this node currently have the focus for keyboard input
events?  use tab / alt tab and clicking events to update focus -- see
interface on Window
`,
	15: `FullReRender indicates that a full re-render is required due to nature
of update event -- otherwise default is local re-render -- used
internally for nodes to determine what to do on the ReRender step
`,
	16: `ReRenderAnchor: this node has a static size, and repaints its
background -- any children under it that need to dynamically resize on
a ReRender (Update) can refer the update up to rerendering this node,
instead of going further up the tree -- e.g., true of Frame's within a
SplitView
`,
	17: `Invisible means that the node has been marked as invisible by a parent
that has switch-like powers (e.g., layout stacked / tabview or splitter
panel that has been collapsed).  This flag is propagated down to all
child nodes, and rendering or other interaction / update routines
should not run when this flag is set (PushBounds does this for most
cases).  However, it IS a good idea to have styling, layout etc all
take place as normal, so that when the flag is cleared, rendering can
proceed directly.
`,
	18: `Inactive disables interaction with widgets or other nodes (i.e., they
are read-only) -- they should indicate this inactive state in an
appropriate way, and each node should interpret events appropriately
based on this state (select and context menu events should still be
generated)
`,
	19: `Selected indicates that this node has been selected by the user --
widely supported across different nodes
`,
	20: `MouseHasEntered indicates that the MouseFocusEvent Enter was previously
registered on this node
`,
	21: `DNDHasEntered indicates that the DNDFocusEvent Enter was previously
registered on this node
`,
	22: `NodeDragging indicates this node is currently dragging -- win.Dragging
set to this node
`,
	23: `InstaDrag indicates this node should start dragging immediately when
clicked -- otherwise there is a time-and-distance threshold to the
start of dragging -- use this for controls that are small and are
primarily about dragging (e.g., the Splitter handle)
`,
	24: `can extend node flags from here
`,
}

func (i NodeFlags) Desc() string {
	if str, ok := _NodeFlags_descMap[i]; ok {
		return str
	}
	return "NodeFlags(" + strconv.FormatInt(int64(i), 10) + ")"
}
