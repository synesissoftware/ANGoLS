package angols_test

import (
	"github.com/stretchr/testify/require"
	angols "github.com/synesissoftware/ANGoLS"

	"testing"
)

const (
	Expected_VersionMajor uint16 = 0
	Expected_VersionMinor uint16 = 7
	Expected_VersionPatch uint16 = 0
	Expected_VersionAB    uint16 = 0x8001
)

func Test_Version_Elements(t *testing.T) {
	require.Equal(t, Expected_VersionMajor, angols.VersionMajor)
	require.Equal(t, Expected_VersionMinor, angols.VersionMinor)
	require.Equal(t, Expected_VersionPatch, angols.VersionPatch)
	require.Equal(t, Expected_VersionAB, angols.VersionAB)
}

func Test_Version(t *testing.T) {
	require.Equal(t, uint64(0x0000_0007_0000_8001), angols.Version)
}

func Test_Version_String(t *testing.T) {
	require.Equal(t, "0.7.0-beta1", angols.VersionString())
}
