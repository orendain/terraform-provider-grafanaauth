# Terraform Provider for Grafana Auth API

A very simple Terraform Provider to manage the creation of [Grafana API Keys](https://grafana.com/docs/grafana/latest/http_api/auth/).

## Overview

This project exists for several reasons:
- The official [Terraform Provider for Grafana](https://github.com/grafana/terraform-provider-grafana) does not support managing API key resources.
- The official [Grafana API Client](https://github.com/grafana/grafana-api-golang-client), which the official provider above uses, does not expose methods for calls against the Authorization API.
- I needed a working provider for this purpose the next day, and could not want to wait for contributions to the two official projects above to be merged in.
- I do indeed plan on contributing these changes to the official projects above very soon.
- I wanted practice with Terraform Plugin SDK v2 :smirk:.

Corresponding changes to the upstream grafana-api-golang-client are [staged here](https://github.com/orendain/grafana-api-golang-client/pull/1).

## Development

To build the binary, and move it to one of the expected [plugin locations](https://www.terraform.io/docs/extend/how-terraform-works.html#plugin-locations),
modify the `OS_ARCH` variable in the makefile and run:
```
make install
```

To run acceptance tests:
1. Make sure an instance of Grafana is running and accessible.
2. Set appropriate env vars `GRAFANA_URL` (required), and `GRAFANA_USERNAME`, `GRAFANA_PASSWORD`, or `GRAFANA_API_TOKEN` if necessary.
3. Run `make testacc`.

## Use

Please see the `example` directory for some simple examples.

## Limitations / Known Issues

- Due to a bug in previous versions of Grafana, the underlying API calls work as expected only with versions of Grafana `>= v6.6`.
- Support for importing API keys (via `terraform import`) is possible but not planned.
