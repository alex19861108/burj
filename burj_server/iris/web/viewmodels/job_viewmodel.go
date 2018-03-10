package viewmodels

import (
	"github.com/alex19861108/burj/burj_center/iris/proto"
	"github.com/kataras/iris/context"
)

type Job struct {
	proto.Job
}

func (j Job) IsValid() bool {
	/* do some checks and return true if it's valid... */
	return j.Id != ""
}

func (j Job) Dispatch(ctx context.Context) {
	if !j.IsValid() {
		ctx.NotFound()
		return
	}
	ctx.JSON(j, context.JSON{Indent: " "})
}
