package service

import (
	"testing"

	"github.com/ditramadia/lockleaf/internal/vault"
	"github.com/stretchr/testify/require"
)

func TestCreateCredential(t *testing.T) {
	cases := []struct {
		name              string
		vaultName         string
		credentialName    string
		credentialFields  map[string]vault.Field
		isCredentialExist bool
		wantErr           bool
	}{
		{"create credential successful", "titus", "bolt-pistol", map[string]vault.Field{}, false, false},
		{"credential exists", "titus", "bolt-pistol", map[string]vault.Field{}, true, true},
	}

	for _, c := range cases {
		tc := c
		t.Run(tc.name, func(t *testing.T) {
			s := newTestStorage(t)
			require.NoError(t, s.Save(newTestVault(tc.vaultName, nil)))

			if tc.isCredentialExist {
				v, err := s.Load(tc.vaultName)
				require.NoError(t, err)
				v.Credentials[tc.credentialName] = vault.Credential{
					Name:   tc.credentialName,
					Fields: tc.credentialFields,
				}
				require.NoError(t, s.Save(v))
			}

			err := newTestService(s).CreateCredential(tc.vaultName, tc.credentialName)
			if tc.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			v, err := s.Load(tc.vaultName)
			require.NoError(t, err)

			got := v.Credentials[tc.credentialName]
			require.Equal(t, tc.credentialName, got.Name)
			require.Equal(t, tc.credentialFields, got.Fields)
		})
	}
}

func TestListCredentials(t *testing.T) {
	cases := []struct {
		name        string
		vaultName   string
		credentials []string
	}{
		{"list credentials successful", "titus", []string{"bolt-pistol", "chainsword", "power-armour"}},
		{"list empty credentials", "titus", []string{}},
	}

	for _, c := range cases {
		tc := c
		t.Run(tc.name, func(t *testing.T) {
			s := newTestStorage(t)
			v := newTestVault(tc.vaultName, nil)
			for _, credName := range tc.credentials {
				v.Credentials[credName] = vault.Credential{
					Name:   credName,
					Fields: map[string]vault.Field{},
				}
			}
			require.NoError(t, s.Save(v))

			got, err := newTestService(s).ListCredentials(tc.vaultName)
			require.NoError(t, err)

			require.ElementsMatch(t, tc.credentials, got)
		})
	}
}

func TestRenameCredential(t *testing.T) {
	cases := []struct {
		name               string
		vaultName          string
		oldCredentialName  string
		newCredentialName  string
		fields             map[string]vault.Field
		isVaultExist       bool
		isCredentialExist  bool
		isNewNameAvailable bool
		wantErr            bool
	}{
		{"rename credential successful", "titus", "bolt-pistol", "chainsword", newTestFields(), true, true, true, false},
		{"vault does not exist", "titus", "bolt-pistol", "chainsword", newTestFields(), false, true, true, true},
		{"rename missing credential", "guilliman", "bolt-pistol", "chainsword", newTestFields(), true, false, true, true},
		{"new credential name already exists", "titus", "bolt-pistol", "chainsword", newTestFields(), true, true, false, true},
	}

	for _, c := range cases {
		tc := c
		t.Run(tc.name, func(t *testing.T) {
			s := newTestStorage(t)

			if tc.isVaultExist {
				v := newTestVault(tc.vaultName, nil)

				if tc.isCredentialExist {
					v.Credentials[tc.oldCredentialName] = vault.Credential{
						Name:   tc.oldCredentialName,
						Fields: tc.fields,
					}
				}

				if !tc.isNewNameAvailable {
					v.Credentials[tc.newCredentialName] = vault.Credential{
						Name:   tc.newCredentialName,
						Fields: tc.fields,
					}
				}

				require.NoError(t, s.Save(v))
			}

			err := newTestService(s).RenameCredential(tc.vaultName, tc.oldCredentialName, tc.newCredentialName)
			if tc.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			loadedV, err := s.Load(tc.vaultName)
			require.NoError(t, err)

			got := loadedV.Credentials[tc.newCredentialName]
			require.Equal(t, tc.newCredentialName, got.Name)
			require.Equal(t, tc.fields, got.Fields)
		})
	}
}
