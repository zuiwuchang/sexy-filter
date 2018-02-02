package js

type Node struct {
	//url 地址
	Url string `xorm:"unique"`
	//標題
	Title string `xorm:"index"`
	//插件 Id
	PluginsId string `xorm:"index"`
	//插件 名稱
	PluginsName string `xorm:"index"`
}
