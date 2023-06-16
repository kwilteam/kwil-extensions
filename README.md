# Kwil Extensions

This repo is a prototype for creating Kwil extensions using Docker.  It may be deleted, archived, or unsupported at any time.

## Extension Overview

The main purpose of an extension is to provide a way to connect arbitrary applications to Kwil databases.  It does this by exposing functionality defined in the extension to a Kwil dataset, allowing applications to use abitrary code in a database.

Kwil extensions run as Docker containers, and act side cars for Kwil databases.  They should run on the same machine as the Kwil database.  **Kwil extensions should be stateless and deterministic**.

## Getting Started

To build a Kwil extension, users will need to import the [server](./server/) package and the [types](./types/) package.   These packages provide tools to create new Kwil extensions.  Currently, there are three parts to a Kwil extension:

- Configuration
- Metadata
- Methods

### Configuration

Extension configuration is used by node operators provide sensitive metadata to extensions.  This includes things like private keys, api keys, or endpoints (e.g. an Infura endpoint).  When node operators start, they will need to include their own configurations for each extension on their network.

### Metadata

Metadata is used to provide "context" for extension calls.  Typically, extension metadata is loaded when a user creates their Kuneiform schema; it can be thought of as the logical equivalent to a constructor, making modules reusable in different datasets running on the same Kwil node.

Since Kwil extensions should be stateless, configured metadata is provided with every call to an extension method.

### Methods

Methods execute code and return data back to Kwil.  They take a list of arguments (scalar values), as well as metadata, and return scalar values.

## Example

An example of a useful Kwil extension would be an ERC20 extension that could read token balances from Ethereum.

#### ERC20 Configuration

An ERC20 extension would likely take an ETH provider (like an Infura or Alchemy endpoint) as its config.  Each node operator would be responsible for providing its own endpoint.

#### ERC20 Metadata

An ERC20 extension would likely take a token address as metadata.  When users create their Kuneiform schemas, they would include the metadata that should be used for this extension.

#### Methods

An ERC20 extension would likely have a method for reading token balances.  It would take one argument (a wallet address), and would return two values: the amount, and thr block height.
