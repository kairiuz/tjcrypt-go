# tjcrypt-go

Golang port of the previous ndk based project [tjcrypt](https://git.teknik.io/wobm/tjcrypt)

## Build

This module uses CGO extension, used by LZ4 decompression/compression.

```
CGO_ENABLED=1 go build ./cmd/...
```

There is an alredy prebuilt binaries in [releases](https://git.teknik.io/wobm/tjcrypt-go/releases).
Currently only support linux, windows, and android. Support for OSX is planned
in the near future.

## Usage
If the output path isn't supplied, the output will be directly written to
stdout.

Decrypt:
```
tjdecrypt /path/to/encrypted/file [/path/to/output]
```

Encrypt:
```
tjencrypt /path/to/encrypted/file [/path/to/output]
```

## API

*Currently only support text type encryption, no blob files like png, jpg, etc.*

```go
// Decrypt supplied data and return the decrypted data
tjcrypt.Decrypt(data []byte) ([]byte, error)
```

```go
// Encrypt supplied data and return the encrypted data, returned []byte also
// includes the tjcrypt header
tjcrypt.Encrypt(data []byte) ([]byte, error)
```

