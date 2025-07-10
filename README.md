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
    return func(ctx context.Context, req interface{}, next aop.InterceptorFunc) (interface{}, error) {
        defer func() {
            if r := recover(); r != nil {
                fmt.Println("Recovered from panic:", r)
            }
        }()
        return next(ctx, req)
    }
}
```


### Execute

#### 不合理的使用
```go
func main() {
	aop.Execute(context.Background(), run)
}

func run() error {
	painc("here is your code")
}

```

#### 合理的使用
```go
// 通过中间件方式引入拦截器插件
http.Handle("/api/xxx", AOPMiddleware(myHandler))
```