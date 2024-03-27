package main

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

func filterCACertificate(pCaCertificate *x509.Certificate, cmpCertificates []*x509.Certificate) (result bool, err error) {

	result = false

	if time.Now().After(pCaCertificate.NotAfter) {
		return
	}

	for _, c := range cmpCertificates {
		if c.Equal(pCaCertificate) {
			return
		}
	}

	result = true
	return
}

func readCACertificates(pOutCertificates *[]*x509.Certificate) (err error) {

	data, err := os.ReadFile(cabakFile)
	if err != nil {
		return
	}

	for {
		block, rest := pem.Decode(data)
		data = rest
		if block == nil {
			break
		}

		switch block.Type {
		case "CERTIFICATE":
			pCertificate, cerr := x509.ParseCertificate(block.Bytes)
			if cerr != nil {
				err = errors.Join(
					err,
					cerr)
				continue
			}
			filter, cerr := filterCACertificate(pCertificate, *pOutCertificates)
			if cerr != nil {
				err = errors.Join(
					err,
					cerr)
			}
			if filter {
				*pOutCertificates = append(*pOutCertificates, pCertificate)
			}

		default:
			err = errors.Join(
				err,
				errors.New(fmt.Sprintf("pem type %s not supported.\n", block.Type)))

		}
	}

	return
}

func writeCACertificates(certificates []*x509.Certificate) (err error) {

	file, err := os.Create(caFile)
	if err != nil {
		return
	}
	defer (func() { err = file.Close() })()

	for _, c := range certificates {
		block := pem.Block{
			Type:  "CERTIFICATE",
			Bytes: c.Raw}
		cerr := pem.Encode(file, &block)
		if cerr != nil {
			err = errors.Join(
				err,
				cerr)
		}
	}

	return
}

func main() {

	err := initConstants()
	if err != nil {
		log.Fatal(err)
	}

	// back up copy ca file
	err = os.Rename(caFile, cabakFile)
	if err != nil {
		log.Fatal(err)
	}

	var caCertificates []*x509.Certificate

	err = readCACertificates(&caCertificates)
	if err != nil {
		log.Fatal(err)
	}

	err = writeCACertificates(caCertificates)
	if err != nil {
		log.Fatal(err)
	}

}
