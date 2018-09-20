// Copyright 2015 The sero.cash Authors
// This file is part of the go-sero library.
//
// The go-sero library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-sero library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-sero library. If not, see <http://www.gnu.org/licenses/>.

package witness

import (
	"github.com/sero-cash/go-sero/zero/witness/merkle"
	"testing"
)

var genLeaf = merkle.GenLeaf
var getLeaf = merkle.GetLeaf

func TestMerkle(t *testing.T) {
	tree1 := merkle.Tree{}
	w1 := Witness{Tree: tree1.Clone()}
	m1 := make(map[merkle.Leaf]int)
	for i := 0; i < 100000; i++ {
		tree1.Append(merkle.Leaf{1})
		w1.Append(merkle.Leaf{1})
		t_root := tree1.Root()
		w_root := w1.Root()
		if t_root != w_root {
			panic("")
		}
		if v, ok := m1[t_root]; ok {
			panic(v)
		} else {
			m1[t_root] = i
		}
	}

	var tree merkle.Tree
	tree.Append(genLeaf())
	tree.Append(genLeaf())
	tree.Append(genLeaf())
	tree.Append(genLeaf())
	tree.Append(genLeaf())
	tree.Append(genLeaf())
	w := Witness{Tree: tree.Clone()}
	w.Append(genLeaf())
	tree.Append(getLeaf())
	t_root := tree.Root()
	w_root := w.Root()
	t.Log(t_root, w_root)

	w.Append(genLeaf())
	tree.Append(getLeaf())
	t_root = tree.Root()
	w_root = w.Root()

	w.Append(genLeaf())
	tree.Append(getLeaf())
	t_root = tree.Root()
	w_root = w.Root()

	w.Append(genLeaf())
	tree.Append(getLeaf())
	t_root = tree.Root()
	w_root = w.Root()

	w.Append(genLeaf())
	tree.Append(getLeaf())
	t_root = tree.Root()
	w_root = w.Root()

	w.Append(genLeaf())
	tree.Append(getLeaf())
	t_root = tree.Root()
	w_root = w.Root()

	w.Append(genLeaf())
	tree.Append(getLeaf())
	t_root = tree.Root()
	w_root = w.Root()

	w.Append(genLeaf())
	tree.Append(getLeaf())
	t_root = tree.Root()
	w_root = w.Root()

	w.Append(genLeaf())
	tree.Append(getLeaf())
	t_root = tree.Root()
	w_root = w.Root()

	w.Append(genLeaf())
	tree.Append(getLeaf())
	t_root = tree.Root()
	w_root = w.Root()

	root := w.Root()
	t.Logf("root:%v", root)
	path := PartialPath(w)
	t.Logf("path:%v", path)
	last := merkle.Last(w.Tree)
	t.Logf("last:%v", last)
	elem := w.Element()
	t.Logf("elem:%v", elem)
	p, index := w.Path()
	t.Logf("path:%v", p)
	t.Logf("index:%v", index)
}