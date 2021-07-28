package main
import(
"custom_package/person"
"fmt"
)

func main() {
	p:=person.Description("milap")
	fmt.Println(p)
}
//自定义包的使用
//新建文件夹，创建同名包文件，然后 go install 安装以备使用，在 调用地方使用 import 导入
