package cdn

import (
	"os"
	testing2 "testing"
)

func TestGetCDNReturnsScalewayWhenScalewayConfigured(t *testing2.T) {
	os.Setenv("STORAGE_MODE", "s3")
	defer os.Unsetenv("STORAGE_MODE")
	os.Setenv("S3_BUCKET_NAME", "test-bucket")
	defer os.Unsetenv("S3_BUCKET_NAME")
	os.Setenv("SCALEWAY_CDN_DOMAIN", "https://test.edge.scw.cloud")
	defer os.Unsetenv("SCALEWAY_CDN_DOMAIN")
	ResetCDNInstance()
	c := GetCDN()
	if c == nil {
		t.Fatalf("expected CDN instance, got nil")
	}
	if _, ok := c.(*ScalewayCDN); !ok {
		t.Fatalf("expected *ScalewayCDN, got %T", c)
	}
}

func TestScalewayCDNComputeURL(t *testing2.T) {
	os.Setenv("SCALEWAY_CDN_DOMAIN", "https://test.edge.scw.cloud")
	defer os.Unsetenv("SCALEWAY_CDN_DOMAIN")
	os.Unsetenv("S3_KEY_PREFIX")

	cdn := &ScalewayCDN{}
	url, err := cdn.ComputeRedirectionURLForAsset("main", "1.0.0", "update-123", "bundle.js")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "https://test.edge.scw.cloud/main/1.0.0/update-123/bundle.js"
	if url != expected {
		t.Fatalf("expected %s, got %s", expected, url)
	}
}

func TestScalewayCDNComputeURLWithKeyPrefix(t *testing2.T) {
	os.Setenv("SCALEWAY_CDN_DOMAIN", "https://test.edge.scw.cloud")
	defer os.Unsetenv("SCALEWAY_CDN_DOMAIN")
	os.Setenv("S3_KEY_PREFIX", "myapp/")
	defer os.Unsetenv("S3_KEY_PREFIX")

	cdn := &ScalewayCDN{}
	url, err := cdn.ComputeRedirectionURLForAsset("main", "1.0.0", "update-123", "bundle.js")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "https://test.edge.scw.cloud/myapp/main/1.0.0/update-123/bundle.js"
	if url != expected {
		t.Fatalf("expected %s, got %s", expected, url)
	}
}

func TestGetCDNReturnsGCSDirectWhenGCSConfigured(t *testing2.T) {
	os.Setenv("STORAGE_MODE", "gcs")
	os.Setenv("GCS_BUCKET_NAME", "test-bucket")
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS_B64", "e3ZhbHVlOiAxfQ==")
	ResetCDNInstance()
	c := GetCDN()
	if c == nil {
		t.Fatalf("expected CDN instance, got nil")
	}
	if _, ok := c.(*GCSDirectCDN); !ok {
		t.Fatalf("expected *GCSDirectCDN, got %T", c)
	}
}
