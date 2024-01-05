```
❯ go run .
panic: failed to get API group resources: unable to retrieve the complete list of server APIs: data.packaging.carvel.dev/__internal: the server could not find the requested resource

goroutine 1 [running]:
k8s.io/apimachinery/pkg/util/runtime.Must(...)
        /Users/*****/go/pkg/mod/k8s.io/apimachinery@v0.29.0/pkg/util/runtime/runtime.go:175
main.main()
        /Users/*********workspace/list-packages/main.go:47 +0x56b
exit status 2
```
