package main
import "fmt"
func main(){
	/*
	Go语言中默认没有char类型，需要利用用unicode编码值来表示，即用整数类型存储
	对于ASCII码字符，通常用byte即可
	对于汉字，通常用int，具体取决于编码方式
	*/

	var c1 byte = 'A'
	var c2 byte = ' '
	var c3 int = '中'

	/*
	字符类型无法直接打印到控制台，需要借助格式化打印
	*/

	fmt.Printf("%c,%c,%c\n",c1,c2,c3)


	/*
	转义字符：
	\n:换行
	\r:回车，将光标回到首列，并且会将后续的输出从首列开始覆盖
	\t:制表符，补满8格
	\b:退格，删除前一个字符
	\":双引号
	\':单引号
	*/
	fmt.Printf("aaaaaaaa,bbb\t,bbbb\t,end\n")
	fmt.Printf("aaaaaa\rbbb\n")

}