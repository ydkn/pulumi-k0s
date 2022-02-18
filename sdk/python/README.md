[![Actions Status](https://github.com/ydkn/pulumi-k0s/workflows/master/badge.svg)](https://github.com/ydkn/pulumi-k0s/actions)
[![NPM version](https://badge.fury.io/js/%40pulumi%2Faws.svg)](https://www.npmjs.com/package/@pulumi/k0s)
[![Python version](https://badge.fury.io/py/pulumi-k0s.svg)](https://pypi.org/project/pulumi-k0s)
[![NuGet version](https://badge.fury.io/nu/pulumi.k0s.svg)](https://badge.fury.io/nu/pulumi.k0s)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/ydkn/pulumi-k0s/sdk/go)](https://pkg.go.dev/github.com/ydkn/pulumi-k0s/sdk/go)

# k0s Pulumi Provider

The Amazon Web Services (AWS) resource provider for Pulumi lets you use AWS resources in your cloud programs. To use
this package, please [install the Pulumi CLI first](https://pulumi.com/). For a streamlined Pulumi walkthrough, including language runtime installation and AWS configuration, click "Get Started" below.

## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (Java/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @pulumi/k0s

or `yarn`:

    $ yarn add @pulumi/k0s

### Python

To use from Python, install using `pip`:

    $ pip install pulumi_k0s

### Go

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/pulumi/pulumi-k0s/sdk

### .NET

To use from .NET, install using `dotnet add package`:

    $ dotnet add package Pulumi.K0s

## Configuration

The following configuration points are available:

- `k0s:skipDowngradeCheck` - Do not check if a downgrade would be performed.
- `k0s:noDrain` - Do not drain nodes before upgrades/updates.
