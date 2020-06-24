package main

import (
	"bufio"
	"fmt"
	"github.com/antchfx/htmlquery"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	home := "https://www.gamersky.com/ent/202006/1298826.shtml"
	//url2 := "https://www.gamersky.com/ent/202006/1298826_2.shtml"
	//url3 := "https://www.gamersky.com/ent/202006/1298826_3.shtml"
	doc, err := htmlquery.LoadURL(home)
	if err != nil {
		panic(err)
	}
	// Find all news item.
	list, err := htmlquery.QueryAll(doc, "//p/a")
	if err != nil {
		panic(err)
	}
	//图片链接存放
	openFile, err := os.OpenFile("./link.html", os.O_WRONLY|os.O_RDONLY|os.O_CREATE, 0611)
	if err != nil {
		fmt.Println(err)
	}

	if err != nil {
		fmt.Println(err)
		return
	}
	for _, n := range list {
		a := htmlquery.FindOne(n, "//img")
		uurl := htmlquery.SelectAttr(a, "src") //链接
		fmt.Println(uurl)
		sprintf := fmt.Sprintf("%s\n", uurl)
		//fmt.Println(sprintf)
		openFile.WriteString(sprintf) //写入图片链接
	}
	getPic()
}

//读取文件并下载
func getPic() {
	open, err := os.Open("./link.html")
	if err != nil {
		fmt.Println(err)
	}
	reader := bufio.NewReader(open)
	for i := 0; i < 4; i++ {
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println(err)
		}

		resp, err := http.Get(string(line)) //
		//fmt.Println("xxxxxx", string(line))
		if err != nil {
			fmt.Println(err)
		}
		data, err := ioutil.ReadAll(resp.Body) //拿到图片二进制信息
		if err != nil {
			fmt.Println("读取图片信息失败")
		}
		//
		str := fmt.Sprintf("./PIC%d.png", i)
		file, err := os.OpenFile(str, os.O_CREATE|os.O_RDONLY|os.O_WRONLY, 0644)
		file.Write(data) //写入图片到文件
	}
	fmt.Println("Done")
}
