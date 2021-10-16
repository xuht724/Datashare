package log

import (
	"log"

	"github.com/beego/beego/v2/core/logs"
)

var l *log.Logger

func Init() {
	logs.SetLogger("console")
	l = logs.GetLogger()
}
