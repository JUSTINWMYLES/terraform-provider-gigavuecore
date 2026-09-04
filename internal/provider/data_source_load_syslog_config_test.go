package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadSyslogConfigDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadSyslogConfigDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadSyslogConfigDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadSyslogConfigDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadSyslogConfigDataSourceMetadata(t *testing.T) {
	d := NewLoadSyslogConfigDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_syslog_config" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_syslog_config")
	}
}
