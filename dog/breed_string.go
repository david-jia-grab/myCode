// Code generated by "stringer -type Breed"; DO NOT EDIT.

package dog

import "strconv"

const _Breed_name = "PoodleBeagleLabradorPug"

var _Breed_index = [...]uint8{0, 6, 12, 20, 23}

func (i Breed) String() string {
	if i < 0 || i >= Breed(len(_Breed_index)-1) {
		return "Breed(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Breed_name[_Breed_index[i]:_Breed_index[i+1]]
}
