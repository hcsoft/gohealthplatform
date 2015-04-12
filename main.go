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
	"log"
	"fmt"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	m := getMartini()
	store := sessions.NewCookieStore([]byte("gohealthplatform_secret"))
	m.Use(sessions.Sessions("gohealthplatform_session", store))
	m.Use(render.Renderer())
	//配置文件
	configFile := flag.String("configfile", "config.ini", "配置文件")
	inicfg, _ := config.ReadDefault(*configFile)
	m.Map(inicfg)
	//数据库
	dbsession, err := mgo.Dial("localhost")
	util.CheckErr(err)
	dbsession.SetMode(mgo.Monotonic, true)
	m.Map(dbsession)
	//TODO 缓存
	//	cache := make(map[string]interface{})
	//	district, _ := inicfg.String("base", "district")
	//	cache["district"] = auth.GetDistrict(session, district)
	//	m.Map(cache)
	//TODO 登录
	//	m.Any("/login", auth.Login)
	//	m.Get("/logout", auth.Logout)
	m.Get("/", index)

	//	m.Group("/admin", admin.Router, auth.Auth)
	m.Run()
}

func index(dbsession *mgo.Session, r render.Render, req *http.Request, inicfg *config.Config , log *log.Logger) {
	ret := make(map[string]interface{})

//	c := dbsession.DB("healthplatform").C("test")
	test := bson.M{"a": 222, "b": "333"}
	fmt.Println(test)
	r.HTML(200, "index", ret)
}

func getMartini() *martini.ClassicMartini {
	//与martini.Classic相比,去掉了martini.Logger() handler  去掉讨厌的http请求日志
	base := martini.New()
	router := martini.NewRouter()
	base.Use(martini.Recovery())
	base.Use(martini.Static("static"))
	base.MapTo(router, (*martini.Routes)(nil))
	base.Action(router.Handle)
	return &martini.ClassicMartini{base, router}
}
