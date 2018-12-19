// Copyright (c) 2018, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syms

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"sort"
	"strings"

	"github.com/goki/gi/complete"
	"github.com/goki/pi/lex"
	"github.com/goki/pi/token"
)

// SymMap is a map between symbol names and their full information.
// A given project will have a top-level SymMap and perhaps local
// maps for individual files, etc.  Namespaces / packages can be
// created and elements added to them to create appropriate
// scoping structure etc.  Note that we have to use pointers
// for symbols b/c otherwise it is very expensive to re-assign
// values all the time -- https://github.com/golang/go/issues/3117
type SymMap map[string]*Symbol

// Alloc ensures that map is made
func (sm *SymMap) Alloc() {
	if *sm == nil {
		*sm = make(SymMap)
	}
}

// Add adds symbol to map
func (sm *SymMap) Add(sy *Symbol) {
	sm.Alloc()
	(*sm)[sy.Name] = sy
}

// AddNew adds a new symbol to the map with the basic info
func (sm *SymMap) AddNew(name string, kind token.Tokens, fname string, reg lex.Reg) *Symbol {
	sy := NewSymbol(name, kind, fname, reg)
	sm.Alloc()
	(*sm)[name] = sy
	return sy
}

// Reset resets the symbol map
func (sm *SymMap) Reset() {
	*sm = make(SymMap)
}

// CopyFrom copies all the symbols from given source map into this one,
// including merging everything from common elements
func (sm *SymMap) CopyFrom(src SymMap) {
	sm.Alloc()
	for nm, ssy := range src {
		dsy, has := (*sm)[nm]
		if !has {
			(*sm)[nm] = ssy
			continue
		}
		if dsy.IsTemp() {
			ssy.Children.CopyFrom(dsy.Children)
			(*sm)[nm] = ssy
		} else {
			dsy.Children.CopyFrom(ssy.Children)
		}
	}
}

// FindNameScoped looks for given symbol name within this map and any children on the map
// that are of subcategory token.NameScope (i.e., namespace, module, package, library)
func (sm *SymMap) FindNameScoped(nm string) (*Symbol, bool) {
	sy, has := (*sm)[nm]
	if has {
		return sy, has
	}
	for _, ss := range *sm {
		if ss.Kind.SubCat() == token.NameScope {
			sy, has = ss.Children.FindNameScoped(nm)
			if has {
				return sy, has
			}
		}
	}
	return nil, false
}

// FindKindScoped looks for given symbol kind within this map and any children on the map
// that are of subcategory token.NameScope (i.e., namespace, module, package, library).
// Returns all instances found
func (sm *SymMap) FindKindScoped(kind token.Tokens) SymMap {
	sys := make(SymMap)
	for _, sy := range *sm {
		if sy.Kind == kind {
			sys[sy.Name] = sy
		}
	}
	for _, ss := range *sm {
		if ss.Kind.SubCat() == token.NameScope {
			csys := ss.Children.FindKindScoped(kind)
			if len(csys) > 0 {
				sys.CopyFrom(csys)
			}
		}
	}
	return sys
}

// OpenJSON opens from a JSON-formatted file.
func (sm *SymMap) OpenJSON(filename string) error {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, sm)
}

// SaveJSON saves to a JSON-formatted file.
func (sm *SymMap) SaveJSON(filename string) error {
	b, err := json.MarshalIndent(sm, "", "  ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, b, 0644)
	return err
}

// Names returns a slice of the names in this map, optionally sorted
func (sm *SymMap) Names(sorted bool) []string {
	nms := make([]string, len(*sm))
	idx := 0
	for _, sy := range *sm {
		nms[idx] = sy.Name
		idx++
	}
	if sorted {
		sort.StringSlice(nms).Sort()
	}
	return nms
}

// KindNames returns a slice of the kind:names in this map, optionally sorted
func (sm *SymMap) KindNames(sorted bool) []string {
	nms := make([]string, len(*sm))
	idx := 0
	for _, sy := range *sm {
		nms[idx] = sy.Kind.String() + ":" + sy.Name
		idx++
	}
	if sorted {
		sort.StringSlice(nms).Sort()
	}
	return nms
}

// WriteDoc writes basic doc info, sorted by kind and name
func (sm *SymMap) WriteDoc(out io.Writer, depth int) {
	nms := sm.KindNames(true)
	for _, nm := range nms {
		ci := strings.Index(nm, ":")
		sy := (*sm)[nm[ci+1:]]
		sy.WriteDoc(out, depth)
	}
}

//////////////////////////////////////////////////////////////////////
// Partial lookups

// Slice returns a slice of the elements in the map, optionally sorted by name
func (sm *SymMap) Slice(sorted bool) []*Symbol {
	sys := make([]*Symbol, len(*sm))
	idx := 0
	for _, sy := range *sm {
		sys[idx] = sy
		idx++
	}
	if sorted {
		sort.Slice(sys, func(i, j int) bool {
			return sys[i].Name < sys[j].Name
		})
	}
	return sys
}

// FindNamePrefix looks for given symbol name prefix within this map
// and any children on the map that are of subcategory
// token.NameScope (i.e., namespace, module, package, library)
// adds to given matches map (which can be nil), for more efficient recursive use
func (sm *SymMap) FindNamePrefix(seed string, matches *SymMap) {
	noCase := true
	if complete.HasUpperCase(seed) {
		noCase = false
	}
	if *matches == nil {
		*matches = make(SymMap)
	}
	for _, sy := range *sm {
		nm := sy.Name
		if noCase {
			nm = strings.ToLower(nm)
		}
		if strings.HasPrefix(nm, seed) {
			(*matches)[sy.Name] = sy
		}
	}
	for _, ss := range *sm {
		if ss.Kind.SubCat() == token.NameScope {
			ss.Children.FindNamePrefix(seed, matches)
		}
	}
}
