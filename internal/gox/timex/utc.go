package timex

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"hash"
	"time"
)

const UTCLayout = "2006-01-02T15:04:05.000Z"

type UTCTime time.Time

func (t UTCTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(t).UTC().Format(UTCLayout))
}

func (t *UTCTime) UnmarshalJSON(data []byte) (err error) {
	if data[0] != []byte(`"`)[0] || data[len(data)-1] != []byte(`"`)[0] {
		return errors.New("Not quoted")
	}
	*t, err = ParseUTCTime(string(data[1 : len(data)-1]))
	return
}

func (t UTCTime) Hash32(h hash.Hash32) error {
	err := binary.Write(h, binary.LittleEndian, time.Time(t).UnixNano())
	return err
}

func (t UTCTime) String() string {
	return time.Time(t).Format(UTCLayout)
}

func ParseUTCTime(timespec string) (UTCTime, error) {
	t, err := time.Parse(UTCLayout, timespec)
	return UTCTime(t), err
}

func MustParseUTCTime(timespec string) UTCTime {
	ts, err := ParseUTCTime(timespec)
	if err != nil {
		panic(err)
	}
	return ts
}
