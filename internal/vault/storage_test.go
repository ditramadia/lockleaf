package vault

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsVaultExists(t *testing.T) {
	cases := []struct {
		name         string
		vaultName    string
		isVaultExist bool
		want         bool
	}{
		{"vault exists", "titus", true, true},
		{"vault missing", "metaurus", false, false},
	}

	for _, c := range cases {
		tc := c
		t.Run(tc.name, func(t *testing.T) {
			s := newTestStorage(t)

			if tc.isVaultExist {
				require.NoError(t, s.Save(newTestVault(tc.vaultName, nil)))
			}

			got, err := s.IsVaultExist(tc.vaultName)
			require.NoError(t, err)

			require.Equal(t, tc.want, got)
		})
	}
}

func TestSaveAndLoad(t *testing.T) {
	cases := []struct {
		name         string
		vaultName    string
		credentials  map[string]Credential
		isVaultExist bool
		wantErr      bool
	}{
		{"save and load vault successful", "titus", newTestCredentials(), true, false},
		{"load missing vault", "guilliman", nil, false, true},
	}

	for _, c := range cases {
		tc := c
		t.Run(tc.name, func(t *testing.T) {
			s := newTestStorage(t)

			if tc.isVaultExist {
				require.NoError(t, s.Save(newTestVault(tc.vaultName, tc.credentials)))
			}

			got, err := s.Load(tc.vaultName)

			if tc.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tc.vaultName, got.Name)
			if tc.credentials != nil {
				require.Equal(t, tc.credentials, got.Credentials)
			}
		})
	}
}

func TestList(t *testing.T) {
	cases := []struct {
		name string
		want []string
	}{
		{"list vaults successful", []string{"metaurus", "icaron", "levantus", "titus"}},
	}

	for _, c := range cases {
		tc := c
		t.Run(tc.name, func(t *testing.T) {
			s := newTestStorage(t)
			for _, vName := range tc.want {
				require.NoError(t, s.Save(newTestVault(vName, nil)))
			}

			got, err := s.List()
			require.NoError(t, err)

			require.ElementsMatch(t, tc.want, got)
		})
	}
}

func TestRename(t *testing.T) {
	cases := []struct {
		name         string
		oldVaultName string
		newVaultName string
		credentials  map[string]Credential
	}{
		{"rename vault successful", "metasaur", "titus", newTestCredentials()},
	}

	for _, c := range cases {
		tc := c
		t.Run(tc.name, func(t *testing.T) {
			s := newTestStorage(t)
			require.NoError(t, s.Save(newTestVault(tc.oldVaultName, tc.credentials)))

			err := s.Rename(tc.oldVaultName, tc.newVaultName)

			require.NoError(t, err)

			loadedVault, err := s.Load(tc.newVaultName)
			require.NoError(t, err)

			require.Equal(t, tc.newVaultName, loadedVault.Name)
			require.Equal(t, tc.credentials, loadedVault.Credentials)

			_, err = s.Load(tc.oldVaultName)
			require.Error(t, err)
		})
	}
}

func TestRemove(t *testing.T) {
	cases := []struct {
		name         string
		vaultName    string
		isVaultExist bool
		wantErr      bool
	}{
		{"remove vault successful", "titus", true, false},
		{"remove missing vault", "titus", false, true},
	}

	for _, c := range cases {
		tc := c
		t.Run(tc.name, func(t *testing.T) {
			s := newTestStorage(t)

			if tc.isVaultExist {
				require.NoError(t, s.Save(newTestVault(tc.vaultName, nil)))
			}

			err := s.Remove(tc.vaultName)

			if tc.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			got, err := s.IsVaultExist(tc.vaultName)
			require.NoError(t, err)
			require.False(t, got)
		})
	}
}

// Internal helpers

func newTestStorage(t *testing.T) *Storage {
	tmpDir := t.TempDir()
	storage := NewStorage(tmpDir)

	return storage
}

func newTestVault(vaultName string, credentials map[string]Credential) *Vault {
	v := NewVault(vaultName)
	if credentials != nil {
		v.Credentials = credentials
	}

	return v
}

func newTestCredentials() map[string]Credential {
	c := &Credential{
		Name: "Github",
		Fields: map[string]Field{
			"username": {Label: "username", Value: "gopher", IsSecret: false},
			"password": {Label: "password", Value: "s3cr3t", IsSecret: true},
		},
	}

	return map[string]Credential{
		c.Name: *c,
	}
}
