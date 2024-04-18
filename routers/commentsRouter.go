package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["go-olin/controllers:ObjectController"] = append(beego.GlobalControllerRouter["go-olin/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/soal1`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-olin/controllers:ObjectController"] = append(beego.GlobalControllerRouter["go-olin/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Post2",
            Router: `/soal2`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-olin/controllers:ObjectController"] = append(beego.GlobalControllerRouter["go-olin/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Post3",
            Router: `/soal3`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
