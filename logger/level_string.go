package logger

import "fmt"

const _Level_name = "TRACEDEBUGINFOWARNERRORFATAL"

var _Level_index = [...]uint8{0, 5, 10, 14, 18, 23, 28}

func (i Level) String() string {
	if i < 0 || i+1 >= Level(len(_Level_index)) {
		return fmt.Sprintf("Level(%d)", i)
	}
	return _Level_name[_Level_index[i]:_Level_index[i+1]]
}
