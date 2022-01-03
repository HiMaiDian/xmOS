package config

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
)

func mapingTo(key string, s interface{}) {
	sectionSlice := getSection(key)
	fmt.Println(sectionSlice)
	if len(sectionSlice) > 0 {
		el := reflect.ValueOf(s).Elem()
		for _, line := range sectionSlice {
			ll := strings.Split(line, " = ")
			if len(ll) == 2 && ll[1] != " " {
				el.FieldByName(strings.TrimSpace(ll[0])).SetString(ll[1])
			}
		}
	}
}

func getSection(key string) []string {
	var (
		fp           *os.File
		contentSlice []string
	)
	contentSlice = make([]string, 0)
	fp, _ = os.Open(CONF_URL)
	defer fp.Close()
	w := bufio.NewReader(fp)
	sk := false
	for {
		readerString, err := w.ReadString('\n')

		if strings.Index(readerString, "["+key+"]") == 0 {
			sk = true
			continue
		}
		if err == io.EOF || readerString == "\n" {
			break
		}
		fmt.Println(readerString)
		if sk {
			contentSlice = append(contentSlice, strings.Trim(readerString, "\n"))
		}

	}
	return contentSlice
}
