package tools

import (
	"fmt"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"os"
)

func RenderHome() []byte {
	markdown, _ := ioutil.ReadFile("README.md")
	homepage := blackfriday.MarkdownBasic(markdown)
	return homepage
}

func GetPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "4747"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}
