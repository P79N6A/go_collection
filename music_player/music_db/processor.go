package music_db

import (
	"fmt"
	"container/list"
	"bufio"
	"os"
	"io"
	"io/ioutil"
	"strings"
)

type Processor struct {
	filePath string
	FileList *list.List
	musicList *list.List
}

func NewProcessor(path string) *Processor {
	infoList,err := ioutil.ReadDir(path)
	fList := list.New()
	mList := list.New()

	if err != nil {
		fmt.Println("read dir err")
	}
	for _,info := range infoList {
		fList.PushBack(info.Name())
	}

	return &Processor{
		filePath:path,
		FileList:fList,
		musicList:mList,
	}
}

func (processor *Processor) ParseContent(name string) []string {
	info := []string{}
	musicName := "name:"+strings.Split(name,".")[0]
	fileType := "type:"+strings.Split(name,".")[1]
	info = append(info,musicName,fileType)

	file,err := os.Open(processor.filePath+name)
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

func (processor *Processor)  GenerateMusic(name string) *Music {
	info := processor.ParseContent(name)
	music := NewMusic() //新建一个Music类型的指针
    music.Name = info[0]
	music.FileType = info[1]
	music.Singer = info[2]
	music.Album = info[3]
	music.Time = info[4]
	for i := 5 ; i < len(info) ; i++ {
		music.Lyric += info[i]
	}
	processor.musicList.PushBack(music)
	return music
}

func (processor *Processor) LoadAllMusic() *list.List {
	for e := processor.FileList.Front() ; e != nil ; e = e.Next() {
		//interface{}类型可用于向函数传递任意类型的变量，但对于函数内部，该变量仍然为interface{}类型
		//e.Value是interface{}类型，需要将接口类型向普通类型转型，称为类型断言（运行期确定）
		processor.GenerateMusic(e.Value.(string))
	}
	return processor.musicList
}

