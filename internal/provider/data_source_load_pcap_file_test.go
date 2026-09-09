package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadPcapFileDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadPcapFileDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadPcapFileDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadPcapFileDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadPcapFileDataSourceMetadata(t *testing.T) {
	d := NewLoadPcapFileDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_pcap_file" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_pcap_file")
	}
}
