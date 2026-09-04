package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetPolicyDeploymentStatusDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetPolicyDeploymentStatusDataSourceSchemaValidation(t *testing.T) {
	d := NewGetPolicyDeploymentStatusDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetPolicyDeploymentStatusDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetPolicyDeploymentStatusDataSourceMetadata(t *testing.T) {
	d := NewGetPolicyDeploymentStatusDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_policy_deployment_status" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_policy_deployment_status")
	}
}
