package viewmodels

import (
	"github.com/alex19861108/burj/burj_center/iris/proto"
	"github.com/kataras/iris/context"
)

type Node struct {
	proto.Node
}

func (n Node) IsValid() bool {
	/* do some checks and return true if it's valid... */
	return n.Id != ""
}

func (n Node) Dispatch(ctx context.Context) {
	if !n.IsValid() {
		ctx.NotFound()
		return
	}
	ctx.JSON(n, context.JSON{Indent: " "})
}
