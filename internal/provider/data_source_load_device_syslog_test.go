package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadDeviceSyslogDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadDeviceSyslogDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadDeviceSyslogDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadDeviceSyslogDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadDeviceSyslogDataSourceMetadata(t *testing.T) {
	d := NewLoadDeviceSyslogDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_device_syslog" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_device_syslog")
	}
}
