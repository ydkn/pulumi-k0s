[![Actions Status](https://github.com/ydkn/pulumi-k0s/workflows/master/badge.svg)](https://github.com/ydkn/pulumi-k0s/actions)
[![NPM version](https://badge.fury.io/js/%40ydkn%2Fpulumi-k0s.svg)](https://www.npmjs.com/package/@ydkn/pulumi-k0s)
[![Python version](https://badge.fury.io/py/pulumi-k0s.svg)](https://pypi.org/project/pulumi-k0s)
[![NuGet version](https://badge.fury.io/nu/pulumi.k0s.svg)](https://badge.fury.io/nu/pulumi.k0s)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/ydkn/pulumi-k0s/sdk/go)](https://pkg.go.dev/github.com/ydkn/pulumi-k0s/sdk/go)

# k0s Pulumi Provider

Pulumi provider for [k0s](https://k0sproject.io) based on [k0sctl](https://github.com/k0sproject/k0sctl).

## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (Java/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @ydkn/pulumi-k0s

or `yarn`:

    $ yarn add @ydkn/pulumi-k0s

### Python

To use from Python, install using `pip`:

    $ pip install pulumi_k0s

### Go

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/ydkn/pulumi-k0s/sdk

### .NET

To use from .NET, install using `dotnet add package`:

    $ dotnet add package Pulumi.K0s

## Configuration

The following configuration points are available:

- `k0s:skipDowngradeCheck` - Do not check if a downgrade would be performed.
- `k0s:noDrain` - Do not drain nodes before upgrades/updates.
