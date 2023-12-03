package handlers

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func registerCompanyHandler() {
	http.HandleFunc("/company", hGetCompanyList)
	http.HandleFunc("/company/", hGetCompany)
}

func hGetCompanyList(w http.ResponseWriter, r *http.Request) {
	companies := []string{"MS", "Oracle", "Google"}
	w.Write([]byte(strings.Join(companies, ",")))
}

func hGetCompany(w http.ResponseWriter, r *http.Request) {
	// 创建正则表达式
	pattern, _ := regexp.Compile(`/company/(\d+)`)
	// 正则匹配
	match := pattern.FindStringSubmatch(r.URL.Path)
	if len(match) > 0 {
		// 将字符串转为数字
		companyId, _ := strconv.Atoi(match[1])
		fmt.Fprintln(w, companyId)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
