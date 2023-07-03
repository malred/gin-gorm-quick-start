# 依赖安装

```shell
go get github.com/joho/godotenv
go get -u gorm.io/gorm
go get -u gorm.io/driver/数据库系统名(mysql/sqlite/...)
go get -u github.com/gin-gonic/gin
```

# 开发时运行

```shell
# command后面是go.mod的模块名
CompileDaemon -command="./quick-start"
```

# script 里是启动脚本

```shell
cd scripts
dev
# 或
./dev.sh
```

# 功能

- gin-validate 校验参数
- gin web 框架
- gorm 操作数据库
- jwt 生成和解析 jwt
- bcrypt 密码加盐
- godotenv 读取.env 文件
