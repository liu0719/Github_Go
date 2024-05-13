// 2023/1/4,14:58
package main

import "fmt"

func modifyUser(users map[string]map[string]string, name string, nickname string) {
	if users[name] != nil {
		users[name]["password"] = "88888"
	} else {
		users[name] = make(map[string]string)
		users[name]["password"] = "22222"
		users[name]["nickname"] = nickname
	}
}
func main() {
	user := make(map[string]map[string]string)
	user["smith"] = make(map[string]string)
	user["smith"]["password"] = "66666"
	user["smith"]["nickname"] = "SM"
	modifyUser(user, "smith", "昵称1")
	modifyUser(user, "mary", "昵称2")
	//modifyUser(user, "tom", "昵称1")
	fmt.Println(user)
}
