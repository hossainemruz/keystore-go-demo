package pkg

import (
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	keystore "github.com/pavel-v-chernykh/keystore-go"
	"github.com/pkg/errors"
)

const (
	CERT_TYPE_PRIVATE_KEY = "PRIVATE KEY"
	CERT_TYPE_CERTIFICATE = "CERTIFICATE"
)

func ImportCert(certFile, keyStoreName, pass, alias string) error {

	// read the cert file
	pke, err := ioutil.ReadFile(certFile)
	if err != nil {
		return errors.Wrapf(err, "failed to read cert file: %s", certFile)
	}

	// decode certificate
	p, _ := pem.Decode(pke)
	fmt.Println("Cert type: ", p.Type)

	var ks keystore.KeyStore

	switch p.Type {
	case CERT_TYPE_PRIVATE_KEY:
		fmt.Println("type privatekey")
		ks = keystore.KeyStore{
			alias: &keystore.PrivateKeyEntry{
				Entry: keystore.Entry{
					CreationDate: time.Now(),
				},
				PrivKey: p.Bytes,
			},
		}
	case CERT_TYPE_CERTIFICATE:
		fmt.Println("type certificate")
		ks = keystore.KeyStore{
			alias: &keystore.TrustedCertificateEntry{
				Entry: keystore.Entry{
					CreationDate: time.Now(),
				},
				Certificate: keystore.Certificate{
					Type:    "x509",
					Content: p.Bytes,
				},
			},
		}
	default:
		fmt.Println("unknown certificate type")

	}

	// now write cert file in file on keystore
	err = writeInKeyStore(ks, keyStoreName, pass)
	if err != nil {
		return errors.Wrapf(err, "failed to write certificate in keystore")
	}

	return nil
}

func writeInKeyStore(keyStore keystore.KeyStore, filename string, password string) error {
	o, err := os.Create(filename)
	defer o.Close()
	if err != nil {
		return err
	}
	err = keystore.Encode(o, keyStore, []byte(password))
	if err != nil {
		return err
	}
	return nil
}
