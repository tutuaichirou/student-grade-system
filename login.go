package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
		println("欢迎光临，按1登录，按2注册，按q退出")
		reader := bufio.NewReader(os.Stdin)
		res, _ := reader.ReadString('\n')
		res = strings.TrimSpace(res)
		if strings.ToUpper(res)=="Q"{break}
		if res=="1"{
			println("请输入用户名:(按q退出)")
			reader := bufio.NewReader(os.Stdin)
			res, _ := reader.ReadString('\n')
			res = strings.TrimSpace(res)
			if strings.ToUpper(res)=="Q"{break}
			if v, ok := user[res]; ok {
				for i := 2; i >= 0; i-- {
					println("请输入密码:（按q退出）")
					reader := bufio.NewReader(os.Stdin)
					password, _ := reader.ReadString('\n')
					password = strings.TrimSpace(password)
					if strings.ToUpper(password)=="Q"{break}
					if v.password == password {
						println("登录成功")
						break
					} else {
						if i == 0 {
							print("用户被锁定")
							break
						}
						fmt.Printf("密码错误，你还有%v次机会\n", i)
					}
				}
			} else {
				println("账户不存在，请重新输入")
			}
		}
		if res=="2"{
			println("请")
		}

	}
	
}
