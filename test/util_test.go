package test

import (
	"testing"
	"github.com/smartystreets/goconvey/convey"
	"codegen/util"
)

func TestToCamelBak(t *testing.T) {
	convey.Convey("下划线转驼峰测试", t, func(){
		convey.Convey("_ors_user to OrsUser ", func(){
			convey.So(util.ToCamelBak("_ors_user"), convey.ShouldEqual, "OrsUser")
		})

		convey.Convey("ors_user_user", func(){
			convey.So(util.ToCamelBak("ors_user_user"), convey.ShouldEqual, "OrsUserUser")
		})

		convey.Convey("ors_user_user_", func(){
			convey.So(util.ToCamelBak("ors_user_user_"), convey.ShouldEqual, "OrsUserUser")
		})
	})
}
