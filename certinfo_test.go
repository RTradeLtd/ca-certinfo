package certinfo

import (
	"bytes"
	"encoding/pem"
	"io/ioutil"
	"testing"

	"github.com/RTradeLtd/ca-cli/pkg/x509"
)

type InputType int

const (
	tCertificate InputType = iota
	tCertificateRequest
)

// Compares a PEM-encoded certificate to a reference file.
func testPair(t *testing.T, certFile, refFile string, inputType InputType) {
	// Read and parse the certificate
	pemData, err := ioutil.ReadFile(certFile)
	if err != nil {
		t.Fatal(err)
	}
	block, rest := pem.Decode([]byte(pemData))
	if block == nil || len(rest) > 0 {
		t.Fatal("Certificate decoding error")
	}
	var result string
	switch inputType {
	case tCertificate:
		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			t.Fatal(err)
		}
		result, err = CertificateText(cert)
		if err != nil {
			t.Fatal(err)
		}
	case tCertificateRequest:
		cert, err := x509.ParseCertificateRequest(block.Bytes)
		if err != nil {
			t.Fatal(err)
		}
		result, err = CertificateRequestText(cert)
		if err != nil {
			t.Fatal(err)
		}
	}
	resultData := []byte(result)

	// Read the reference output
	refData, err := ioutil.ReadFile(refFile)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(resultData, refData) {
		t.Logf("'%s' did not match reference '%s'\n", certFile, refFile)
		t.Errorf("Dump follows:\n%s\n", result)
	}
}

// Compares a PEM-encoded certificate to a reference file.
func testPairShort(t *testing.T, certFile, refFile string, inputType InputType) {
	// Read and parse the certificate
	pemData, err := ioutil.ReadFile(certFile)
	if err != nil {
		t.Fatal(err)
	}
	block, rest := pem.Decode([]byte(pemData))
	if block == nil || len(rest) > 0 {
		t.Fatal("Certificate decoding error")
	}
	var result string
	switch inputType {
	case tCertificate:
		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			t.Fatal(err)
		}
		result, err = CertificateShortText(cert)
		if err != nil {
			t.Fatal(err)
		}
	case tCertificateRequest:
		cert, err := x509.ParseCertificateRequest(block.Bytes)
		if err != nil {
			t.Fatal(err)
		}
		result, err = CertificateRequestShortText(cert)
		if err != nil {
			t.Fatal(err)
		}
	}
	resultData := []byte(result)

	// Read the reference output
	refData, err := ioutil.ReadFile(refFile)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(resultData, refData) {
		t.Logf("'%s' did not match reference '%s'\n", certFile, refFile)
		t.Errorf("Dump follows:\n%s\n", result)
	}
}

// Test the root CA certificate
func TestCertInfoRoot(t *testing.T) {
	testPair(t, "test_certs/root1.cert.pem", "test_certs/root1.cert.text", tCertificate)
	testPair(t, "test_certs/root1.csr.pem", "test_certs/root1.csr.text", tCertificateRequest)
	testPairShort(t, "test_certs/root1.cert.pem", "test_certs/root1.cert.short", tCertificate)
	testPairShort(t, "test_certs/root1.csr.pem", "test_certs/root1.csr.short", tCertificateRequest)
}

// Test the leaf (user) RSA certificate
func TestCertInfoLeaf1(t *testing.T) {
	testPair(t, "test_certs/leaf1.cert.pem", "test_certs/leaf1.cert.text", tCertificate)
	testPair(t, "test_certs/leaf1.csr.pem", "test_certs/leaf1.csr.text", tCertificateRequest)
	testPairShort(t, "test_certs/leaf1.cert.pem", "test_certs/leaf1.cert.short", tCertificate)
	testPairShort(t, "test_certs/leaf1.csr.pem", "test_certs/leaf1.csr.short", tCertificateRequest)
}

// Test the leaf (user) DSA certificate
func TestCertInfoLeaf2(t *testing.T) {
	testPair(t, "test_certs/leaf2.cert.pem", "test_certs/leaf2.cert.text", tCertificate)
	testPair(t, "test_certs/leaf2.csr.pem", "test_certs/leaf2.csr.text", tCertificateRequest)
	testPairShort(t, "test_certs/leaf2.cert.pem", "test_certs/leaf2.cert.short", tCertificate)
	testPairShort(t, "test_certs/leaf2.csr.pem", "test_certs/leaf2.csr.short", tCertificateRequest)
}

// Test the leaf (user) ECDSA certificate
func TestCertInfoLeaf3(t *testing.T) {
	testPair(t, "test_certs/leaf3.cert.pem", "test_certs/leaf3.cert.text", tCertificate)
	testPair(t, "test_certs/leaf3.csr.pem", "test_certs/leaf3.csr.text", tCertificateRequest)
	testPairShort(t, "test_certs/leaf3.cert.pem", "test_certs/leaf3.cert.short", tCertificate)
	testPairShort(t, "test_certs/leaf3.csr.pem", "test_certs/leaf3.csr.short", tCertificateRequest)
}

// Test the leaf (user) with multiple sans
func TestCertInfoLeaf4(t *testing.T) {
	testPair(t, "test_certs/leaf4.cert.pem", "test_certs/leaf4.cert.text", tCertificate)
	testPair(t, "test_certs/leaf4.csr.pem", "test_certs/leaf4.csr.text", tCertificateRequest)
	testPairShort(t, "test_certs/leaf4.cert.pem", "test_certs/leaf4.cert.short", tCertificate)
	testPairShort(t, "test_certs/leaf4.csr.pem", "test_certs/leaf4.csr.short", tCertificateRequest)
}
