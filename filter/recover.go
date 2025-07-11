package filter

import (
	"context"
	"fmt"
	"github.com/baisiyi/aop"
)

// 通过init自动组册
func init() {
	aop.Register(RecoverFilter())
}

// RecoverFilter filter实现
func RecoverFilter() aop.Interceptor {
	return func(ctx context.Context, req interface{}, next aop.InterceptorFunc) (interface{}, error) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered from panic:", r)
			}
		}()
		return next(ctx, req)
	}
}
