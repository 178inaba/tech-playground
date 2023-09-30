package main

import (
	"bytes"

	"golang.org/x/text/transform"
)

type Replacer struct {
	old, new []byte

	// 前回書き込めなかった分。
	preDst []byte

	// 前回余ったold分。
	preSrc []byte
}

func (r *Replacer) Transform(dst, src []byte, atEOF bool) (int, int, error) {
	// srcの前方にpreSrcを付加する。
	newSrc := src
	if len(r.preSrc) > 0 {
		newSrc = make([]byte, len(r.preSrc)+len(src))
		copy(newSrc, r.preSrc)
		copy(newSrc[len(r.preSrc):], src)
	}

	nDst, nSrc, preSrc, err := r.transform(dst, newSrc, atEOF)

	// 読み込んだ長さより退避していた長さが長い場合。
	if nSrc < len(r.preSrc) {
		r.preSrc = r.preSrc[nSrc:]
		nSrc = 0
	} else {
		nSrc -= len(r.preSrc)

		// 新しく退避させる。
		r.preSrc = preSrc
	}

	return nDst, nSrc, err
}
func (r *Replacer) transform(dst, src []byte, atEOF bool) (nDst, nSrc int, preSrc []byte, err error) {
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

			// srcの末尾がr.oldの前方部分で終わる場合。
			var w int
			if !atEOF { // まだ次で読み込める余地がある。
				// srcの末尾とr.oldが同じ分の長さ。
				w = overlapWidth(src[nSrc:], r.old)
				if w > 0 {
					// コピーする分から一旦除外しておく。
					n -= w
					err = transform.ErrShortSrc
				}
			}

			m := copy(dst[nDst:], src[nSrc:nSrc+n])
			nDst += m
			nSrc += m

			// 全部書き込めなかった場合
			if m < n {
				err = transform.ErrShortDst
				return
			}

			// 次のsrcの先頭に追加しておく。
			preSrc = r.old[:w]

			// 読み込んだことにする。
			nSrc += w

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
	r.preSrc = nil
}

// aの末尾とbの先頭がマッチする長さ。
// 例: a:[0, 1, 2], b:[1, 2] -> 2
func overlapWidth(a, b []byte) int {
	// aとbで短い方の長さ。
	w := len(a)
	if w > len(b) {
		w = len(b)
	}

	// wを短くしながらマッチするまで。
	for ; w > 0; w-- {
		if bytes.Equal(a[len(a)-w:], b[:w]) {
			return w
		}
	}

	// まったくマッチしなかった。
	return 0
}

func main() {
}
