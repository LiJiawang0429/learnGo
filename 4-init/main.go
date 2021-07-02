package main

import (
	// "learnGo/4-init/lib1"  //基本gopath导包
	// "learnGo/4-init/lib2"
	"fmt"
	_ "learnGo/4-init/lib1"      //不引用强制导包
	mylib2 "learnGo/4-init/lib2" //起别名
	// ."learnGo/4-init/lib2" //.package，使得包中所有方法导入本包，使得其放法可以直接使用（一般不建议用，易重名）
)

func main() {
	defer fmt.Println("defer 在程序结束前使用")
	mylib2.Lib2test() //用别名调用
}
