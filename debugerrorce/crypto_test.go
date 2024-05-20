package debugerrorce

import (
	"encoding/pem"
	"errors"
	"fmt"
	"strings"
	"testing"

	"crypto/x509"

	"github.com/stretchr/testify/assert"
)

func TestStrSha256(t *testing.T) {
	const test = "bla"
	const testExpected = "4df3c3f68fcc83b27e9d42c90431a72499f17875c81a599b566c9889b9696703"
	assert.Equal(t, testExpected, fmt.Sprintf("%x", Str2sha256(test)), "Values not equal")
}

func TestStrSha256Base64(t *testing.T) {
	const test = "bla"
	//CondDebugSet(true)
	const testExpected = "TfPD9o/Mg7J+nULJBDGnJJnxeHXIGlmbVmyYiblpZwM="
	assert.Equal(t, testExpected, Bytea2b64(Str2sha256(test)), "Values not equal")
}

func TestString2Md5(t *testing.T) {
	// a0cdcc72e656541f132bf96747ee17dc	# echo -n 'blablue'| md5sum
	const test = "blablue"
	assert.Equal(t, "a0cdcc72e656541f132bf96747ee17dc", String2md5(test))
}

func TestPemToPubKey(t *testing.T) {
	const publicKey = `
		MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEA9sc3FCBC3x9kTPpKTzl/
		qQct3NcjfrTsrqNloaTCncXtSAln2X+yClCmaVIrpQOyL7TbCXniKojmJhMOdhfH
		V6sWmFiFQV3XJyzQFXbCX3dE/v5uy21l4xrZtkLcX4JzsS6fpsf2avO48OM6ZCNO
		MHO5ifXUoHYVG5ApP4P5B4j0AVg7rSb4HWIX2cv+K6+p47dYgV5N2XO0z6g+ZsK6
		yAuklaHU5b1yhrYjpRdXgCeukwaNHI8YqiDpSWrSxE5pmBsL2EP3z5jLydgwacPJ
		x1MEI+4a4ta0ivsr1sgrd5UwvmrnVRhn/3Vl8Q5AKie3zpOhtiH3mhZOwhxndlsG
		5T0v6RY1/ZEdMYSl/DSaYYZQgEqsiJJQLpsgfNZZJI4fPfHiaRvhDVB8O78CwNzj
		20mHCymY9pgFStdsdneFsZr6dFwyCtDCI9uXv1jNnr+x3GSqlR4fIsZOzNGOkR15
		yXjbSYwCeegJJsvUp15jGaKt6QVKQSaXjfKVG2wOzIiNJCrjrme1k4p2Fte+/Qkl
		xPmL0nPjvIuyLZmeNRVNy8SroSvC5YoGyvWWQkl5QOQtRM/nA84jriVw0q2/YacN
		QQ5cLFehoFQqJB2wn+x7wSrSDgeOHC2S2QQXd1GTkRMPNfgMBIQrgprGmcnkD5Uv
		RaYRL1gjPNuOwGW0lLt/lDsCAwEAAQ==
	`
	str := "-----BEGIN PUBLIC KEY-----" + // line feeds not required
		strings.Replace(strings.Replace(publicKey, " ", "", -1), "\t", "", -1) + // remove spaces + tabs
		"-----END PUBLIC KEY-----"
	block, _ := pem.Decode([]byte(str))
	assert.NotNil(t, block, "pem.Decode error")

	// fmt.Print(strings.Replace(strings.Replace(publicKey, " ", "", -1), "\t", "", -1))
	_, err := Pem2RsaPublicKey([]byte(str))

	assert.Nil(t, err, "could not decipher public key")
}

func TestACorrectCSR(t *testing.T) {
	const pemCsr = `
-----BEGIN CERTIFICATE REQUEST-----
MIIEbTCCAlUCAQAwKDELMAkGA1UEBhMCWloxGTAXBgNVBAMMEG9wZW5zc2xyZXFf
dW5lbmMwggIiMA0GCSqGSIb3DQEBAQUAA4ICDwAwggIKAoICAQD2xzcUIELfH2RM
+kpPOX+pBy3c1yN+tOyuo2WhpMKdxe1ICWfZf7IKUKZpUiulA7IvtNsJeeIqiOYm
Ew52F8dXqxaYWIVBXdcnLNAVdsJfd0T+/m7LbWXjGtm2QtxfgnOxLp+mx/Zq87jw
4zpkI04wc7mJ9dSgdhUbkCk/g/kHiPQBWDutJvgdYhfZy/4rr6njt1iBXk3Zc7TP
qD5mwrrIC6SVodTlvXKGtiOlF1eAJ66TBo0cjxiqIOlJatLETmmYGwvYQ/fPmMvJ
2DBpw8nHUwQj7hri1rSK+yvWyCt3lTC+audVGGf/dWXxDkAqJ7fOk6G2IfeaFk7C
HGd2WwblPS/pFjX9kR0xhKX8NJphhlCASqyIklAumyB81lkkjh898eJpG+ENUHw7
vwLA3OPbSYcLKZj2mAVK12x2d4Wxmvp0XDIK0MIj25e/WM2ev7HcZKqVHh8ixk7M
0Y6RHXnJeNtJjAJ56Akmy9SnXmMZoq3pBUpBJpeN8pUbbA7MiI0kKuOuZ7WTinYW
1779CSXE+YvSc+O8i7ItmZ41FU3LxKuhK8LligbK9ZZCSXlA5C1Ez+cDziOuJXDS
rb9hpw1BDlwsV6GgVCokHbCf7HvBKtIOB44cLZLZBBd3UZOREw81+AwEhCuCmsaZ
yeQPlS9FphEvWCM8247AZbSUu3+UOwIDAQABoAAwDQYJKoZIhvcNAQELBQADggIB
ADD0YH1pW2lOyqoT3n3cGeM4iPt6MMtHek4T6+lVImEXzfoioU5GEv+pfZBdG9wA
waQAZ+cb/x7BNStM6ZpdvZYZKX81jHcrx85sprk5oMgcrCTkawVZnvG5SC01FsUD
0BmowXRpEM/5h/wFdpDRfg3lvR65pNtCZadXydCtumQISo7IKbLHxWNF/be07zVy
QCL2c6wR1LHJyfH9GOeCLUyCHifjuOzNdVTvpuqnnHBSK1v86XW/zBHxiPsKPsP1
gLt8u5Da7/gtFZkYHAPDKbkY9wljDMIY7k7BOy0r7wxq2Bx1vyIGdt4RDlDLg3yC
zJp4eiWmLRAJ9xUFR1Zm9uXhJ2MSaSPsOH6ctoK47KP09um8hUUKwXFvmsnoH0St
WsJJtHkvKMxfJe7qQKO6efWqkEZcLQJf0NWeNbFWpMBx5P+b89eKXvl2U3ESp1am
vSXfHySld88mTB3jEucwBtm7NelhVhj3kwxB0VFrsGUnnW6gOFQlM+U/MJH23mjE
oqjJjUQ8ErPfVWZXsV7aBAakFgcyPSefCafENlKfDRI4KVQxWtWXOZeGp0/pzpEw
SrwZoQ0iFZlO0osLRb+A3S6Jwf4Ls55eZU9HgicgWNs3xhyuEgUoJ8QrIyFst1HH
pQ3dQDkR/QCqbIaO7P/1J8YDCkBqCUUV7xWp74dmSzbP
-----END CERTIFICATE REQUEST-----
	`
	csr, err := Pem2CSR([]byte(pemCsr))
	assert.Nil(t, err, "error in Pem2CSR")
	_, err = Any2RsaPublicKey(csr.PublicKey)
	assert.Nil(t, err, "error in Any2RsaPublicKey")
}

func TestUnformattedFailingCSR(t *testing.T) {
	const pemCsr = `
		-----BEGIN CERTIFICATE REQUEST-----
		MIIEbTCCAlUCAQAwKDELMAkGA1UEBhMCWloxGTAXBgNVBAMMEG9wZW5zc2xyZXFf
		dW5lbmMwggIiMA0GCSqGSIb3DQEBAQUAA4ICDwAwggIKAoICAQD2xzcUIELfH2RM
		+kpPOX+pBy3c1yN+tOyuo2WhpMKdxe1ICWfZf7IKUKZpUiulA7IvtNsJeeIqiOYm
		Ew52F8dXqxaYWIVBXdcnLNAVdsJfd0T+/m7LbWXjGtm2QtxfgnOxLp+mx/Zq87jw
		4zpkI04wc7mJ9dSgdhUbkCk/g/kHiPQBWDutJvgdYhfZy/4rr6njt1iBXk3Zc7TP
		qD5mwrrIC6SVodTlvXKGtiOlF1eAJ66TBo0cjxiqIOlJatLETmmYGwvYQ/fPmMvJ
		2DBpw8nHUwQj7hri1rSK+yvWyCt3lTC+audVGGf/dWXxDkAqJ7fOk6G2IfeaFk7C
		HGd2WwblPS/pFjX9kR0xhKX8NJphhlCASqyIklAumyB81lkkjh898eJpG+ENUHw7
		vwLA3OPbSYcLKZj2mAVK12x2d4Wxmvp0XDIK0MIj25e/WM2ev7HcZKqVHh8ixk7M
		0Y6RHXnJeNtJjAJ56Akmy9SnXmMZoq3pBUpBJpeN8pUbbA7MiI0kKuOuZ7WTinYW
		1779CSXE+YvSc+O8i7ItmZ41FU3LxKuhK8LligbK9ZZCSXlA5C1Ez+cDziOuJXDS
		rb9hpw1BDlwsV6GgVCokHbCf7HvBKtIOB44cLZLZBBd3UZOREw81+AwEhCuCmsaZ
		yeQPlS9FphEvWCM8247AZbSUu3+UOwIDAQABoAAwDQYJKoZIhvcNAQELBQADggIB
		ADD0YH1pW2lOyqoT3n3cGeM4iPt6MMtHek4T6+lVImEXzfoioU5GEv+pfZBdG9wA
		waQAZ+cb/x7BNStM6ZpdvZYZKX81jHcrx85sprk5oMgcrCTkawVZnvG5SC01FsUD
		0BmowXRpEM/5h/wFdpDRfg3lvR65pNtCZadXydCtumQISo7IKbLHxWNF/be07zVy
		QCL2c6wR1LHJyfH9GOeCLUyCHifjuOzNdVTvpuqnnHBSK1v86XW/zBHxiPsKPsP1
		gLt8u5Da7/gtFZkYHAPDKbkY9wljDMIY7k7BOy0r7wxq2Bx1vyIGdt4RDlDLg3yC
		zJp4eiWmLRAJ9xUFR1Zm9uXhJ2MSaSPsOH6ctoK47KP09um8hUUKwXFvmsnoH0St
		WsJJtHkvKMxfJe7qQKO6efWqkEZcLQJf0NWeNbFWpMBx5P+b89eKXvl2U3ESp1am
		vSXfHySld88mTB3jEucwBtm7NelhVhj3kwxB0VFrsGUnnW6gOFQlM+U/MJH23mjE
		oqjJjUQ8ErPfVWZXsV7aBAakFgcyPSefCafENlKfDRI4KVQxWtWXOZeGp0/pzpEw
		SrwZoQ0iFZlO0osLRb+A3S6Jwf4Ls55eZU9HgicgWNs3xhyuEgUoJ8QrIyFst1HH
		pQ3dQDkR/QCqbIaO7P/1J8YDCkBqCUUV7xWp74dmSzbP
		-----END CERTIFICATE REQUEST-----
	`
	_, err := Pem2CSR([]byte(pemCsr))
	assert.NotNil(t, err, "got not expected error")
}

func TestCSR2Sha256OfPubKey(t *testing.T) {
	const pemCsr = `
-----BEGIN CERTIFICATE REQUEST-----
MIIEbTCCAlUCAQAwKDELMAkGA1UEBhMCWloxGTAXBgNVBAMMEG9wZW5zc2xyZXFf
dW5lbmMwggIiMA0GCSqGSIb3DQEBAQUAA4ICDwAwggIKAoICAQD2xzcUIELfH2RM
+kpPOX+pBy3c1yN+tOyuo2WhpMKdxe1ICWfZf7IKUKZpUiulA7IvtNsJeeIqiOYm
Ew52F8dXqxaYWIVBXdcnLNAVdsJfd0T+/m7LbWXjGtm2QtxfgnOxLp+mx/Zq87jw
4zpkI04wc7mJ9dSgdhUbkCk/g/kHiPQBWDutJvgdYhfZy/4rr6njt1iBXk3Zc7TP
qD5mwrrIC6SVodTlvXKGtiOlF1eAJ66TBo0cjxiqIOlJatLETmmYGwvYQ/fPmMvJ
2DBpw8nHUwQj7hri1rSK+yvWyCt3lTC+audVGGf/dWXxDkAqJ7fOk6G2IfeaFk7C
HGd2WwblPS/pFjX9kR0xhKX8NJphhlCASqyIklAumyB81lkkjh898eJpG+ENUHw7
vwLA3OPbSYcLKZj2mAVK12x2d4Wxmvp0XDIK0MIj25e/WM2ev7HcZKqVHh8ixk7M
0Y6RHXnJeNtJjAJ56Akmy9SnXmMZoq3pBUpBJpeN8pUbbA7MiI0kKuOuZ7WTinYW
1779CSXE+YvSc+O8i7ItmZ41FU3LxKuhK8LligbK9ZZCSXlA5C1Ez+cDziOuJXDS
rb9hpw1BDlwsV6GgVCokHbCf7HvBKtIOB44cLZLZBBd3UZOREw81+AwEhCuCmsaZ
yeQPlS9FphEvWCM8247AZbSUu3+UOwIDAQABoAAwDQYJKoZIhvcNAQELBQADggIB
ADD0YH1pW2lOyqoT3n3cGeM4iPt6MMtHek4T6+lVImEXzfoioU5GEv+pfZBdG9wA
waQAZ+cb/x7BNStM6ZpdvZYZKX81jHcrx85sprk5oMgcrCTkawVZnvG5SC01FsUD
0BmowXRpEM/5h/wFdpDRfg3lvR65pNtCZadXydCtumQISo7IKbLHxWNF/be07zVy
QCL2c6wR1LHJyfH9GOeCLUyCHifjuOzNdVTvpuqnnHBSK1v86XW/zBHxiPsKPsP1
gLt8u5Da7/gtFZkYHAPDKbkY9wljDMIY7k7BOy0r7wxq2Bx1vyIGdt4RDlDLg3yC
zJp4eiWmLRAJ9xUFR1Zm9uXhJ2MSaSPsOH6ctoK47KP09um8hUUKwXFvmsnoH0St
WsJJtHkvKMxfJe7qQKO6efWqkEZcLQJf0NWeNbFWpMBx5P+b89eKXvl2U3ESp1am
vSXfHySld88mTB3jEucwBtm7NelhVhj3kwxB0VFrsGUnnW6gOFQlM+U/MJH23mjE
oqjJjUQ8ErPfVWZXsV7aBAakFgcyPSefCafENlKfDRI4KVQxWtWXOZeGp0/pzpEw
SrwZoQ0iFZlO0osLRb+A3S6Jwf4Ls55eZU9HgicgWNs3xhyuEgUoJ8QrIyFst1HH
pQ3dQDkR/QCqbIaO7P/1J8YDCkBqCUUV7xWp74dmSzbP
-----END CERTIFICATE REQUEST-----
	`
	csr, err := Pem2CSR([]byte(pemCsr))
	assert.Nil(t, err, "error in Pem2CSR")
	pubkey, err := Any2RsaPublicKey(csr.PublicKey)
	assert.Nil(t, err, "error in Any2RsaPublicKey")
	_, err = x509.MarshalPKIXPublicKey(pubkey) // DER format
	assert.Nil(t, err, "keyDer marshall")
}

// MustMarshalPublicPEMToDER reads a PEM-encoded public key and returns it in DER encoding.
// If an error occurs, it panics.
func mustMarshalPublicPEMToDER(keyPEM string) ([]byte, error) {
	block, _ := pem.Decode([]byte(keyPEM))
	if block == nil {
		return nil, errors.New(CurrentFunctionName() + ":ERROR decoding PEM block")
	}
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, errors.New(CurrentFunctionName() + ":ERROR parsinPkiXPublicKey")
	}

	keyDER, err := x509.MarshalPKIXPublicKey(key)
	if err != nil {
		return nil, errors.New(CurrentFunctionName() + ":ERROR MarshalPKIXPublicKey")

	}
	return keyDER, nil
}

func TestPubKeyDigest(t *testing.T) {
	const publicKey = `
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEA9sc3FCBC3x9kTPpKTzl/
qQct3NcjfrTsrqNloaTCncXtSAln2X+yClCmaVIrpQOyL7TbCXniKojmJhMOdhfH
V6sWmFiFQV3XJyzQFXbCX3dE/v5uy21l4xrZtkLcX4JzsS6fpsf2avO48OM6ZCNO
MHO5ifXUoHYVG5ApP4P5B4j0AVg7rSb4HWIX2cv+K6+p47dYgV5N2XO0z6g+ZsK6
yAuklaHU5b1yhrYjpRdXgCeukwaNHI8YqiDpSWrSxE5pmBsL2EP3z5jLydgwacPJ
x1MEI+4a4ta0ivsr1sgrd5UwvmrnVRhn/3Vl8Q5AKie3zpOhtiH3mhZOwhxndlsG
5T0v6RY1/ZEdMYSl/DSaYYZQgEqsiJJQLpsgfNZZJI4fPfHiaRvhDVB8O78CwNzj
20mHCymY9pgFStdsdneFsZr6dFwyCtDCI9uXv1jNnr+x3GSqlR4fIsZOzNGOkR15
yXjbSYwCeegJJsvUp15jGaKt6QVKQSaXjfKVG2wOzIiNJCrjrme1k4p2Fte+/Qkl
xPmL0nPjvIuyLZmeNRVNy8SroSvC5YoGyvWWQkl5QOQtRM/nA84jriVw0q2/YacN
QQ5cLFehoFQqJB2wn+x7wSrSDgeOHC2S2QQXd1GTkRMPNfgMBIQrgprGmcnkD5Uv
RaYRL1gjPNuOwGW0lLt/lDsCAwEAAQ==
`

	str := "-----BEGIN PUBLIC KEY-----" + // line feeds not required
		strings.Replace(strings.Replace(publicKey, " ", "", -1), "\t", "", -1) + // remove spaces + tabs
		"-----END PUBLIC KEY-----"
	block, _ := pem.Decode([]byte(str))
	assert.NotNil(t, block, "pem.Decode error")

	// fmt.Print(strings.Replace(strings.Replace(publicKey, " ", "", -1), "\t", "", -1))
	key, err := Pem2RsaPublicKey([]byte(str))
	assert.Nil(t, err, "convert to RSA pub key")

	digest, err := RsaPublicKey2Sha256Digest(key)
	assert.Nil(t, err, "digest calculation")

	assert.Equal(t, "1c509c2b33c41cb370ab02d8f8af0ce3fd2f05c5272e5b2b848487b56dfc51fa", digest, "not expected digest")
}

func TestPubKeyFromFileDigest(t *testing.T) {
	key, err := LoadRsaPublicKey("./TestFiles/opensslreq_pkcs8_unenc.pub")
	assert.Nil(t, err, "error loading RSA public key from file")
	digest, err := RsaPublicKey2Sha256Digest(key)
	assert.Nil(t, err, "digest calculation")
	assert.Equal(t, "1c509c2b33c41cb370ab02d8f8af0ce3fd2f05c5272e5b2b848487b56dfc51fa", digest, "not expected digest")
}

func TestPrvKeyFromFileDigest(t *testing.T) {
	key, err := LoadPrivateKey("./TestFiles/opensslreq_pkcs8_unenc.key")
	assert.Nil(t, err, "error loading RSA private key from file")
	digest, err := RsaPrivateKey2Sha256Digest(key)
	assert.Nil(t, err, "digest calculation")
	assert.Equal(t, "1c509c2b33c41cb370ab02d8f8af0ce3fd2f05c5272e5b2b848487b56dfc51fa", digest, "not expected digest")
}

// EOF
