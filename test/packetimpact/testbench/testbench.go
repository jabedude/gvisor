// Copyright 2020 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package testbench

import (
	"flag"
	"time"
)

var (
	// Device is the local device on the test network.
	Device = ""
	// LocalIPv4 is the local IPv4 address on the test network.
	LocalIPv4 = ""
	// LocalIPv6 is the local IPv6 address on the test network.
	LocalIPv6 = ""
	// LocalMAC is the local MAC address on the test network.
	LocalMAC = ""
	// POSIXServerIP is the POSIX server's IP address on the control network.
	POSIXServerIP = ""
	// POSIXServerPort is the UDP port the POSIX server is bound to on the
	// control network.
	POSIXServerPort = 40000
	// RemoteIPv4 is the DUT's IPv4 address on the test network.
	RemoteIPv4 = ""
	// RemoteIPv6 is the DUT's IPv6 address on the test network.
	RemoteIPv6 = ""
	// RemoteMAC is the DUT's MAC address on the test network.
	RemoteMAC = ""
	// RPCKeepalive is the gRPC keepalive.
	RPCKeepalive = 10 * time.Second
	// RPCTimeout is the gRPC timeout.
	RPCTimeout = 100 * time.Millisecond
)

// RegisterFlags defines flags and associates them with the package-level
// exported variables above. It should be called by tests in their init
// functions.
func RegisterFlags(fs *flag.FlagSet) {
	fs.StringVar(&POSIXServerIP, "posix_server_ip", POSIXServerIP, "ip address to listen to for UDP commands")
	fs.IntVar(&POSIXServerPort, "posix_server_port", POSIXServerPort, "port to listen to for UDP commands")
	fs.DurationVar(&RPCTimeout, "rpc_timeout", RPCTimeout, "gRPC timeout")
	fs.DurationVar(&RPCKeepalive, "rpc_keepalive", RPCKeepalive, "gRPC keepalive")
	fs.StringVar(&LocalIPv4, "local_ipv4", LocalIPv4, "local IPv4 address for test packets")
	fs.StringVar(&RemoteIPv4, "remote_ipv4", RemoteIPv4, "remote IPv4 address for test packets")
	fs.StringVar(&LocalIPv6, "local_ipv6", LocalIPv6, "local IPv6 address for test packets")
	fs.StringVar(&RemoteIPv6, "remote_ipv6", RemoteIPv6, "remote IPv6 address for test packets")
	fs.StringVar(&LocalMAC, "local_mac", LocalMAC, "local mac address for test packets")
	fs.StringVar(&RemoteMAC, "remote_mac", RemoteMAC, "remote mac address for test packets")
	fs.StringVar(&Device, "device", Device, "local device for test packets")
}
