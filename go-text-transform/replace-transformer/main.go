package main

import (
	"bytes"

	"golang.org/x/text/transform"
)

type Replacer struct {
	transform.NopResetter
	old, new []byte
}

func (r *Replacer) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	// r.oldが空であれば、そのままコピー。
	if len(r.old) == 0 {
		n := copy(dst[nDst:], src)
		nDst += n
		nSrc += n
		return
	}

	for {
		// srcのnSrc番目からr.oldを探す。
		i := bytes.Index(src[nSrc:], r.old)

		// 見つからなかった場合。
		if i == -1 {
			// TODO
			return
		}

		// 見つけたところまでをコピーして書き込む。
		n := copy(dst[nDst:], src[nSrc:nSrc+i])
		nDst += n
		nSrc += n
		if n < i {
			err = transform.ErrShortDst
			return
		}

		// 置換する文字をコピーして書き込む。
		n = copy(dst[nDst:], r.new)
		nDst += n
		nSrc += len(r.old)
	}
}

func main() {
}
