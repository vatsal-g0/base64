# Base64

Simple terminal Base64 encoder/decoder written in Go.

## Usage

```bash
base64 <e|d> <text>
```

## Examples

Encode text:

```bash
base64 e hello
```

Output:

```text
aGVsbG8=
```

Decode text:

```bash
base64 d aGVsbG8=
```

Output:

```text
hello
```

## Build

```bash
go build
```

## Notes

- `e` = encode
- `d` = decode
