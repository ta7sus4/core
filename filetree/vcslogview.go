// Copyright (c) 2020, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package filetree

import (
	"log/slog"

	"cogentcore.org/core/events"
	"cogentcore.org/core/gi"
	"cogentcore.org/core/giv"
	"cogentcore.org/core/glop/dirs"
	"cogentcore.org/core/grr"
	"cogentcore.org/core/icons"
	"cogentcore.org/core/mimedata"
	"cogentcore.org/core/states"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/texteditor"
	"cogentcore.org/core/vci"
)

// VCSLogView is a view of the VCS log data
type VCSLogView struct {
	gi.Layout

	// current log
	Log vci.Log

	// file that this is a log of -- if blank then it is entire repository
	File string

	// date expression for how long ago to include log entries from
	Since string

	// version control system repository
	Repo vci.Repo `json:"-" xml:"-" copier:"-"`

	// revision A -- defaults to HEAD
	RevA string `set:"-"`

	// revision B -- blank means current working copy
	RevB string `set:"-"`

	// double-click will set the A revision -- else B
	SetA bool
}

func (lv *VCSLogView) OnInit() {
	lv.Style(func(s *styles.Style) {
		s.Grow.Set(1, 1)
	})
	lv.OnWidgetAdded(func(w gi.Widget) {
		switch w.PathFrom(lv) {
		case "a-tf", "b-tf":
			w.Style(func(s *styles.Style) {
				s.Min.X.Em(12)
			})
		}
	})
}

// ConfigRepo configures to given repo, log and file (file could be empty)
func (lv *VCSLogView) ConfigRepo(repo vci.Repo, lg vci.Log, file, since string) {
	lv.Repo = repo
	lv.Log = lg
	lv.File = file
	lv.Since = since
	if lv.HasChildren() {
		return
	}
	lv.Style(func(s *styles.Style) {
		s.Direction = styles.Column
	})
	gi.NewToolbar(lv, "toolbar")
	tv := giv.NewTableView(lv, "log")
	tv.SetReadOnly(true)
	tv.SetSlice(&lv.Log)
	lv.RevA = "HEAD"
	lv.RevB = ""
	lv.SetA = true
	lv.ConfigToolbar()
	tv.AddContextMenu(func(m *gi.Scene) {
		gi.NewButton(m).SetText("Set Revision A").
			SetTooltip("Set Buffer A's revision to this").
			OnClick(func(e events.Event) {
				cmt := lv.Log[tv.SelectedIndex]
				lv.SetRevA(cmt.Rev)
			})
		gi.NewButton(m).SetText("Set Revision B").
			SetTooltip("Set Buffer B's revision to this").
			OnClick(func(e events.Event) {
				cmt := lv.Log[tv.SelectedIndex]
				lv.SetRevB(cmt.Rev)
			})
		gi.NewButton(m).SetText("Copy Revision ID").
			SetTooltip("Copies the revision number / hash for this ").
			OnClick(func(e events.Event) {
				cmt := lv.Log[tv.SelectedIndex]
				tv.Clipboard().Write(mimedata.NewText(cmt.Rev))
			})
		gi.NewButton(m).SetText("Checkout Revision").
			SetTooltip("Checks out this revision").
			OnClick(func(e events.Event) {
				cmt := lv.Log[tv.SelectedIndex]
				grr.Log(repo.UpdateVersion(cmt.Rev))
			})
	})
	tv.OnDoubleClick(func(e events.Event) {
		idx := tv.SelectedIndex
		if idx < 0 || idx >= len(lv.Log) {
			return
		}
		cmt := lv.Log[idx]
		if lv.File != "" {
			if lv.SetA {
				lv.SetRevA(cmt.Rev)
			} else {
				lv.SetRevB(cmt.Rev)
			}
			lv.ToggleRev()
		}
		cinfo, err := lv.Repo.CommitDesc(cmt.Rev, false)
		if err != nil {
			slog.Error(err.Error())
			return
		}
		d := gi.NewBody().AddTitle("Commit Info: " + cmt.Rev)
		buf := texteditor.NewBuffer()
		buf.Filename = gi.Filename(lv.File)
		buf.Opts.LineNos = true
		buf.Stat()
		texteditor.NewEditor(d).SetBuffer(buf)
		buf.SetText(cinfo)
		d.AddBottomBar(func(pw gi.Widget) {
			gi.NewButton(pw).SetText("Copy to clipboard").SetIcon(icons.ContentCopy).
				OnClick(func(e events.Event) {
					d.Clipboard().Write(mimedata.NewTextBytes(cinfo))
				})
			d.AddOk(pw)
		})
		d.NewFullDialog(lv).Run()
	})
}

// SetRevA sets the revision to use for buffer A
func (lv *VCSLogView) SetRevA(rev string) {
	lv.RevA = rev
	tb := lv.Toolbar()
	tfi := tb.ChildByName("a-tf", 2)
	if tfi == nil {
		return
	}
	tfi.(*gi.TextField).SetText(rev)
}

// SetRevB sets the revision to use for buffer B
func (lv *VCSLogView) SetRevB(rev string) {
	lv.RevB = rev
	tb := lv.Toolbar()
	tfi := tb.ChildByName("b-tf", 2)
	if tfi == nil {
		return
	}
	tfi.(*gi.TextField).SetText(rev)
}

// ToggleRev switches the active revision to set
func (lv *VCSLogView) ToggleRev() {
	tb := lv.Toolbar()
	cba := tb.ChildByName("a-rev", 2).(*gi.Switch)
	cbb := tb.ChildByName("b-rev", 2).(*gi.Switch)
	lv.SetA = !lv.SetA
	cba.SetState(lv.SetA, states.Checked)
	cbb.SetState(!lv.SetA, states.Checked)
}

// Toolbar returns the toolbar
func (lv *VCSLogView) Toolbar() *gi.Toolbar {
	return lv.ChildByName("toolbar", 0).(*gi.Toolbar)
}

// TableView returns the tableview
func (lv *VCSLogView) TableView() *giv.TableView {
	return lv.ChildByName("log", 1).(*giv.TableView)
}

// ConfigToolbar
func (lv *VCSLogView) ConfigToolbar() {
	tb := lv.Toolbar()
	if lv.File != "" {
		gi.NewLabel(tb, "fl", "File: "+dirs.DirAndFile(lv.File))
		gi.NewSeparator(tb, "flsep")
		cba := gi.NewSwitch(tb, "a-rev").SetText("A Rev: ").
			SetTooltip("If selected, double-clicking in log will set this A Revision to use for Diff")
		cba.SetState(true, states.Checked)
		tfa := gi.NewTextField(tb, "a-tf").SetText(lv.RevA)
		tfa.OnChange(func(e events.Event) {
			lv.RevA = tfa.Text()
		})
		gi.NewButton(tb, "view-a").SetText("View A").SetIcon(icons.Document).
			SetTooltip("View file at revision A").
			OnClick(func(e events.Event) {
				FileAtRevDialog(lv, lv.Repo, lv.File, lv.RevA)
			})

		gi.NewSeparator(tb, "absep")

		cbb := gi.NewSwitch(tb, "b-rev").SetText("B Rev: ").
			SetTooltip("If selected, double-clicking in log will set this B Revision to use for Diff")
		cbb.OnClick(func(e events.Event) {
			lv.SetA = !cbb.IsChecked()
			cba.SetState(lv.SetA, states.Checked)
			cba.NeedsRender()
		})
		cba.OnClick(func(e events.Event) {
			lv.SetA = cba.IsChecked()
			cbb.SetState(!lv.SetA, states.Checked)
			cbb.NeedsRender()
		})

		tfb := gi.NewTextField(tb, "b-tf").SetText(lv.RevB)
		tfb.OnChange(func(e events.Event) {
			lv.RevB = tfb.Text()
		})
		gi.NewButton(tb, "view-b").SetText("View B").SetIcon(icons.Document).
			SetTooltip("View file at revision B").
			OnClick(func(e events.Event) {
				FileAtRevDialog(lv, lv.Repo, lv.File, lv.RevB)
			})

		gi.NewSeparator(tb, "dsep")

		gi.NewButton(tb, "diff").SetText("Diff").SetIcon(icons.Difference).
			SetTooltip("Show the diffs between two revisions -- if blank, A is current HEAD, and B is current working copy").
			OnClick(func(e events.Event) {
				texteditor.DiffViewDialogFromRevs(lv, lv.Repo, lv.File, nil, lv.RevA, lv.RevB)
			})
	}

}

// VCSLogViewDialog returns a VCS Log View for given repo, log and file (file could be empty)
func VCSLogViewDialog(ctx gi.Widget, repo vci.Repo, lg vci.Log, file, since string) *gi.Body {
	title := "VCS Log: "
	if file == "" {
		title += "All files"
	} else {
		title += dirs.DirAndFile(file)
	}
	if since != "" {
		title += " since: " + since
	}
	d := gi.NewBody().AddTitle(title)
	lv := NewVCSLogView(d, "vcslog")
	lv.ConfigRepo(repo, lg, file, since)
	d.NewFullDialog(ctx).SetNewWindow(true).Run()
	return d
}

// FileAtRevDialog Shows a file at a given revision in a new dialog window
func FileAtRevDialog(ctx gi.Widget, repo vci.Repo, file, rev string) *gi.Body {
	fb, err := repo.FileContents(file, rev)
	if err != nil {
		gi.ErrorDialog(ctx, err)
		return nil
	}
	if rev == "" {
		rev = "HEAD"
	}
	title := "File at VCS Revision: " + dirs.DirAndFile(file) + "@" + rev
	d := gi.NewBody().AddTitle(title)

	tb := texteditor.NewBuffer().SetText(fb).SetFilename(file) // file is key for getting lang
	texteditor.NewEditor(d).SetBuffer(tb).SetReadOnly(true)
	d.NewFullDialog(ctx).SetNewWindow(true).Run()
	tb.StartDelayedReMarkup() // update markup
	return d
}
