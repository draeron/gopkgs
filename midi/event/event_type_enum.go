// Code generated by go-enum
// DO NOT EDIT!

package event

import (
	"fmt"
)

const (
	// Aftertouch is a Type of type Aftertouch.
	Aftertouch Type = iota
	// ControlChange is a Type of type ControlChange.
	ControlChange
	// NoteOn is a Type of type NoteOn.
	NoteOn
	// NoteOff is a Type of type NoteOff.
	NoteOff
	// PitchBend is a Type of type PitchBend.
	PitchBend
	// PolyAftertouch is a Type of type PolyAftertouch.
	PolyAftertouch
	// ProgramChange is a Type of type ProgramChange.
	ProgramChange
	// SysEx is a Type of type SysEx.
	SysEx
)

const _TypeName = "AftertouchControlChangeNoteOnNoteOffPitchBendPolyAftertouchProgramChangeSysEx"

var _TypeMap = map[Type]string{
	0: _TypeName[0:10],
	1: _TypeName[10:23],
	2: _TypeName[23:29],
	3: _TypeName[29:36],
	4: _TypeName[36:45],
	5: _TypeName[45:59],
	6: _TypeName[59:72],
	7: _TypeName[72:77],
}

// String implements the Stringer interface.
func (x Type) String() string {
	if str, ok := _TypeMap[x]; ok {
		return str
	}
	return fmt.Sprintf("Type(%d)", x)
}

var _TypeValue = map[string]Type{
	_TypeName[0:10]:  0,
	_TypeName[10:23]: 1,
	_TypeName[23:29]: 2,
	_TypeName[29:36]: 3,
	_TypeName[36:45]: 4,
	_TypeName[45:59]: 5,
	_TypeName[59:72]: 6,
	_TypeName[72:77]: 7,
}

// ParseType attempts to convert a string to a Type
func ParseType(name string) (Type, error) {
	if x, ok := _TypeValue[name]; ok {
		return x, nil
	}
	return Type(0), fmt.Errorf("%s is not a valid Type", name)
}
