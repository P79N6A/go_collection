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
	fileList *list.List
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

	for e := fList.Front() ; e != nil ; e = e.Next() {
		fmt.Println(e.Value)
	}

	return &Processor{
		filePath:path,
		fileList:fList,
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
	music := NewMusic()
    music.name = info[0]
	music.fileType = info[1]
	music.singer = info[2]
	music.album = info[3]
	music.time = info[4]
	for i := 5 ; i < len(info) ; i++ {
		music.lyric += info[i]
	}
	processor.musicList.PushBack(&music)
	return music
}

