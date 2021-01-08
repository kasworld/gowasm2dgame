// Copyright 2015,2016,2017,2018,2019,2020,2021 SeukWon Kang (kasworld@gmail.com)
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

type iv struct {
	Index int
	Value int
}

var SpriteFilter = [TeamType_Count]struct {
	Name string
	IV   []iv
}{
	Red:    {"red", []iv{{0, 255}}},
	Blue:   {"blue", []iv{{1, 255}}},
	Green:  {"green", []iv{{2, 255}}},
	RRed:   {"rred", []iv{{0, 0}}},
	RBlue:  {"rblue", []iv{{1, 0}}},
	RGreen: {"rgreen", []iv{{2, 0}}},
}
