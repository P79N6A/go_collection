package main

import (
	"fmt"
	"container/list"
	"../music_db"
	"../player"
)

type Manager struct {
	processor *music_db.Processor
	MusicPlayer *player.MusicPlayer
	MusicList *list.List
	DeletedList *list.List
	PlayList *list.List
}

func NewManager(path string) *Manager {
	manager := Manager{}
	p := music_db.NewProcessor(path)
	ml := p.LoadAllMusic()
	pl := list.New()
	dl := list.New()
	manager.processor = p
	manager.MusicList = ml
	manager.PlayList = pl
	manager.DeletedList = dl
	return &manager
}

func (manager *Manager) PlayMusic() {
	if manager.PlayList.Len() == 0 {
		fmt.Println("no music in play list")
	}
	for e := manager.PlayList.Front(); e != nil; e = e.Next() {
		music := e.Value.(*music_db.Music)
		p := player.Factory(music.FileType)
		player.PlayMusic(p,music)
	}
}

func (manager *Manager) AddMusic(name string) {
	flag := 0
	for e := manager.processor.FileList.Front(); e != nil; e = e.Next() {
		if name == e.Value.(string) {
			flag = 1
			break
		}
	}
	if flag == 1 {
		music := manager.processor.GenerateMusic(name)
		manager.PlayList.PushBack(music)
		fmt.Println("you've added "+music.Name)
	} else {
		fmt.Println("no such music in lib")
	}
}

func (manager *Manager) ShowPlayList() {
	if manager.PlayList.Len() == 0 {
		fmt.Println("no music in play list")
		return
	}
	fmt.Println("==========show play list==========")
	for e := manager.PlayList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(*music_db.Music).Name)
	}
}

