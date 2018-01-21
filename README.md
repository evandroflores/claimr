
# Claimr - A slack bot to manage containers use

[![Build Status](https://travis-ci.org/evandroflores/claimr.svg?branch=master)](https://travis-ci.org/evandroflores/claimr)
[![codecov](https://codecov.io/gh/evandroflores/claimr/branch/master/graph/badge.svg)](https://codecov.io/gh/evandroflores/claimr)
[![Go Report Card](https://goreportcard.com/badge/github.com/evandroflores/claimr)](https://goreportcard.com/report/github.com/evandroflores/claimr)
[![Maintainability](https://api.codeclimate.com/v1/badges/b9c9833444e3012fadf4/maintainability)](https://codeclimate.com/github/evandroflores/claimr/maintainability)
[![License: BSD-3](https://img.shields.io/badge/License-BSD3-green.svg)](https://opensource.org/licenses/BSD-3-Clause)

---

## Commands

* **help** - *Shows this command list.*
* **add** `container-name` - *Adds a container to your channel.*
* **claim** `container-name` `reason` - *Claims a container for your use.*
* **free** `container-name` - *Makes a container available for use.*
* **list** - *List all containers.*
* **log-level** `level` - _Change the current log level. admin-only_
* **purge**  - _Purge soft delete models from the database. admin-only_
* **remove** `container-name` - *Removes a container from your channel.*
* **show** `container-name` - *Shows a container details.*

Notice that commands tagged as _admin-only_ will be shown and are enabled only
for the super user (or 'admins' as soon I create a better way of doing).

---
Make sure to have your credentials on environment set in order to connect
to Slack `CLAIMR_TOKEN`, MySQL `CLAIMR_DATABASE`, and `CLAIMR_SUPERUSER`
