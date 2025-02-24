package angols_test

import (
	strings "github.com/synesissoftware/ANGoLS/strings"
	stegol "github.com/synesissoftware/STEGoL"

	"testing"
)

func Test_StringChomp_1(t *testing.T) {

	stegol.CheckStringEqual(t, "", strings.StringChomp(""))
	stegol.CheckStringEqual(t, "", strings.StringChomp("\r"))
	stegol.CheckStringEqual(t, "", strings.StringChomp("\n"))
	stegol.CheckStringEqual(t, "", strings.StringChomp("\r\n"))
	stegol.CheckStringEqual(t, "\r", strings.StringChomp("\r\r"))
	stegol.CheckStringEqual(t, "\n", strings.StringChomp("\n\r"))
	stegol.CheckStringEqual(t, "\n", strings.StringChomp("\n\n"))
	stegol.CheckStringEqual(t, "\n", strings.StringChomp("\n\r\n"))

	stegol.CheckStringEqual(t, "abc", strings.StringChomp("abc"))
	stegol.CheckStringEqual(t, "abc", strings.StringChomp("abc\r"))
	stegol.CheckStringEqual(t, "abc", strings.StringChomp("abc\n"))
	stegol.CheckStringEqual(t, "abc", strings.StringChomp("abc\r\n"))
	stegol.CheckStringEqual(t, "abc\r", strings.StringChomp("abc\r\r"))
	stegol.CheckStringEqual(t, "abc\n", strings.StringChomp("abc\n\r"))
	stegol.CheckStringEqual(t, "abc\n", strings.StringChomp("abc\n\n"))
	stegol.CheckStringEqual(t, "abc\n", strings.StringChomp("abc\n\r\n"))
}

func Test_StringChompAll_1(t *testing.T) {

	stegol.CheckStringEqual(t, "", strings.StringChompAll(""))
	stegol.CheckStringEqual(t, "", strings.StringChompAll("\r"))
	stegol.CheckStringEqual(t, "", strings.StringChompAll("\n"))
	stegol.CheckStringEqual(t, "", strings.StringChompAll("\r\n"))
	stegol.CheckStringEqual(t, "", strings.StringChompAll("\r\r"))
	stegol.CheckStringEqual(t, "", strings.StringChompAll("\n\r"))
	stegol.CheckStringEqual(t, "", strings.StringChompAll("\n\n"))
	stegol.CheckStringEqual(t, "", strings.StringChompAll("\n\r\n"))

	stegol.CheckStringEqual(t, "abc", strings.StringChompAll("abc"))
	stegol.CheckStringEqual(t, "abc", strings.StringChompAll("abc\r"))
	stegol.CheckStringEqual(t, "abc", strings.StringChompAll("abc\n"))
	stegol.CheckStringEqual(t, "abc", strings.StringChompAll("abc\r\n"))
	stegol.CheckStringEqual(t, "abc", strings.StringChompAll("abc\r\r"))
	stegol.CheckStringEqual(t, "abc", strings.StringChompAll("abc\n\r"))
	stegol.CheckStringEqual(t, "abc", strings.StringChompAll("abc\n\n"))
	stegol.CheckStringEqual(t, "abc", strings.StringChompAll("abc\n\r\n"))
}
