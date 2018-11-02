package test

import (
	"testing"
	"github.com/smartystreets/goconvey/convey"
	"codegen/generated/models/apimodel"
	"time"
	"codegen/generated/service/apisrv"
)

func TestGetColumns2Update(t *testing.T){
	obj := apimodel.LevelItem{}

	obj.Type = "1"
	now := time.Now()
	obj.Datetime = &now
	srv := apisrv.BaseService{}
	convey.Convey("根据结构体的属性值是否为默认值来确定那些属性需要更新", t, func() {
		convey.Convey(srv.GetColumns2Update(obj), convey.ShouldContain("Type", "Datetime"))
	})
}
