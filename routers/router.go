package routers

import (
	"rentFrontend/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/property-details", &controllers.PropertyController{})
}
