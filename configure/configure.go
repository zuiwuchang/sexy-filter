package configure

import (
	"fmt"
	kStrings "github.com/zuiwuchang/king-go/strings"
	"io/ioutil"
	"os"
	"path/filepath"
	"sexy-filter/log"
	"sort"
	"strings"
)

//語言定義
type Lang struct {
	Key   string
	Name  string
	Index int
}
type configure struct {
	//可用 語言包
	keyLangs  map[string]*Lang
	arrsLangs []string

	//使用的 語言包
	lang string

	//可用 界面風格
	arrsStyle []string
	//使用的 界面 風格
	style int
}

var g_Configure *configure

func Init(appPath string) {
	cnf := &configure{
		keyLangs:  make(map[string]*Lang),
		arrsStyle: []string{"Material Dark", "Material Light", "Universal Dark", "Universal Light"},
	}
	if str, e := filepath.Abs(appPath); e == nil {
		appPath = str
	} else {
		if log.Warn != nil {
			log.Warn.Println(e)
		}
	}
	dir := filepath.Dir(appPath) + "/locale"
	log.Trace.Println(dir)
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if path == dir {
			return nil
		}
		if info.IsDir() {
			return filepath.SkipDir
		}
		if !strings.HasSuffix(path, ".qm") {
			return nil
		}
		key := filepath.Base(path)
		key = key[:len(key)-3]
		if strings.ToLower(key) == "zh_cn" {
			if log.Fault != nil {
				log.Fault.Println("unknow locale :", key)
			}
			return nil
		} else if strings.ToLower(key) == "zh_cn" {
			if log.Warn != nil {
				log.Warn.Println("Unable to cover zh_TW")
			}
			return nil
		}

		var name string
		b, e := ioutil.ReadFile(dir + "/" + key + ".name")
		if e == nil {
			name = kStrings.BytesToString(b)
			if len(([]rune)(name)) > 10 {
				if log.Warn != nil {
					log.Warn.Printf("locale %v's name is too long (%v)", key, name)
				}
				name = ""
			}
		} else {
			if log.Warn != nil {
				log.Warn.Println(e)
			}
		}

		cnf.keyLangs[key] = &Lang{
			Key:  key,
			Name: name,
		}
		return nil
	})

	arrs := make([]string, len(cnf.keyLangs)+1)
	arrs[0] = "zh_TW"
	i := 1
	for k := range cnf.keyLangs {
		arrs[i] = k
		i++
	}
	sort.Sort(sort.StringSlice(arrs[1:]))
	for i, key := range arrs[1:] {
		cnf.keyLangs[key].Index = i + 1
	}
	cnf.arrsLangs = arrs

	cnf.keyLangs["zh_TW"] = &Lang{
		Key:   "zh_TW",
		Name:  "中文",
		Index: 0,
	}

	if log.Trace != nil {
		log.Trace.Println(arrs)
	}
	g_Configure = cnf
}
func InitLocale(name string) {
	cnf := g_Configure
	if _, ok := cnf.keyLangs[name]; ok {
		cnf.lang = name
		return
	}

	cnf.lang = "zh_TW"
	if log.Warn != nil {
		log.Warn.Println("unknow locale", name)
	}
}
func InitStyle(style int) {
	cnf := g_Configure
	if style < 0 || style > len(cnf.arrsStyle) {
		if log.Error != nil {
			log.Error.Println("unknow style", style)
		}
		return
	}
	cnf.style = style
}

//返回 風格 列表
func GetStyles() []string {
	return g_Configure.arrsStyle
}

//返回 當前 風格
func GetStyle() int {
	return g_Configure.style
}

//返回 當前 語言
func GetLocale() string {
	return g_Configure.lang
}

//返回 語言 名稱
func GetLocaleName(key string) string {
	if lang, ok := g_Configure.keyLangs[key]; ok {
		return lang.Name
	}
	return ""
}

//返回 語言 列表
func GetLocales() []string {
	return g_Configure.arrsLangs
}

//返回 當前 選擇的語言
func GetLocaleIndex() int {
	keys := g_Configure.keyLangs
	return keys[g_Configure.lang].Index
}

//驗證 語言 風格 設置
func Verification2(local, style int) error {
	cnf := g_Configure
	if local < 0 || local > len(cnf.arrsLangs) {
		return fmt.Errorf("unknow locale (%v)", local)
	} else if style < 0 || style > len(cnf.arrsStyle) {
		return fmt.Errorf("unknow style (%v)", local)
	}
	return nil
}

//設置 語言 風格 設置
func Set2(local, style int) {
	cnf := g_Configure
	cnf.lang = cnf.arrsLangs[local]
	cnf.style = style
}
