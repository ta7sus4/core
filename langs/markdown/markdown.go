// Copyright (c) 2018, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pi

import (
	"github.com/goki/pi/complete"
	"github.com/goki/pi/filecat"
	"github.com/goki/pi/lex"
	"github.com/goki/pi/pi"
	"github.com/goki/pi/syms"
)

// MarkdownLang implements the Lang interface for the Markdown language
type MarkdownLang struct {
	Pr *pi.Parser
}

// TheMarkdownLang is the instance variable providing support for the Markdown language
var TheMarkdownLang = MarkdownLang{}

func init() {
	pi.StdLangProps[filecat.Markdown].Lang = &TheMarkdownLang
}

func (ml *MarkdownLang) Parser() *pi.Parser {
	if ml.Pr != nil {
		return ml.Pr
	}
	lp, _ := pi.LangSupport.Props(filecat.Markdown)
	if lp.Parser == nil {
		pi.LangSupport.OpenStd()
	}
	ml.Pr = lp.Parser
	if ml.Pr == nil {
		return nil
	}
	return ml.Pr
}

func (ml *MarkdownLang) ParseFile(fss *pi.FileStates, txt []byte) {
	pr := ml.Parser()
	if pr == nil {
		return
	}
	pfs := fss.StartProc(txt) // current processing one
	pr.LexAll(pfs)
	fss.EndProc() // now done
	// no parser
}

func (ml *MarkdownLang) LexLine(fs *pi.FileState, line int, txt []rune) lex.Line {
	pr := ml.Parser()
	if pr == nil {
		return nil
	}
	return pr.LexLine(fs, line, txt)
}

func (ml *MarkdownLang) ParseLine(fs *pi.FileState, line int) *pi.FileState {
	// n/a
	return nil
}

func (ml *MarkdownLang) HiLine(fss *pi.FileStates, line int, txt []rune) lex.Line {
	fs := fss.Done()
	return ml.LexLine(fs, line, txt)
}

func (ml *MarkdownLang) CompleteLine(fss *pi.FileStates, str string, pos lex.Pos) (md complete.Matches) {
	// n/a
	return md
}

// Lookup is the main api called by completion code in giv/complete.go to lookup item
func (gl *MarkdownLang) Lookup(fss *pi.FileStates, str string, pos lex.Pos) (ld complete.Lookup) {
	return
}

func (ml *MarkdownLang) CompleteEdit(fs *pi.FileStates, text string, cp int, comp complete.Completion, seed string) (ed complete.Edit) {
	// n/a
	return ed
}

func (ml *MarkdownLang) ParseDir(path string, opts pi.LangDirOpts) *syms.Symbol {
	// n/a
	return nil
}
