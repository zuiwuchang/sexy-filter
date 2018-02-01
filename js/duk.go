package js

import (
	"errors"
	"github.com/axgle/mahonia"
	"gopkg.in/olebedev/go-duktape.v3"
	"os"
	"sexy-filter/log"
	"strings"
)

type Duktape struct {
	duk *duktape.Context
}

//銷毀 js 環境
func (d *Duktape) Close() {
	if d.duk != nil {
		d.duk.DestroyHeap()
		d.duk = nil
	}
}

//初始化 插件 系統
func (d *Duktape) initPlugins() {
	duk := d.duk

	e := duk.PevalString(`
(function(){
	"use strict";

	//所有 已加載的 插件
	var pPlugins = {};
	//默認 插件
	var pDefault = null;
	//加載 一個 插件
	var loadPlugins = function(obj){
		//驗證 id
		var id = obj.Id();
		if(!id){
			throw "bad plugins id";
		}
		//驗證 存在
		if(pPlugins[id]){
			throw "plugins already exists";
		}

		//驗證 名稱
		var name = obj.Name();
		if(!name){
			throw "bad plugins name";
		}

		//驗證 分析 函數
		if(obj.Analyze("Analyze") != "Analyze"){
			throw "bad plugins Analyze func";
		}


		//加載 插件
		pPlugins[id] = obj;
		if(!pDefault){
			pDefault = obj;
		}
	};
	//返回 插件 實例
	var getPlugins = function(id){
		if(id){
			return pPlugins[id];
		}
		return pDefault;
	}
	return {
		//加載 一個 插件
		LoadPlugins:function(obj){
			return loadPlugins(obj);
		},
		//打印 所有 插件信息
		DisplayPlugins:function(){
			for(var k in pPlugins){
				var obj = pPlugins[k];			
				alert(k + "\t" +obj.Name());
			}
		},
		//使用 指定 插件 解析 str
		//如果 id 爲空 false 使用 (第一個被加載的 插件)
		Analyze:function(id,str){
			var obj = getPlugins(id);
			if(!obj){
				throw "unknow plugins" + id;
			}
			return obj.Analyze(str);
		},
		//返回 插件 名稱
		GetPluginsName:function(id){
			var obj = getPlugins(id);
			if(!obj){
				return "unknow";
			}
			return obj.Name();
		},
		//返回 插件 id
		GetPluginsId:function(id){
			var obj = getPlugins(id);
			if(!obj){
				return "unknow";
			}
			return obj.Id();
		},
		//返回 url 地址
		GetUrl(id,i){
			var obj = getPlugins(id);
			if(!obj){
				return "";
			}
			var url = obj.GetUrl(i);
			if(!url){
				return "";
			}
			return url;
		}
	};
})();
`)
	if e != nil {
		if log.Fault != nil {
			log.Fault.Println(e)
		}
		os.Exit(1)
	}
}

//初始化 duk 環境
func NewDuktape() (duk *Duktape) {
	duk = &Duktape{
		duk: duktape.New(),
	}
	duk.initTools()
	duk.initPlugins()
	return
}

//初始化 輔助 函數
func (d *Duktape) initTools() {
	duk := d.duk
	duk.PushGlobalObject()

	duk.PushGoFunction(func(c *duktape.Context) int {
		if !c.IsString(-1) {
			return 1
		}
		str := c.ToString(-1)

		dec := mahonia.NewDecoder("gbk")
		str = dec.ConvertString(str)
		c.Pop()
		c.PushString(str)
		return 1
	})
	duk.PutPropString(-2, "GBKtoUTF8")
	duk.Pop()
}

//將 一個 js 作爲 插件 加載
func (d *Duktape) LoadPluginsJs(jsPath string) (e error) {
	duk := d.duk

	//push 函數
	duk.GetPropString(0, "LoadPlugins")

	//push 參數
	//加載 js
	e = duk.PevalFile(jsPath)
	if e != nil {
		duk.Pop2()
		return
	}

	//call
	if duk.Pcall(1) != 0 {
		e = errors.New(duk.SafeToString(-1))

		duk.Pop()
		return
	}
	duk.Pop()
	return
}

//打印所有 插件 信息
func (d *Duktape) DisplayPlugins() {
	duk := d.duk
	duk.GetPropString(0, "DisplayPlugins")
	if duk.Pcall(0) != 0 {
		if log.Warn != nil {
			log.Warn.Println(duk.SafeToString(-1))
		}
	}
	duk.Pop()
}

//返回 插件名稱
func (d *Duktape) GetPluginsName(id string) (name string) {
	duk := d.duk
	duk.GetPropString(0, "GetPluginsName")
	duk.PushString(id)
	duk.Pcall(1)
	name = duk.SafeToString(-1)
	duk.Pop()
	return
}

//返回 插件id
func (d *Duktape) GetPluginsId(id string) (name string) {
	duk := d.duk
	duk.GetPropString(0, "GetPluginsId")
	duk.PushString(id)
	duk.Pcall(1)
	name = duk.SafeToString(-1)
	duk.Pop()
	return
}

//使用 指定 插件 解析 字符串
func (d *Duktape) Analyze(id, str string) (nodes []*Node, e error) {
	duk := d.duk
	duk.GetPropString(0, "Analyze")
	duk.PushString(id)
	duk.PushString(str)
	if duk.Pcall(2) != 0 {
		e = errors.New(duk.SafeToString(-1))
		duk.Pop()
		return
	}
	if !duk.IsArray(-1) {
		duk.Pop()
		return
	}

	id = d.GetPluginsId(id)
	name := d.GetPluginsName(id)

	nodes = make([]*Node, 0, 20)
	for i := 0; i < duk.GetLength(-1); i++ {
		duk.GetPropIndex(-1, uint(i))
		{
			if duk.IsObject(-1) {
				node := &Node{
					PluginsId:   id,
					PluginsName: name,
				}
				duk.GetPropString(-1, "Title")
				node.Title = duk.SafeToString(-1)
				duk.GetPropString(-2, "Url")
				node.Url = duk.SafeToString(-1)

				duk.Pop2()

				nodes = append(nodes, node)
			}
		}
		duk.Pop()
	}
	duk.Pop()
	return
}

//返回 url 地址
func (d *Duktape) GetUrl(id string, i int) (url string) {
	duk := d.duk
	duk.GetPropString(0, "GetUrl")
	duk.PushString(id)
	duk.PushInt(i)
	if duk.Pcall(2) != 0 {
		if log.Error != nil {
			log.Error.Println(duk.SafeToString(-1))
		}
		duk.Pop()
		return
	}
	if !duk.IsString(-1) {
		if log.Error != nil {
			log.Error.Println("unknow plugins rs :", duk.SafeToString(-1))
		}
		duk.Pop()
		return
	}
	url = strings.TrimSpace(duk.ToString(-1))
	duk.Pop()
	return
}
