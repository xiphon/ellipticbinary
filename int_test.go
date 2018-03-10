/*
Basic Elliptic Curve primitives over Binary Field GF(2ⁿ)

Copyright (C) 2018 Xiphon

Greatly inspired by Kurt Rose's python implementation
https://gist.github.com/kurtbrose/4423605

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package ellipticbinary

import (
	"encoding/hex"
	"math/big"
	"testing"
)

func TestInt(t *testing.T) {
	bytes, _ := hex.DecodeString("04debb88a0aac5502feb6bb4ff6f0d16a0ad1c21d28e63202fd2c13bea637c04f1da862b")
	value := newBianryFieldInt(big.NewInt(0).SetBytes(bytes))

	{
		other := newBianryFieldInt(big.NewInt(0).SetBytes(bytes))
		if value.cmp(other) != 0 {
			t.FailNow()
		}
	}

	{
		bytes, _ = hex.DecodeString("04debb88a0aac5502feb6bb4ff6f0d16a0ad1c21d28e63202fd2c13bea637c04f1da862a")
		other := newBianryFieldInt(big.NewInt(0).SetBytes(bytes))
		if value.cmp(other) == 0 {
			t.FailNow()
		}
	}

	value.mul(value, value)
	{
		bytes, _ = hex.DecodeString("1051544545404044004444501111000455544514454510555514550051011444004451015004015104405414050400045551045001054554441405155000105501514440140445")
		valid := newBianryFieldInt(big.NewInt(0).SetBytes(bytes))
		if value.cmp(valid) != 0 {
			t.FailNow()
		}
	}

	bytes, _ = hex.DecodeString("0800000000000000000000000000000000000000000000000000000000000000000010a1")
	order := newBianryFieldInt(big.NewInt(0).SetBytes(bytes))
	bytes, _ = hex.DecodeString("05a2b1423829c50c19d078e9712379fa7234e078178958a5924c8e16091cee9c4dab87d9")
	valid := newBianryFieldInt(big.NewInt(0).SetBytes(bytes))

	value.mod(value, order)
	if value.cmp(valid) != 0 {
		t.FailNow()
	}

	{
		bytes, _ = hex.DecodeString("aa160a283c315eaa6456d0156ec97d7332402abb709d4abf75031fdd0aa7861c84d35c")
		other := newBianryFieldInt(big.NewInt(0).SetBytes(bytes))
		bytes, _ = hex.DecodeString("0508a7481015f452b3b42e39644db0870106a052acf9c5ef2d398d09d416491a512f5485")
		valid := newBianryFieldInt(big.NewInt(0).SetBytes(bytes))

		value.add(value, other)
		if value.cmp(valid) != 0 {
			t.FailNow()
		}
	}

	{
		bytes, _ = hex.DecodeString("04debb88a0aac5502feb6bb4ff6f0d16a0ad1c21d28e63202fd2c13bea637c04f1da862b")
		other := newBianryFieldInt(big.NewInt(0).SetBytes(bytes))
		bytes, _ = hex.DecodeString("16b75d4fa88a16b3181053c7fdc885f8001412234d4f345eeb259c46ab500b204100e5179da1ae641e51e3042a2da52e6990c68fb9ef58b6ad921d767f9b88605c558596184923")
		valid := newBianryFieldInt(big.NewInt(0).SetBytes(bytes))

		value.divmod(value, other, order)
		if value.cmp(valid) != 0 {
			t.FailNow()
		}
	}
}
