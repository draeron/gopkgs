package event

// go tool github.com/abice/go-enum -f=$GOFILE --noprefix

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
