package strings_test

import (
	strings "github.com/synesissoftware/ANGoLS/strings"

	"github.com/stretchr/testify/require"

	"testing"
)

func Test_SplitAfterByte_1(t *testing.T) {

	{
		require.Equal(t, []string{}, strings.SplitAfterByte("", ';'))
		require.Equal(t, []string{}, strings.SplitAfterByte("", '/'))
	}

	{
		require.Equal(t, []string{
			";",
		}, strings.SplitAfterByte(";", ';'))
	}

	{
		require.Equal(t, []string{
			"a;",
		}, strings.SplitAfterByte("a;", ';'))
	}

	{
		require.Equal(t, []string{
			";",
			"a",
		}, strings.SplitAfterByte(";a", ';'))
	}

	{
		require.Equal(t, []string{
			"a;",
			"a",
		}, strings.SplitAfterByte("a;a", ';'))
	}

	{
		require.Equal(t, []string{
			"abcdef:;",
			"ghi",
		}, strings.SplitAfterByte("abcdef:;ghi", ';'))
	}

	{
		require.Equal(t, []string{
			"🐻;",
			"🐻‍❄️",
		}, strings.SplitAfterByte("🐻;🐻‍❄️", ';'))
	}

	{
		require.Equal(t, []string{
			"a;", "b;", "c;", "d;", "e;", "f;", "g;", "h;", "i;", "j;", "k;", "l;", "m;", "n;", "o;", "p;", "q;", "r;", "s;", "t;", "u;", "v;", "w;", "x;", "y;", "z",
		}, strings.SplitAfterByte("a;b;c;d;e;f;g;h;i;j;k;l;m;n;o;p;q;r;s;t;u;v;w;x;y;z", ';'))
	}
}

func Test_SplitAfterByteN_1(t *testing.T) {

	{
		require.Equal(t, []string{}, strings.SplitAfterByteN("", ';', 3))
		require.Equal(t, []string{}, strings.SplitAfterByteN("", '/', 3))
	}

	{
		require.Equal(t, []string{
			";",
		}, strings.SplitAfterByteN(";", ';', -1))
		require.Equal(t, []string{}, strings.SplitAfterByteN(";", ';', 0))
		require.Equal(t, []string{
			";",
		}, strings.SplitAfterByteN(";", ';', 1))
		require.Equal(t, []string{
			";",
		}, strings.SplitAfterByteN(";", ';', 2))
	}

	{
		require.Equal(t, []string{
			"a;",
		}, strings.SplitAfterByteN("a;", ';', -1))
		require.Equal(t, []string{}, strings.SplitAfterByteN("a;", ';', 0))
		require.Equal(t, []string{
			"a;",
		}, strings.SplitAfterByteN("a;", ';', 1))
		require.Equal(t, []string{
			"a;",
		}, strings.SplitAfterByteN("a;", ';', 2))
	}

	{
		require.Equal(t, []string{
			";",
			"a",
		}, strings.SplitAfterByteN(";a", ';', -1))
		require.Equal(t, []string{}, strings.SplitAfterByteN(";a", ';', 0))
		require.Equal(t, []string{
			";a",
		}, strings.SplitAfterByteN(";a", ';', 1))
		require.Equal(t, []string{
			";",
			"a",
		}, strings.SplitAfterByteN(";a", ';', 2))
		require.Equal(t, []string{
			";",
			"a",
		}, strings.SplitAfterByteN(";a", ';', 3))
	}

	{
		require.Equal(t, []string{
			";",
			"a;",
		}, strings.SplitAfterByteN(";a;", ';', -1))
		require.Equal(t, []string{}, strings.SplitAfterByteN(";a;", ';', 0))
		require.Equal(t, []string{
			";a;",
		}, strings.SplitAfterByteN(";a;", ';', 1))
		require.Equal(t, []string{
			";",
			"a;",
		}, strings.SplitAfterByteN(";a;", ';', 2))
		require.Equal(t, []string{
			";",
			"a;",
		}, strings.SplitAfterByteN(";a;", ';', 3))
	}

	{
		require.Equal(t, []string{
			"a;",
			"a",
		}, strings.SplitAfterByteN("a;a", ';', -1))
		require.Equal(t, []string{}, strings.SplitAfterByteN("a;a", ';', 0))
		require.Equal(t, []string{
			"a;a",
		}, strings.SplitAfterByteN("a;a", ';', 1))
		require.Equal(t, []string{
			"a;",
			"a",
		}, strings.SplitAfterByteN("a;a", ';', 2))
		require.Equal(t, []string{
			"a;",
			"a",
		}, strings.SplitAfterByteN("a;a", ';', 3))
	}

	{
		require.Equal(t, []string{
			"abcdef:;",
			"ghi",
		}, strings.SplitAfterByteN("abcdef:;ghi", ';', -1))
		require.Equal(t, []string{}, strings.SplitAfterByteN("abcdef:;ghi", ';', 0))
		require.Equal(t, []string{
			"abcdef:;ghi",
		}, strings.SplitAfterByteN("abcdef:;ghi", ';', 1))
		require.Equal(t, []string{
			"abcdef:;",
			"ghi",
		}, strings.SplitAfterByteN("abcdef:;ghi", ';', 2))
		require.Equal(t, []string{
			"abcdef:;",
			"ghi",
		}, strings.SplitAfterByteN("abcdef:;ghi", ';', 3))
	}

	{
		require.Equal(t, []string{
			"🐻;",
			"🐻‍❄️",
		}, strings.SplitAfterByteN("🐻;🐻‍❄️", ';', -1))
		require.Equal(t, []string{}, strings.SplitAfterByteN("🐻;🐻‍❄️", ';', 0))
		require.Equal(t, []string{
			"🐻;🐻‍❄️",
		}, strings.SplitAfterByteN("🐻;🐻‍❄️", ';', 1))
		require.Equal(t, []string{
			"🐻;",
			"🐻‍❄️",
		}, strings.SplitAfterByteN("🐻;🐻‍❄️", ';', 2))
		require.Equal(t, []string{
			"🐻;",
			"🐻‍❄️",
		}, strings.SplitAfterByteN("🐻;🐻‍❄️", ';', 3))
	}

	{
		require.Equal(t, []string{
			"a;", "b;", "c;", "d;", "e;", "f;", "g;", "h;", "i;", "j;", "k;", "l;", "m;", "n;", "o;", "p;", "q;", "r;", "s;", "t;", "u;", "v;", "w;", "x;", "y;", "z",
		}, strings.SplitAfterByteN("a;b;c;d;e;f;g;h;i;j;k;l;m;n;o;p;q;r;s;t;u;v;w;x;y;z", ';', -1))
		require.Equal(t, []string{}, strings.SplitAfterByteN("a;b;c;d;e;f;g;h;i;j;k;l;m;n;o;p;q;r;s;t;u;v;w;x;y;z", ';', 0))
		require.Equal(t, []string{
			"a;b;c;d;e;f;g;h;i;j;k;l;m;n;o;p;q;r;s;t;u;v;w;x;y;z",
		}, strings.SplitAfterByteN("a;b;c;d;e;f;g;h;i;j;k;l;m;n;o;p;q;r;s;t;u;v;w;x;y;z", ';', 1))
		require.Equal(t, []string{
			"a;",
			"b;c;d;e;f;g;h;i;j;k;l;m;n;o;p;q;r;s;t;u;v;w;x;y;z",
		}, strings.SplitAfterByteN("a;b;c;d;e;f;g;h;i;j;k;l;m;n;o;p;q;r;s;t;u;v;w;x;y;z", ';', 2))
		require.Equal(t, []string{
			"a;",
			"b;",
			"c;d;e;f;g;h;i;j;k;l;m;n;o;p;q;r;s;t;u;v;w;x;y;z",
		}, strings.SplitAfterByteN("a;b;c;d;e;f;g;h;i;j;k;l;m;n;o;p;q;r;s;t;u;v;w;x;y;z", ';', 3))

		require.Equal(t, []string{
			"a;", "b;", "c;", "d;", "e;", "f;", "g;", "h;", "i;", "j;", "k;", "l;", "m;", "n;", "o;", "p;", "q;", "r;", "s;", "t;", "u;", "v;", "w;", "x;", "y;z",
		}, strings.SplitAfterByteN("a;b;c;d;e;f;g;h;i;j;k;l;m;n;o;p;q;r;s;t;u;v;w;x;y;z", ';', 25))
		require.Equal(t, []string{
			"a;", "b;", "c;", "d;", "e;", "f;", "g;", "h;", "i;", "j;", "k;", "l;", "m;", "n;", "o;", "p;", "q;", "r;", "s;", "t;", "u;", "v;", "w;", "x;", "y;", "z",
		}, strings.SplitAfterByteN("a;b;c;d;e;f;g;h;i;j;k;l;m;n;o;p;q;r;s;t;u;v;w;x;y;z", ';', 26))
		require.Equal(t, []string{
			"a;", "b;", "c;", "d;", "e;", "f;", "g;", "h;", "i;", "j;", "k;", "l;", "m;", "n;", "o;", "p;", "q;", "r;", "s;", "t;", "u;", "v;", "w;", "x;", "y;", "z",
		}, strings.SplitAfterByteN("a;b;c;d;e;f;g;h;i;j;k;l;m;n;o;p;q;r;s;t;u;v;w;x;y;z", ';', 27))
	}
}

func Test_SplitAfterRune_1(t *testing.T) {

	{
		require.Equal(t, []string{}, strings.SplitAfterRune("", ';'))
		require.Equal(t, []string{}, strings.SplitAfterRune("", '/'))
	}

	{
		require.Equal(t, []string{
			";",
		}, strings.SplitAfterRune(";", ';'))
	}

	{
		require.Equal(t, []string{
			"a;",
		}, strings.SplitAfterRune("a;", ';'))
	}

	{
		require.Equal(t, []string{
			";",
			"a",
		}, strings.SplitAfterRune(";a", ';'))
	}

	{
		require.Equal(t, []string{
			"a;",
			"a",
		}, strings.SplitAfterRune("a;a", ';'))
	}

	{
		require.Equal(t, []string{
			"abcdef:;",
			"ghi",
		}, strings.SplitAfterRune("abcdef:;ghi", ';'))
	}

	{
		require.Equal(t, []string{
			"🐻;",
			"🐻‍❄️",
		}, strings.SplitAfterRune("🐻;🐻‍❄️", ';'))
	}

	{
		require.Equal(t, []string{
			"a;", "b;", "c;", "d;", "e;", "f;", "g;", "h;", "i;", "j;", "k;", "l;", "m;", "n;", "o;", "p;", "q;", "r;", "s;", "t;", "u;", "v;", "w;", "x;", "y;", "z",
		}, strings.SplitAfterRune("a;b;c;d;e;f;g;h;i;j;k;l;m;n;o;p;q;r;s;t;u;v;w;x;y;z", ';'))
	}
}

func Test_SplitAfterRuneN_1(t *testing.T) {

	{
		require.Equal(t, []string{}, strings.SplitAfterRuneN("", ';', 3))
		require.Equal(t, []string{}, strings.SplitAfterRuneN("", '/', 3))
	}

	{
		require.Equal(t, []string{
			";",
		}, strings.SplitAfterRuneN(";", ';', -1))
		require.Equal(t, []string{}, strings.SplitAfterRuneN(";", ';', 0))
		require.Equal(t, []string{
			";",
		}, strings.SplitAfterRuneN(";", ';', 1))
		require.Equal(t, []string{
			";",
		}, strings.SplitAfterRuneN(";", ';', 2))
	}

	{
		require.Equal(t, []string{
			"a;",
		}, strings.SplitAfterRuneN("a;", ';', -1))
		require.Equal(t, []string{}, strings.SplitAfterRuneN("a;", ';', 0))
		require.Equal(t, []string{
			"a;",
		}, strings.SplitAfterRuneN("a;", ';', 1))
		require.Equal(t, []string{
			"a;",
		}, strings.SplitAfterRuneN("a;", ';', 2))
	}

	{
		require.Equal(t, []string{
			";",
			"a",
		}, strings.SplitAfterRuneN(";a", ';', -1))
		require.Equal(t, []string{}, strings.SplitAfterRuneN(";a", ';', 0))
		require.Equal(t, []string{
			";a",
		}, strings.SplitAfterRuneN(";a", ';', 1))
		require.Equal(t, []string{
			";",
			"a",
		}, strings.SplitAfterRuneN(";a", ';', 2))
		require.Equal(t, []string{
			";",
			"a",
		}, strings.SplitAfterRuneN(";a", ';', 3))
	}

	{
		require.Equal(t, []string{
			";",
			"a;",
		}, strings.SplitAfterRuneN(";a;", ';', -1))
		require.Equal(t, []string{}, strings.SplitAfterRuneN(";a;", ';', 0))
		require.Equal(t, []string{
			";a;",
		}, strings.SplitAfterRuneN(";a;", ';', 1))
		require.Equal(t, []string{
			";",
			"a;",
		}, strings.SplitAfterRuneN(";a;", ';', 2))
		require.Equal(t, []string{
			";",
			"a;",
		}, strings.SplitAfterRuneN(";a;", ';', 3))
	}

	{
		require.Equal(t, []string{
			"a;",
			"a",
		}, strings.SplitAfterRuneN("a;a", ';', -1))
		require.Equal(t, []string{}, strings.SplitAfterRuneN("a;a", ';', 0))
		require.Equal(t, []string{
			"a;a",
		}, strings.SplitAfterRuneN("a;a", ';', 1))
		require.Equal(t, []string{
			"a;",
			"a",
		}, strings.SplitAfterRuneN("a;a", ';', 2))
		require.Equal(t, []string{
			"a;",
			"a",
		}, strings.SplitAfterRuneN("a;a", ';', 3))
	}

	{
		require.Equal(t, []string{
			"abcdef:;",
			"ghi",
		}, strings.SplitAfterRuneN("abcdef:;ghi", ';', -1))
		require.Equal(t, []string{}, strings.SplitAfterRuneN("abcdef:;ghi", ';', 0))
		require.Equal(t, []string{
			"abcdef:;ghi",
		}, strings.SplitAfterRuneN("abcdef:;ghi", ';', 1))
		require.Equal(t, []string{
			"abcdef:;",
			"ghi",
		}, strings.SplitAfterRuneN("abcdef:;ghi", ';', 2))
		require.Equal(t, []string{
			"abcdef:;",
			"ghi",
		}, strings.SplitAfterRuneN("abcdef:;ghi", ';', 3))
	}

	{
		require.Equal(t, []string{
			"🐻;",
			"🐻‍❄️",
		}, strings.SplitAfterRuneN("🐻;🐻‍❄️", ';', -1))
		require.Equal(t, []string{}, strings.SplitAfterRuneN("🐻;🐻‍❄️", ';', 0))
		require.Equal(t, []string{
			"🐻;🐻‍❄️",
		}, strings.SplitAfterRuneN("🐻;🐻‍❄️", ';', 1))
		require.Equal(t, []string{
			"🐻;",
			"🐻‍❄️",
		}, strings.SplitAfterRuneN("🐻;🐻‍❄️", ';', 2))
		require.Equal(t, []string{
			"🐻;",
			"🐻‍❄️",
		}, strings.SplitAfterRuneN("🐻;🐻‍❄️", ';', 3))
	}

	{
		require.Equal(t, []string{
			"a;", "b;", "c;", "d;", "e;", "f;", "g;", "h;", "i;", "j;", "k;", "l;", "m;", "n;", "o;", "p;", "q;", "r;", "s;", "t;", "u;", "v;", "w;", "x;", "y;", "z",
		}, strings.SplitAfterRuneN("a;b;c;d;e;f;g;h;i;j;k;l;m;n;o;p;q;r;s;t;u;v;w;x;y;z", ';', -1))
		require.Equal(t, []string{}, strings.SplitAfterRuneN("a;b;c;d;e;f;g;h;i;j;k;l;m;n;o;p;q;r;s;t;u;v;w;x;y;z", ';', 0))
		require.Equal(t, []string{
			"a;b;c;d;e;f;g;h;i;j;k;l;m;n;o;p;q;r;s;t;u;v;w;x;y;z",
		}, strings.SplitAfterRuneN("a;b;c;d;e;f;g;h;i;j;k;l;m;n;o;p;q;r;s;t;u;v;w;x;y;z", ';', 1))
		require.Equal(t, []string{
			"a;",
			"b;c;d;e;f;g;h;i;j;k;l;m;n;o;p;q;r;s;t;u;v;w;x;y;z",
		}, strings.SplitAfterRuneN("a;b;c;d;e;f;g;h;i;j;k;l;m;n;o;p;q;r;s;t;u;v;w;x;y;z", ';', 2))
		require.Equal(t, []string{
			"a;",
			"b;",
			"c;d;e;f;g;h;i;j;k;l;m;n;o;p;q;r;s;t;u;v;w;x;y;z",
		}, strings.SplitAfterRuneN("a;b;c;d;e;f;g;h;i;j;k;l;m;n;o;p;q;r;s;t;u;v;w;x;y;z", ';', 3))

		require.Equal(t, []string{
			"a;", "b;", "c;", "d;", "e;", "f;", "g;", "h;", "i;", "j;", "k;", "l;", "m;", "n;", "o;", "p;", "q;", "r;", "s;", "t;", "u;", "v;", "w;", "x;", "y;z",
		}, strings.SplitAfterRuneN("a;b;c;d;e;f;g;h;i;j;k;l;m;n;o;p;q;r;s;t;u;v;w;x;y;z", ';', 25))
		require.Equal(t, []string{
			"a;", "b;", "c;", "d;", "e;", "f;", "g;", "h;", "i;", "j;", "k;", "l;", "m;", "n;", "o;", "p;", "q;", "r;", "s;", "t;", "u;", "v;", "w;", "x;", "y;", "z",
		}, strings.SplitAfterRuneN("a;b;c;d;e;f;g;h;i;j;k;l;m;n;o;p;q;r;s;t;u;v;w;x;y;z", ';', 26))
		require.Equal(t, []string{
			"a;", "b;", "c;", "d;", "e;", "f;", "g;", "h;", "i;", "j;", "k;", "l;", "m;", "n;", "o;", "p;", "q;", "r;", "s;", "t;", "u;", "v;", "w;", "x;", "y;", "z",
		}, strings.SplitAfterRuneN("a;b;c;d;e;f;g;h;i;j;k;l;m;n;o;p;q;r;s;t;u;v;w;x;y;z", ';', 27))
	}
}

func Test_SplitAfterAny_1(t *testing.T) {

	{
		require.Equal(t, []string{}, strings.SplitAfterAny("", ""))
		require.Equal(t, []string{}, strings.SplitAfterAny("", `/\`))
	}

	{
		require.Equal(t, []string{
			";",
		}, strings.SplitAfterAny(";", ""))
		require.Equal(t, []string{
			";",
		}, strings.SplitAfterAny(";", ";:"))
	}

	{
		require.Equal(t, []string{
			"a;",
		}, strings.SplitAfterAny("a;", ";:"))
	}

	{
		require.Equal(t, []string{
			";",
			"a",
		}, strings.SplitAfterAny(";a", ";:"))
	}

	{
		require.Equal(t, []string{
			"a;",
			"a",
		}, strings.SplitAfterAny("a;a", ";:"))
	}

	{
		require.Equal(t, []string{
			"abcdef:;ghi",
		}, strings.SplitAfterAny("abcdef:;ghi", ""))
		require.Equal(t, []string{
			"abcdef:",
			";",
			"ghi",
		}, strings.SplitAfterAny("abcdef:;ghi", ";:"))
	}

	{
		require.Equal(t, []string{
			"🐻;",
			"🐻‍❄️",
		}, strings.SplitAfterAny("🐻;🐻‍❄️", ";:"))
	}

	{
		require.Equal(t, []string{
			"a;", "b:", "c;", "d:", "e;", "f:", "g;", "h:", "i;", "j:", "k;", "l:", "m;", "n:", "o;", "p:", "q;", "r:", "s;", "t:", "u;", "v:", "w;", "x:", "y;", "z",
		}, strings.SplitAfterAny("a;b:c;d:e;f:g;h:i;j:k;l:m;n:o;p:q;r:s;t:u;v:w;x:y;z", ":;"))
	}
}

func Test_SplitAfterAnyN_1(t *testing.T) {

	{
		require.Equal(t, []string{}, strings.SplitAfterAnyN("", "", -1))
		require.Equal(t, []string{}, strings.SplitAfterAnyN("", `/\`, -1))
	}

	{
		require.Equal(t, []string{
			";",
		}, strings.SplitAfterAnyN(";", ";:", -1))
		require.Equal(t, []string{}, strings.SplitAfterAnyN(";", ";:", 0))
		require.Equal(t, []string{
			";",
		}, strings.SplitAfterAnyN(";", ";:", 1))
		require.Equal(t, []string{
			";",
		}, strings.SplitAfterAnyN(";", ";:", 2))
	}

	{
		require.Equal(t, []string{
			"a;",
		}, strings.SplitAfterAnyN("a;", ";:", -1))
		require.Equal(t, []string{}, strings.SplitAfterAnyN("a;", ";:", 0))
		require.Equal(t, []string{
			"a;",
		}, strings.SplitAfterAnyN("a;", ";:", 1))
		require.Equal(t, []string{
			"a;",
		}, strings.SplitAfterAnyN("a;", ";:", 2))
	}

	{
		require.Equal(t, []string{
			";",
			"a",
		}, strings.SplitAfterAnyN(";a", ";:", -1))
		require.Equal(t, []string{}, strings.SplitAfterAnyN(";a", ";:", 0))
		require.Equal(t, []string{
			";a",
		}, strings.SplitAfterAnyN(";a", ";:", 1))
		require.Equal(t, []string{
			";",
			"a",
		}, strings.SplitAfterAnyN(";a", ";:", 2))
		require.Equal(t, []string{
			";",
			"a",
		}, strings.SplitAfterAnyN(";a", ";:", 3))
	}

	{
		require.Equal(t, []string{
			"a;",
			"a",
		}, strings.SplitAfterAnyN("a;a", ";:", -1))
		require.Equal(t, []string{}, strings.SplitAfterAnyN("a;a", ";:", 0))
		require.Equal(t, []string{
			"a;a",
		}, strings.SplitAfterAnyN("a;a", ";:", 1))
		require.Equal(t, []string{
			"a;",
			"a",
		}, strings.SplitAfterAnyN("a;a", ";:", 2))
		require.Equal(t, []string{
			"a;",
			"a",
		}, strings.SplitAfterAnyN("a;a", ";:", 3))
	}

	{
		require.Equal(t, []string{
			"abcdef:;ghi",
		}, strings.SplitAfterAnyN("abcdef:;ghi", "", -1))

		require.Equal(t, []string{
			"abcdef:",
			";",
			"ghi",
		}, strings.SplitAfterAnyN("abcdef:;ghi", ";:", -1))
		require.Equal(t, []string{}, strings.SplitAfterAnyN("abcdef:;ghi", ";:", 0))
		require.Equal(t, []string{
			"abcdef:;ghi",
		}, strings.SplitAfterAnyN("abcdef:;ghi", ";:", 1))
		require.Equal(t, []string{
			"abcdef:",
			";ghi",
		}, strings.SplitAfterAnyN("abcdef:;ghi", ";:", 2))
		require.Equal(t, []string{
			"abcdef:",
			";",
			"ghi",
		}, strings.SplitAfterAnyN("abcdef:;ghi", ";:", 3))
		require.Equal(t, []string{
			"abcdef:",
			";",
			"ghi",
		}, strings.SplitAfterAnyN("abcdef:;ghi", ";:", 4))
	}

	{
		require.Equal(t, []string{
			"🐻;",
			"🐻‍❄️",
		}, strings.SplitAfterAnyN("🐻;🐻‍❄️", ";:", -1))
	}

	{
		require.Equal(t, []string{
			"a;", "b:", "c;", "d:", "e;", "f:", "g;", "h:", "i;", "j:", "k;", "l:", "m;", "n:", "o;", "p:", "q;", "r:", "s;", "t:", "u;", "v:", "w;", "x:", "y;", "z",
		}, strings.SplitAfterAnyN("a;b:c;d:e;f:g;h:i;j:k;l:m;n:o;p:q;r:s;t:u;v:w;x:y;z", ":;", -1))
		require.Equal(t, []string{}, strings.SplitAfterAnyN("a;b:c;d:e;f:g;h:i;j:k;l:m;n:o;p:q;r:s;t:u;v:w;x:y;z", ":;", 0))
		require.Equal(t, []string{
			"a;b:c;d:e;f:g;h:i;j:k;l:m;n:o;p:q;r:s;t:u;v:w;x:y;z",
		}, strings.SplitAfterAnyN("a;b:c;d:e;f:g;h:i;j:k;l:m;n:o;p:q;r:s;t:u;v:w;x:y;z", ":;", 1))
		require.Equal(t, []string{
			"a;",
			"b:c;d:e;f:g;h:i;j:k;l:m;n:o;p:q;r:s;t:u;v:w;x:y;z",
		}, strings.SplitAfterAnyN("a;b:c;d:e;f:g;h:i;j:k;l:m;n:o;p:q;r:s;t:u;v:w;x:y;z", ":;", 2))
		require.Equal(t, []string{
			"a;",
			"b:",
			"c;d:e;f:g;h:i;j:k;l:m;n:o;p:q;r:s;t:u;v:w;x:y;z",
		}, strings.SplitAfterAnyN("a;b:c;d:e;f:g;h:i;j:k;l:m;n:o;p:q;r:s;t:u;v:w;x:y;z", ":;", 3))
		require.Equal(t, []string{
			"a;", "b:", "c;", "d:", "e;", "f:", "g;", "h:", "i;", "j:", "k;", "l:", "m;", "n:", "o;", "p:", "q;", "r:", "s;", "t:", "u;", "v:", "w;", "x:", "y;z",
		}, strings.SplitAfterAnyN("a;b:c;d:e;f:g;h:i;j:k;l:m;n:o;p:q;r:s;t:u;v:w;x:y;z", ":;", 25))
		require.Equal(t, []string{
			"a;", "b:", "c;", "d:", "e;", "f:", "g;", "h:", "i;", "j:", "k;", "l:", "m;", "n:", "o;", "p:", "q;", "r:", "s;", "t:", "u;", "v:", "w;", "x:", "y;", "z",
		}, strings.SplitAfterAnyN("a;b:c;d:e;f:g;h:i;j:k;l:m;n:o;p:q;r:s;t:u;v:w;x:y;z", ":;", 26))
		require.Equal(t, []string{
			"a;", "b:", "c;", "d:", "e;", "f:", "g;", "h:", "i;", "j:", "k;", "l:", "m;", "n:", "o;", "p:", "q;", "r:", "s;", "t:", "u;", "v:", "w;", "x:", "y;", "z",
		}, strings.SplitAfterAnyN("a;b:c;d:e;f:g;h:i;j:k;l:m;n:o;p:q;r:s;t:u;v:w;x:y;z", ":;", 27))
	}
}

func Test_SplitAfterAnyBytes_1(t *testing.T) {

	{
		require.Equal(t, []string{}, strings.SplitAfterAnyBytes("", []byte{}))
		require.Equal(t, []string{}, strings.SplitAfterAnyBytes("", []byte{'/', '\\'}))
	}

	{
		require.Equal(t, []string{
			";",
		}, strings.SplitAfterAnyBytes(";", []byte{';', ':'}))
	}

	{
		require.Equal(t, []string{
			"a;",
		}, strings.SplitAfterAnyBytes("a;", []byte{';', ':'}))
	}

	{
		require.Equal(t, []string{
			";",
			"a",
		}, strings.SplitAfterAnyBytes(";a", []byte{';', ':'}))
	}

	{
		require.Equal(t, []string{
			"a;",
			"a",
		}, strings.SplitAfterAnyBytes("a;a", []byte{';', ':'}))
	}

	{
		require.Equal(t, []string{
			"abcdef:;ghi",
		}, strings.SplitAfterAnyBytes("abcdef:;ghi", []byte{}))
		require.Equal(t, []string{
			"abcdef:",
			";",
			"ghi",
		}, strings.SplitAfterAnyBytes("abcdef:;ghi", []byte{';', ':'}))
	}

	{
		require.Equal(t, []string{
			"🐻;",
			"🐻‍❄️",
		}, strings.SplitAfterAnyBytes("🐻;🐻‍❄️", []byte{';', ':'}))
	}

	{
		require.Equal(t, []string{
			"a;", "b:", "c;", "d:", "e;", "f:", "g;", "h:", "i;", "j:", "k;", "l:", "m;", "n:", "o;", "p:", "q;", "r:", "s;", "t:", "u;", "v:", "w;", "x:", "y;", "z",
		}, strings.SplitAfterAnyBytes("a;b:c;d:e;f:g;h:i;j:k;l:m;n:o;p:q;r:s;t:u;v:w;x:y;z", []byte{';', ':'}))
	}
}

func Test_SplitAfterAnyBytesN_1(t *testing.T) {

	{
		require.Equal(t, []string{}, strings.SplitAfterAnyBytesN("", []byte{}, -1))
		require.Equal(t, []string{}, strings.SplitAfterAnyBytesN("", []byte{'/', '\\'}, -1))
	}

	{
		require.Equal(t, []string{
			";",
		}, strings.SplitAfterAnyBytesN(";", []byte{';', ':'}, -1))
		require.Equal(t, []string{}, strings.SplitAfterAnyBytesN(";", []byte{';', ':'}, 0))
		require.Equal(t, []string{
			";",
		}, strings.SplitAfterAnyBytesN(";", []byte{';', ':'}, 1))
		require.Equal(t, []string{
			";",
		}, strings.SplitAfterAnyBytesN(";", []byte{';', ':'}, 2))
	}

	{
		require.Equal(t, []string{
			"a;",
		}, strings.SplitAfterAnyBytesN("a;", []byte{';', ':'}, -1))
		require.Equal(t, []string{}, strings.SplitAfterAnyBytesN("a;", []byte{';', ':'}, 0))
		require.Equal(t, []string{
			"a;",
		}, strings.SplitAfterAnyBytesN("a;", []byte{';', ':'}, 1))
		require.Equal(t, []string{
			"a;",
		}, strings.SplitAfterAnyBytesN("a;", []byte{';', ':'}, 2))
	}

	{
		require.Equal(t, []string{
			";",
			"a",
		}, strings.SplitAfterAnyBytesN(";a", []byte{';', ':'}, -1))
		require.Equal(t, []string{}, strings.SplitAfterAnyBytesN(";a", []byte{';', ':'}, 0))
		require.Equal(t, []string{
			";a",
		}, strings.SplitAfterAnyBytesN(";a", []byte{';', ':'}, 1))
		require.Equal(t, []string{
			";",
			"a",
		}, strings.SplitAfterAnyBytesN(";a", []byte{';', ':'}, 2))
		require.Equal(t, []string{
			";",
			"a",
		}, strings.SplitAfterAnyBytesN(";a", []byte{';', ':'}, 3))
	}

	{
		require.Equal(t, []string{
			"a;",
			"a",
		}, strings.SplitAfterAnyBytesN("a;a", []byte{';', ':'}, -1))
		require.Equal(t, []string{}, strings.SplitAfterAnyBytesN("a;a", []byte{';', ':'}, 0))
		require.Equal(t, []string{
			"a;a",
		}, strings.SplitAfterAnyBytesN("a;a", []byte{';', ':'}, 1))
		require.Equal(t, []string{
			"a;",
			"a",
		}, strings.SplitAfterAnyBytesN("a;a", []byte{';', ':'}, 2))
		require.Equal(t, []string{
			"a;",
			"a",
		}, strings.SplitAfterAnyBytesN("a;a", []byte{';', ':'}, 3))
	}

	{
		require.Equal(t, []string{
			"abcdef:;ghi",
		}, strings.SplitAfterAnyBytesN("abcdef:;ghi", []byte{}, -1))
		require.Equal(t, []string{
			"abcdef:;ghi",
		}, strings.SplitAfterAnyBytesN("abcdef:;ghi", []byte{}, 1))

		require.Equal(t, []string{
			"abcdef:",
			";",
			"ghi",
		}, strings.SplitAfterAnyBytesN("abcdef:;ghi", []byte{';', ':'}, -1))
		require.Equal(t, []string{}, strings.SplitAfterAnyBytesN("abcdef:;ghi", []byte{';', ':'}, 0))
		require.Equal(t, []string{
			"abcdef:;ghi",
		}, strings.SplitAfterAnyBytesN("abcdef:;ghi", []byte{';', ':'}, 1))
		require.Equal(t, []string{
			"abcdef:",
			";ghi",
		}, strings.SplitAfterAnyBytesN("abcdef:;ghi", []byte{';', ':'}, 2))
		require.Equal(t, []string{
			"abcdef:",
			";",
			"ghi",
		}, strings.SplitAfterAnyBytesN("abcdef:;ghi", []byte{';', ':'}, 3))
		require.Equal(t, []string{
			"abcdef:",
			";",
			"ghi",
		}, strings.SplitAfterAnyBytesN("abcdef:;ghi", []byte{';', ':'}, 4))
	}

	{
		require.Equal(t, []string{
			"🐻;",
			"🐻‍❄️",
		}, strings.SplitAfterAnyBytesN("🐻;🐻‍❄️", []byte{';', ':'}, -1))
	}

	{
		require.Equal(t, []string{
			"a;", "b:", "c;", "d:", "e;", "f:", "g;", "h:", "i;", "j:", "k;", "l:", "m;", "n:", "o;", "p:", "q;", "r:", "s;", "t:", "u;", "v:", "w;", "x:", "y;", "z",
		}, strings.SplitAfterAnyBytesN("a;b:c;d:e;f:g;h:i;j:k;l:m;n:o;p:q;r:s;t:u;v:w;x:y;z", []byte{':', ';'}, -1))
		require.Equal(t, []string{}, strings.SplitAfterAnyBytesN("a;b:c;d:e;f:g;h:i;j:k;l:m;n:o;p:q;r:s;t:u;v:w;x:y;z", []byte{':', ';'}, 0))
		require.Equal(t, []string{
			"a;b:c;d:e;f:g;h:i;j:k;l:m;n:o;p:q;r:s;t:u;v:w;x:y;z",
		}, strings.SplitAfterAnyBytesN("a;b:c;d:e;f:g;h:i;j:k;l:m;n:o;p:q;r:s;t:u;v:w;x:y;z", []byte{':', ';'}, 1))
		require.Equal(t, []string{
			"a;",
			"b:c;d:e;f:g;h:i;j:k;l:m;n:o;p:q;r:s;t:u;v:w;x:y;z",
		}, strings.SplitAfterAnyBytesN("a;b:c;d:e;f:g;h:i;j:k;l:m;n:o;p:q;r:s;t:u;v:w;x:y;z", []byte{':', ';'}, 2))
		require.Equal(t, []string{
			"a;",
			"b:",
			"c;d:e;f:g;h:i;j:k;l:m;n:o;p:q;r:s;t:u;v:w;x:y;z",
		}, strings.SplitAfterAnyBytesN("a;b:c;d:e;f:g;h:i;j:k;l:m;n:o;p:q;r:s;t:u;v:w;x:y;z", []byte{':', ';'}, 3))
		require.Equal(t, []string{
			"a;", "b:", "c;", "d:", "e;", "f:", "g;", "h:", "i;", "j:", "k;", "l:", "m;", "n:", "o;", "p:", "q;", "r:", "s;", "t:", "u;", "v:", "w;", "x:", "y;z",
		}, strings.SplitAfterAnyBytesN("a;b:c;d:e;f:g;h:i;j:k;l:m;n:o;p:q;r:s;t:u;v:w;x:y;z", []byte{':', ';'}, 25))
		require.Equal(t, []string{
			"a;", "b:", "c;", "d:", "e;", "f:", "g;", "h:", "i;", "j:", "k;", "l:", "m;", "n:", "o;", "p:", "q;", "r:", "s;", "t:", "u;", "v:", "w;", "x:", "y;", "z",
		}, strings.SplitAfterAnyBytesN("a;b:c;d:e;f:g;h:i;j:k;l:m;n:o;p:q;r:s;t:u;v:w;x:y;z", []byte{':', ';'}, 26))
		require.Equal(t, []string{
			"a;", "b:", "c;", "d:", "e;", "f:", "g;", "h:", "i;", "j:", "k;", "l:", "m;", "n:", "o;", "p:", "q;", "r:", "s;", "t:", "u;", "v:", "w;", "x:", "y;", "z",
		}, strings.SplitAfterAnyBytesN("a;b:c;d:e;f:g;h:i;j:k;l:m;n:o;p:q;r:s;t:u;v:w;x:y;z", []byte{':', ';'}, 27))
	}
}

func Test_SplitAfterAnyRunes_1(t *testing.T) {

	{
		require.Equal(t, []string{}, strings.SplitAfterAnyRunes("", []rune{}))
		require.Equal(t, []string{}, strings.SplitAfterAnyRunes("", []rune{'/', '\\'}))
	}

	{
		require.Equal(t, []string{
			";",
		}, strings.SplitAfterAnyRunes(";", []rune{';', ':'}))
	}

	{
		require.Equal(t, []string{
			"a;",
		}, strings.SplitAfterAnyRunes("a;", []rune{';', ':'}))
	}

	{
		require.Equal(t, []string{
			";",
			"a",
		}, strings.SplitAfterAnyRunes(";a", []rune{';', ':'}))
	}

	{
		require.Equal(t, []string{
			"a;",
			"a",
		}, strings.SplitAfterAnyRunes("a;a", []rune{';', ':'}))
	}

	{
		require.Equal(t, []string{
			"abcdef:;ghi",
		}, strings.SplitAfterAnyRunes("abcdef:;ghi", []rune{}))

		require.Equal(t, []string{
			"abcdef:",
			";",
			"ghi",
		}, strings.SplitAfterAnyRunes("abcdef:;ghi", []rune{';', ':'}))
	}

	{
		require.Equal(t, []string{
			"🐻;",
			"🐻‍❄️",
		}, strings.SplitAfterAnyRunes("🐻;🐻‍❄️", []rune{';', ':'}))
	}

	{
		require.Equal(t, []string{
			"a;", "b:", "c;", "d:", "e;", "f:", "g;", "h:", "i;", "j:", "k;", "l:", "m;", "n:", "o;", "p:", "q;", "r:", "s;", "t:", "u;", "v:", "w;", "x:", "y;", "z",
		}, strings.SplitAfterAnyRunes("a;b:c;d:e;f:g;h:i;j:k;l:m;n:o;p:q;r:s;t:u;v:w;x:y;z", []rune{';', ':'}))
	}
}

func Test_SplitAfterAnyRunesN_1(t *testing.T) {

	{
		require.Equal(t, []string{}, strings.SplitAfterAnyRunesN("", []rune{}, -1))
		require.Equal(t, []string{}, strings.SplitAfterAnyRunesN("", []rune{'/', '\\'}, -1))
	}

	{
		require.Equal(t, []string{
			";",
		}, strings.SplitAfterAnyRunesN(";", []rune{';', ':'}, -1))
		require.Equal(t, []string{}, strings.SplitAfterAnyRunesN(";", []rune{';', ':'}, 0))
		require.Equal(t, []string{
			";",
		}, strings.SplitAfterAnyRunesN(";", []rune{';', ':'}, 1))
		require.Equal(t, []string{
			";",
		}, strings.SplitAfterAnyRunesN(";", []rune{';', ':'}, 2))
	}

	{
		require.Equal(t, []string{
			"a;",
		}, strings.SplitAfterAnyRunesN("a;", []rune{';', ':'}, -1))
		require.Equal(t, []string{}, strings.SplitAfterAnyRunesN("a;", []rune{';', ':'}, 0))
		require.Equal(t, []string{
			"a;",
		}, strings.SplitAfterAnyRunesN("a;", []rune{';', ':'}, 1))
		require.Equal(t, []string{
			"a;",
		}, strings.SplitAfterAnyRunesN("a;", []rune{';', ':'}, 2))
	}

	{
		require.Equal(t, []string{
			";",
			"a",
		}, strings.SplitAfterAnyRunesN(";a", []rune{';', ':'}, -1))
		require.Equal(t, []string{}, strings.SplitAfterAnyRunesN(";a", []rune{';', ':'}, 0))
		require.Equal(t, []string{
			";a",
		}, strings.SplitAfterAnyRunesN(";a", []rune{';', ':'}, 1))
		require.Equal(t, []string{
			";",
			"a",
		}, strings.SplitAfterAnyRunesN(";a", []rune{';', ':'}, 2))
		require.Equal(t, []string{
			";",
			"a",
		}, strings.SplitAfterAnyRunesN(";a", []rune{';', ':'}, 3))
	}

	{
		require.Equal(t, []string{
			"a;",
			"a",
		}, strings.SplitAfterAnyRunesN("a;a", []rune{';', ':'}, -1))
		require.Equal(t, []string{}, strings.SplitAfterAnyRunesN("a;a", []rune{';', ':'}, 0))
		require.Equal(t, []string{
			"a;a",
		}, strings.SplitAfterAnyRunesN("a;a", []rune{';', ':'}, 1))
		require.Equal(t, []string{
			"a;",
			"a",
		}, strings.SplitAfterAnyRunesN("a;a", []rune{';', ':'}, 2))
		require.Equal(t, []string{
			"a;",
			"a",
		}, strings.SplitAfterAnyRunesN("a;a", []rune{';', ':'}, 3))
	}

	{
		require.Equal(t, []string{
			"abcdef:;ghi",
		}, strings.SplitAfterAnyRunesN("abcdef:;ghi", []rune{}, -1))
		require.Equal(t, []string{
			"abcdef:;ghi",
		}, strings.SplitAfterAnyRunesN("abcdef:;ghi", []rune{}, 3))

		require.Equal(t, []string{
			"abcdef:",
			";",
			"ghi",
		}, strings.SplitAfterAnyRunesN("abcdef:;ghi", []rune{';', ':'}, -1))
		require.Equal(t, []string{}, strings.SplitAfterAnyRunesN("abcdef:;ghi", []rune{';', ':'}, 0))
		require.Equal(t, []string{
			"abcdef:;ghi",
		}, strings.SplitAfterAnyRunesN("abcdef:;ghi", []rune{';', ':'}, 1))
		require.Equal(t, []string{
			"abcdef:",
			";ghi",
		}, strings.SplitAfterAnyRunesN("abcdef:;ghi", []rune{';', ':'}, 2))
		require.Equal(t, []string{
			"abcdef:",
			";",
			"ghi",
		}, strings.SplitAfterAnyRunesN("abcdef:;ghi", []rune{';', ':'}, 3))
		require.Equal(t, []string{
			"abcdef:",
			";",
			"ghi",
		}, strings.SplitAfterAnyRunesN("abcdef:;ghi", []rune{';', ':'}, 4))
	}

	{
		require.Equal(t, []string{
			"🐻;",
			"🐻‍❄️",
		}, strings.SplitAfterAnyRunesN("🐻;🐻‍❄️", []rune{';', ':'}, -1))
	}

	{
		require.Equal(t, []string{
			"a;", "b:", "c;", "d:", "e;", "f:", "g;", "h:", "i;", "j:", "k;", "l:", "m;", "n:", "o;", "p:", "q;", "r:", "s;", "t:", "u;", "v:", "w;", "x:", "y;", "z",
		}, strings.SplitAfterAnyRunesN("a;b:c;d:e;f:g;h:i;j:k;l:m;n:o;p:q;r:s;t:u;v:w;x:y;z", []rune{':', ';'}, -1))
		require.Equal(t, []string{}, strings.SplitAfterAnyRunesN("a;b:c;d:e;f:g;h:i;j:k;l:m;n:o;p:q;r:s;t:u;v:w;x:y;z", []rune{':', ';'}, 0))
		require.Equal(t, []string{
			"a;b:c;d:e;f:g;h:i;j:k;l:m;n:o;p:q;r:s;t:u;v:w;x:y;z",
		}, strings.SplitAfterAnyRunesN("a;b:c;d:e;f:g;h:i;j:k;l:m;n:o;p:q;r:s;t:u;v:w;x:y;z", []rune{':', ';'}, 1))
		require.Equal(t, []string{
			"a;",
			"b:c;d:e;f:g;h:i;j:k;l:m;n:o;p:q;r:s;t:u;v:w;x:y;z",
		}, strings.SplitAfterAnyRunesN("a;b:c;d:e;f:g;h:i;j:k;l:m;n:o;p:q;r:s;t:u;v:w;x:y;z", []rune{':', ';'}, 2))
		require.Equal(t, []string{
			"a;",
			"b:",
			"c;d:e;f:g;h:i;j:k;l:m;n:o;p:q;r:s;t:u;v:w;x:y;z",
		}, strings.SplitAfterAnyRunesN("a;b:c;d:e;f:g;h:i;j:k;l:m;n:o;p:q;r:s;t:u;v:w;x:y;z", []rune{':', ';'}, 3))
		require.Equal(t, []string{
			"a;", "b:", "c;", "d:", "e;", "f:", "g;", "h:", "i;", "j:", "k;", "l:", "m;", "n:", "o;", "p:", "q;", "r:", "s;", "t:", "u;", "v:", "w;", "x:", "y;z",
		}, strings.SplitAfterAnyRunesN("a;b:c;d:e;f:g;h:i;j:k;l:m;n:o;p:q;r:s;t:u;v:w;x:y;z", []rune{':', ';'}, 25))
		require.Equal(t, []string{
			"a;", "b:", "c;", "d:", "e;", "f:", "g;", "h:", "i;", "j:", "k;", "l:", "m;", "n:", "o;", "p:", "q;", "r:", "s;", "t:", "u;", "v:", "w;", "x:", "y;", "z",
		}, strings.SplitAfterAnyRunesN("a;b:c;d:e;f:g;h:i;j:k;l:m;n:o;p:q;r:s;t:u;v:w;x:y;z", []rune{':', ';'}, 26))
		require.Equal(t, []string{
			"a;", "b:", "c;", "d:", "e;", "f:", "g;", "h:", "i;", "j:", "k;", "l:", "m;", "n:", "o;", "p:", "q;", "r:", "s;", "t:", "u;", "v:", "w;", "x:", "y;", "z",
		}, strings.SplitAfterAnyRunesN("a;b:c;d:e;f:g;h:i;j:k;l:m;n:o;p:q;r:s;t:u;v:w;x:y;z", []rune{':', ';'}, 27))
	}
}
