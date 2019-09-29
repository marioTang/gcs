package admin

import (
	"fmt"
	"gcs/utils/base"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gcache"
	"github.com/gogf/gf/os/glog"
)

type TmpAction struct {
	base.BaseRouter
}

func (o *TmpAction) Test(r *ghttp.Request) {
	glog.Println("#### /tmp/test")
	r.Response.Write("test")
}

func (o *TmpAction) Access(r *ghttp.Request) {
	glog.Println("#### /tmp/access")
	r.Response.Writeln("请在运行终端查看日志输出")
}

func (o *TmpAction) Error(r *ghttp.Request) {
	glog.Println("#### /tmp/error")
	panic("异常信息")
}

func (o *TmpAction) Cache(r *ghttp.Request) {
	glog.Println("#### /tmp/cache")
	gcache.Set("a", "b", 0)
}

func (o *TmpAction) Mysql(r *ghttp.Request) {
	glog.Println("#### /tmp/mysql")
	db := g.Database()
	if r, err := db.Table("test1").Where("id=?", 1).One(); err == nil {
		fmt.Printf("goods    id: %d\n", r["id"].Int())
		fmt.Printf("goods task_id: %s\n", r["task_id"].Int8())
		fmt.Printf("goods name: %.2f\n", r["name"].String())
	} else {
		glog.Error(err)
	}
}
