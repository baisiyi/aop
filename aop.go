package aop

import (
	"context"
	"sync"
)

// 定义拦截器类型
type Interceptor func(ctx context.Context, next func(ctx context.Context) error) error

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
func Execute(ctx context.Context, target func(ctx context.Context) error) error {
	var chain func(context.Context) error
	idx := 0
	chain = func(c context.Context) error {
		if idx < len(interceptors) {
			current := interceptors[idx]
			idx++
			return current(c, chain)
		}
		return target(c)
	}
	return chain(ctx)
}
