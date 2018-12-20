package main

import (
	"github.com/pavel-v-chernykh/keystore-go"
	"log"
	"os"
)

func readKeyStore(filename string, password []byte) keystore.KeyStore {
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	keyStore, err := keystore.Decode(f, password)
	if err != nil {
		log.Fatal(err)
	}
	return keyStore
}

func writeKeyStore(keyStore keystore.KeyStore, filename string, password []byte) {
	o, err := os.Create(filename)
	defer o.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = keystore.Encode(o, keyStore, password)
	if err != nil {
		log.Fatal(err)
	}
}

func zeroing(s []byte) {
	for i := 0; i < len(s); i++ {
		s[i] = 0
	}
}

func main() {
	password := []byte{'e', 'm', 'r', 'u', 'z', '0'}
	defer zeroing(password)
	ks:=keystore.KeyStore{}


	writeKeyStore(ks,"emruz.jks",password)
	//ks1 := readKeyStore("keystore.jks", password)
	//fmt.Println(reflect.TypeOf(ks1))
	//oneliners.PrettyJson(ks1)
	//writeKeyStore(ks1, "keystore2.jks", password)
	//
	//ks2 := readKeyStore("keystore2.jks", password)
	//fmt.Println(ks2)
	//log.Printf("Is equal: %v\n", reflect.DeepEqual(ks1, ks2))
}