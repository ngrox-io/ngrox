//go:build !release
// +build !release

package client

var rootCrtPaths = []string{"tls/ngrokroot.crt", "tls/snakeoilca.crt"}

func useInsecureSkipVerify() bool {
	return true
}
