package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/Jeffail/gabs"
	"github.com/ghodss/yaml"
)

func main() {
	// https://editor.swagger.io/#
	// 去這個網站上可以線上編輯swagger json
	// 主要概念應該就是產生swagger文件, json to yaml, 或者是用yaml to json
	// json and yaml 都有文件了
	// 這個我使用之前幫Thomas找的娃娃機文件當範例

	// 讀取json檔案, 並轉成yaml (你也可以讀取yaml轉成json)
	jsonSwaggerFile, err := os.Open("swagger.json")
	defer jsonSwaggerFile.Close()
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	swaggerByte, _ := ioutil.ReadAll(jsonSwaggerFile)

	jsonTestFile, err := os.Open("test.json")
	defer jsonTestFile.Close()
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	testByte, _ := ioutil.ReadAll(jsonTestFile)

	// Search or Merge JSON https://github.com/Jeffail/gabs
	jsonParsed1, _ := gabs.ParseJSON(swaggerByte)
	jsonParsed2, _ := gabs.ParseJSON(testByte)
	jsonParsed1.Merge(jsonParsed2)

	yaml, err := yaml.JSONToYAML(jsonParsed1.Bytes())
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	f, err := os.Create("swagger.yaml")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	_, err2 := f.Write(yaml)

	if err2 != nil {
		log.Fatal(err2)
	}
}
