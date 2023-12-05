package liblog

import "log"

func init() {
	log.SetPrefix("TRACE: ")
	// 设置输出日志的各部分：日期，带微秒的时间，输出日志的文件
	// Ldate: 日期 2023/01/01
	// Ltime: 时间 02:23:32
	// Lmicroseconds 毫秒级时间 02:23:32.123123，会覆盖 Ltime
	// Llongfile: 完整路径的文件名和行号 /a/b/c/d.go:30
	// Lshortfile: 最终的文件名和行号 d.go:30，会覆盖 Llongfile
	// LstdFlags: 标准日志记录器的初始值，Ldate | Ltime
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func TestStart() {
	// Println 写到标准日志记录器
	log.Println("message")

	// Fatalln 在调用 Println 之后接着调用 os.Exit(1)，程序正常结束时的 code 是 0
	log.Fatalln("fatal message")

	// Panicln 在调用 Println 之后接着调用 panic()
	//log.Panicln("panic message")
}
