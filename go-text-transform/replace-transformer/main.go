package main

import "golang.org/x/text/transform"

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

	return
}

func main() {
}
