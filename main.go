package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	// Загрузим страницы
	for url, path := range getInfo() {
		if err := download(url, path); err != nil {
			fmt.Println(err)
		}
	}

}

func getInfo() map[string]string {
	return map[string]string{
		"https://dentalik.webflow.io":          "D:\\PROG\\GoProjects\\src\\WFDL\\out\\index.html",
		"https://dentalik.webflow.io/products": "D:\\PROG\\GoProjects\\src\\WFDL\\out\\products.html",
	}
}
func download(url string, path string) (err error) {

	var resp *http.Response
	var file *os.File

	resp, err = http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err = os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	io.Copy(file, resp.Body)

	return err
}
