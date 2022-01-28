package aws

import "fmt"

func GetSNSStringAttribute(attribute interface{}) (string, error) {
	v, ok := attribute.(map[string]interface{})
	if ok {
		s, ok := v["Value"].(string)
		if ok {
			return s, nil
		}
		return "", fmt.Errorf("cannot read string value of spring property: %v", attribute)
	}
	return "", fmt.Errorf("cannot read string property: %v", attribute)
}
