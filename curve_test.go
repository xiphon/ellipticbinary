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

func TestPoint(t *testing.T) {
	sect283k1 := &Curve{}
	sect283k1.Name = "sect283k1"
	sect283k1.P, _ = new(big.Int).SetString("0800000000000000000000000000000000000000000000000000000000000000000010a1", 16)
	sect283k1.N, _ = new(big.Int).SetString("01FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFE9AE2ED07577265DFF7F94451E061E163C61", 16)
	sect283k1.A, _ = new(big.Int).SetString("0", 10)
	sect283k1.B, _ = new(big.Int).SetString("1", 10)
	sect283k1.Gx, _ = new(big.Int).SetString("0503213f78ca44883f1a3b8162f188e553cd265f23c1567a16876913b0c2ac2458492836", 16)
	sect283k1.Gy, _ = new(big.Int).SetString("01ccda380f1c9e318d90f95d07e5426fe87e45c0e8184698e45962364e34116177dd2259", 16)
	sect283k1.BitSize = 283

	var xVal, yVal *big.Int

	{
		bytes, _ := hex.DecodeString("aa160a283c315eaa6456d0156ec97d7332402abb709d4abf75031fdd0aa7861c84d35c")
		xVal = big.NewInt(0).SetBytes(bytes)

		bytes, _ = hex.DecodeString("04debb88a0aac5502feb6bb4ff6f0d16a0ad1c21d28e63202fd2c13bea637c04f1da862b")
		yVal = big.NewInt(0).SetBytes(bytes)

		if !sect283k1.IsOnCurve(xVal, yVal) {
			t.FailNow()
		}
		if sect283k1.IsOnCurve(yVal, xVal) {
			t.FailNow()
		}
	}

	{
		x, y := sect283k1.Add(xVal, yVal, sect283k1.Params().Gx, sect283k1.Params().Gy)

		bytes, _ := hex.DecodeString("02b4b7a7117eb95c19a9d5365bd5e319039b2cd348a29dd5d62a059a13bda65da7826185")
		xVaild := big.NewInt(0).SetBytes(bytes)

		bytes, _ = hex.DecodeString("043ca8ebfc777ecc668736e51eac7ff7a463dddc91aa8a0f61334ce161298646c0a5a19e")
		yVaild := big.NewInt(0).SetBytes(bytes)

		if x.Cmp(xVaild) != 0 || y.Cmp(yVaild) != 0 {
			t.FailNow()
		}
	}

	{
		x, y := sect283k1.Add(xVal, yVal, big.NewInt(0), big.NewInt(0))
		if x.Cmp(xVal) != 0 || y.Cmp(yVal) != 0 {
			t.FailNow()
		}

		x, y = sect283k1.Add(big.NewInt(0), big.NewInt(0), xVal, yVal)
		if x.Cmp(xVal) != 0 || y.Cmp(yVal) != 0 {
			t.FailNow()
		}
	}

	{
		num := big.NewInt(21910281)
		x, y := sect283k1.ScalarBaseMult(num.Bytes())

		bytes, _ := hex.DecodeString("04185e5d0268dca3823cdd9bd2bf71be3a5138c65a198983eb73c2258d9a8eb8a604fe6e")
		xVaild := big.NewInt(0).SetBytes(bytes)

		bytes, _ = hex.DecodeString("06653509ffa1715ea5767f74f4647803ca404cd32c5502929a4eccf4382c5be00dc14c47")
		yVaild := big.NewInt(0).SetBytes(bytes)

		if x.Cmp(xVaild) != 0 || y.Cmp(yVaild) != 0 {
			t.FailNow()
		}
	}

}
