---
page_title: "gigavuecore Provider"
subcategory: ""
description: |-
  This provider is generated using eidos from the Gigamon GigaVUE-FM OpenAPI specification.
---

# gigavuecore Provider

This provider is generated using [eidos](https://github.com/signalbreak-labs/eidos) from the Gigamon GigaVUE-FM [OpenAPI](https://docs.gigamon.com/ref-api/Content/apiref_514onwards/doc/webapp/release/6.14.00%20GigaVUE-FM%20Core/openapi.fm.yaml) specification.

## Example Usage

```terraform
provider "gigavuecore" {
  endpoint = "example"
  username = "example"
  password = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `endpoint` (String, optional) - Overrides the default API base URL derived from the OpenAPI servers. Useful for directing the provider at a test or mock server.
* `username` (String, optional) - Username for HTTP basic authentication.
* `password` (String, optional) - Password for HTTP basic authentication.
* `log_file` (String, optional) - Path to a file that receives HTTP request/response trace logs. When unset, trace logging is disabled.
* `log_capture_request_headers` (Boolean, optional) - Capture request headers in the trace log. Sensitive headers are redacted.
* `log_capture_request_body` (Boolean, optional) - Capture request bodies in the trace log. Disabled by default to avoid writing sensitive payloads to disk.
* `log_capture_response_headers` (Boolean, optional) - Capture response headers in the trace log. Sensitive headers are redacted.
* `log_capture_response_body` (Boolean, optional) - Capture response bodies in the trace log. Disabled by default to avoid writing sensitive payloads to disk.
* `log_max_body_bytes` (Number, optional) - Maximum number of body bytes captured per log entry before truncation. Defaults to 4096.
* `tls_skip_verify` (Boolean, optional) - Disable TLS certificate verification for API requests. Defaults to false; enable only against endpoints with self-signed or otherwise untrusted certificates.


