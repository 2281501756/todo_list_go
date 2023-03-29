# Todo List 备忘录
这是一个使用 Gin+Gorm,基于REASTful API实现的一个备忘录项目

主要是熟悉go开发web应用以及react入门

## 项目主要依赖：
Golang V1.20
- Gin
- Gorm
- mysql
- redis
- ini
- jwt-go
- logrus
- go-swagger


## 项目结构

```
TodoList/
├── api
├── cache
├── conf
├── middleware
├── model
├── pkg
│  ├── e
│  └──  util
├── routes
├── serializer
├── web
└── service
```
- api 后端接口
- cache redis缓存
- conf 文件配置
- middleware 中间件
- model 数据库设计
- pkg 工具包
- routes 路由管理
- serializer 返回结果处理
- web 前端页面
- service 接口逻辑

## 运行项目
下载依赖
```shell
go mod tity

```
运行项目
```shell
go run main.go
```