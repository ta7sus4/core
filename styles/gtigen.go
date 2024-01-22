// Code generated by "core generate"; DO NOT EDIT.

package styles

import (
	"cogentcore.org/core/gti"
)

var _ = gti.AddType(&gti.Type{Name: "cogentcore.org/core/styles.Border", IDName: "border", Doc: "Border contains style parameters for borders", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Fields: []gti.Field{{Name: "Style", Doc: "how to draw the border"}, {Name: "Width", Doc: "width of the border"}, {Name: "Radius", Doc: "radius (rounding) of the corners"}, {Name: "Color", Doc: "color of the border"}}})

var _ = gti.AddType(&gti.Type{Name: "cogentcore.org/core/styles.Shadow", IDName: "shadow", Doc: "style parameters for shadows", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Fields: []gti.Field{{Name: "HOffset", Doc: "horizontal offset of shadow; positive = right side, negative = left side"}, {Name: "VOffset", Doc: "vertical offset of shadow; positive = below, negative = above"}, {Name: "Blur", Doc: "blur radius; higher numbers = more blurry"}, {Name: "Spread", Doc: "spread radius; positive number increases size of shadow, negative decreases size"}, {Name: "Color", Doc: "color of the shadow"}, {Name: "Inset", Doc: "if true, shadow is inset within box instead of outset outside of box;\nTODO: implement"}}})

var _ = gti.AddType(&gti.Type{Name: "cogentcore.org/core/styles.Font", IDName: "font", Doc: "Font contains all font styling information.\nMost of font information is inherited.\nFont does not include all information needed\nfor rendering -- see [FontRender] for that.", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Fields: []gti.Field{{Name: "Size", Doc: "size of font to render (inhereted); converted to points when getting font to use"}, {Name: "Family", Doc: "font family (inhereted): ordered list of comma-separated names from more general to more specific to use; use split on , to parse"}, {Name: "Style", Doc: "style (inhereted): normal, italic, etc"}, {Name: "Weight", Doc: "weight (inhereted): normal, bold, etc"}, {Name: "Stretch", Doc: "font stretch / condense options (inhereted)"}, {Name: "Variant", Doc: "normal or small caps (inhereted)"}, {Name: "Deco", Doc: "underline, line-through, etc (not inherited)"}, {Name: "Shift", Doc: "super / sub script (not inherited)"}, {Name: "Face", Doc: "full font information including enhanced metrics and actual font codes for drawing text; this is a pointer into FontLibrary of loaded fonts"}}})

var _ = gti.AddType(&gti.Type{Name: "cogentcore.org/core/styles.FontRender", IDName: "font-render", Doc: "FontRender contains all font styling information\nthat is needed for SVG text rendering. It is passed to\nPaint and Style functions. It should typically not be\nused by end-user code -- see [Font] for that.\nIt stores all values as pointers so that they correspond\nto the values of the style object it was derived from.", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Embeds: []gti.Field{{Name: "Font"}}, Fields: []gti.Field{{Name: "Color", Doc: "text color (inhereted)"}, {Name: "Background", Doc: "background color (not inherited, transparent by default)"}, {Name: "Opacity", Doc: "alpha value between 0 and 1 to apply to the foreground and background of this element and all of its children"}}})

var _ = gti.AddType(&gti.Type{Name: "cogentcore.org/core/styles.FontFace", IDName: "font-face", Doc: "FontFace is our enhanced Font Face structure which contains the enhanced computed\nmetrics in addition to the font.Face face", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Fields: []gti.Field{{Name: "Name", Doc: "The full FaceName that the font is accessed by"}, {Name: "Size", Doc: "The integer font size in raw dots"}, {Name: "Face", Doc: "The system image.Font font rendering interface"}, {Name: "Metrics", Doc: "enhanced metric information for the font"}}})

var _ = gti.AddType(&gti.Type{Name: "cogentcore.org/core/styles.FontMetrics", IDName: "font-metrics", Doc: "FontMetrics are our enhanced dot-scale font metrics compared to what is available in\nthe standard font.Metrics lib, including Ex and Ch being defined in terms of\nthe actual letter x and 0", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Fields: []gti.Field{{Name: "Height", Doc: "reference 1.0 spacing line height of font in dots -- computed from font as ascent + descent + lineGap, where lineGap is specified by the font as the recommended line spacing"}, {Name: "Em", Doc: "Em size of font -- this is NOT actually the width of the letter M, but rather the specified point size of the font (in actual display dots, not points) -- it does NOT include the descender and will not fit the entire height of the font"}, {Name: "Ex", Doc: "Ex size of font -- this is the actual height of the letter x in the font"}, {Name: "Ch", Doc: "Ch size of font -- this is the actual width of the 0 glyph in the font"}}})

var _ = gti.AddType(&gti.Type{Name: "cogentcore.org/core/styles.AlignSet", IDName: "align-set", Doc: "AlignSet specifies the 3 levels of Justify or Align: Content, Items, and Self", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Fields: []gti.Field{{Name: "Content", Doc: "Content specifies the distribution of the entire collection of items within\nany larger amount of space allocated to the container.  By contrast, Items\nand Self specify distribution within the individual element's allocated space."}, {Name: "Items", Doc: "Items specifies the distribution within the individual element's allocated space,\nas a default for all items within a collection."}, {Name: "Self", Doc: "Self specifies the distribution within the individual element's allocated space,\nfor this specific item.  Auto defaults to containers Items setting."}}})

var _ = gti.AddType(&gti.Type{Name: "cogentcore.org/core/styles.Paint", IDName: "paint", Doc: "Paint provides the styling parameters for SVG-style rendering", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Fields: []gti.Field{{Name: "Off", Doc: "prop: display:none -- node and everything below it are off, non-rendering"}, {Name: "Display", Doc: "todo big enum of how to display item -- controls layout etc"}, {Name: "StrokeStyle", Doc: "stroke (line drawing) parameters"}, {Name: "FillStyle", Doc: "fill (region filling) parameters"}, {Name: "FontStyle", Doc: "font also has global opacity setting, along with generic color, background-color settings, which can be copied into stroke / fill as needed"}, {Name: "TextStyle", Doc: "font also has global opacity setting, along with generic color, background-color settings, which can be copied into stroke / fill as needed"}, {Name: "VecEff", Doc: "various rendering special effects settings"}, {Name: "Transform", Doc: "our additions to transform -- pushed to render state"}, {Name: "UnContext", Doc: "units context -- parameters necessary for anchoring relative units"}, {Name: "StyleSet", Doc: "have the styles already been set?"}, {Name: "PropsNil"}, {Name: "dotsSet"}, {Name: "lastUnCtxt"}}})

var _ = gti.AddType(&gti.Type{Name: "cogentcore.org/core/styles.Sides", IDName: "sides", Doc: "Sides contains values for each side or corner of a box.\nIf Sides contains sides, the struct field names correspond\ndirectly to the side values (ie: Top = top side value).\nIf Sides contains corners, the struct field names correspond\nto the corners as follows: Top = top left, Right = top right,\nBottom = bottom right, Left = bottom left.", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Fields: []gti.Field{{Name: "Top", Doc: "top/top-left value"}, {Name: "Right", Doc: "right/top-right value"}, {Name: "Bottom", Doc: "bottom/bottom-right value"}, {Name: "Left", Doc: "left/bottom-left value"}}})

var _ = gti.AddType(&gti.Type{Name: "cogentcore.org/core/styles.SideValues", IDName: "side-values", Doc: "SideValues contains units.Value values for each side/corner of a box", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Embeds: []gti.Field{{Name: "Sides"}}})

var _ = gti.AddType(&gti.Type{Name: "cogentcore.org/core/styles.SideFloats", IDName: "side-floats", Doc: "SideFloats contains float32 values for each side/corner of a box", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Embeds: []gti.Field{{Name: "Sides"}}})

var _ = gti.AddType(&gti.Type{Name: "cogentcore.org/core/styles.SideColors", IDName: "side-colors", Doc: "SideColors contains color values for each side/corner of a box", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Embeds: []gti.Field{{Name: "Sides"}}})

var _ = gti.AddType(&gti.Type{Name: "cogentcore.org/core/styles.Style", IDName: "style", Doc: "Style has all the CSS-based style elements -- used for widget-type GUI objects.", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Fields: []gti.Field{{Name: "State", Doc: "State holds style-relevant state flags, for convenient styling access,\ngiven that styles typically depend on element states."}, {Name: "Abilities", Doc: "Abilities specifies the abilities of this element, which determine\nwhich kinds of states the element can express.\nThis is used by the goosi/events system.  Putting this info next\nto the State info makes it easy to configure and manage."}, {Name: "Cursor", Doc: "the cursor to switch to upon hovering over the element (inherited)"}, {Name: "Padding", Doc: "Padding is the transparent space around central content of box,\nwhich is _included_ in the size of the standard box rendering."}, {Name: "Margin", Doc: "Margin is the outer-most transparent space around box element,\nwhich is _excluded_ from standard box rendering."}, {Name: "Display", Doc: "Display controls how items are displayed, in terms of layout"}, {Name: "Direction", Doc: "Direction specifies the way in which elements are laid out, or\nthe dimension on which an element is longer / travels in."}, {Name: "Wrap", Doc: "Wrap causes elements to wrap around in the CrossAxis dimension\nto fit within sizing constraints (on by default)."}, {Name: "Justify", Doc: "Justify specifies the distribution of elements along the main axis,\ni.e., the same as Direction, for Flex Display.  For Grid, the main axis is\ngiven by the writing direction (e.g., Row-wise for latin based languages)."}, {Name: "Align", Doc: "Align specifies the cross-axis alignment of elements, orthogonal to the\nmain Direction axis. For Grid, the cross-axis is orthogonal to the\nwriting direction (e.g., Column-wise for latin based languages)."}, {Name: "Min", Doc: "Min is the minimum size of the actual content, exclusive of additional space\nfrom padding, border, margin; 0 = default is sum of Min for all content\n(which _includes_ space for all sub-elements).\nThis is equivalent to the Basis for the CSS flex styling model."}, {Name: "Max", Doc: "Max is the maximum size of the actual content, exclusive of additional space\nfrom padding, border, margin; 0 = default provides no Max size constraint"}, {Name: "Grow", Doc: "Grow is the proportional amount that the element can grow (stretch)\nif there is more space available.  0 = default = no growth.\nExtra available space is allocated as: Grow / sum (all Grow).\nImportant: grow elements absorb available space and thus are not\nsubject to alignment (Center, End)."}, {Name: "GrowWrap", Doc: "GrowWrap is a special case for Text elements where it grows initially\nin the horizontal axis to allow for longer, word wrapped text to fill\nthe available space, but then it does not grow thereafter, so that alignment\noperations still work (Grow elements do not align because they absorb all\navailable space)."}, {Name: "FillMargin", Doc: "FillMargin determines is whether to fill the margin with\nthe surrounding background color before rendering the element itself.\nThis is typically necessary to prevent text, border, and box shadow from rendering\nover themselves. It should be kept at its default value of true\nin most circumstances, but it can be set to false when the element\nis fully managed by something that is guaranteed to render the\nappropriate background color for the element."}, {Name: "Overflow", Doc: "Overflow determines how to handle overflowing content in a layout.\nDefault is OverflowVisible.  Set to OverflowAuto to enable scrollbars."}, {Name: "Gap", Doc: "For layouts, extra space added between elements in the layout."}, {Name: "Columns", Doc: "For grid layouts, the number of columns to use.\nIf > 0, number of rows is computed as N elements / Columns.\nUsed as a constraint in layout if individual elements\ndo not specify their row, column positions"}, {Name: "ObjectFit", Doc: "If this object is a replaced object (image, video, etc)\nor has a background image, ObjectFit specifies the way\nin which the replaced object should be fit into the element."}, {Name: "Border", Doc: "Border is a line border around the box element"}, {Name: "MaxBorder", Doc: "MaxBorder is the largest border that will ever be rendered\naround the element, the size of which is used for computing\nthe effective margin to allocate for the element."}, {Name: "BoxShadow", Doc: "BoxShadow is the box shadows to render around box (can have multiple)"}, {Name: "MaxBoxShadow", Doc: "MaxBoxShadow contains the largest shadows that will ever be rendered\naround the element, the size of which are used for computing the\neffective margin to allocate for the element."}, {Name: "Color", Doc: "Color specifies the text / content color, and it is inherited."}, {Name: "Background", Doc: "Background specifies the background of the element. It is not inherited,\nand it is nil (transparent) by default."}, {Name: "Opacity", Doc: "alpha value between 0 and 1 to apply to the foreground and background of this element and all of its children"}, {Name: "StateLayer", Doc: "StateLayer, if above zero, indicates to create a state layer over the element with this much opacity (on a scale of 0-1) and the\ncolor Color (or StateColor if it defined). It is automatically set based on State, but can be overridden in stylers."}, {Name: "StateColor", Doc: "StateColor, if not the zero color, is the color to use for the StateLayer instead of Color. If you want to disable state layers\nfor an element, do not use this; instead, set StateLayer to 0."}, {Name: "ActualBackground", Doc: "ActualBackground is the computed actual background rendered for the element,\ntaking into account its Background, Opacity, StateLayer, and parent\nActualBackground. It is automatically computed and should not be set manually."}, {Name: "Pos", Doc: "position is only used for Layout = Nil cases"}, {Name: "ZIndex", Doc: "ordering factor for rendering depth -- lower numbers rendered first.\nSort children according to this factor"}, {Name: "Row", Doc: "specifies the row that this element should appear within a grid layout"}, {Name: "Col", Doc: "specifies the column that this element should appear within a grid layout"}, {Name: "RowSpan", Doc: "specifies the number of sequential rows that this element should occupy\nwithin a grid layout (todo: not currently supported)"}, {Name: "ColSpan", Doc: "specifies the number of sequential columns that this element should occupy\nwithin a grid layout"}, {Name: "ScrollBarWidth", Doc: "width of a layout scrollbar"}, {Name: "Font", Doc: "font parameters -- no xml prefix -- also has color, background-color"}, {Name: "Text", Doc: "text parameters -- no xml prefix"}, {Name: "UnContext", Doc: "units context -- parameters necessary for anchoring relative units"}}})

var _ = gti.AddType(&gti.Type{Name: "cogentcore.org/core/styles.Text", IDName: "text", Doc: "Text is used for layout-level (widget, html-style) text styling --\nFontStyle contains all the lower-level text rendering info used in SVG --\nmost of these are inherited", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Fields: []gti.Field{{Name: "Align", Doc: "how to align text, horizontally (inhereted).\nThis *only* applies to the text within its containing element,\nand is typically relevant only for multi-line text:\nfor single-line text, if element does not have a specified size\nthat is different from the text size, then this has *no effect*."}, {Name: "AlignV", Doc: "vertical alignment of text (inhereted).\nThis is only applicable for SVG styling, not regular CSS / Cogent Core,\nwhich uses the global Align.Y.  It *only* applies to the text within\nits containing element: if that element does not have a specified size\nthat is different from the text size, then this has *no effect*."}, {Name: "Anchor", Doc: "for svg rendering only (inhereted):\ndetermines the alignment relative to text position coordinate.\nFor RTL start is right, not left, and start is top for TB"}, {Name: "LetterSpacing", Doc: "spacing between characters and lines"}, {Name: "WordSpacing", Doc: "extra space to add between words (inhereted)"}, {Name: "LineHeight", Doc: "specified height of a line of text (inhereted); text is centered within the overall lineheight;\nthe standard way to specify line height is in terms of em"}, {Name: "WhiteSpace", Doc: "prop: white-space (*not* inherited) specifies how white space is processed,\nand how lines are wrapped.  If set to WhiteSpaceNormal (default) lines are wrapped.\nSee info about interactions with Grow.X setting for this and the NoWrap case."}, {Name: "UnicodeBidi", Doc: "determines how to treat unicode bidirectional information (inhereted)"}, {Name: "Direction", Doc: "bidi-override or embed -- applies to all text elements (inhereted)"}, {Name: "WritingMode", Doc: "overall writing mode -- only for text elements, not span (inhereted)"}, {Name: "OrientationVert", Doc: "for TBRL writing mode (only), determines orientation of alphabetic characters (inhereted);\n90 is default (rotated); 0 means keep upright"}, {Name: "OrientationHoriz", Doc: "for horizontal LR/RL writing mode (only), determines orientation of all characters (inhereted);\n0 is default (upright)"}, {Name: "Indent", Doc: "how much to indent the first line in a paragraph (inhereted)"}, {Name: "ParaSpacing", Doc: "extra spacing between paragraphs (inhereted); copied from [Style.Margin] per CSS spec\nif that is non-zero, else can be set directly with para-spacing"}, {Name: "TabSize", Doc: "tab size, in number of characters (inhereted)"}}})

var _ = gti.AddType(&gti.Type{Name: "cogentcore.org/core/styles.XY", IDName: "xy", Doc: "XY represents X,Y values", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Fields: []gti.Field{{Name: "X", Doc: "X is the horizontal axis value"}, {Name: "Y", Doc: "Y is the vertical axis value"}}})
