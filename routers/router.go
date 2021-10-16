package beegos

import (
	"goProject/controllers"
	"goProject/models/frontend"

	"github.com/beego/beego/v2/server/web/context"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Options("/uploadFile", func(ctx *context.Context) { ctx.Output.SetStatus(200) })

	beego.Router("/", &controllers.MainController{})

	beego.Get("/getDataNum", getDataNum)
	beego.Get("/getDataList", getDataList)

	beego.Post("/download", download)
	beego.Post("/getDownloadURL", getDownloadURL)
	beego.Post("/uploadFile", uploadFile)

	beego.Delete("/delete", delete)
}

func delete(ctx *context.Context) {
	frontend.Delete(ctx)
}

/*获取当前账户在存储服务器上存储了的数据集的值*/
func getDataNum(ctx *context.Context) {
	frontend.GetDataNum(ctx)
}

func uploadFile(ctx *context.Context) {
	frontend.UploadFile(ctx)
}

func getDataList(ctx *context.Context) {
	frontend.GetDataList(ctx)
}

func getDownloadURL(ctx *context.Context) {
	frontend.GetDownloadURL(ctx)
}

func download(ctx *context.Context) {
	frontend.Download(ctx)
}
