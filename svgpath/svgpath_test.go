package svgpath_test

import (
	"testing"

	"github.com/alecthomas/assert/v2"

	"github.com/twpayne/go-svg/svgpath"
)

func TestString(t *testing.T) {
	for _, tc := range []struct {
		name     string
		path     *svgpath.Path
		expected string
	}{
		{
			name: "empty",
		},
		{
			name: "simple",
			path: svgpath.New().
				MoveToAbs([]float64{200, 300}).
				LineToAbs([]float64{400, 50}).
				ClosePath(),
			expected: "M200,300 L400,50 z",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.path.String())
		})
	}
}
