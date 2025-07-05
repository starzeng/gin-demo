# gin-demo

### 一个基于 Gin 的模块化单体项目示例，支持按业务模块分层（Controller / Service / Repository / Model），适用于中小型后端项目开发。

## 项目结构
```go
gin-demo/
├── main.go                      # 启动入口（如 main.go）
│
├── internal/                    # 内部业务模块，不能被外部调用
│   ├── user/                    # user 模块
│   │   ├── controller/          # 路由处理器（HTTP 层）
│   │   ├── service/             # 业务逻辑层（定义接口+实现）
│   │   ├── repository/          # 数据访问层（DAO 层）
│   │   └── model/               # 数据结构（DB模型、DTO）
│   │
│   ├── order/                   # order 模块（结构与 user 一致）
│   ├── product/
│   └── ...
│
├── docs/                        # OpenAPI（Swagger）文档定义（可选）
│
├── common/                      # 通用模块（响应封装、错误定义等）
│   ├── response/
│   ├── errors/
│   ├── constants/
│   └── base_context.go
│
├── config/                      # 配置文件加载
│   ├── config.yaml
│   └── config.go
│
├── middleware/                  # gin中间件，如日志、JWT、跨域等
│   └── auth.go
│
├── pkg/                         # 外部依赖的封装（redis、gorm等）
│   ├── db/
│   ├── redis/
│   └── logger/
│
├── router/                      # 路由统一注册入口
│   └── router.go
│
├── utils/                       # 工具函数（通用的、无状态）
│   ├── crypto.go
│   ├── time.go
│   └── string.go
│
├── go.mod
└── README.md

```