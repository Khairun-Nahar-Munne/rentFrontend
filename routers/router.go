package routers

import (
	"rentFrontend/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/property-details", &controllers.PropertyController{})
	beego.Router("/api/list/fetch", &controllers.PropertyClientController{})
	beego.Router("/api/details/fetch", &controllers.PropertyDetailsClientController{})
	beego.Router("/api/list/fetch/*", &controllers.PropertySpecificController{})
	beego.Router("/*", &controllers.PropertyViewController{})
}
