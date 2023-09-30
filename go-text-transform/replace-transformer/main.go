package main

import (
	"bytes"

	"golang.org/x/text/transform"
)

type Replacer struct {
	old, new []byte
	// 前回書き込めなかった分。
	preDst []byte
}

func (r *Replacer) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	// 前回書き込めなかった分を書き込む。
	if len(r.preDst) > 0 {
		n := copy(dst, r.preDst)
		nDst += n
		r.preDst = r.preDst[n:]

		// それでもまだ足りない場合。
		if len(r.preDst) > 0 {
			err = transform.ErrShortDst
			return
		}
	}

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
			n := len(src[nSrc:])
			m := copy(dst[nDst:], src[nSrc:nSrc+n])
			nDst += m
			nSrc += m

			// 全部書き込めなかった場合
			if m < n {
				err = transform.ErrShortDst
				return
			}

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

		// r.newが長くてdstが足りない場合は次回に持ち越し。
		if n < len(r.new) {
			r.preDst = r.new[n:]
			err = transform.ErrShortDst
			return
		}
	}
}

func (r *Replacer) Reset() {
	r.preDst = nil
}

func main() {
}
