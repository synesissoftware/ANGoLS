package angols_test

import (
	angols "github.com/synesissoftware/ANGoLS"

	"github.com/stretchr/testify/require"

	"testing"
)

const (
	Expected_VersionMajor uint16 = 0
	Expected_VersionMinor uint16 = 9
	Expected_VersionPatch uint16 = 1
	Expected_VersionAB    uint16 = 0xFFFF
)

func Test_Version_Elements(t *testing.T) {
	require.Equal(t, Expected_VersionMajor, angols.VersionMajor)
	require.Equal(t, Expected_VersionMinor, angols.VersionMinor)
	require.Equal(t, Expected_VersionPatch, angols.VersionPatch)
	require.Equal(t, Expected_VersionAB, angols.VersionAB)
}

func Test_Version(t *testing.T) {
	require.Equal(t, uint64(0x0000_0009_0001_FFFF), angols.Version)
}

func Test_Version_String(t *testing.T) {
	require.Equal(t, "0.9.1", angols.VersionString())
}
