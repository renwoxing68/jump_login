# 跳板机自动登录 
获取动态密码(例如谷歌身份验证),登录跳板机

### 第一步 配置token实现动态密码  Go服务

1. 服务介绍 
   - 前缀+动态密码+后缀组成 前后缀都可为空
   - 针对一下不同场景:
     - 固定密码+动态密码 例如:rwx324343
     - 动态密码+固定密码 例如:324343xwr
     - 以及 固定密码+动态+固定密码 例如:rwx543334xwr
2. token/conf/conf 配置维护
```
#前缀
PREFIX=rwx(替换 可为空)
#后缀
SUFFIX=xwr(替换 可为空)
#动态密码 秘钥
SECRET=renwoxing68
```
3. token包对应 code.go 如修改code.go 重新构建

### 第二步 配置expect脚本
1. sh/sv_test.sh 对应替换成个人的信息
2. **如果所需的服务是 先输入密码 成功后输入动态口令,可在脚本中自行添加一列**
2. **本脚本中还包括登录跳板机以后相关的操作,如不需要可遗弃**
3. **如需配置跳板机登录后跳转其他服务可配置多个sv_test.sh脚本** 

### 第三步 执行
1. expect 绝对路径/jump_login/sh/sv_test.sh  
2. 添加别名执行 
```
#linux mac
~/.zshrc  或 ~/.bashrc 中增加下面一行
source  正确路径/jump_login/sh/.alias_login

```


   


