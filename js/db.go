package js

import (
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"sexy-filter/log"
	"sync"
)

var g_engine *xorm.Engine
var g_mutex sync.Mutex

func InitDB() (e error) {
	g_engine, e = xorm.NewEngine("sqlite3", WorkDir+"/my.db")
	if e != nil {
		if log.Fault != nil {
			log.Fault.Println(e)
		}
		return
	}
	engine := g_engine

	var node Node
	var ok bool
	if ok, e = engine.IsTableExist(&node); e != nil {
		if log.Fault != nil {
			log.Fault.Println(e)
		}
		return
	} else if !ok {
		if e = engine.CreateTables(&node); e != nil {
			if log.Fault != nil {
				log.Fault.Println(e)
			}
			return e
		}
		if e = engine.CreateUniques(&node); e != nil {
			if log.Fault != nil {
				log.Fault.Println(e)
			}
			return
		}
		if e = engine.CreateIndexes(&node); e != nil {
			if log.Fault != nil {
				log.Fault.Println(e)
			}
			return
		}
	}

	return
}
func InsertNodes(nodes []*Node) {
	g_mutex.Lock()
	g_engine.Insert(nodes)
	g_mutex.Unlock()
}
