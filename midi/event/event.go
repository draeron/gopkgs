package event

import (
	"fmt"

	"github.com/pkg/errors"
	"gitlab.com/gomidi/midi/v2"
	"gitlab.com/gomidi/midi/v2/sysex"
)

type Event struct {
	Type           Type
	Channel        uint8
	Controller     uint8
	ControllerName string
	Value          uint8
	valid          bool
}

func (e Event) IsValid() bool {
	return e.valid
}

func From(msg midi.Message) (Event, error) {

	var err error
	evt := Event{
		valid: true,
	}

	// var ignored8 uint8
	var ignored16 int16
	var ignoredu16 uint16

	switch {
	case msg.Is(midi.SysExMsg):
		evt.Type = SysEx
		bytes := []byte{}
		msg.GetSysEx(&bytes)
		info, err := sysex.Parse(bytes)
		if err != nil {
			return evt, errors.WithStack(err)
		}
		log.Infof("%s")
		evt.ControllerName = info.ManufacturerID.String()
		// evt.Controller = info.ModelID

	case msg.GetControlChange(&evt.Channel, &evt.Controller, &evt.Value):
		evt.Type = ControlChange

	case msg.GetNoteOn(&evt.Channel, &evt.Controller, &evt.Value):
		evt.Type = NoteOn

	case msg.GetNoteOff(&evt.Channel, &evt.Controller, &evt.Value):
		evt.Type = NoteOff

	case msg.GetAfterTouch(&evt.Channel, &evt.Value):
		evt.Type = Aftertouch

	case msg.GetPolyAfterTouch(&evt.Channel, &evt.Controller, &evt.Value):
		evt.Type = PolyAftertouch

	case msg.GetPitchBend(&evt.Channel, &ignored16, &ignoredu16):
		evt.Type = PitchBend

	default:
		log.Warnf("unsupported message: %s\n", msg.String())
		err = errors.New("unsupported message")
		evt.valid = false
	}

	return evt, err
}

func (e Event) String() string {
	return fmt.Sprintf("%s - Channel: %d, Controller: %s (%d), Value: %d",
		e.Type,
		e.Channel,
		e.ControllerName,
		e.Controller,
		e.Value,
	)
}
