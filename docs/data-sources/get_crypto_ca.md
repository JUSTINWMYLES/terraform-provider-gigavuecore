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
  algorithm = null
  cluster_id = null
  comment = null
  issuer_name = null
  name = null
  page = null
  sort = null
  subject_name = null
  valid_from = null
  valid_till = null
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

* `items` (List(Object({comment, default_cert, issuer, name, pem, private_key, serial_number, sha1_fingerprint, signature_algorithm, subject, subject_public_key_algorithm, subject_public_key_length, validity, version})), computed)

