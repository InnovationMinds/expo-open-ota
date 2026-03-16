package cdn

import (
	"expo-open-ota/config"
	"fmt"
	"strings"
)

type ScalewayCDN struct{}

func getScalewayCDNDomain() string {
	return config.GetEnv("SCALEWAY_CDN_DOMAIN")
}

func (c *ScalewayCDN) isCDNAvailable() bool {
	return config.GetEnv("STORAGE_MODE") == "s3" &&
		config.GetEnv("S3_BUCKET_NAME") != "" &&
		getScalewayCDNDomain() != ""
}

func (c *ScalewayCDN) ComputeRedirectionURLForAsset(branch, runtimeVersion, updateId, asset string) (string, error) {
	domain := getScalewayCDNDomain()
	if domain == "" {
		return "", fmt.Errorf("Scaleway CDN configuration is incomplete")
	}

	domain = strings.TrimRight(domain, "/")

	keyPrefix := strings.Trim(config.GetEnv("S3_KEY_PREFIX"), "/")
	if keyPrefix != "" {
		keyPrefix += "/"
	}

	path := fmt.Sprintf("%s%s/%s/%s/%s", keyPrefix, branch, runtimeVersion, updateId, asset)
	return fmt.Sprintf("%s/%s", domain, path), nil
}
