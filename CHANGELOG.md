<!--
  ~ Copyright 2022 kwanhur
  ~
  ~ Licensed under the Apache License, Version 2.0 (the "License");
  ~ you may not use this file except in compliance with the License.
  ~ You may obtain a copy of the License at
  ~
  ~ http://www.apache.org/licenses/LICENSE-2.0
  ~
  ~ Unless required by applicable law or agreed to in writing, software
  ~ distributed under the License is distributed on an "AS IS" BASIS,
  ~ WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  ~ See the License for the specific language governing permissions and
  ~ limitations under the License.
  ~
-->

<!--
This changelog should always be read on `master` branch. Its contents on other branches
does not necessarily reflect the changes.
-->

# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

- ipvsctl sub-command `local-address`
- support import from stdin, compatible with ipvsadm dump format

## [v1.2.0] - 2022-03-21

### Added

- ipvsctl sub-command `connection`

### Fixed

- show daemon title ends with newline
- show version newline behind ip_vs information

## [v1.1.0] - 2022-03-20

### Added

- ipvsctl sub-command `flush`, `daemon`

## [v1.0.0] - 2022-03-19

### Added

- ipvsctl sub-commands `service`, `server`, `timeout`, `zero`
- Documents README.md

### Fixed

- functions' annotation
