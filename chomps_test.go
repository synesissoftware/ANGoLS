
package angols_test

import (

	angols "github.com/synesissoftware/ANGoLS"
	stegol "github.com/synesissoftware/STEGoL"

	"testing"
)

func Test_StringChomp_1(t *testing.T) {

	stegol.CheckStringEqual(t, "", angols.StringChomp(""))
	stegol.CheckStringEqual(t, "", angols.StringChomp("\r"))
	stegol.CheckStringEqual(t, "", angols.StringChomp("\n"))
	stegol.CheckStringEqual(t, "", angols.StringChomp("\r\n"))
	stegol.CheckStringEqual(t, "\r", angols.StringChomp("\r\r"))
	stegol.CheckStringEqual(t, "\n", angols.StringChomp("\n\r"))
	stegol.CheckStringEqual(t, "\n", angols.StringChomp("\n\n"))
	stegol.CheckStringEqual(t, "\n", angols.StringChomp("\n\r\n"))

	stegol.CheckStringEqual(t, "abc", angols.StringChomp("abc"))
	stegol.CheckStringEqual(t, "abc", angols.StringChomp("abc\r"))
	stegol.CheckStringEqual(t, "abc", angols.StringChomp("abc\n"))
	stegol.CheckStringEqual(t, "abc", angols.StringChomp("abc\r\n"))
	stegol.CheckStringEqual(t, "abc\r", angols.StringChomp("abc\r\r"))
	stegol.CheckStringEqual(t, "abc\n", angols.StringChomp("abc\n\r"))
	stegol.CheckStringEqual(t, "abc\n", angols.StringChomp("abc\n\n"))
	stegol.CheckStringEqual(t, "abc\n", angols.StringChomp("abc\n\r\n"))
}

func Test_StringChompAll_1(t *testing.T) {

	stegol.CheckStringEqual(t, "", angols.StringChompAll(""))
	stegol.CheckStringEqual(t, "", angols.StringChompAll("\r"))
	stegol.CheckStringEqual(t, "", angols.StringChompAll("\n"))
	stegol.CheckStringEqual(t, "", angols.StringChompAll("\r\n"))
	stegol.CheckStringEqual(t, "", angols.StringChompAll("\r\r"))
	stegol.CheckStringEqual(t, "", angols.StringChompAll("\n\r"))
	stegol.CheckStringEqual(t, "", angols.StringChompAll("\n\n"))
	stegol.CheckStringEqual(t, "", angols.StringChompAll("\n\r\n"))

	stegol.CheckStringEqual(t, "abc", angols.StringChompAll("abc"))
	stegol.CheckStringEqual(t, "abc", angols.StringChompAll("abc\r"))
	stegol.CheckStringEqual(t, "abc", angols.StringChompAll("abc\n"))
	stegol.CheckStringEqual(t, "abc", angols.StringChompAll("abc\r\n"))
	stegol.CheckStringEqual(t, "abc", angols.StringChompAll("abc\r\r"))
	stegol.CheckStringEqual(t, "abc", angols.StringChompAll("abc\n\r"))
	stegol.CheckStringEqual(t, "abc", angols.StringChompAll("abc\n\n"))
	stegol.CheckStringEqual(t, "abc", angols.StringChompAll("abc\n\r\n"))
}

