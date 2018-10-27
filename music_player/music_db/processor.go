package music_db

import (
	"fmt"
	"container/list"
	"bufio"
	"os"
	"io"
	"io/ioutil"
)

type Processor struct {
	filePath string
	fileList *list.List
}

func NewProcessor(path string) *Processor {
	infoList,err := ioutil.ReadDir(path)
	fList := list.New()

	if err != nil {
		fmt.Println("read dir err")
	}
	for _,info := range infoList {
		fList.PushBack(info.Name())
	}

	for e := fList.Front() ; e != nil ; e = e.Next() {
		fmt.Println(e.Value)
	}

	return &Processor{
		filePath:path,
		fileList:fList,
	}
}

func (processor *Processor) ParseContent(name string) []string {
	info := []string{}

	file,err := os.Open("/Users/wangzijie/code/golang/src/music_player/data/"+name)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)
	for {
		line,err := reader.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
	info = append(info,line)
	}
	return info
}

