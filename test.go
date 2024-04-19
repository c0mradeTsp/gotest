// test.go
package testpackage // 包名通常是小写字母，以区分标准库和第三方库

import (
	"fmt"
)

// SayHello 打印 "Hello from testpackage!"
func SayHello() {
	fmt.Println("Hello from testpackage!")
}
