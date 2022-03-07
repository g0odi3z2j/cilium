// Code generated by "stringer -type=FSMState"; DO NOT EDIT.

package bgp

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[BGP_FSM_IDLE-0]
	_ = x[BGP_FSM_CONNECT-1]
	_ = x[BGP_FSM_ACTIVE-2]
	_ = x[BGP_FSM_OPENSENT-3]
	_ = x[BGP_FSM_OPENCONFIRM-4]
	_ = x[BGP_FSM_ESTABLISHED-5]
}

const _FSMState_name = "BGP_FSM_IDLEBGP_FSM_CONNECTBGP_FSM_ACTIVEBGP_FSM_OPENSENTBGP_FSM_OPENCONFIRMBGP_FSM_ESTABLISHED"

var _FSMState_index = [...]uint8{0, 12, 27, 41, 57, 76, 95}

func (i FSMState) String() string {
	if i < 0 || i >= FSMState(len(_FSMState_index)-1) {
		return "FSMState(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _FSMState_name[_FSMState_index[i]:_FSMState_index[i+1]]
}
