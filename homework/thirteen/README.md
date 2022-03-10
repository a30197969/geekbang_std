### 毕业项目：对当下自己项目中的业务，进行一个微服务改造，需要考虑如下技术点：

微服务架构（BFF、Service、Admin、Job、Task 分模块）

API 设计（包括 API 定义、错误码规范、Error 的使用）

gRPC 的使用

Go 项目工程化（项目结构、DI、代码分层、ORM 框架）

并发的使用（errgroup 的并行链路请求）

微服务中间件的使用（ELK、Opentracing、Prometheus、Kafka）

缓存的使用优化（一致性处理、Pipeline 优化）

###个人理解：

把整套微服务分成多个go项目，web服务可以使用beego、gin框架，接口服务等使用kratos框架

使用Nginx upstream负载均衡到BFF层，入口为BFF层，专门对数据进行组装和分发，使用errorgroup并行请求

BFF后的多个微服务之间使用gRPC进行接口通信

数据库引入gorm框架

拿到的数据使用redis cluster缓存，防止redis击穿、穿透，可随机设置缓存时间或者使用数据库直接查询，再配合消息队列归并查询记录缓存

并发写入过高时使用消息队列kafak，进行去峰处理，降低高并发时的写库压力


#### kratos框架配合使用gorm的DEMO如下链接：https://github.com/a30197969/kratos_study

#### beego 处理web端服务的DEMO如下链接：https://github.com/a30197969/beego_study

