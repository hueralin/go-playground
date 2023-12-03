package main

import (
	"encoding/json"
	"net/http"
)

func main10() {
	// JSON
	// 类型映射
	// Go float64: JSON 数值
	// Go string: JSON 字符串
	// Go bool: JSON 布尔值
	// Go nil: JSON null
	// 在 Go 结构体上使用 Tag 的形式，将 JSON 的数据映射到结构体上
	// 对于未知结构的 JSON，可以使用以下两种方式
	// map[string]interface{}: 存储任意 JSON 对象
	// []interface{}: 存储任意 JSON 数组
	// 读取 JSON:
	// 先创建一个解码器 dec := json.NewDecoder(r.Body)，NewDecoder 的参数是一个 Reader 接口值（表明可从里面读数据）
	// 再使用解码器将 JSON 内容输出到变量 dec.decode(&xxx)
	// 写入 JSON:
	// 先创建一个编码器 enc := json.NewEncoder(w), NewEncoder 的参数是一个 Writer 接口值（表明可往里面写数据）
	// 再使用编码器将数据变成 JSON 格式写入 w
	//server := http.Server{Addr: "localhost:8888"}
	//http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
	//	switch r.Method {
	//	case http.MethodPost:
	//		company := Company{}
	//		dec := json.NewDecoder(r.Body)
	//		err := dec.Decode(&company)
	//		if err != nil {
	//			log.Println(err.Error())
	//			w.WriteHeader(http.StatusInternalServerError)
	//			return
	//		}
	//
	//		enc := json.NewEncoder(w)
	//		err = enc.Encode(company)
	//		if err != nil {
	//			log.Println(err.Error())
	//			w.WriteHeader(http.StatusInternalServerError)
	//			return
	//		}
	//	default:
	//		w.WriteHeader(http.StatusMethodNotAllowed)
	//	}
	//})
	//err := server.ListenAndServe()
	//if err != nil {
	//	return
	//}

	server := http.Server{Addr: "localhost:8888"}
	http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		jsonStr := `{
			"id": 456,
			"name": "Oracle",
			"country": "USA"
		}`
		c := Company{}
		// 将 JSON 字符串解码到 c
		err := json.Unmarshal([]byte(jsonStr), &c)
		if err != nil {
			return
		}

		// 将 c 编码为字节数组
		bytes, err := json.Marshal(c)
		if err != nil {
			return
		}
		w.Write(bytes)
	})
	err := server.ListenAndServe()
	if err != nil {
		return
	}

	// 第二种方式是使用 Marshal（编码） 和 Unmarshal（解码）
	// Marshal：把 go struct 转为 json 格式
	// Unmarshal：把 json 格式转化为 go struct

	// 两种方式的区别
	// 针对 string 和 byte，使用 Marshal 和 Unmarshal
	// Marshal => string
	// Unmarshal <= string
	// 针对 stream，使用 Encoder 和 Decoder
	// Encode => Stream，把数据写入 io.Writer
	// Decoder <= Stream，从 io.Reader 读取数据
}
