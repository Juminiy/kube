package safe_json

import (
	stdjson "encoding/json"
	"errors"
	"io"
)

var ErrNeedObjectBegin = errors.New("need JSON object begin delimiter '{'")
var ErrNeedObjectKey = errors.New("need JSON object key string")
var ErrNeedObjectEnd = errors.New("need JSON object end delimiter '}'")
var ErrNeedArrayBegin = errors.New("need JSON array begin delimiter '['")
var ErrNeedArrayEnd = errors.New("need JSON array end delimiter ']'")

type NextValuer interface {
	Next(valuePtr any) (bool, error)
}

type StreamObject struct {
	Reader io.Reader
	*streamed
}

func (d *StreamObject) Next(valuePtr any) (bool, error) {
	if d.streamed == nil {
		d.streamed = newStreamed(d.Reader)
	}
	if !d.notFirstToken {
		d.notFirstToken = true
		tkn, err := d.dcr.Token()
		if err != nil {
			return false, err
		} else if tkn != stdjson.Delim('{') {
			return false, ErrNeedObjectBegin
		}
	}
	if d.lastToken {
		return false, nil
	}
	if d.dcr.More() {
		if objKey, err := d.dcr.Token(); err != nil {
			return false, err
		} else if _, ok := objKey.(string); !ok {
			return false, ErrNeedObjectKey
		}
		if err := d.dcr.Decode(valuePtr); err == nil {
			return true, nil
		} else if err == io.EOF {
			return false, nil
		} else {
			return false, err
		}
	} else {
		d.lastToken = true
		tkn, err := d.dcr.Token()
		if err != nil {
			return false, err
		} else if tkn != stdjson.Delim('}') {
			return false, ErrNeedObjectEnd
		}
		return false, nil
	}
}

func (d *StreamObject) HasNext(valuePtr any) bool {
	hasValue, err := d.Next(valuePtr)
	if err != nil {
		d.err = err
	}
	return hasValue
}

type StreamArray struct {
	Reader io.Reader
	*streamed
}

func (d *StreamArray) Next(valuePtr any) (bool, error) {
	if d.streamed == nil {
		d.streamed = newStreamed(d.Reader)
	}
	if !d.notFirstToken {
		d.notFirstToken = true
		tkn, err := d.dcr.Token()
		if err != nil {
			return false, err
		} else if tkn != stdjson.Delim('[') {
			return false, ErrNeedArrayBegin
		}
	}
	if d.lastToken {
		return false, nil
	}
	if d.dcr.More() {
		if err := d.dcr.Decode(valuePtr); err == nil {
			return true, nil
		} else if err == io.EOF {
			return false, nil
		} else {
			return false, err
		}
	} else {
		d.lastToken = true
		tkn, err := d.dcr.Token()
		if err != nil {
			return false, err
		} else if tkn != stdjson.Delim(']') {
			return false, ErrNeedArrayEnd
		}
		return false, nil
	}
}

func (d *StreamArray) HasNext(valuePtr any) bool {
	hasValue, err := d.Next(valuePtr)
	if err != nil {
		d.err = err
	}
	return hasValue
}

type StreamValue struct {
	Reader io.Reader
	*streamed
}

func (d *StreamValue) Next(valuePtr any) (bool, error) {
	if d.streamed == nil {
		d.streamed = newStreamed(d.Reader)
	}
	if !d.notFirstToken {
		d.notFirstToken = true
	}
	if d.lastToken {
		return false, nil
	}
	if d.dcr.More() {
		if err := d.dcr.Decode(valuePtr); err == nil {
			return true, nil
		} else if err == io.EOF {
			return false, nil
		} else {
			return false, err
		}
	} else {
		d.lastToken = true
		return false, nil
	}
}

func (d *StreamValue) HasNext(valuePtr any) bool {
	hasValue, err := d.Next(valuePtr)
	if err != nil {
		d.err = err
	}
	return hasValue
}

type streamed struct {
	notFirstToken bool
	lastToken     bool
	dcr           *stdjson.Decoder
	err           error
}

func newStreamed(rdr io.Reader) *streamed {
	return &streamed{
		notFirstToken: false,
		lastToken:     false,
		dcr:           stdjson.NewDecoder(rdr),
		err:           nil,
	}
}

func (s *streamed) Err() error {
	return s.err
}
