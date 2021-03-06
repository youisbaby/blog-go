package admin

import (
	"blog/fox/db"
	"net/url"
	"blog/fox"
	"strings"
	"blog/model"
	"fmt"
	"blog/service/conf"
)
//站点配置
type Site struct {
	Config map[string]string
}
//初始化 站点配置
func NewSiteService() *Site {
	return new(Site)
}
//列表
func (t *Site)Query() (*db.Paginator, error) {
	return NewTypeService().Query(conf.SITE_ID)
}
//更新
func (t *Site)Update(form url.Values) (bool, error) {
	if len(form) < 1 {
		return false,fox.NewError("站点信息 不能为空")
	}

	o := db.NewDb()
	for key, v := range form {
		key = strings.TrimSpace(key)
		val := strings.TrimSpace(v[0])
		if key == "SITE_NAME" && len(val) < 1 {
			return false,fox.NewError("站点名称 不能为空")
		}
		if val != "" {
			//更新
			mod := model.NewType()
			mod.Content = val
			num, err := o.Where("type_id=? and mark=? ", conf.SITE_ID, key).Update(mod)
			if err != nil {
				return false, err
			}
			fmt.Println("更新 " + key + "=>" + val, num)
		} else {
			fmt.Println(key + "值为空 跳过更新")
		}

	}
	return true, nil
}
//获取站点配置
func (c *Site)SiteConfig() map[string]string {
	tp := make([]model.Type, 0)
	o := db.NewDb()
	tps := make(map[string]string)
	err := o.Where("type_id=?", conf.SITE_ID).Find(&tp)
	if err != nil {
		fmt.Println(err)
		return tps
	}
	for _, v := range tp {
		if v.Mark != "" {
			tps[v.Mark] = v.Content
		}
	}
	return tps
}
//站点配置赋值
func (c *Site)SetSiteConfig() {
	c.Config = c.SiteConfig()
}
//获取信息
func (c *Site)GetString(key string) string {
	return c.Config[key]
}