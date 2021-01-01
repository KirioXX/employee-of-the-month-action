package file

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"text/template"
	"time"
)

const (
	Dir          = "/tmp/repo"
	startComment = "<!---employee-of-the-month-action:Start--->"
	endComment   = "<!---employee-of-the-month-action:End--->"
	regex        = `(<!\-\-\-.*:Start\-\-\->(.*\s)*<!\-\-\-.*:End\-\-\->)`
)

type content struct {
	StartComment string
	Title        string
	ImageURL     string
	EndComment   string
}
type titleContent struct {
	Month string
}

func GenTemplate(title string, imageURL string) ([]byte, error) {
	b, err := template.New("block").Parse(`{{.StartComment}}
## {{.Title}}

![Employee of the month]({{.ImageURL}})
{{.EndComment}}`)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	t, err := template.New("title").Parse(title)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var tplTitle bytes.Buffer
	if err := t.Execute(&tplTitle, titleContent{time.Now().Month().String()}); err != nil {
		fmt.Println(err)
		return nil, err
	}
	var tpl bytes.Buffer
	if err := b.Execute(&tpl, content{startComment, tplTitle.String(), imageURL, endComment}); err != nil {
		fmt.Println(err)
		return nil, err
	}
	result := tpl.Bytes()
	return result, nil
}

type marker struct {
	found bool
	line  int
}

func HasMarkers(str []byte) bool {
	// Splits on newlines by default.
	scanner := bufio.NewScanner(bytes.NewReader(str))

	line := 1
	startMarker := marker{found: false, line: 0}
	endMarker := marker{found: false, line: 0}
	// https://golang.org/pkg/bufio/#Scanner.Scan
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), startComment) {
			startMarker = marker{found: true, line: line}
		} else if strings.Contains(scanner.Text(), endComment) {
			endMarker = marker{found: true, line: line}
		}

		line++
	}

	return startMarker.found && endMarker.found && startMarker.line < endMarker.line
}

func ReplaceMarker(original string, newString string) string {
	re := regexp.MustCompile(regex)
	s := re.ReplaceAllString(original, newString)
	return s
}

func ReadFile(fileName string) []byte {
	file, err := ioutil.ReadFile(Dir + "/" + fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return file
}

func WriteFile(fileName string, content []byte) {
	if err := ioutil.WriteFile(Dir+"/"+fileName, content, 0644); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Writing file success")
}

func CleanDir() {
	if err := os.RemoveAll(Dir); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Repository removed")
}
