package scheduler

import "strings"

// getPrimaryDomain 用于获取给定主机名的主域名。
func getPrimaryDomain(host string) (string, error) {
	host = strings.TrimSpace(host)
	if host == "" {
		return "", genError("empty host")
	}
	if regexpForIP.MatchString(host) {
		return host, nil
	}
	var suffixIndex int
	for _, re := range regexpForDomains {
		pos := re.FindStringIndex(host)
		if pos != nil {
			suffixIndex = pos[0]
			break
		}
	}
	if suffixIndex > 0 {
		var pdIndex int
		firstPart := host[:suffixIndex]
		index := strings.LastIndex(firstPart, ".")
		if index < 0 {
			pdIndex = 0
		} else {
			pdIndex = index + 1
		}
		return host[pdIndex:], nil
	}
	return "", genError("unrecognized host")
}
