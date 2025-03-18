package static_test

import (
	"testing"

	"github.com/stellaraf/edl/static"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_GetList(t *testing.T) {
	t.Parallel()
	list, err := static.GetList("test")
	require.NoError(t, err)
	exp := []byte(`test1
test2
`)
	assert.Equal(t, exp, list)
}
