package main

import "golang.org/x/text/transform"

type Upper struct{ transform.NopResetter }

func (Upper) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	// 末尾ではないのにdstが足りない場合はErrShortDstを返す。
	if len(dst) == 0 && !atEOF {
		err = transform.ErrShortDst
		return
	}

	for {
		// srcをすべて処理し終えた、またはdstが足りなくなった場合。
		if len(src) <= nSrc || len(dst) <= nDst {
			return
		}

		// 小文字から大文字への変換。
		if 'a' <= src[nSrc] && src[nSrc] <= 'z' {
			dst[nDst] = src[nSrc] - 'a' + 'A'
		} else {
			dst[nDst] = src[nSrc]
		}

		// 処理したバイト数分だけ進める。
		nSrc++
		nDst++
	}
}

func main() {
}
