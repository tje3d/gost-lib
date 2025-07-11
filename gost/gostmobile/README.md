# GOST for Android (gomobile)

This directory contains the Go code specifically prepared for use with `gomobile`. The primary purpose is to compile the GOST core functionalities into an Android Archive (AAR) library.

## Building the AAR

To build the AAR file, you will need to have the Android NDK and `gomobile` installed and configured correctly.

Once your environment is set up, you can generate the AAR by running the following command from the `gost/gostmobile` directory:

```bash
gomobile bind -target=android
```

This will produce a `.aar` file that can be included as a dependency in your Android projects.