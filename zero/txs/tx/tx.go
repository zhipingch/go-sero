// copyright 2018 The sero.cash Authors
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

package tx

import (
	"github.com/sero-cash/go-czero-import/keys"
	"github.com/sero-cash/go-sero/zero/txs/assets"
	"github.com/sero-cash/go-sero/zero/utils"
)

type In struct {
	Root keys.Uint256
}

type Out struct {
	Addr  keys.Uint512
	Asset assets.Asset
	Memo  keys.Uint512
	Z     OutType
}

type OutType int

const (
	TYPE_N = OutType(0)
	TYPE_O = OutType(1)
	TYPE_Z = OutType(2)
)

type T struct {
	FromRnd *keys.Uint256
	Ehash   keys.Uint256
	Fee     utils.U256
	Ins     []In
	Outs    []Out
}

func (self *T) TokenCost() (ret map[keys.Uint256]utils.U256) {
	ret = make(map[keys.Uint256]utils.U256)
	seroCy := utils.StringToUint256("SERO")
	ret[seroCy] = self.Fee
	if len(self.Outs) > 0 {
		for _, out := range self.Outs {
			if out.Asset.Tkn != nil {
				if cost, ok := ret[out.Asset.Tkn.Currency]; ok {
					cost.AddU(&out.Asset.Tkn.Value)
					ret[out.Asset.Tkn.Currency] = cost
				} else {
					ret[out.Asset.Tkn.Currency] = out.Asset.Tkn.Value
				}
			}
		}
	}
	return
}

func (self *T) TikectCost() (ret map[keys.Uint256][]keys.Uint256) {
	ret = make(map[keys.Uint256][]keys.Uint256)
	if len(self.Outs) > 0 {
		for _, out := range self.Outs {
			if out.Asset.Tkt != nil {
				if tkts, ok := ret[out.Asset.Tkt.Category]; ok {
					tkts = append(tkts, out.Asset.Tkt.Value)
					ret[out.Asset.Tkt.Category] = tkts
				} else {
					ret[out.Asset.Tkt.Category] = []keys.Uint256{out.Asset.Tkt.Value}
				}
			}
		}
	}
	return
}
