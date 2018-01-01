
# Claimr - A slack bot to manage containers use

[![Build Status](https://travis-ci.org/evandroflores/claimr.svg?branch=master)](https://travis-ci.org/evandroflores/claimr)
[![codecov](https://codecov.io/gh/evandroflores/claimr/branch/master/graph/badge.svg)](https://codecov.io/gh/evandroflores/claimr)
[![Go Report Card](https://goreportcard.com/badge/github.com/evandroflores/claimr)](https://goreportcard.com/report/github.com/evandroflores/claimr)
[![License: BSD-3](https://img.shields.io/badge/License-BSD3-green.svg)](https://opensource.org/licenses/BSD-3-Clause)

---

## Commands

* **help** - *Shows commands list.*
* **add** `container-name` - *Adds a container to your channel.*
* **claim** `container-name` `reason` - *Claims a container for your use.*
* **free** `container-name` - *Makes a container available for use.*
* **list** - *List all containers.*
* **remove** `container-name` - *Removes a container from your channel.*
* **show** `container-name` - *Shows a container details.*

---
Make sure to have your credentials on environment set in order to connect
to Slack `CLAIMR_TOKEN` and MySQL `CLAIMR_DATABASE`.