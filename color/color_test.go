package color_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/draeron/gopkgs/color"
)

func Test_Int32Convertion(t *testing.T) {

	org := color.Red.RGB()

	packed := color.ToInt32(org)

	t.Logf("packed: 0x%x", uint32(packed))

	unpacked, ok := color.FromInt32(packed).(color.RGB)
	require.True(t, ok)
	require.Equal(t, org, unpacked)
}
