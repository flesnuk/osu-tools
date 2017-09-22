package osu

import "time"

// OsuReplay stores
type Replay struct {
	GameMode    byte
	GameVersion uint32
	BeatmapHash string
	PlayerName  string
	ReplayHash  string
	N300        uint16
	N100        uint16
	N50         uint16
	Geki        uint16
	Katu        uint16
	Misses      uint16
	Score       uint32
	Combo       uint16
	Perfect     bool
	Mods        uint32
	ModTime     time.Time
}

// Beatmap info
type Beatmap struct {
	ID       uint32
	Filename string
}
