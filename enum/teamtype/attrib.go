// Copyright 2015,2016,2017,2018,2019,2020 SeukWon Kang (kasworld@gmail.com)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package teamtype

import (
	"github.com/kasworld/htmlcolors"
)

var Attrib = [TeamType_Count]struct {
	Color24 htmlcolors.Color24
}{
	Red:    {htmlcolors.Red},
	Blue:   {htmlcolors.Blue},
	Green:  {htmlcolors.Green},
	RRed:   {htmlcolors.Red.Neg()},
	RBlue:  {htmlcolors.Blue.Neg()},
	RGreen: {htmlcolors.Green.Neg()},
}
