// Code generated by "stringer -linecomment -type TLSModel"; DO NOT EDIT.

package ir

import "strconv"

const _TLSModel_name = "noneinitialexeclocaldynamiclocalexec"

var _TLSModel_index = [...]uint8{0, 4, 15, 27, 36}

func (i TLSModel) String() string {
	if i >= TLSModel(len(_TLSModel_index)-1) {
		return "TLSModel(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TLSModel_name[_TLSModel_index[i]:_TLSModel_index[i+1]]
}
