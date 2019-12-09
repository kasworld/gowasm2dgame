// Copyright 2015,2016,2017,2018,2019 SeukWon Kang (kasworld@gmail.com)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vector2f

import "math"

type Vector2f struct {
	X float64
	Y float64
}

var VtZero = Vector2f{0, 0}

func NewVectorLenAngle(l, a float64) Vector2f {
	return Vector2f{
		X: l * math.Cos(a),
		Y: l * math.Sin(a),
	}
}

func (v Vector2f) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vector2f) Add(v2 Vector2f) Vector2f {
	return Vector2f{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
	}
}

func (v Vector2f) Sub(v2 Vector2f) Vector2f {
	return Vector2f{
		X: v.X - v2.X,
		Y: v.Y - v2.Y,
	}
}

func (v Vector2f) MulF(f float64) Vector2f {
	return Vector2f{
		X: v.X * f,
		Y: v.Y * f,
	}
}

func (v Vector2f) DivF(f float64) Vector2f {
	return Vector2f{
		X: v.X / f,
		Y: v.Y / f,
	}
}

func (v Vector2f) Normalize() Vector2f {
	return v.DivF(v.Abs())
}

func (v Vector2f) Neg() Vector2f {
	return Vector2f{
		X: -v.X,
		Y: -v.Y,
	}
}

func (v Vector2f) NegX() Vector2f {
	return Vector2f{
		X: -v.X,
		Y: v.Y,
	}
}

func (v Vector2f) NegY() Vector2f {
	return Vector2f{
		X: v.X,
		Y: -v.Y,
	}
}

func (v Vector2f) LenTo(v2 Vector2f) float64 {
	return v2.Sub(v).Abs()
}

func (v Vector2f) Phase() float64 {
	return math.Atan2(v.Y, v.X)
}

func (v Vector2f) AddAngle(angle float64) Vector2f {
	return NewVectorLenAngle(v.Abs(), v.Phase()+angle)
}

func (v Vector2f) RotateBy(center Vector2f, angle float64) Vector2f {
	return v.Sub(center).AddAngle(angle).Add(center)
}

func (v Vector2f) Dot(v2 Vector2f) float64 {
	return v.X*v2.X + v.Y*v2.Y
}

func (v Vector2f) Cross() Vector2f {
	return Vector2f{
		X: v.Y,
		Y: -v.X,
	}
}

func (v Vector2f) IsIn(rt Rect) bool {
	return v.X >= rt.X1() && v.X <= rt.X2() &&
		v.Y >= rt.Y1() && v.Y <= rt.Y2()
}
