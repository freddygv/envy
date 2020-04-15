package safe

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNamespace_String(t *testing.T) {
	t.Parallel()

	ns := &Namespace{
		Name: "ns1",
		Content: map[string]Encrypted{
			"foo": []byte{1, 1, 1, 1, 1},
			"bar": []byte{2, 2, 2, 2, 2},
		},
	}
	s := ns.String()
	exp := "(ns1 [foo bar])"
	require.Equal(t, exp, s)
}
