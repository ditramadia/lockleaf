package service

import (
	"testing"

	"github.com/ditramadia/lockleaf/internal/vault"
	"github.com/stretchr/testify/require"
)

func TestCreateVault(t *testing.T) {
	cases := []struct {
		name         string
		vaultName    string
		isVaultExist bool
		wantErr      bool
	}{
		{"create vault successful", "titus", false, false},
		{"vault exists", "titus", true, true},
	}

	for _, c := range cases {
		tc := c
		t.Run(tc.name, func(t *testing.T) {
			s := newTestStorage(t)

			if tc.isVaultExist {
				require.NoError(t, s.Save(newTestVault(tc.vaultName, nil)))
			}

			err := newTestService(s).CreateVault(tc.vaultName)
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

			got, err := newTestService(s).ListVaults()
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

			err := newTestService(s).RenameVault(tc.oldVaultName, tc.newVaultName)

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

func TestRemoveVault(t *testing.T) {
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

			err := newTestService(s).RemoveVault(tc.vaultName)
			if tc.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			got, err := s.IsVaultExist(tc.name)
			require.NoError(t, err)
			require.False(t, got)
		})
	}
}

func TestIsVaultExist(t *testing.T) {
	cases := []struct {
		name         string
		vaultName    string
		isVaultExist bool
	}{
		{"vault exists", "titus", true},
		{"vault missing", "guilliman", false},
	}

	for _, c := range cases {
		tc := c
		t.Run(tc.name, func(t *testing.T) {
			s := newTestStorage(t)

			if tc.isVaultExist {
				require.NoError(t, s.Save(newTestVault(tc.vaultName, nil)))
			}

			got, err := newTestService(s).IsVaultExist(tc.vaultName)

			require.NoError(t, err)
			require.Equal(t, tc.isVaultExist, got)
		})
	}
}
