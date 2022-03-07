// Code generated by "stringer -type=nexthopFlag"; DO NOT EDIT.

package zebra

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[nexthopFlagActive-1]
	_ = x[nexthopFlagFIB-2]
	_ = x[nexthopFlagRecursive-4]
	_ = x[nexthopFlagOnlink-8]
	_ = x[nexthopFlagMatched-16]
	_ = x[nexthopFlagFiltered-32]
	_ = x[nexthopFlagDuplicate-64]
	_ = x[nexthopFlagEvpnRvtep-128]
}

const (
	_nexthopFlag_name_0 = "nexthopFlagActivenexthopFlagFIB"
	_nexthopFlag_name_1 = "nexthopFlagRecursive"
	_nexthopFlag_name_2 = "nexthopFlagOnlink"
	_nexthopFlag_name_3 = "nexthopFlagMatched"
	_nexthopFlag_name_4 = "nexthopFlagFiltered"
	_nexthopFlag_name_5 = "nexthopFlagDuplicate"
	_nexthopFlag_name_6 = "nexthopFlagEvpnRvtep"
)

var (
	_nexthopFlag_index_0 = [...]uint8{0, 17, 31}
)

func (i nexthopFlag) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _nexthopFlag_name_0[_nexthopFlag_index_0[i]:_nexthopFlag_index_0[i+1]]
	case i == 4:
		return _nexthopFlag_name_1
	case i == 8:
		return _nexthopFlag_name_2
	case i == 16:
		return _nexthopFlag_name_3
	case i == 32:
		return _nexthopFlag_name_4
	case i == 64:
		return _nexthopFlag_name_5
	case i == 128:
		return _nexthopFlag_name_6
	default:
		return "nexthopFlag(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
