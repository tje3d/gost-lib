# GOST Mobile Library

[![Build Status](https://github.com/tje3d/gost-lib/workflows/Build%20Android%20AAR%20and%20Create%20Release/badge.svg)](https://github.com/tje3d/gost-lib/actions)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Android API](https://img.shields.io/badge/API-21%2B-brightgreen.svg?style=flat)](https://android-arsenal.com/api?level=21)

A Go mobile library that provides basic GOST tunneling functionality for embedding into mobile applications. This library wraps the [GOST (GO Simple Tunnel)](https://github.com/ginuerzh/gost) project using Google's `gomobile` tool.

**Note: Functionality is currently very limited. Pull requests are welcome to expand capabilities.**

## What is this?

This repository provides a simplified mobile binding for GOST, allowing developers to embed basic tunneling functionality into their applications. The current implementation (see `gostlib.go`) offers minimal functionality with room for significant expansion.

## Current Features

The library currently provides basic functionality through three main functions:

- `StartTunnel(transport, addr, username, password string)` - Start a SOCKS5 tunnel
- `StopTunnel()` - Stop the running tunnel
- Support for basic transports: WebSocket (ws), WebSocket Secure (wss), SSH

**Limitations:**

- Only SOCKS5 proxy support
- Limited transport options
- No advanced GOST features (chaining, load balancing, etc.)
- Basic error handling

## Contributing

This project needs significant development to unlock GOST's full potential on mobile platforms. Areas where contributions are especially welcome:

- Expanding transport protocol support
- Adding proxy chaining capabilities
- Implementing advanced GOST features
- Improving error handling and logging
- Adding configuration options
- Platform-specific optimizations

See `gostlib.go` for the current implementation.

## Platform Support

| Platform    | Status    | Library Format | API Language      |
| ----------- | --------- | -------------- | ----------------- |
| **Android** | Available | AAR            | Java/Kotlin       |
| **iOS**     | Planned   | Framework      | Swift/Objective-C |

## Installation

### Option 1: Download Pre-built Library

#### Android (AAR)

1. Go to [Releases](https://github.com/tje3d/gost-lib/releases)
2. Download the latest `gost.aar` file
3. Add it to your Android project:

```gradle
// In your app's build.gradle
dependencies {
    implementation files('libs/gost.aar')
}
```

### Option 2: Build from Source

**Prerequisites:**

- Go 1.21+
- Android SDK with NDK (for Android builds)
- Java 17+

**Build steps:**

```bash
# Install gomobile
go install golang.org/x/mobile/cmd/gomobile@latest
gomobile init

# Build Android AAR
gomobile bind -target=android -o gost.aar -androidapi 21 -ldflags="-s -w" .
```

## Usage Examples

### Android (Java/Kotlin)

#### Basic Tunnel Setup

```java
import gostmobile.Gostmobile;

public class TunnelManager {

    public void startSecureTunnel() {
        try {
            // Start a SOCKS5 proxy tunnel
            Gostmobile.startTunnel(
                "socks5",
                "user:password@proxy.example.com:1080",
                "myusername",
                "mypassword"
            );

            Log.i("Tunnel", "Secure tunnel started successfully!");

        } catch (Exception e) {
            Log.e("Tunnel", "Failed to start tunnel: " + e.getMessage());
        }
    }

    public void stopTunnel() {
        try {
            Gostmobile.stopTunnel();
            Log.i("Tunnel", "Tunnel stopped");
        } catch (Exception e) {
            Log.e("Tunnel", "Error stopping tunnel: " + e.getMessage());
        }
    }
}
```

#### Advanced Configuration

```java
// HTTP proxy with authentication
Gostmobile.startTunnel(
    "http",
    "corporate-proxy.company.com:8080",
    "employee_id",
    "secure_password"
);

// SSH tunnel for secure communication
Gostmobile.startTunnel(
    "ssh",
    "ssh-server.example.com:22",
    "ssh_username",
    "ssh_password"
);
```

#### Integration with Android Network Stack

```java
public class NetworkService {

    public void setupProxiedConnection() {
        // Start GOST tunnel
        startSecureTunnel();

        // Configure your HTTP client to use the local proxy
        OkHttpClient client = new OkHttpClient.Builder()
            .proxy(new Proxy(Proxy.Type.SOCKS,
                   new InetSocketAddress("127.0.0.1", 1080)))
            .build();

        // Now all requests go through the GOST tunnel!
        Request request = new Request.Builder()
            .url("https://api.example.com/data")
            .build();

        client.newCall(request).enqueue(callback);
    }
}
```

### iOS (Coming Soon)

```swift
// Future iOS API (planned)
import GostMobile

class TunnelManager {
    func startSecureTunnel() {
        do {
            try GostMobile.startTunnel(
                protocol: "socks5",
                endpoint: "user:password@proxy.example.com:1080",
                username: "myusername",
                password: "mypassword"
            )
            print("Secure tunnel started successfully!")
        } catch {
            print("Failed to start tunnel: \(error)")
        }
    }
}
```

## Requirements

### Android Runtime

- Android API Level 21+ (Android 5.0+)
- Permissions: `INTERNET` (add to AndroidManifest.xml)

### Build Requirements

- Go 1.21+
- Android SDK with NDK (for Android builds)
- Java 17+

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

The original GOST project is also MIT licensed. See the [GOST repository](https://github.com/ginuerzh/gost) for more information.
