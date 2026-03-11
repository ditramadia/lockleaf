package service

import (
	"testing"

	"github.com/ditramadia/lockleaf/internal/vault"
	"github.com/stretchr/testify/require"
)

func TestCreateCredential(t *testing.T) {
	cases := []struct {
		name             string
		vaultName        string
		credentialName   string
		credentialFields map[string]vault.Field
	}{
		{"create credential successful", "titus", "bolt-pistol", map[string]vault.Field{}},
	}

	for _, c := range cases {
		tc := c
		t.Run(tc.name, func(t *testing.T) {
			s := newTestStorage(t)
			require.NoError(t, s.Save(newTestVault(tc.vaultName, nil)))

			err := newTestService(s).CreateCredential(tc.vaultName, tc.credentialName)
			require.NoError(t, err)

			v, err := s.Load(tc.vaultName)
			require.NoError(t, err)

			got := v.Credentials[tc.credentialName]
			require.Equal(t, tc.credentialName, got.Name)
			require.Equal(t, tc.credentialFields, got.Fields)
		})
	}
}
