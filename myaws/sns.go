package myaws

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func GetSNSStringAttribute(log *logrus.Entry, attribute interface{}) (string, error) {
	log.Infof("read sns attribute value: %v", attribute)
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
