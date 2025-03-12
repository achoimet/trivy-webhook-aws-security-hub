package tools

import (
	"fmt"
	"log"
	"strconv"

	"github.com/aquasecurity/trivy-operator/pkg/apis/aquasecurity/v1alpha1"
	"k8s.io/utils/env"
)

func GetVulnScore(d v1alpha1.Vulnerability) float64 {
	if d.Score != nil {
		return *d.Score
	}
	return 0.0
}

func ParseEnvBool(key string, defaultVal bool) bool {
	val := env.GetString(key, fmt.Sprintf("%v", defaultVal))
	parsed, err := strconv.ParseBool(val)
	if err != nil {
		log.Printf("Invalid boolean for %s: %s, using default: %v", key, val, defaultVal)
		return defaultVal
	}
	return parsed
}
