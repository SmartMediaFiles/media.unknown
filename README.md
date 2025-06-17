# ❓ media.unknown

[![Work in Progress](https://img.shields.io/badge/Status-Work%20in%20Progress-yellow)](https://shields.io)
[![Go Report Card](https://goreportcard.com/badge/github.com/SmartMediaFiles/media.unknown)](https://goreportcard.com/report/github.com/SmartMediaFiles/media.unknown)
[![GoDoc](https://pkg.go.dev/badge/github.com/SmartMediaFiles/media.unknown)](https://pkg.go.dev/github.com/SmartMediaFiles/media.unknown)
[![Release](https://img.shields.io/github/release/SmartMediaFiles/media.unknown.svg?style=flat)](https://github.com/SmartMediaFiles/media.unknown/releases)

## Overview

`media.unknown` is a specialized library within the **SmartMediaFiles ecosystem**. Its purpose is to handle files whose type cannot be determined from their extension. It acts as a fallback for the main `@/media` library.

## Features

- **Fallback Mechanism**: Provides the `Unknown` media type definition for files that do not match any other known media type.
- **Future-proofing**: This library could be enhanced with logic to analyze file headers or content to try to identify the file type more accurately.

## Installation

```bash
go get -u github.com/smartmediafiles/media.unknown
```

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

---

⚠️ **Note:** This README will be updated regularly as the project progresses. Check back often for the latest information!
