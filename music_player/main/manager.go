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
	} else {
		fmt.Println("no such music in lib")
	}
}
