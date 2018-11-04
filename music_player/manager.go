package main

import (
	//"fmt"
	"container/list"
	"./music_db"
	"./player"
)

type Manager struct {
	processor *music_db.Processor
	musicPlayer *player.MusicPlayer
	musicList *list.List
	deletedList *list.List
}

func NewManager(path string) *Manager {
	manager := Manager{}
	processor := music_db.NewProcessor("/Users/wangzijie/code/golang/src/music_player/data/")
	manager.musicList = processor.LoadAllMusic()
	return &manager
}

func (manager *Manager) PlayMusic(name string) {
	for e := manager.musicList.Front() ; e != nil ; e = e.Next() {
		player.PlayMusic(*manager.musicPlayer,e.Value.(*music_db.Music))
	}
}
