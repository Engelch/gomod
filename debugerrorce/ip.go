package debugerrorce

import (
	"errors"
	"regexp"
)

func ValidIPv4Address(ipv4 string) error {
	matched, err := regexp.MatchString(`^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$`, ipv4)
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("IP Address regexpr not matching for " + ipv4)
	}
	return nil
}
