//go:build !linux && !windows

package archive // import "github.com/docker/docker/pkg/archive"

import "golang.org/x/sys/unix"

var noattr = unix.ENOATTR
