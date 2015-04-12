package main

import (
	"flag"
	"github.com/go-martini/martini"
//	"github.com/hcsoft/gohealthplatform/auth"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	_ "github.com/mattn/go-adodb"
	"net/http"
	"github.com/hcsoft/gohealthplatform/util"
	"github.com/larspensjo/config"
	//	"log"
		"fmt"
	 mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	m := martini.Classic()
	store := sessions.NewCookieStore([]byte("secret123"))
	m.Use(sessions.Sessions("gohealthplatform_session", store))
	m.Use(render.Renderer())
	//配置文件
	configFile := flag.String("configfile", "config.ini", "配置文件")
	inicfg, _ := config.ReadDefault(*configFile)
	m.Map(inicfg)
	//数据库
	session, err := mgo.Dial("localhost")
	util.CheckErr(err)

	m.Map(session)
	//设置日志

	//缓存
//	cache := make(map[string]interface{})
//	district, _ := inicfg.String("base", "district")
//	cache["district"] = auth.GetDistrict(session, district)
//	m.Map(cache)

//	m.Any("/login", auth.Login)
//	m.Get("/logout", auth.Logout)
	m.Get("/", index)
//	m.Get("/cats/:catid", helpmaker.Cats)
//	m.Get("/pages/:id", helpmaker.Pages)

	//静态内容
	m.Use(martini.Static("static"))
	//需要权限的内容
//	m.Group("/admin", admin.Router, auth.Auth)
	m.Run()
}

func index( session *mgo.Session, r render.Render, req *http.Request, inicfg *config.Config) {
	ret := make(map[string]interface{})
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("healthplatform").C("test")
	test:=bson.D{{"a", 1}, {"b", true}}
	err := c.Insert(test)
	util.CheckErr(err)
	str ,_ :=bson.Marshal(test)
	fmt.Println(string(str))
	//	catid := req.FormValue("catid")
	//	ret["cats"] = helpmaker.GetCats(catid, db)
	r.HTML(200, "index", ret)
}
