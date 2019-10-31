# golang-Gin

**1.Gin：go-web框架**

    1.前置知识：
        ·了解go基本语法
        ·了解go协程基本知识
    2.优势：
        ·简单原则
        ·并发高
        ·分配内存少

**2.课程内容：**
    
    1.gin基础知识：
        1.安装gin及快速开始            D:\go-20191030\golang-Gin\start
        2.请求路由：
            1.设置多种请求类型：         D:\go-20191030\golang-Gin\router_type\main.go
            2.绑定静态文件夹:              D:\go-20191030\golang-Gin\router_static
            3.参数作为uri:                  D:\go-20191030\golang-Gin\router_uri
            4.泛绑定：类似于nginx的请求匹配:D:\go-20191030\golang-Gin\router_generic
        3.获取请求参数：
            1.获取get请求参数         D:\go-20191030\golang-Gin\param_get
            2.获取post请求参数        D:\go-20191030\golang-Gin\param_body
            3.获取body值              D:\go-20191030\golang-Gin\param_body
            4.获取参数bind绑定结构     D:\go-20191030\golang-Gin\param_struct
        4.验证请求参数：//请求参数验证：文档：https://godoc.org/gopkg.in/go-playground/validator.v8
            1.结构体binding验证      D:\go-20191030\golang-Gin\vaild_binding
            2.自定义验证             D:\go-20191030\golang-Gin\valid_custom
            3.支持多语言错误信息     D:\go-20191030\golang-Gin\valid_v9
        5.中间件：拦截操作-请求拦截，日志打印
            1.使用gin中间件          D:\go-20191030\golang-Gin\middleware_gin
            2.自定义ip白名单中间件   D:\go-20191030\golang-Gin\middleware_whitelist
        6.其他补充：
            1.优雅关停：保证服务在关闭时程序运行一半   D:\go-20191030\golang-Gin\other_shutdown
            2.模板渲染：                             D:\go-20191030\golang-Gin\other_template 具体使用参照官网
            3.自动证书：自动配置证书，过期可以自动续约
    2.搭建企业级脚手架：
        1.功能展示
        2.文件分层
        3.引入轻量级golang类库
        4.输出格式统一封装
        5.定义中间件链路日志打印
        6.请求数据绑定结构体与校验
    3.开发用户管理系统
    
**3.开发环境：**
    
    go version go1.13.3 windows/amd64
    gin 1.4.0    
    idea
    
**4.搭建测试环境：**

    1.新建github仓库：golang-Gin
    2.拉取此仓库到本地，并使用idea打开
    3.进入目录：并init
        D:\go-20191030\golang-Gin>go mod init
        go: cannot determine module path for source directory D:\go-20191030\golang-Gin (outside GOPATH, module path must be specified)
    4.出现以上错误时：添加模块名重新init
        D:\go-20191030\golang-Gin>go mod init golang-Gin
        go: creating new go.mod: module golang-Gin
    5.下载Gin：
        D:\go-20191030\golang-Gin>go get -v github.com/gin-gonic/gin@v1.4
        dial tcp 216.58.200.49:443: connectex: A connection attempt
    6.出现以上错误的原因如下：
        在Go 1.13中，我们可以通过GOPROXY来控制代理，以及通过GOPRIVATE控制私有库不走代理。
        设置GOPROXY代理：
        go env -w GOPROXY=https://goproxy.cn,direct
        设置GOPRIVATE来跳过私有库，比如常用的Gitlab或Gitee，中间使用逗号分隔：
        go env -w GOPRIVATE=*.gitlab.com,*.gitee.com
        如果在运行go mod vendor时，提示Get https://sum.golang.org/lookup/xxxxxx: dial tcp 216.58.200.49:443: i/o timeout，则是因为Go 1.13设置了
        默认的GOSUMDB=sum.golang.org，这个网站是被墙了的，用于验证包的有效性，可以通过如下命令关闭：
        go env -w GOSUMDB=off
        可以设置 GOSUMDB="sum.golang.google.cn"， 这个是专门为国内提供的sum 验证服务。
        go env -w GOSUMDB="sum.golang.google.cn"
    7.解决5里的错误，重新下载：
        D:\go-20191030\golang-Gin>go env -w GOSUMDB="sum.golang.google.cn"
        D:\go-20191030\golang-Gin>go get -v github.com/gin-gonic/gin@v1.4
    8.配置go moudles：
        Settings》》Languages & Frameworks 》》 go >>go moudles >> 将Enable打上勾，这样才能引用数据
    9.创建start文件夹，创建main.go
