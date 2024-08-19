package ginx

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func WrapBody[Req any](
	fn func(ctx *gin.Context, req Req) (Result, error),
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req Req
		if err := ctx.Bind(&req); err != nil {
			zap.L().Error("输入参数错误", zap.Error(err))
			return
		}

		zap.L().Debug("输入参数", zap.Any("req", req))

		res, err := fn(ctx, req)
		if err != nil {
			zap.L().Error("执行业务逻辑失败", zap.Error(err))
		}

		ctx.JSON(http.StatusOK, res)
	}
}
