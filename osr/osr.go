package osr

import (
	"io"

	"github.com/bnch/osubinary"
	"github.com/flesnuk/osu-tools/osu"
)

func NewReplay(r io.Reader) osu.Replay {
	var (
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
	)
	osuR := osubinary.New(r)

	osuR.OsuRead(&GameMode, &GameVersion, &BeatmapHash,
		&PlayerName, &ReplayHash, &N300, &N100, &N50,
		&Geki, &Katu, &Misses, &Score,
		&Combo, &Perfect, &Mods)

	return osu.Replay{
		GameMode:    GameMode,
		GameVersion: GameVersion,
		BeatmapHash: BeatmapHash,
		PlayerName:  PlayerName,
		ReplayHash:  ReplayHash,
		N300:        N300,
		N100:        N100,
		N50:         N50,
		Geki:        Geki,
		Katu:        Katu,
		Misses:      Misses,
		Score:       Score,
		Combo:       Combo,
		Perfect:     Perfect,
		Mods:        Mods,
	}
}
