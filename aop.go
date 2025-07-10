package aop

import (
	"context"
	"sync"
)

// InterceptorFunc 定义拦截器内部方法
type InterceptorFunc func(ctx context.Context, req interface{}) (rsp interface{}, err error)

// Interceptor 定义拦截器类型
type Interceptor func(ctx context.Context, req interface{}, next InterceptorFunc) (rsp interface{}, err error)

var (
	interceptors []Interceptor
	mu           sync.Mutex
)

// Register 注册拦截器
func Register(interceptor Interceptor) {
	mu.Lock()
	defer mu.Unlock()
	interceptors = append(interceptors, interceptor)
}

// Execute 执行拦截器链
func Execute(ctx context.Context, req interface{}, next InterceptorFunc) (interface{}, error) {
	var chain InterceptorFunc
	idx := 0
	chain = func(ctx context.Context, req interface{}) (rsp interface{}, err error) {
		if idx < len(interceptors) {
			current := interceptors[idx]
			idx++
			return current(ctx, req, chain)
		}
		return next(ctx, req)
	}
	return chain(ctx, req)
}
