#!/usr/bin/expect
#跳板机地址
set jump_server relay00.rwx.com(替换成自己的)
#用户
set jump_user rwx(替换成自己的)
#跳转的服务
set ip_server 10.1.0.0(替换成自己的)
#动态密码
set token [exec sh -c {cd 绝对路径/jump_login/token/ && ./token}]
spawn ssh $jump_server -l $jump_user
#根据捕获到每一步的显示情况 执行相应的操作
expect {
    "*Password:*" {#动态密码  屏幕出现 "Password:"以后输入token密码  支持除动态密码以外,所需要的前后缀
        send "$token\n";
        exp_continue;
    }
    "*交互终端*" {#登录跳板机以后跳转的服务   屏幕出现 "交互终端"证明跳板机登录成功,输入跳转的服务
        send "$ip_server\n";
        exp_continue;
    }
    "*Last login:*" {#登录服务以后切换的用户
        send "sudo -iu work\n";
    }
}
interact
