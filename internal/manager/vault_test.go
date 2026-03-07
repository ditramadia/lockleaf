package manager

import (
	"testing"

	"github.com/ditramadia/lockleaf/internal/vault"
	"github.com/stretchr/testify/require"
)

func TestCreateVault(t *testing.T) {
	cases := []struct {
		name      string
		vaultName string
		exists    bool
		wantErr   bool
	}{
		{"create vault successful", "titus", false, false},
		{"vault exists", "titus", true, true},
	}

	for _, c := range cases {
		tc := c
		t.Run(tc.name, func(t *testing.T) {
			s := newTestStorage(t)

			if tc.exists {
				require.NoError(t, s.Save(newTestVault(tc.vaultName, nil)))
			}

			err := newTestManager(s).CreateVault(tc.vaultName)
			if tc.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			got, err := s.Load(tc.vaultName)
			require.NoError(t, err)
			require.Equal(t, tc.vaultName, got.Name)
		})
	}
}

func TestListVaults(t *testing.T) {
	cases := []struct {
		name   string
		vaults []string
	}{
		{"list vaults successful", []string{"metaurus", "icaron", "levantus", "titus"}},
		{"list empty vaults", []string{}},
	}

	for _, c := range cases {
		tc := c
		t.Run(tc.name, func(t *testing.T) {
			s := newTestStorage(t)
			for _, vName := range tc.vaults {
				require.NoError(t, s.Save(newTestVault(vName, nil)))
			}

			got, err := newTestManager(s).ListVaults()
			require.NoError(t, err)

			require.ElementsMatch(t, tc.vaults, got)
		})
	}
}

func TestRenameVaults(t *testing.T) {
	cases := []struct {
		name               string
		oldVaultName       string
		newVaultName       string
		credentials        map[string]vault.Credential
		isVaultExist       bool
		isNewNameAvailable bool
		wantErr            bool
	}{
		{"rename vault successful", "metaurus", "titus", newTestCredentials(), true, true, false},
		{"rename missing vault", "guilliman", "titus", newTestCredentials(), false, true, true},
		{"new name unavailable", "metaurus", "titus", newTestCredentials(), true, false, true},
	}

	for _, c := range cases {
		tc := c
		t.Run(tc.name, func(t *testing.T) {
			s := newTestStorage(t)

			if tc.isVaultExist {
				require.NoError(t, s.Save(newTestVault(tc.oldVaultName, tc.credentials)))
			}

			if !tc.isNewNameAvailable {
				require.NoError(t, s.Save(newTestVault(tc.newVaultName, tc.credentials)))
			}

			err := newTestManager(s).RenameVault(tc.oldVaultName, tc.newVaultName)

			if tc.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			got, err := s.Load(tc.newVaultName)
			require.NoError(t, err)
			require.Equal(t, tc.newVaultName, got.Name)
			require.Equal(t, tc.credentials, got.Credentials)
			_, err = s.Load(tc.oldVaultName)
			require.Error(t, err)
		})
	}
}

// Internal helpers

func newTestStorage(t *testing.T) *vault.Storage {
	return vault.NewStorage(t.TempDir())
}

func newTestManager(s *vault.Storage) *Manager {
	return NewManager(s)
}

func newTestVault(vaultName string, credentials map[string]vault.Credential) *vault.Vault {
	v := vault.NewVault(vaultName)
	if credentials != nil {
		v.Credentials = credentials
	}

	return v
}

func newTestCredentials() map[string]vault.Credential {
	c := &vault.Credential{
		Name: "Github",
		Fields: map[string]vault.Field{
			"username": {Label: "username", Value: "gopher", IsSecret: false},
			"password": {Label: "password", Value: "s3cr3t", IsSecret: true},
		},
	}

	return map[string]vault.Credential{
		c.Name: *c,
	}
}
