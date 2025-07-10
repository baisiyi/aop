package filter

import (
	"context"
	"github.com/baisiyi/aop"
)

// 通过init自动组册
func init() {
	aop.Register(TimeOutFilter())
}

// TimeOutFilter filter实现
func TimeOutFilter() aop.Interceptor {
	return func(ctx context.Context, req interface{}, next aop.InterceptorFunc) (interface{}, error) {
		
		return next(ctx, req)
	}
}
