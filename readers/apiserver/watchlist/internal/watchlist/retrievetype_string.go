// Code generated by "stringer -type=RetrieveType -linecomment"; DO NOT EDIT.

package watchlist

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[RTNode-1]
	_ = x[RTPod-2]
	_ = x[RTNamespace-4]
	_ = x[RTPersistentVolume-8]
	_ = x[RTRBAC-16]
}

const (
	_RetrieveType_name_0 = "NodePod"
	_RetrieveType_name_1 = "Namespace"
	_RetrieveType_name_2 = "PersistentVolume"
	_RetrieveType_name_3 = "RBAC"
)

var (
	_RetrieveType_index_0 = [...]uint8{0, 4, 7}
)

func (i RetrieveType) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _RetrieveType_name_0[_RetrieveType_index_0[i]:_RetrieveType_index_0[i+1]]
	case i == 4:
		return _RetrieveType_name_1
	case i == 8:
		return _RetrieveType_name_2
	case i == 16:
		return _RetrieveType_name_3
	default:
		return "RetrieveType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
