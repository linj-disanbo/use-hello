# 多版本

## 定义

```
# go.mod
module github.com/linj-disanbo/hello/v2


```

## 使用

```
import helloV2 "github.com/linj-disanbo/hello/v2"
# 版本出现在位置
# site/who/repo/v?/sub-dir/sub-dir
```

## 下载

```
go get github.com/my/mod/v2@v2.0.1
```

## 语义版本

v2.3.4
 1. 2主版本，有不兼容
 1. 3次版本，有特性
 1. 4patch，fix bug