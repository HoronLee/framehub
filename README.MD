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