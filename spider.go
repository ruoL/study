/*****
* 文件名: spider.go
* 创建日期: 2015-3-19
* 作者:	张天其
* 功能:　实现了简单的网页下载，将网页的url-body以key-value的形式存到redis数据库
*****/

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hoisie/redis"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

var reg *regexp.Regexp
var pattern string = `href="http://www.litrin.net/[^\s]*"`
var urls []string

//数据库连接相关的参数
var host string = "localhost"
var port string = "3306"
var user string = "root"
var passwd string = "zhangtianqi"
var database string = "db1"

//连接数据库
func connect(host, port, user, passwd, database string) *sql.DB {
	conn := user + ":" + passwd + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8"
	db, err := sql.Open("mysql", conn)
	CheckError(err)
	fmt.Println("connection success !")
	return db
}

//错误检查
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//根据url获取相应的网页，以字符串的方式返回
func getBody(url string) string {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	buff, err_1 := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err_1 != nil {
		log.Fatal(err_1)
	}
	some := fmt.Sprintf("%s", buff)
	return some
}

//将获取的url插入到urls列表里面
func insertUrlToUrls(url string) {
	for i := 0; i < len(urls); i++ {
		if urls[i] == url {
			return
		}
	}
	urls = append(urls, url)
}

//从网页body里面提取目标类型的url，插入到urls列表里面
func getUrlFromBody(body string) {
	reg = regexp.MustCompile(pattern)
	strs := reg.FindAllString(body, -1)
	for _, v := range strs {
		insertUrlToUrls(v[6 : len(v)-1])
	}
}

var index int

//主函数
func main() {
	index = 1
	db := connect(host, port, user, passwd, database)
	table := "html"
	sql := "insert into " + table + "(id, body) values(?, ?)"
	stmt, err := db.Prepare(sql)
	CheckError(err)

	var num int = 0

	if urls == nil {
		urls = make([]string, 0, 10)
		urls = append(urls, "http://www.litrin.net")
	}

	var client redis.Client

	for j := 0; j < 4; j++ {
		n := len(urls)
		for i := 0; i < n; i++ {
			body := getBody(urls[i])
			client.Set(urls[i], []byte(body))
			stmt.Exec(num, body)
			fmt.Printf("%d, %s\n", index, body[:20])
			index++
			num++
			getUrlFromBody(body)

		}
		for i, v := range urls {
			fmt.Println(i, v)
		}
	}
}
