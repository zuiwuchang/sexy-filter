package js

import (
	kStrings "github.com/zuiwuchang/king-go/strings"
	"io/ioutil"
	"os"
	"path/filepath"
	"sexy-filter/log"
	"strings"
)

var WorkDir string

func init() {
	abs, e := filepath.Abs(os.Args[0])
	if e != nil {
		if log.Warn != nil {
			log.Warn.Println(e)
		}
		WorkDir = filepath.Dir(os.Args[0])
		return
	}
	WorkDir = filepath.Dir(abs)

}

//返回 所有 的 插件 檔案
func GetPluginsFiles() (rs []string) {
	rs = make([]string, 0, 20)

	dir := WorkDir + "/plugins-js"
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if path == dir {
			return nil
		} else if info.IsDir() {
			return filepath.SkipDir
		}
		name := info.Name()
		if len(name) > 3 && strings.HasSuffix(name, ".js") {
			rs = append(rs, name[:len(name)-3])
		}
		return nil
	})
	return
}

//返回 所有 的 測試 檔案
func GetTestFiles() (rs []string) {
	rs = make([]string, 0, 20)

	dir := WorkDir + "/test-js"
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if path == dir {
			return nil
		} else if info.IsDir() {
			return filepath.SkipDir
		}
		name := info.Name()
		rs = append(rs, name)
		return nil
	})
	return
}

//測試 插件 是否可以 正確 解析 檔案
func TestFile(jsFile, tFile string) (string, error) {
	return testFile(
		WorkDir+"/plugins-js/"+jsFile+".js",
		WorkDir+"/test-js/"+tFile,
	)
}
func testFile(jsPath, testPath string) (msg string, e error) {
	//讀取 檔案
	b, ef := ioutil.ReadFile(testPath)
	if ef != nil {
		e = ef
		if log.Warn != nil {
			log.Warn.Println(e)
		}
		return
	}

	//插件 js 環境
	duk := NewDuktape()
	defer duk.Close()
	defer duk.duk.DumpContextStdout()

	//加載 插件
	e = duk.LoadPluginsJs(jsPath)
	if e != nil {
		if log.Warn != nil {
			log.Warn.Println(e)
		}
		return
	}

	//解析 數據
	e = duk.Analyze("", kStrings.BytesToString(b))
	if e != nil {
		if log.Warn != nil {
			log.Warn.Println(e)
		}
		return
	}

	duk.DisplayPlugins()
	return
}
