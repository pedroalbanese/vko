package main

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"log"
	"io"
	"math/big"
	"os"

	"github.com/pedroalbanese/gogost/gost3410"
)

var (
	paramset = flag.String("paramset", "A", "ParamSet: A, B, C or D.")
	pubHex   = flag.String("pub", "", "Remote's side public key.")
	prvHex   = flag.String("key", "", "Our private key.")
	keygen   = flag.Bool("keygen", false, "Generate keypair.")
)

func main() {
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage of", os.Args[0]+":")
		flag.PrintDefaults()
		os.Exit(1)
	}

	var curve *gost3410.Curve
	switch *paramset {
	case "A":
		curve = gost3410.CurveIdtc26gost34102012256paramSetA()
	case "B":
		curve = gost3410.CurveIdtc26gost34102012256paramSetD()
	case "C":
		curve = gost3410.CurveIdtc26gost34102012256paramSetC()
	case "D":
		curve = gost3410.CurveIdtc26gost34102012256paramSetD()
	default:
		panic(errors.New("unknown curve specified"))
	}

	var err error
	var prvRaw []byte
	var pubRaw []byte
	var prv *gost3410.PrivateKey
	var pub *gost3410.PublicKey

	if *keygen {
		prvRaw = make([]byte, 256/8)
		_, err = io.ReadFull(rand.Reader, prvRaw)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintln(os.Stderr, "Private:", hex.EncodeToString(prvRaw))
		prv, err = gost3410.NewPrivateKey(curve, prvRaw)
		if err != nil {
			panic(err)
		}

		pub, err = prv.PublicKey()
		if err != nil {
			log.Fatal(err)
		}
		pubRaw = pub.Raw()
		fmt.Fprintln(os.Stderr, "Public:", hex.EncodeToString(pubRaw))
		os.Exit(0)
	}

	prvRaw, err = hex.DecodeString(*prvHex)
	if err != nil {
		log.Fatal(err)
	}
	if len(prvRaw) != 256/8 {
		log.Fatal(errors.New("private key has wrong length"))
	}
	prv, err = gost3410.NewPrivateKey(curve, prvRaw)
	if err != nil {
		log.Fatal(err)
	}
	pubRaw, err = hex.DecodeString(*pubHex)
	if err != nil {
		log.Fatal(err)
	}
	if len(pubRaw) != 2*256/8 {
		log.Fatal(errors.New("public key has wrong length"))
	}
	pub, err = gost3410.NewPublicKey(curve, pubRaw)
	if err != nil {
		log.Fatal(err)
	}

	shared, err := prv.KEK2012256(pub, big.NewInt(1))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Shared:", hex.EncodeToString(shared))
}
