package doh_test

import (
	"testing"

	"github.com/stellaraf/edl/lib/doh"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_New(t *testing.T) {
	t.Run("client setup", func(t *testing.T) {
		t.Parallel()
		client := doh.New()
		assert.NotNil(t, client)
	})
}

func TestClient_A(t *testing.T) {
	t.Run("query", func(t *testing.T) {
		expected := []string{"1.1.1.1", "1.0.0.1"}
		client := doh.New()
		result, err := client.A("one.one.one.one")
		require.NoError(t, err)
		assert.IsType(t, &doh.DOHResponse{}, result)
		answers := []string{}
		for _, answer := range result.Answer {
			answers = append(answers, answer.Data)
		}
		for _, answer := range answers {
			assert.Contains(t, expected, answer)
		}
	})
}
