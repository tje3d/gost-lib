# GOST for Android

[![Build Status](https://travis-ci.org/ginuerzh/gost.svg?branch=master)](https://travis-ci.org/ginuerzh/gost)
[![Go Report Card](https://goreportcard.com/badge/github.com/ginuerzh/gost)](https://goreportcard.com/report/github.com/ginuerzh/gost)

[English](README_en.md) | [中文](README.md)

GOST (GO Simple Tunnel) is a versatile and secure tunnel program written in Go.

This repository is a fork of the original GOST project, specifically tailored for creating an Android Archive (AAR) library using `gomobile`. This allows you to embed GOST's powerful networking capabilities directly into your Android applications.

## Features

*   Multi-platform support (Windows, Linux, macOS, Android, iOS)
*   Multi-protocol support (various proxy and tunneling protocols)
*   Chainable proxy nodes for complex routing
*   TLS/SSL encryption
*   Multiplexing and connection pooling

## Building for Android

This repository is configured to be built into an AAR library for Android using `gomobile`.

### Prerequisites

*   [Go](https://golang.org/dl/) (version 1.13+)
*   [Android NDK](https://developer.android.com/ndk/downloads)
*   `gomobile` tool

Install `gomobile` with the following command:

```bash
go get golang.org/x/mobile/cmd/gomobile
gomobile init
```

### Build Command

To build the AAR file, run the following command from the root of this repository:

```bash
gomobile bind -target=android -o gost.aar -androidapi 21 -ldflags="-s -w" ./gost/gostmobile
```

This command will generate `gost.aar` in the root directory. You can then import this AAR file into your Android Studio project.

## Usage in Android

Once the AAR is included in your project, you can call the exported functions from the `gostmobile` package.

**Example:**

```java
import gostmobile.Gostmobile;

// Start the tunnel
try {
    Gostmobile.startTunnel("ssh", "user:pass@host:port", "username", "password");
} catch (Exception e) {
    // Handle error
}

// Stop the tunnel
try {
    Gostmobile.stopTunnel();
} catch (Exception e) {
    // Handle error
}
```

For more details on the original GOST project, please see the official [gost repository](https://github.com/ginuerzh/gost).