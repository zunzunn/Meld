# MELD

A universal, local-first file utility that combines file conversion, compression, optimization, and other file operations into one application.

## Goals

- Provide a powerful Unix-inspired CLI for developers.
- Provide a beautiful, minimal, keyboard-friendly TUI for non-technical users.
- Share a unified processing engine between both interfaces.
- Support macOS, Linux, and Windows.
- Support images, videos, audio, documents, archives, and other file types.
- Automatically detect and install dependencies.
- Support batch processing and concurrent execution.
- Provide an extensible plugin system.
- Ensure safe file processing that preserves original files by default.

## Planned Features

- CLI with file conversion, compression, and optimization commands.
- TUI with keyboard-driven navigation.
- Plugin system for extending file type support.
- Batch processing and concurrent execution.
- Cross-platform support (macOS, Linux, Windows).

## Current Development Status

Initial foundation laid. CLI framework, configuration, and error handling are implemented.
TUI and file processing features are planned for future iterations.

## Building and Running

```bash
go build ./cmd/meld
./meld help
./meld version
```