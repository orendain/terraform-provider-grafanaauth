package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProvider *schema.Provider
var testAccProviderFactories map[string]func() (*schema.Provider, error)

func init() {
	testAccProvider = New()
	testAccProviderFactories = map[string]func() (*schema.Provider, error){
		"grafanaauth": func() (*schema.Provider, error) { return testAccProvider, nil },
	}
}

func TestProvider(t *testing.T) {
	if err := New().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = New()
}

func testAccPreCheck(t *testing.T) {
	if err := os.Getenv("GRAFANA_URL"); err == "" {
		t.Fatal("GRAFANA_URL must be set for acceptance tests.")
	}
}
