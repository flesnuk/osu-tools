package osudb

import (
	"io"

	"github.com/flesnuk/osu-tools/osu"
	"github.com/pkg/errors"
)

// GetBeatmaps returns a map with the beatmap hash being the key
func GetBeatmaps(r io.Reader) (map[string]osu.Beatmap, error) {
	or := osu.SafeReader{r, nil}

	var hm = make(map[string]osu.Beatmap)
	var bmsetid uint32
	var hash, filename string
	var length uint32
	or.SkipBytes(17)
	or.SkipString()
	or.ReadInt(&length)
	if or.Err != nil {
		return nil, errors.Wrap(or.Err, "Failed reading osu!db header")
	}

	for i := uint32(0); i < length; i++ {
		or.SkipBytes(4)
		for j := 0; j < 7; j++ {
			or.SkipString()
		}

		or.ReadString(&hash)
		or.ReadString(&filename)
		or.SkipBytes(39)

		for k := 0; k < 4; k++ {
			var size uint32
			or.ReadInt(&size)
			or.SkipBytes(14 * int64(size))
		}
		or.SkipBytes(12)
		var ntimings uint32
		or.ReadInt(&ntimings)
		or.SkipBytes(int64(ntimings) * 17)
		or.SkipBytes(4)

		or.ReadInt(&bmsetid)

		hm[hash] = osu.Beatmap{
			ID:       bmsetid,
			Filename: filename,
		}

		or.SkipBytes(15)
		or.SkipString()
		or.SkipString()
		or.SkipBytes(2)
		or.SkipString()
		or.SkipBytes(10)
		or.SkipString()
		or.SkipBytes(18)

		if or.Err != nil {
			return nil, errors.Wrapf(or.Err, "Failed at reading osu!db entry i: %d", i)
		}

	}

	return hm, nil
}
