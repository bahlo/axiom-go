// Code generated by "stringer -type=UserRole -linecomment -output=users_string.go"; DO NOT EDIT.

package axiom

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[RoleReadOnly-1]
	_ = x[RoleUser-2]
	_ = x[RoleAdmin-3]
	_ = x[RoleOwner-4]
}

const _UserRole_name = "read-onlyuseradminowner"

var _UserRole_index = [...]uint8{0, 9, 13, 18, 23}

func (i UserRole) String() string {
	i -= 1
	if i >= UserRole(len(_UserRole_index)-1) {
		return "UserRole(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _UserRole_name[_UserRole_index[i]:_UserRole_index[i+1]]
}