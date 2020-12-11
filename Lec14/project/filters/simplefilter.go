package filters

import (
	"fmt"
	"time"

	context "github.com/astaxie/beego/server/web/context"
)

var LogManager = func(ctx *context.Context) {
	fmt.Println("IP :: " + ctx.Request.RemoteAddr + ", Time :: " + time.Now().Format(time.RFC850))
}
