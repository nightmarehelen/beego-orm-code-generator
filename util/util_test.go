package util

import (
	"testing"
	"github.com/smartystreets/goconvey/convey"
)

func TestToCamelBak(t *testing.T) {
	convey.Convey("下划线转驼峰测试", t, func(){
		convey.Convey("_ors_user to OrsUser ", func(){
			convey.So(ToCamelBak("_ors_user"), convey.ShouldEqual, "OrsUser")
		})

		convey.Convey("ors_user_user", func(){
			convey.So(ToCamelBak("ors_user_user"), convey.ShouldEqual, "OrsUserUser")
		})

		convey.Convey("ors_user_user_", func(){
			convey.So(ToCamelBak("ors_user_user_"), convey.ShouldEqual, "OrsUserUser")
		})
	})
}
