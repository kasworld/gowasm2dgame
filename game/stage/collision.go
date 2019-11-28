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

package stage

import (
	"github.com/kasworld/gowasm2dgame/game/gameconst"
	"github.com/kasworld/gowasm2dgame/lib/quadtreef"
	"github.com/kasworld/gowasm2dgame/lib/rectf"
)

func (stg *Stage) newQtree() *quadtreef.QuadTree {
	maxr := 32.0
	qtree := quadtreef.New(rectf.Rect{
		0 - maxr, 0 - maxr,
		gameconst.StageW + maxr*2, gameconst.StageH + maxr*2,
	})
	return qtree
}

func (stg *Stage) checkCollision() []*GameObj {
	toDeleteList := make([]*GameObj, 0)
	qtree := stg.newQtree()
	obj2check := make([]*GameObj, 0)
	for _, bt := range stg.Teams {
		if qtree.Insert(bt.Ball) {
			obj2check = append(obj2check, bt.Ball)
		}
		for _, v := range bt.Objs {
			if v.toDelete {
				continue
			}
			if qtree.Insert(v) {
				obj2check = append(obj2check, v)
			}
		}
	}
	for _, v := range obj2check {
		if v.toDelete {
			continue
		}
		qtree.QueryByRect(
			func(qo quadtreef.QuadTreeObjI) bool {
				targetObj := qo.(*GameObj)
				if targetObj.toDelete {
					return false
				}
				_ = targetObj
				if targetObj.teamType == v.teamType {
					return false
				}
				if !targetObj.toDelete {
					targetObj.toDelete = true
					toDeleteList = append(toDeleteList, targetObj)
				}
				if !v.toDelete {
					v.toDelete = true
					toDeleteList = append(toDeleteList, v)
					return true
				}
				return false
			},
			v.GetRect(),
		)
	}
	return toDeleteList
}
