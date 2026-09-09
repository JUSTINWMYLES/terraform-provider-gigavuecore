package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAlertPoliciesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadAlertPoliciesDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadAlertPoliciesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadAlertPoliciesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadAlertPoliciesDataSourceMetadata(t *testing.T) {
	d := NewLoadAlertPoliciesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_alert_policies" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_alert_policies")
	}
}
