package angols_test

import (
	angols "github.com/synesissoftware/ANGoLS"

	"github.com/stretchr/testify/require"

	"testing"
)

const (
	Expected_VersionMajor uint16 = 0
	Expected_VersionMinor uint16 = 10
	Expected_VersionPatch uint16 = 0
	Expected_VersionAB    uint16 = 0x8002
)

func Test_Version_Elements(t *testing.T) {
	require.Equal(t, Expected_VersionMajor, angols.VersionMajor)
	require.Equal(t, Expected_VersionMinor, angols.VersionMinor)
	require.Equal(t, Expected_VersionPatch, angols.VersionPatch)
	require.Equal(t, Expected_VersionAB, angols.VersionAB)
}

func Test_Version(t *testing.T) {
	require.Equal(t, uint64(0x0000_000A_0000_8002), angols.Version)
}

func Test_Version_String(t *testing.T) {
	require.Equal(t, "0.10.0-beta2", angols.VersionString())
}
