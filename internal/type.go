package internal

import (
	"strconv"
	"strings"
)

type PerfConfig struct {
	NacosAddr               string
	ServiceCount            int
	InstanceCountPerService int
	ClientCount             int
	NamingRegTps            int
	NamingQueryQps          int
	PerfMode                string
	ConfigPubTps            int
	ConfigGetTps            int
	ConfigContentLength     int
	ConfigCount             int
	NamingMetadataLength    int
	PerfTimeSec             int
	PerfApi                 string
}

// parseAddrAndPort parses an address in the format "ip:port" or "ip" and returns ip and port
// If no port is specified, returns the default port 8848
func parseAddrAndPort(addr string) (string, uint64) {
	parts := strings.Split(strings.TrimSpace(addr), ":")
	if len(parts) == 2 {
		port, err := strconv.ParseUint(parts[1], 10, 64)
		if err != nil {
			// If port parsing fails, use default port
			return parts[0], 8848
		}
		return parts[0], port
	}
	// If no port specified, use default port 8848
	return parts[0], 8848
}
