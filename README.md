# aop
> 一个最小入侵业务的aop实现

## 使用方式

### Register
```go
// 通过init自动组册
func init() {
	aop.Register(RecoverFilter())
}

// RecoverFilter filter实现
func RecoverFilter() aop.Interceptor {
	return func(ctx context.Context, next func(ctx context.Context) error) error {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered from panic:", r)
			}
		}()
		return next(ctx)
	}
}
```


### Execute

```go
package main

import (
    "github.com/baisiyi/aop"
    _ "github.com/baisiyi/aop/filter"
)


func main() {
	aop.Execute(context.Background(), run)
}

func run() error {
	painc("here is your code")
}

```

#### 更小粒度的代理每个接口
```go

http.Handle("/api/xxx", AOPMiddleware(myHandler))

func AOPMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        err := aop.Execute(r.Context(), func(ctx context.Context) error {
            next.ServeHTTP(w, r)
            return nil
		})
        if err != nil {
        // 错误处理
        }
    })
}
```