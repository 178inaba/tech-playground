package main

import "golang.org/x/text/transform"

type Upper struct{ transform.NopResetter }

func (Upper) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	// 末尾ではないのにdstが足りない場合はErrShortDstを返す。
	if len(dst) == 0 && !atEOF {
		err = transform.ErrShortDst
		return
	}

	// TODO
	return
}

func main() {
}
