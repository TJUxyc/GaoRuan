package main
 
import "fmt"

type dealFunc func()

type DataNode struct{
	cmd string
	desc string
	handler dealFunc
	next *DataNode
}

var head [3]DataNode = [3]DataNode{
	{"help", "this is help cmd!", Help, nil},
	{"version", "menu program v1.0", nil, nil},
	{"quit", "Quit from menu", Quit, nil},
}

// 因为golang不支持静态变量，无法在Help函数中直接读取head中的变量(cmd和desc)
// 因此将字符串部分单独提取出来给Help做变量
var info [3]string = [3]string{
	"help - this is help cmd!",
	"version - menu program v1.0",
	"quit - Quit from menu",
}

func Help(){
	fmt.Println("Menu List:")
	fmt.Println("=========================")
	var p string
	for i:=0;i<3;i++ {
		p = info[i]
		fmt.Println(p)
	}
	fmt.Println("=========================")
}

func Quit() {
	panic("Quit!")
}

func main() {
	// 同理，由于golang不支持静态变量，为了避免head出现循环初始化的情况
	// head初始化需要将next全部置为nil，在主函数中（此时head已经完成了初始化）再设置next的值
	for i:=0; i<len(head)-1; i++ {
		head[i].next = &head[i+1]
	}
	for {
		var cmd string
		fmt.Print("Input a cmd number(Input \"help\" for tips) > ")
		fmt.Scan(&cmd)
		var p *DataNode = &head[0]
		for p != nil {
			if p.cmd == cmd{
				fmt.Println(p.cmd + " - " + p.desc)
				if p.handler != nil {
					p.handler()
				}
				break
			}
			p = p.next
		}
		if p == nil{
			fmt.Println("This is a wrong cmd!")
		}
	}
}