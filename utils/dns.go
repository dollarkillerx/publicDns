/**
 * @Author: DollarKillerX
 * @Description: dns.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 上午10:26 2019/11/29
 */
package utils

import "strings"

func NotIpv6(ip string) bool {
	if strings.Index(ip, ":") != -1 {
		return false
	}
	return true
}
