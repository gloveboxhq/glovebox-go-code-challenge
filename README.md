# carrierproxy-poc

This repository is used for creating a proof of concept around the `PolicyProvider` interface.

## Project Setup

Before attempting to run any tests provided within this repo, please copy the file `.env.example` (new file should be named `.env`) and input valid Netflix account credentials in corresponding fields. This new file (`.env`) is used by the repo Makefile
to load testing credentials.

## Running Tests

Run `make test` to execute tests of the Netflix policy provider. Please be aware that depending on your network conditions
(and other factors), *Netflix may present a captcha on the login interface.* If this happens, expect all tests to fail completely as there is no easy way to automate captcha completion.

## Useful Links

- [Rod](https://go-rod.github.io/#/) Library used for web automation.
