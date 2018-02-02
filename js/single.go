package js

import (
	kStrings "github.com/zuiwuchang/king-go/strings"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"sexy-filter/log"
	"sexy-filter/net"
)

const (
	//空閒
	StatusNone = iota
	//啓動
	StatusRun
	//正在運行
	StatusRuning
	//正在停止
	StatusStoping
)

type Single struct {
	duks []*Duktape
	//狀態
	status int
	//狀態 改變爲 StatusNone 或 StatusRuning
	OnStatusChanged func(int)

	//需要運行 標記
	run bool
}

var g_Single Single

func GetSingle() *Single {
	return &g_Single
}
func InitSingle() {
	single := &g_Single
	single.status = StatusNone

	//初始化 duk
	n := runtime.NumCPU()
	if log.Info != nil {
		log.Info.Println(n, "goroutine")
	}
	n++
	duks := make([]*Duktape, n)
	for i := 0; i < n; i++ {
		duk := NewDuktape()

		//加載 插件
		dir := WorkDir + "/plugins-js"
		filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if path == dir {
				return nil
			} else if info.IsDir() {
				return filepath.SkipDir
			}
			if i == 0 && log.Info != nil {
				log.Info.Println("load plugins", info.Name())
			}

			if e := duk.LoadPluginsJs(path); e != nil {
				if i == 0 && log.Warn != nil {
					log.Warn.Panicln(e)
				}
				return nil
			}
			return nil
		})

		duks[i] = duk
	}
	single.duks = duks
}
func (s *Single) Status() int {
	return s.status
}
func (s *Single) GetPlugins() []string {
	return s.duks[0].GetPluginsNames()
}
func (s *Single) Start(style int, addr, user, pwd string, pos int) int {
	if pos == -1 {
		return s.status
	}
	if s.status != StatusNone {
		return s.status
	}
	//返回 插件 id
	id := s.duks[0].GetPluginsIdByPos(pos)
	if id == "" {
		return s.status
	}
	name := s.duks[0].GetPluginsName(id)
	if log.Trace != nil {
		log.Trace.Println("plugins :", id, name)
	}

	s.status = StatusRun
	s.run = true
	go func() {
		s.status = StatusRuning
		if s.OnStatusChanged != nil {
			s.OnStatusChanged(StatusRuning)
		}

		ch := make(chan int)
		chRequest := make(chan int)
		//runing
		i := 1
		for ; i < len(s.duks); i++ {
			go s.work(
				style, addr, user, pwd,
				chRequest, ch, i,
				id, name,
			)
		}
		go func() {
			defer func() {
				if e := recover(); e != nil {
				}
			}()
			requestI := 0
			for true {
				chRequest <- requestI
				requestI++
			}
		}()

		//wait stop
		for i != 1 {
			<-ch
			i--
		}

		close(chRequest)

		s.status = StatusNone
		if s.OnStatusChanged != nil {
			s.OnStatusChanged(StatusNone)
		}

	}()
	return s.status
}
func (s *Single) Stop() int {
	if s.status != StatusRuning {
		return s.status
	}
	s.status = StatusStoping
	s.run = false
	return s.status
}
func (s *Single) work(style int, addr, user, pwd string,
	chRequest, ch chan int, i int,
	id, name string,
) {
	if log.Info != nil {
		log.Info.Println("goroutine begin", i)
	}
	duk := s.duks[i]
	if c, e := net.CreateClient(style, addr, user, pwd); e == nil {
		for s.run {
			if s.requestUrl(c, chRequest, duk, id, name) {
				break
			}
		}
	} else {
		if log.Error != nil {
			log.Error.Println(e)
		}
	}
	if log.Info != nil {
		log.Info.Println("goroutine end", i)
	}
	ch <- 1
}
func (s *Single) requestUrl(c *http.Client, chRequest chan int, duk *Duktape, id, name string) (exit bool) {
	i := <-chRequest
	url := duk.GetUrl(id, i)
	if url == "" {
		exit = true
		return
	}
	if log.Trace != nil {
		log.Trace.Println("request", url)
	}

	r, e := c.Get(url)
	if e != nil {
		if log.Error != nil {
			log.Error.Println(e)
		}
		return
	}
	var b []byte
	b, e = ioutil.ReadAll(r.Body)
	if e != nil {
		if log.Warn != nil {
			log.Warn.Println(e)
		}
		return
	}
	msg := kStrings.BytesToString(b)

	var nodes []*Node
	nodes, e = duk.Analyze(id, msg)
	if e != nil {
		if log.Warn != nil {
			log.Warn.Println(e)
		}
		return
	}
	if len(nodes) < 1 {
		return
	}

	InsertNodes(nodes)
	return
}
