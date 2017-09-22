package osudb

import (
	"io"

	"github.com/bnch/osubinary"
	"github.com/flesnuk/osu-tools/osu"
)

var r io.Reader

func bytes(n uint) {
	r.Read(make([]byte, n))
}

func GetBeatmaps(re io.Reader) map[string]osu.Beatmap {
	r = re
	or := osubinary.New(r)

	var hm = make(map[string]osu.Beatmap)
	var bmsetid uint32
	var temps, hash, filename string
	var length, tmp uint32
	bytes(17)
	or.OsuRead(&temps)
	or.OsuRead(&length)

	for i := uint32(0); i < length; i++ {
		or.OsuRead(&tmp)
		for j := 0; j < 7; j++ {
			or.OsuRead(&temps)
		}

		or.OsuRead(&hash)
		or.OsuRead(&filename)
		bytes(39)

		for k := 0; k < 4; k++ {
			var size uint32
			or.OsuRead(&size)
			bytes(14 * uint(size))
		}
		bytes(12)
		var ntimings uint32
		or.OsuRead(&ntimings)
		bytes(uint(ntimings) * 17)
		or.OsuRead(&tmp)

		or.OsuRead(&bmsetid)

		hm[hash] = osu.Beatmap{
			ID:       bmsetid,
			Filename: filename,
		}

		bytes(15)
		or.OsuRead(&temps)
		or.OsuRead(&temps)
		bytes(2)
		or.OsuRead(&temps)
		bytes(10)
		or.OsuRead(&temps)
		bytes(18)

	}

	return hm
}
