// Code generated by "genprotocol.exe -ver=8ce3a4010a59de778695c59389c0fd9a3938197ee346a449f005b536d94d0e60 -basedir=protocol_w2d -prefix=w2d -statstype=int"

package w2d_connbytemanager

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
