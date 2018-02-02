package js

import (
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/zuiwuchang/king-go/go-xorm/params"
	"sexy-filter/log"
	"strconv"
	"strings"
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
	//engine.ShowSQL(true)

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
func Search(title, name, id, limit string) (nodes []*Node) {
	title = strings.TrimSpace(title)
	name = strings.TrimSpace(name)
	id = strings.TrimSpace(id)
	limit = strings.TrimSpace(limit)

	wh := params.NewParams(3)
	first := true
	if title != "" {
		if first {
			wh.WriteWhere(fmt.Sprintf(" %s like ? ", ColNodeTiTle))
			first = false
		} else {
			wh.WriteWhere(fmt.Sprintf(" and %s like ? ", ColNodeTiTle))
		}
		wh.WriteParam("%" + title + "%")
	}

	if name != "" {
		if first {
			wh.WriteWhere(fmt.Sprintf(" %s like ? ", ColNodePluginsName))
			first = false
		} else {
			wh.WriteWhere(fmt.Sprintf(" and %s like ? ", ColNodePluginsName))
		}
		wh.WriteParam("%" + name + "%")
	}

	if id != "" {
		if first {
			wh.WriteWhere(fmt.Sprintf(" %s like ? ", ColNodePluginsId))
			first = false
		} else {
			wh.WriteWhere(fmt.Sprintf(" and %s like ? ", ColNodePluginsId))
		}
		wh.WriteParam("%" + id + "%")
	}

	g_mutex.Lock()
	defer g_mutex.Unlock()
	var e error
	session := g_engine.NewSession()
	defer session.Close()
	if !first {
		wh.Where(session)
	}
	start := 0
	n := 20
	if limit != "" {
		strs := strings.Split(limit, ",")
		if len(strs) > 0 {
			str := strings.TrimSpace(strs[0])
			i, _ := strconv.ParseInt(str, 10, 10)
			if i > 9 && i < 101 {
				n = int(i)
			}
		}
		if len(strs) > 1 {
			str := strings.TrimSpace(strs[1])
			i, _ := strconv.ParseInt(str, 10, 10)
			if i > 0 {
				start = int(i)
			}
		}
	}

	if e = session.Limit(n, start).Find(&nodes); e != nil {
		if log.Error != nil {
			log.Error.Println(e)
		}
	}

	return
}
