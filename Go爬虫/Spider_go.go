package main

import (
	"fmt"
	"go/types"
	"io/ioutil"
	"net/http"
)

func download(url string) types.Document {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://movie.douban.com/top250", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.97 Safari/537.36")
	//配置好Request之后，使用Client对象的Do方法，就可以将Request发送出去
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)

	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)

	}
	return body
}

func main() {

}
