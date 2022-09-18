# 羊了个羊

> 此程序完全免费并仅用作学习与交流，使用过程中出现任何问题本人概不负责。

## 使用

```shell
# 请求10000次接口，每个请求分配一个协程
./ylgy-macos-arm64
Usage of ./ylgy-macos-arm64:
  -count int
        number of requests (default 1000)
  -threshold int
        how many requests is a coroutine responsible for， 1 means one coroutine per request (default 1)
  -token string
        request token
  -v    show version information
```

```shell
# 请求10000次接口，每个请求分配一个协程
./ylgy -token "xxxx.xxxx.xxxx" -count 10000 -threshold 1
```

```shell
# 请求10000次接口，每10个请求分配一个协程
./ylgy -token "xxxx.xxxx.xxxx" -count 10000 -threshold 10
```

