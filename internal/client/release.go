//go:build release
// +build release

package client

var rootCrtPaths = []string{"tls/ngroxroot.crt"}

func useInsecureSkipVerify() bool {
	return false
}
