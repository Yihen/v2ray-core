// Package core provides an entry point to use V2Ray core functionalities.
//
// V2Ray makes it possible to accept incoming network connections with certain
// protocol, process the data, and send them through another connection with
// the same or a difference protocol on demand.
//
// It may be configured to work with multiple protocols at the same time, and
// uses the internal router to tunnel through different inbound and outbound
// connections.
package core

//go:generate go install "v2ray.com/core/common/errors/errorgen"
//go:generate errorgen

import (
	"runtime"

	"v2ray.com/core/common/serial"
)

var (
	version  = "4.22.1"
	build    = "Custom"
	codename = "V2Fly, a community-driven edition of V2Ray."
	intro    = "A unified platform for anti-censorship."
)

// Version returns V2Ray's version as a string, in the form of "x.y.z" where x, y and z are numbers.
// ".z" part may be omitted in regular releases.
func Version() string {
	return version
}

// VersionStatement returns a list of strings representing the full version info.
func VersionStatement() []string {
	return []string{
		serial.Concat("V2Ray ", Version(), " (", codename, ") ", build, " (", runtime.Version(), " ", runtime.GOOS, "/", runtime.GOARCH, ")"),
		intro,
	}
}
