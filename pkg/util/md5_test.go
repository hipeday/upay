package util

import "testing"

func TestMD5SaltHash(t *testing.T) {
	print(MD5SaltHash("a3f0bec59cebeb60553ec80bbfd5dfdf", "GnYchJd4"))
}
