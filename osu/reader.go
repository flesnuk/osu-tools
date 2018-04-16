package osu

import (
	"encoding/binary"
	"io"
	"io/ioutil"

	"github.com/bnch/uleb128"

	"github.com/bnch/osubinary"
	"github.com/pkg/errors"
)

// SafeReader wraps an osubinary.OsuReader with an error field
type SafeReader struct {
	io.Reader
	Err error
}

// ReadString attempts reading a string from the reader
func (r SafeReader) ReadString(s *string) {
	if r.Err != nil {
		return
	}
	b, err := osubinary.ReadString(r)
	if err != nil {
		r.Err = errors.Wrap(err, "ReadString failed")
		return
	}
	*s = string(b)
}

// SkipString skips a string from the reader
func (r SafeReader) SkipString() {
	if r.Err != nil {
		return
	}

	var strHeader uint8
	err := binary.Read(r, binary.LittleEndian, &strHeader)
	if err != nil {
		r.Err = errors.Wrap(err, "SkipString: failed while reading header")
		return
	}
	if strHeader != 11 {
		return
	}

	strlen := uleb128.UnmarshalReader(r)
	io.CopyN(ioutil.Discard, r, int64(strlen))
}

// ReadInt attempts reading a uint32 from the reader
func (r SafeReader) ReadInt(num *uint32) {
	if r.Err != nil {
		return
	}
	err := binary.Read(r, binary.LittleEndian, num)
	if err != nil {
		r.Err = errors.Wrap(err, "ReadInt failed")
	}
}

// SkipBytes attempts skipping n bytes from the reader
func (r SafeReader) SkipBytes(n int64) {
	if r.Err != nil {
		return
	}
	_, err := io.CopyN(ioutil.Discard, r, n)
	if err != nil {
		r.Err = errors.Wrapf(err, "SkipBytes failed while trying to skip %d bytes", n)
	}
}
