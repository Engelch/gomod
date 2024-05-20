package debugerrorce

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
)

// =====================================================================
// sha256, base64 routines
// =====================================================================

func Str2sha256(str string) []byte {
	byteaFromStr := []byte(str) // convert string to []byte aka bytea
	msgHash := sha256.New()
	_, _ = msgHash.Write(byteaFromStr) // todo no error handling, but error is very unlike
	sha256sum := msgHash.Sum(nil)
	CondDebug(fmt.Sprintf("sha256 of string as hex    is %x\n", sha256sum))
	CondDebug(fmt.Sprintf("sha256 of string as binary is %v\n", sha256sum))
	return sha256sum
}

// create the sha256 sum of str, don't treat it as hex but as binary and
// take the base64 of it
func Bytea2b64(ba []byte) string {
	eb := make([]byte, base64.StdEncoding.EncodedLen(len(ba)))
	base64.StdEncoding.Encode(eb, ba)
	CondDebug(fmt.Sprintf("b64 as string is   %s\n", string(eb)))
	CondDebug(fmt.Sprintf("b64 as byte-seq is %x\n", eb))
	return string(eb)
}

func String2md5(str string) string {
	data := []byte(str)
	return fmt.Sprintf("%x", md5.Sum(data))
}

// =====================================================================
// RSA routines
// =====================================================================

// const bitSize = 4096 // RSA keysize

const RSA4096 = 4096
const RSA2048 = 2048

func RsaPrivateKey2Sha256Digest(prvKey *rsa.PrivateKey) (string, error) {
	if prvKey == nil {
		return "", errors.New(CurrentFunctionName() + ":ERROR:supplied key is nil")
	}
	return fmt.Sprintf("%x", string(Bytes2sha256([]byte(fmt.Sprintf("%X,%X", prvKey.PublicKey.E, prvKey.PublicKey.N))))), nil
}

func RsaPublicKey2Sha256Digest(pubKey *rsa.PublicKey) (string, error) {
	if pubKey == nil {
		return "", errors.New(CurrentFunctionName() + ":ERROR:supplied key is nil")
	}
	return fmt.Sprintf("%x", string(Bytes2sha256([]byte(fmt.Sprintf("%X,%X", pubKey.E, pubKey.N))))), nil
}

func Pem2CSR(bytes []byte) (*x509.CertificateRequest, error) {
	block, _ := pem.Decode(bytes)
	if block == nil {
		return nil, errors.New("ERROR decoding PEM block")
	}
	if block.Type != "CERTIFICATE REQUEST" {
		return nil, errors.New("Could not detect certificate request type from PEM block, found:" + block.Type)
	}
	csrBytes, err := x509.ParseCertificateRequest(block.Bytes)
	if err != nil {
		return nil, errors.New("ERROR:CSR could not be read:" + err.Error())
	}
	return csrBytes, nil
}

// TODO: support for at least the additional algorithm Ed25519 https://golang.google.cn/pkg/crypto/x509/#PublicKeyAlgorithm
// Any2RsaPublicKey is required to read public keys from *x509.CertificateRequest
func Any2RsaPublicKey(data any) (*rsa.PublicKey, error) {
	switch data.(type) {
	case *rsa.PublicKey:
		CondDebugln("Any2RsaPublicKey, ok")
		return data.(*rsa.PublicKey), nil
	default:
		CondDebugln(CurrentFunctionName() + ", ERROR:" + fmt.Sprint(reflect.TypeOf(data)))
		return nil, errors.New("Unsupported public key type:" + fmt.Sprint(reflect.TypeOf(data)))
	}
}

// Bytes2sha256 (ex: Sha256bytes2bytes) converts a byte sequence into a SHA-256-based digest of it.
// The output for this application is the same on the commadn line with:
// curl -q localhost:8888 | jq -c .Data | tr -d '\n' | shasum -a256
// The added newline must be removed. Alternatively, gnu-sed can be used instad of tr:
// gsed -Ez 's/\n$//'
// The complete JSON return structure only consists of US-ASCII characters. So potential
// different escaping for special characters do not have to be considered.
func Bytes2sha256(bytes []byte) []byte {
	//return fmt.Sprintf("%x", sha256.Sum256(bytes)) returning type array [32]byte which must usually be converted
	msgHash := sha256.New()
	_, _ = msgHash.Write(bytes) // todo no error handling, but error is very unlike
	return msgHash.Sum(nil)
}

// SignPSSByteArray returns a signature for the given digest or returns an error
func SignPSSByteArray(key *rsa.PrivateKey, digest []byte) ([]byte, error) {
	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto
	if key == nil { // no signing
		return nil, nil
	}
	signature, err := rsa.SignPSS(rand.Reader, key, crypto.SHA256, digest, &opts)
	if err != nil {
		return nil, errors.New(CurrentFunctionName() + ":" + err.Error())
	}
	return signature, nil
}

// SignPSSByteArray2Base64 returns the signature as a base64-encoded string.
func SignPSSByteArray2Base64(key *rsa.PrivateKey, digest []byte) (string, error) {
	sig, err := SignPSSByteArray(key, digest)
	if err != nil {
		return "", errors.New(CurrentFunctionName() + ":" + err.Error())
	}
	return base64.StdEncoding.EncodeToString(sig), nil
}

// VerifyPSSByteArray verifies a digital signature (digest). If no error is returned,
// then the verification was successful. Furthermore, it recalculates the digest of the
// message. It should result in the same digest as the digitally signed one.
func VerifyPSSByteArray(key *rsa.PublicKey, digest []byte, msg []byte) error {
	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto
	if key == nil {
		return errors.New(CurrentFunctionName() + ":Error, public key is nil")
	}
	if digest == nil {
		return errors.New(CurrentFunctionName() + ":Error, digest is nil")
	}
	plaintestDigest := Bytes2sha256(msg)
	CondDebugln(CurrentFunctionName() + ", recalculated digest for msg: " + fmt.Sprintf("%x", plaintestDigest))
	return rsa.VerifyPSS(key, crypto.SHA256, plaintestDigest, digest, &opts)
}

// VerifyPSSBase64String accepts a base64 encoded string as the signature.
// It decodes the signature and calls VerifyByteArray.
func VerifyPSSBase64String(key *rsa.PublicKey, b64 string, msg string) error {
	signatureByte, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return errors.New(CurrentFunctionName() + ":Error, decoding base64 string")
	}
	return VerifyPSSByteArray(key, signatureByte, []byte(msg))
}

// Sign115ByteArray returns a signature for the given digest or returns an error
func Sign115ByteArray(key *rsa.PrivateKey, digest []byte) ([]byte, error) {
	//var opts rsa.PSSOptions
	//opts.SaltLength = rsa.PSSSaltLengthAuto
	if key == nil { // no signing
		return nil, nil
	}
	signature, err := rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA256, digest)
	if err != nil {
		return nil, errors.New(CurrentFunctionName() + ":" + err.Error())
	}
	return signature, nil
}

// Sign115ByteArray2Base64 signs a byte array by calling SignByteArray but returns the signature as a base64-encoded string.
func Sign115ByteArray2Base64(key *rsa.PrivateKey, digest []byte) (string, error) {
	sig, err := Sign115ByteArray(key, digest)
	if err != nil {
		return "", errors.New(CurrentFunctionName() + ":" + err.Error())
	}
	return base64.StdEncoding.EncodeToString(sig), nil
}

// Verify115ByteArray verifies a digital signature (digest). If no error is returned,
// then the verification was successful. Furthermore, it recalculates the digest of the
// message. It should result in the same digest as the digitally signed one.
func Verify115ByteArray(key *rsa.PublicKey, digest []byte, msg []byte) error {
	if key == nil {
		return errors.New(CurrentFunctionName() + ":Error, public key is nil")
	}
	if digest == nil {
		return errors.New(CurrentFunctionName() + ":Error, digest is nil")
	}
	plaintestDigest := Bytes2sha256(msg)
	CondDebugln(CurrentFunctionName() + ", recalculated digest for msg: " + fmt.Sprintf("%x", plaintestDigest))
	return rsa.VerifyPKCS1v15(key, crypto.SHA256, plaintestDigest, digest)
}

// Verify115Base64String accepts a base64 encoded string as the signature.
// It decodes the signature and calls VerifyByteArray.
func Verify115Base64String(key *rsa.PublicKey, b64 string, msg string) error {
	signatureByte, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return errors.New(CurrentFunctionName() + ":Error, decoding base64 string")
	}
	return Verify115ByteArray(key, signatureByte, []byte(msg))
}

// =======================================================================================
// = Key Loading and Signing

// Pem2RsaPrivateKey load a PEM-encoded RSA private key from a buffer. The function does not try
// to read multiple keys from the byte array. Only the first PEM block is processed.
func Pem2RsaPrivateKey(der []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(der)
	if block == nil || !strings.Contains(block.Type, "PRIVATE KEY") {
		if block == nil {
			CondDebugln(CurrentFunctionName() + ":returned decoded block is nil\n")
		}
		CondDebugln(CurrentFunctionName()+":block type is:", block.Type)
		return nil, errors.New(CurrentFunctionName() + ":failed to decode PEM block containing private key")
	}
	prv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		key, err := x509.ParsePKCS8PrivateKey(block.Bytes) // now let's check pkcs8
		if err != nil {
			return nil, errors.New(CurrentFunctionName() + ":failed to parse PEM block as PKCS[18]:" + err.Error())
		}
		switch key.(type) {
		case *rsa.PrivateKey:
			prv = key.(*rsa.PrivateKey)
		default:
			return nil, errors.New(CurrentFunctionName() + ":failed to convert read key of type")
		}
	}
	return prv, nil
}

// LoadPrivateKey load a PEM-encoded RSA private key from a file
func LoadPrivateKey(filename string) (*rsa.PrivateKey, error) {
	buf, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.New(CurrentFunctionName() + ":reading file:" + err.Error())
	}
	return Pem2RsaPrivateKey(buf)
}

// Pem2RsaPublicKey load a PEM-encoded RSA public key from a buffer. The function does not try
// to read multiple keys from the byte array. Only the first PEM block is processed.
func Pem2RsaPublicKey(der []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(der)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, errors.New(CurrentFunctionName() + ":failed to decode PEM block containing public key")
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, errors.New(CurrentFunctionName() + ":failed to parse PEM block:" + err.Error())
	}
	switch pub.(type) {
	case *rsa.PublicKey:
		return pub.(*rsa.PublicKey), nil
	default:
		return nil, errors.New(CurrentFunctionName() + ":Unsupported public key type, not RSA.")
	}
}

// LoadPublicKey load a PEM-encoded RSA public key from a file
func LoadRsaPublicKey(filename string) (*rsa.PublicKey, error) {
	buf, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.New(CurrentFunctionName() + ":reading file:" + err.Error())
	}
	return Pem2RsaPublicKey(buf)
}

func DebugRsaPublicKey(pubKey *rsa.PublicKey) {
	if pubKey == nil {
		fmt.Fprintf(os.Stderr, "ERROR: no key supplied to "+CurrentFunctionName())
		return
	}
	fmt.Println("Exponent", pubKey.E)
	fmt.Println("Modulus", pubKey.N)
}

// TODO VerifySignature
// TODO EncryptAES256
// TODO DecryptAES256

// =======================================================================================
// = Keypair Generation

// WriteRsaPrivateKey converts the key to PEM format and writes them to a file.
func WriteRsaPrivateKey(file *os.File, privKey *rsa.PrivateKey) error {
	var privateKey = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privKey),
	}
	if err := pem.Encode(file, privateKey); err != nil {
		return errors.New(CurrentFunctionName() + ":pem encode+writeFile:" + err.Error())
	}
	if err := os.Chmod(file.Name(), 0600); err != nil {
		return errors.New(CurrentFunctionName() + ":chmod:" + err.Error())
	}
	return nil
}

// WriteRsaPublicKey converts the public key to PEM format and writes them to the file.
func WriteRsaPublicKey(file *os.File, pubKey *rsa.PublicKey) error {
	asn1Bytes, err := x509.MarshalPKIXPublicKey(pubKey)
	if err != nil {
		return errors.New(CurrentFunctionName() + ":1:" + err.Error())
	}
	CondDebugln(fmt.Sprintf("Length of Public Key: %d", len(asn1Bytes)))
	var pemkey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: asn1Bytes,
	}
	if err := pem.Encode(file, pemkey); err != nil {
		return errors.New(CurrentFunctionName() + ":2:" + err.Error())
	}
	return nil
}

// createRSAKeyPair2 creates the keypair and calls the functions to write the keys to the files
func createRSAKeyPair2(privKeyFile *os.File, pubKeyFile *os.File, bitSize int) error {
	privateKey, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		return errors.New(CurrentFunctionName() + "key creation:" + err.Error())
	}
	if err := WriteRsaPrivateKey(privKeyFile, privateKey); err != nil {
		return errors.New(CurrentFunctionName() + "private key writing:" + err.Error())
	}
	if err := WriteRsaPublicKey(pubKeyFile, &privateKey.PublicKey); err != nil {
		return errors.New(CurrentFunctionName() + "public key writing:" + err.Error())
	}
	return nil
}

// CreateRSAKeyPair2File checks if the 2 required files do not exist and can be created sucessfully. Then,
// it transfers control to createKeyPairError2.
func CreateRSAKeyPair2File(outfileName string, bitSize int) error {
	var privKeyFile *os.File
	var pubKeyFile *os.File
	var err error

	const publicKeyFileSuffix = ".pub"

	if _, err = os.Stat(outfileName); err == nil {
		return errors.New("Private key file " + outfileName + " already exists.")
	}
	if _, err = os.Stat(outfileName + publicKeyFileSuffix); err == nil {
		return errors.New("Public key file " + outfileName + " already exists.")
	}
	if privKeyFile, err = os.Create(outfileName); err != nil {
		return errors.New("Error creating private key file " + outfileName + ":" + err.Error())
	}
	if pubKeyFile, err = os.Create(outfileName + publicKeyFileSuffix); err != nil {
		return errors.New("Error creating private key file " + outfileName + ":" + err.Error())
	}
	defer privKeyFile.Close()
	defer pubKeyFile.Close()
	return createRSAKeyPair2(privKeyFile, pubKeyFile, bitSize)
}

// CreateRSAKeyPair creates an RSA bitsize-bit key-pair. This function makes only partly sense,
// as the private key always contains the public key.
func CreateRSAKeyPair(bitSize int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		return nil, nil, errors.New(CurrentFunctionName() + "key creation:" + err.Error())
	}
	return privateKey, &privateKey.PublicKey, nil
}

// EOF
