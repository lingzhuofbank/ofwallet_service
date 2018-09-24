
package params

import (
	"fmt"
)

//2.0.0 okex
const (
	VersionMajor = 2          // Major version component of the current release
	VersionMinor = 0          // Minor version component of the current release
	VersionPatch = 4          // Patch version component of the current release
	VersionMeta  = "zb"	  // Version metadata to append to the version string
)

// Version holds the textual version string.
var Version = func() string {
	v := fmt.Sprintf("%d.%d.%d", VersionMajor, VersionMinor, VersionPatch)
	if VersionMeta != "" {
		v += "-" + VersionMeta
	}
	return v
}()

func VersionWithCommit(gitCommit string) string {
	vsn := Version
	if len(gitCommit) >= 8 {
		vsn += "-" + gitCommit[:8]
	}
	return vsn
}
