package liblog

import (
	"io"
	"log"
	"os"
)

// 自定义日志记录器
var (
	Trace   *log.Logger // 记录所有日志
	Info    *log.Logger // 重要的信息
	Warning *log.Logger // 需要注意的信息
	Error   *log.Logger // 非常严重的问题
)

func init() {
	// 创建日志文件
	file, err := os.OpenFile("./errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file: ", err)
	}

	/**
	io.New(out io.Writer, prefix string, flag int)
	out: 指定日志要写到的目的地
	prefix: 前缀
	flag: 日志标志
	*/

	// Discard: 丢弃，当某个等级的日志不重要时，使用 Discard 变量可以禁用掉这个等级的日志
	Trace = log.New(io.Discard, "TRACE: ", log.Ldate|log.Ltime|log.Llongfile)
	// 输出到标准输出
	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Llongfile)
	// 输出到标准输出
	Warning = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Llongfile)
	// 输出到日志文件和标准输出
	Error = log.New(io.MultiWriter(file, os.Stdout), "ERROR: ", log.Ldate|log.Ltime|log.Llongfile)
}

func TestCustom() {
	Trace.Println("wiwi")
	Info.Println("sama")
	Warning.Printf("tianxian")
	Error.Println("baobao")
}
