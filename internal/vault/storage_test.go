package vault

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsVaultExists(t *testing.T) {
	cases := []struct {
		name      string
		vaultName string
		save      bool
		want      bool
	}{
		{"vault exists", "titus", true, true},
		{"vault missing", "metaurus", false, false},
	}

	for _, c := range cases {
		tc := c
		t.Run(tc.name, func(t *testing.T) {
			s := setupStorage(t)

			if tc.save {
				require.NoError(t, s.Save(setupVault(tc.vaultName, nil)))
			}

			got, err := s.IsVaultExists(tc.vaultName)
			require.NoError(t, err)

			require.Equal(t, tc.want, got)
		})
	}
}

func TestSaveAndLoad(t *testing.T) {
	cases := []struct {
		name        string
		vaultName   string
		credentials map[string]Credential
		save        bool
		wantErr     bool
	}{
		{"save and load vault successful", "titus", setupCredentials(), true, false},
		{"load missing vault", "guilliman", nil, false, true},
	}

	for _, c := range cases {
		tc := c
		t.Run(tc.name, func(t *testing.T) {
			s := setupStorage(t)

			if tc.save {
				require.NoError(t, s.Save(setupVault(tc.vaultName, tc.credentials)))
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
			s := setupStorage(t)
			for _, vName := range tc.want {
				require.NoError(t, s.Save(setupVault(vName, nil)))
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
		save         bool
		exist        bool
		wantErr      bool
	}{
		{"rename vault successful", "metasaur", "titus", setupCredentials(), true, false, false},
		{"rename missing vault", "guilliman", "titus", setupCredentials(), false, false, true},
		{"rename into an existing vault", "metasaur", "titus", setupCredentials(), true, true, true},
	}

	for _, c := range cases {
		tc := c
		t.Run(tc.name, func(t *testing.T) {
			s := setupStorage(t)

			if tc.exist {
				require.NoError(t, s.Save(setupVault(tc.newVaultName, tc.credentials)))
			}

			if tc.save {
				require.NoError(t, s.Save(setupVault(tc.oldVaultName, tc.credentials)))
			}

			err := s.Rename(tc.oldVaultName, tc.newVaultName)

			if tc.wantErr {
				require.Error(t, err)
				return
			}

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

// Internal helpers

func setupStorage(t *testing.T) *Storage {
	tmpDir := t.TempDir()
	storage := NewStorage(tmpDir)

	return storage
}

func setupVault(vaultName string, credentials map[string]Credential) *Vault {
	v := NewVault(vaultName)
	if credentials != nil {
		v.Credentials = credentials
	}

	return v
}

func setupCredentials() map[string]Credential {
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
