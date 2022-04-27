package midi

import (
	"github.com/pkg/errors"
	"gitlab.com/gomidi/midi/v2"
)

type Port struct {
	In   int
	Out  int
	send SendFunc
	stop func()
}

type SendFunc func(msg midi.Message) error
type ListenFunc func(msg midi.Message, timestampms int32)

func (m *Port) Send(msg midi.Message) error {
	if m.send != nil {
		return m.send(msg)
	}
	return errors.New("cannot send message to close port")
}

func (m *Port) Close() {
	if m.stop != nil {
		m.stop()
	}
}

func (m *Port) Open(listen ListenFunc) error {
	var err error

	m.stop, err = midi.ListenTo(m.In, listen)
	if err != nil {
		return errors.WithStack(err)
	}

	m.send, err = midi.SendTo(m.Out)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
