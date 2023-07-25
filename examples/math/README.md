# Kwil Math Extension

This directory contains an example of a math extension written in Kwil.  This extension likely has no real world use cases, and merely exists as an example
of how to build an extension.

## Layout

The application entrypoint (which starts the gRPC server) can be found in [main.go](./main.go).  The logic of the extension can be found in [extension/math.go](./extension/math.go).  This logic is included in the [server builder](./extension/server.go), which utilizes the [Kwil-Go extension toolkit](../../server/)

## Initialization

The extension can be initialized by specifying the direction non-integer results should be rounded.  The initialization function checks what a user passes, and defaults to rounding up.

## Methods

The extension contains the following methods:

- `add`: prompts the user for values to add (in the deployed database)
- `sub`: prompts the user for values to subtract (in the deployed database)
- `mul`: prompts the user for values to multiply (in the deployed database)
- `div`: prompts the user for values to divide (in the deployed database)
