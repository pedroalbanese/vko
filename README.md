# VKO 
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](https://github.com/pedroalbanese/vko/blob/master/LICENSE.md) 
[![GitHub downloads](https://img.shields.io/github/downloads/pedroalbanese/vko/total.svg?logo=github&logoColor=white)](https://github.com/pedroalbanese/vko/releases)
[![GoDoc](https://godoc.org/github.com/pedroalbanese/vko?status.png)](http://godoc.org/github.com/pedroalbanese/vko)
[![Go Report Card](https://goreportcard.com/badge/github.com/pedroalbanese/vko)](https://goreportcard.com/report/github.com/pedroalbanese/vko)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/pedroalbanese/vko)](https://golang.org)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/pedroalbanese/vko)](https://github.com/pedroalbanese/vko/releases)

### GOST R 34.10-2012 VKO (выработка ключа общего)
VKO is an elliptic curve Diffie-Hellman key agreement function using GOST R 34.10-2012. It allows two parties to jointly agree on a shared secret using an insecure channel.
### Command-line VKO Tool:
<pre>Usage of vko:
  -key string
        Our private key.
  -keygen
        Generate keypair.
  -paramset string
        ParamSet. (default "A")
  -pub string
        Remote's side public key.</pre>

### Examples:
```sh
./vko -keygen // 2x
./vko -key $2ndprivatekey -pub $1stpublickey
./vko -key $1stprivatekey -pub $2ndpublickey
```
## License
This project is licensed under the ISC License.
##### Copyright (c) 2020-2021 Pedro Albanese - ALBANESE Research Lab.
