package event

//go:generate go-enum -f=$GOFILE --noprefix

// Event x ENUM(
/*
	Aftertouch
	ControlChange
	NoteOn
	NoteOff
	PitchBend
	PolyAftertouch
	ProgramChange
  SysEx
*/
// )
type Type int
