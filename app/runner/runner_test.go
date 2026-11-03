/*
Copyright © 2025 taksenov@gmail.com
*/

package runner

import (
	"testing"
)

func TestQuoteRemoteCmd(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"whoami", "'whoami'"},
		{"ls -lah ; exit", "'ls -lah ; exit'"},
		{"sudo systemctl list-timers | grep logrotate; exit", "'sudo systemctl list-timers | grep logrotate; exit'"},
		{`echo "a b" && df -h`, `'echo "a b" && df -h'`},
		{"it's ok", `'it'\''s ok'`},
	}
	for _, c := range cases {
		if got := quoteRemoteCmd(c.in); got != c.want {
			t.Errorf("quoteRemoteCmd(%q) = %s, want %s", c.in, got, c.want)
		}
	}
}
