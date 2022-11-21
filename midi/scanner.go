package midi

import (
	"gitlab.com/gomidi/midi/v2"
	"regexp"

	"github.com/pkg/errors"
	"gitlab.com/gomidi/midi/v2/drivers"
	_ "gitlab.com/gomidi/midi/v2/drivers/rtmididrv"
)

func DetectPort(inRx, outRx *regexp.Regexp) (*Port, error) {

	port := Port{}

	ins, err := drivers.Ins()
	if err != nil {
		return nil, errors.WithMessage(err, "failed to list IN ports")
	}
	for _, in := range ins {
		if inRx.MatchString(in.String()) {
			port.In, err = midi.InPort(in.Number())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			log.Infof("matched midi IN: %s", in)
		} else {
			log.Debugf("ignored input device: %s", in)
		}
	}

	outs, err := drivers.Outs()
	if err != nil {
		return nil, errors.WithMessage(err, "failed to list OUT ports")
	}
	for _, out := range outs {
		if outRx.MatchString(out.String()) {
			port.Out, err = midi.OutPort(out.Number())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			log.Infof("matched midi OUT: %s", out)
		} else {
			log.Debugf("ignored output device: %s", out)
		}
	}

	if port.In == nil || port.Out == nil {
		return nil, errors.Errorf("failed to detect midi ports")
	}

	return &port, nil
}
