# 基于Goframe的快开框架

## 编码顺序

```
API(controller) ----> Logic ------> ctrl调用Logic -> 注册控制器
↓                     |             |                 |
api/users/v1/users.go |             |                 |
                      ↓             |                 |
inte.../cont.../users/users_v1_*.go |                 |
                                    ↓                 |
   inte.../cont.../users/users_new.go                 |
                                                      ↓
                                    internal/cmd/cmd.go
```

## 注意事项

在controller层中，需要在`xxx_new.go`中加入对应的logic依赖方法，且需要调用logic层的New方法实例化对象

## 接口文档

因为基于Goframe框架，所以在使用`gf run main.go`之后会自动生成swagger页面，前往查看接口定义即可

如何安装Goframe：https://goframe.org/quick/install

[APIFox接口文档](https://o36k1oz5or.apifox.cn)