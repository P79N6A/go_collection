package music_db

import (
	"fmt"
	"container/list"
	//"bufio"
	//"os"
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

//func (processor *Processor) GetAllFile(filePath string) {

//}
