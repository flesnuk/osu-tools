package osr

import (
	"io"
	"sort"
	"time"

	"github.com/flesnuk/osu-tools/osu"
	"github.com/flesnuk/osubinary"
)

func ReadScoreDB(r io.Reader) []osu.Replay {
	var str string
	var nBM uint32
	var size uint32
	var ret []osu.Replay
	osuR := osubinary.New(r)
	r.Read(make([]byte, 4))
	osuR.OsuRead(&nBM)
	for i := uint32(0); i < nBM; i++ {
		osuR.OsuRead(&str)
		osuR.OsuRead(&size)
		for j := uint32(0); j < size; j++ {
			replay := NewReplay(r)
			if replay.GameMode == 0 {
				ret = append(ret, replay)
			}
			r.Read(make([]byte, 12))
		}
	}
	sort.Slice(ret, func(i, j int) bool { return ret[i].TimeStamp > ret[j].TimeStamp })
	return ret
}

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
		TimeStamp   uint64
	)
	osuR := osubinary.New(r)
	str := ""

	osuR.OsuRead(&GameMode, &GameVersion, &BeatmapHash,
		&PlayerName, &ReplayHash, &N300, &N100, &N50,
		&Geki, &Katu, &Misses, &Score,
		&Combo, &Perfect, &Mods, &str, &TimeStamp)

	localLoc, err := time.LoadLocation("Local")
	if err != nil {
		panic("Error loading local location")
	}

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
		TimeStamp:   TimeStamp,
		ModTime:     TimeFromTicks(int64(TimeStamp)).In(localLoc),
	}
}

func TimeFromTicks(ticks int64) time.Time {
	base := time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC).Unix()
	return time.Unix(ticks/10000000+base, 0).UTC()
}
