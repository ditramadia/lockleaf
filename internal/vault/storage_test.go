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
				v := NewVault(tc.vaultName)
				err := s.Save(v)
				require.NoError(t, err)
			}

			got, err := s.IsVaultExists(tc.vaultName)
			require.NoError(t, err)

			require.Equal(t, tc.want, got)
		})
	}
}

func TestSaveAndLoadVault(t *testing.T) {
	cases := []struct {
		name       string
		vaultName  string
		credential *Credential
		save       bool
		wantErr    bool
	}{
		{"save and load vault successful", "test-vault", &Credential{
			Name: "Github",
			Fields: map[string]Field{
				"username": {Label: "username", Value: "gopher", IsSecret: false},
				"password": {Label: "password", Value: "s3cr3t", IsSecret: true},
			},
		}, true, false},
		{"load missing vault", "missing-vault", nil, false, true},
	}

	for _, c := range cases {
		tc := c
		t.Run(tc.name, func(t *testing.T) {
			s := setupStorage(t)

			if tc.save {
				v := setupVault(tc.vaultName, tc.credential)
				err := s.Save(v)
				require.NoError(t, err)
			}

			loadedV, err := s.Load(tc.vaultName)

			if tc.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.vaultName, loadedV.Name)
				if tc.credential != nil {
					cred, ok := loadedV.Credentials[tc.credential.Name]
					require.True(t, ok)
					require.Equal(t, *tc.credential, cred)
				}
			}
		})
	}
}

func setupStorage(t *testing.T) *Storage {
	tmpDir := t.TempDir()
	storage := NewStorage(tmpDir)

	return storage
}

func setupVault(vaultName string, credential *Credential) *Vault {
	v := NewVault(vaultName)
	v.Credentials[credential.Name] = *credential

	return v
}
