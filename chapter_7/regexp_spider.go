package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("http get error.")
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("resp close error.")
		}
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http read error")
		return
	}

	src := string(body)

	// 将HTML标签全转换成小写
	re, _ := regexp.Compile(`<[\S\s]=?>`)
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	// 去除STYPLE
	re, _ = regexp.Compile(`<style[\S\s]+?</style>`)
	src = re.ReplaceAllString(src, "")
	// 去除HTMLUnscape的STYLE
	re, _ = regexp.Compile(`&lt;style[\S\s]+?&lt;/style&gt;`)
	src = re.ReplaceAllString(src, "")

	// 去除SCRIPT
	re, _ = regexp.Compile(`<script[\S\s]+?</script>`)
	src = re.ReplaceAllString(src, "")
	// 去除HTMLUnsapce的SCRIPT
	re, _ = regexp.Compile(`&lt;script[\S\s]+?&lt;/script&gt;`)
	src = re.ReplaceAllString(src, "")

	// 去除所有尖括号内的HTML代码，并换乘换行符
	re, _ = regexp.Compile(`<[\S\s]+?>`)
	src = re.ReplaceAllString(src, "\n")

	// 去除连续的换行符
	re, _ = regexp.Compile(`\s{2,}`)
	src = re.ReplaceAllString(src, "\n")

	fmt.Println(strings.TrimSpace(src))
}
