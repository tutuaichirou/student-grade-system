package main

import (
	"fmt"
	"nalohoduo/util"
)

func main() {
	type userDetail struct {
		password string
		age string
	}

	user:=map[string]userDetail{
		"tutu":{"123", "18"},
		"tiantian":{"321", "22"},
	}

	for {
		out:
		println("欢迎光临，按1登录，按2注册，按q退出")
		res:=util.KeyboardInput()
		if res=="q"{return}
		if res=="1"{
			for {
				println("请输入用户名:")
				name:=util.KeyboardInput()
				if name=="q"{break}
				if v, ok := user[name]; ok {
					for i := 2; i >= 0; i-- {
						println("请输入密码:")
						password:=util.KeyboardInput()
						if password=="q"{break}
						if v.password == password {
							println("登录成功")
							break
						} else {
							if i == 0 {
								fmt.Printf("用户%v被锁定\n",name)
								goto out
							}
							fmt.Printf("密码错误，你还有%v次机会\n",i)
						}
					}
				} else {
					println("账户不存在，请重新输入")
				}
			}

		}
		if res=="2"{
			for{
				println("请填写用户名：")
				name:=util.KeyboardInput()
				if name=="q"{break}
				if len(name)<3||len(name)>16{
					println("用户名长度必须在3~16位")
				}

				if _,ok:=user[name];ok{
					println("该用户名已经存在!")
					continue
				}

				println("请输入密码：")
				password:=util.KeyboardInput()
				if password=="q"{break}
				println("请再次输入")
				password2:=util.KeyboardInput()
				if password2=="q"{break}
				if password!=password2{
					println("两次密码输入不一致!")
					continue
				}


				break
			}

		}

	}
	
}
