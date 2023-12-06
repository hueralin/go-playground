package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type CommonRes struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func hHome(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Home"))
	if err != nil {
		return
	}
}

func hGetBookDetail(w http.ResponseWriter, r *http.Request) {
	// 读取路径参数
	params := mux.Vars(r)
	// map[id:12]，一个 map，存储路径参数的 k-v
	fmt.Println(params)
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, "Book id: %v\n", params["id"])
	if err != nil {
		return
	}
}

func hGetUserDetail(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	_, err := fmt.Fprintf(w, "User id: %v\n", params["id"])
	if err != nil {
		return
	}
}

func hGetGames(w http.ResponseWriter, r *http.Request) {
	res := CommonRes{Code: "ok", Msg: "", Data: []string{"React", "Vue", "Angular"}}
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		return
	}
}

func main() {
	// r 是 *mux.Router 类型，其实现了 ServeHTTP 方法，所以也是个 Handler
	r := mux.NewRouter()
	// 开启 Path Slash 的严格检查，即访问 /home 会被重定向到 /home/（根据定义的 route 为准），反之亦然
	r.StrictSlash(true)

	r.HandleFunc("/", hHome)
	// Path 中定义变量，也可以为变量指定正则
	// 如果没有指定正则，那么变量可以是 Path 中的任何东西，直到下一个 /
	// 如果指定了正则，那就按正则来
	r.HandleFunc("/book/{id}", hGetBookDetail)

	// 也可以做更多的限制，如只有 Get 请求，符合 Path 的请求才能进入该 Handler
	r.
		Methods(http.MethodGet).
		Path("/user/{id:[0-9]+}").
		HandlerFunc(hGetUserDetail)

	// 路由分组 - 子路由，共享相同的配置
	rGame := r.PathPrefix("/games").Subrouter()
	rGame.
		Methods(http.MethodGet).
		Path("/").
		HandlerFunc(hGetGames)

	// 静态文件服务
	var dir string
	flag.StringVar(&dir, "dir", ".", "desc")
	flag.Parse()
	r.PathPrefix("/static").Handler(http.FileServer(http.Dir(dir)))

	err := http.ListenAndServe(":8888", r)
	if err != nil {
		return
	}
}
