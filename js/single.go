package js

import (
	"os"
	"path/filepath"
	"runtime"
	"sexy-filter/log"
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
func (s *Single) Start(pos int) int {

	return s.status
}
func (s *Single) Stop() int {

	return s.status
}
