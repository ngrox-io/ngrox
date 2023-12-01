//go:build release
// +build release

package client

var rootCrtPaths = []string{"tls/ngrokroot.crt"}

func useInsecureSkipVerify() bool {
	return false
}
