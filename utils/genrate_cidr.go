package utils

import (
	"fmt"
	"net"
)

// Generate /8 或 /16 拆成 /24 子网
func Generate24(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var subnets []string

	maskSize, _ := ipnet.Mask.Size()
	if maskSize > 16 {
		return nil, fmt.Errorf("CIDR too small, must be /16 or larger")
	}

	start := ip.To4()
	if start == nil {
		return nil, fmt.Errorf("invalid IP")
	}

	// 计算需要拆多少个 /24
	total := 1 << (24 - maskSize) // /16 -> 256, /8 -> 65536

	for i := 0; i < total; i++ {
		octet2 := int(start[1]) + ((i >> 8) & 0xFF)
		octet3 := int(start[2]) + (i & 0xFF)
		subnet := fmt.Sprintf("%d.%d.%d.0/24", start[0], octet2, octet3)
		subnets = append(subnets, subnet)
	}

	return subnets, nil
}
