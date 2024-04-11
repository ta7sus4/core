// Code generated by "core generate"; DO NOT EDIT.

package cli

import (
	"cogentcore.org/core/types"
)

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/cli.metaConfigFields", IDName: "meta-config-fields", Doc: "metaConfigFields is the struct used for the implementation\nof [AddMetaConfigFields], and for the usage information for\nmeta configuration options in [Usage].\nNOTE: we could do this through [MetaConfig], but that\ncauses problems with the HelpCmd field capturing\neverything, so it easier to just add through a separate struct.\nTODO: maybe improve the structure of this\nTODO: can we get HelpCmd to display correctly in usage?", Directives: []types.Directive{{Tool: "gti", Directive: "add"}}, Fields: []types.Field{{Name: "Config", Doc: "the file name of the config file to load"}, {Name: "Help", Doc: "whether to display a help message"}, {Name: "HelpCmd", Doc: "the name of the command to display\nhelp information for."}, {Name: "Verbose", Doc: "whether to run the command in verbose mode\nand print more information"}, {Name: "VeryVerbose", Doc: "whether to run the command in very verbose mode\nand print as much information as possible"}, {Name: "Quiet", Doc: "whether to run the command in quiet mode\nand print less information"}}})
