---
page_title: "gigavuecore_get_crypto_ca Data Source - gigavuecore"
subcategory: ""
description: |-
  get a list of configured trusted certificate authorities (CA List)
---

# gigavuecore_get_crypto_ca Data Source

get a list of configured trusted certificate authorities (CA List)

## Example Usage

```terraform
data "gigavuecore_get_crypto_ca" "example" {
  algorithm    = "example"
  cluster_id   = "example"
  comment      = "example"
  issuer_name  = "example"
  name         = "example"
  page         = "example"
  sort         = "example"
  subject_name = "example"
  valid_from   = "example"
  valid_till   = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `algorithm` (String, optional) - Certificate Public key algorithm (filter param)
* `cluster_id` (String, required) - Target Cluster ID
* `comment` (String, optional) - Certificate description (filter param)
* `issuer_name` (String, optional) - Issuer's common Name (filter param)
* `name` (String, optional) - Target Certificate name (filter param)
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)
* `subject_name` (String, optional) - Subject's common Name (filter param)
* `valid_from` (String, optional) - Certificate Validity details (filter param)
* `valid_till` (String, optional) - Certificate Validity details (filter param)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `comment` (String)
* `default_cert` (Boolean) - If true, this certificate is the default server identity
* `issuer` (Attributes) (see [below for nested schema](#nestedatt--items--issuer))
* `name` (String)
* `pem` (String)
* `private_key` (Boolean)
* `serial_number` (String) - lower-case hexadecimal serial number not prefixed with '0x'
* `sha1_fingerprint` (String)
* `signature_algorithm` (String)
* `subject` (Attributes) (see [below for nested schema](#nestedatt--items--subject))
* `subject_public_key_algorithm` (String)
* `subject_public_key_length` (Number)
* `validity` (Attributes) (see [below for nested schema](#nestedatt--items--validity))
* `version` (String)

<a id="nestedatt--items--issuer"></a>
### Nested Schema for `items.issuer`

Read-Only:

* `common_name` (String) - issuer and subject common name (e.g. a domain name)
* `country` (String) - two-alphanumeric-character country code
* `email` (String)
* `locality` (String)
* `org_name` (String)
* `org_unit` (String)
* `state` (String) - state or provence name

<a id="nestedatt--items--subject"></a>
### Nested Schema for `items.subject`

Read-Only:

* `common_name` (String) - issuer and subject common name (e.g. a domain name)
* `country` (String) - two-alphanumeric-character country code
* `email` (String)
* `locality` (String)
* `org_name` (String)
* `org_unit` (String)
* `state` (String) - state or provence name

<a id="nestedatt--items--validity"></a>
### Nested Schema for `items.validity`

Read-Only:

* `not_after` (String) - date and time when certificate stops being valid (\[rfc3339\](https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html\#anchor14))
* `not_before` (String) - date and time when certificate starts being valid (\[rfc3339\](https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html\#anchor14))

