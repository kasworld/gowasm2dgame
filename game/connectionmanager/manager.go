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

package connectionmanager

import (
	"fmt"
	"sync"

	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_serveconnbyte"
)

type Manager struct {
	mutex   sync.RWMutex
	id2Conn map[string]*w2d_serveconnbyte.ServeConnByte
}

func New() *Manager {
	rtn := &Manager{
		id2Conn: make(map[string]*w2d_serveconnbyte.ServeConnByte),
	}
	return rtn
}

func (cm *Manager) Add(id string, c2sc *w2d_serveconnbyte.ServeConnByte) error {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()
	if cm.id2Conn[id] != nil {
		return fmt.Errorf("already exist %v", id)
	}
	cm.id2Conn[id] = c2sc
	return nil
}

func (cm *Manager) Del(id string) error {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()
	if cm.id2Conn[id] == nil {
		return fmt.Errorf("not exist %v", id)
	}
	delete(cm.id2Conn, id)
	return nil
}

func (cm *Manager) Get(id string) *w2d_serveconnbyte.ServeConnByte {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()
	return cm.id2Conn[id]
}

func (cm *Manager) Len() int {
	return len(cm.id2Conn)
}

func (cm *Manager) GetList() []*w2d_serveconnbyte.ServeConnByte {
	rtn := make([]*w2d_serveconnbyte.ServeConnByte, 0, len(cm.id2Conn))
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()
	for _, v := range cm.id2Conn {
		rtn = append(rtn, v)
	}
	return rtn
}
