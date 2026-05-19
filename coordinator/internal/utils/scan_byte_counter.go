package utils

import (
	"bufio"
)

type ScanByteCounter struct {
	BytesRead int64
}

func (s *ScanByteCounter) Wrap(split bufio.SplitFunc) bufio.SplitFunc {
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		adv, tok, err := split(data, atEOF)
		s.BytesRead += int64(adv)
		return adv, tok, err
	}
}
