<!-- Personal fork for learning Terraform provider development. Tracking upstream: citrix/terraform-provider-citrixadc -->
<!-- Note to self: sync with upstream before starting any new feature work -->
<!-- Last synced with upstream: 2026-03-26 -->
<!-- TODO: Look into the lbmonitor respcode drift fix (#1384) - had similar issue at work -->
<!-- TODO: Test the Console Service proxy support (#1384) - could be useful for lab environment -->
<!-- TODO: Test citrixadc_csvserver_lbvserver_binding - need this for multi-vserver setup in lab -->
<!-- TODO: Try the citrixadc_routerdynamicrouting fix (#1388) - was hitting the node id issue in my BGP lab -->

## 2.1.4 (Mar 26, 2026)

ENHANCEMENTS
* **citrixadc_nslaslicense_offline**: Support for Offline LAS Licensing in restricted mode.

## 2.1.3 (Mar 24, 2026)

FEATURES
* **New Resource**: Support for creating and managing citrixadc_csvserver_lbvserver_binding.
* **New Resource**: Support for configuring citrixadc_nslaslicense_offline resource for Offline LAS licensing.

BUG FIXES
* **citrixadc_lbmonitor_sslcertkey_binding**: Fixed `ca` attribute not being passed in the API call when deleting lbmonitor_sslcertkey_binding. [#1377]
* **citrixadc_lbmonitor**: Rectified `respcode` attribute handling to prevent unnecessary drift detection. [#1384]
* **citrixadc_routerdynamicrouting**: Fixed dynamic routing config resource to not require node id for create/apply. [#1388]

UPDATES
* **google.golang.org/grpc**: Version upgrade from 1.75.1 to 1.79.3.
* **github.com/cloudflare/circl**: Version upgrade from 1.6.1 to 1.6.3.

[#1377]: https://github.com/citrix/terraform-provider-citrixadc/issues/1377
[#1384]: https://github.com/citrix/terraform-provider-citrixadc/issues/1384
[#1388]: https://github.com/citrix/terraform-provider-citrixadc/issues/1388

## 2.1.2 (Mar 07, 2026)

BUG FIXES
* **citrixadc_sslcertkey**: Rectified `bundle` attribute to prevent unnecessary drift detection. [#1237]
* **citrixadc_lbparameter**: Rectified `cookiepassphrase` attribute to prevent unnecessary drift detection. [#1361]
* **citrixadc_sslcertkey_update**: Migrated sslcertkey_update resource to Plugin framework. [#1370]
* **citrixadc_sslvserver_sslcertkey_binding**: Fixed nil type assertion error for `ca` and `snicert` attributes.

ENHANCEMENTS
* **provider**: Added support for proxying NetScaler APIs through Console Service.
* **GoLang version**: Updated GoLang toolchain version to 1.24.13.

[#1237]: https://github.com/citrix/terraform-provider-citrixadc/issues/1237
[#1361]: https://github.com/citrix/terraform-provider-citrixadc/issues/1361
[#1370]: https://github.com/citrix/terraform-provider-citrixadc/issues/1370


## 2.1.1 (Feb 17, 2026)

BUG FIXES
* **citrixadc_systemfile**: Fixed systemfile decoding to properly handle already-encoded content. [#1300]
* **citrixadc_lbvserver**: Fixed handling of backupvserver attribute to allow proper removal/unbinding. [#1283]
* **citrixadc_lbparameter**: Rectified int64 conversion for float64 type values from API response. [#1346]
* **citrixadc_sslcertkey_update**: Always including passplain attribute for cer