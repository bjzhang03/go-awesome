package main

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"github.com/pkg/errors"
	"log"
	"math/big"
	"net"
	"os"
	"time"
)

/*存储的是tls证书和对应的私钥,都被编码成PEM的格式*/
type CertKeyPair struct {
	// Cert is the certificate, PEM encoded
	Cert []byte
	// Key is the key corresponding to the certificate, PEM encoded
	Key []byte

	crypto.Signer
	TLSCert *x509.Certificate
}

func generateStoreDir(path string) string {
	// 定义一个文件
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Get pwd path failed! error := %s", err.Error())
	}
	storepath := pwd + path
	log.Printf("store path := %s", storepath)
	_, err = os.Stat(storepath)
	// 如果目录不存在则创建一个
	if err != nil {
		if !os.IsExist(err) {
			os.MkdirAll(storepath, os.ModePerm)
		}
	}
	return storepath
}

/*生成公钥私钥*/
func newPrivKey() (*ecdsa.PrivateKey, []byte, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	privBytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		return nil, nil, err
	}
	return privateKey, privBytes, nil
}

/*生成证书模板*/
func newCertTemplate() (x509.Certificate, error) {
	sn, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
	if err != nil {
		return x509.Certificate{}, err
	}
	return x509.Certificate{
		Subject:      pkix.Name{SerialNumber: sn.String()},
		NotBefore:    time.Now().Add(time.Hour * (-24)),
		NotAfter:     time.Now().Add(time.Hour * 24),
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		SerialNumber: sn,
	}, nil
}

/*生成证书对*/
func newCertKeyPair(isCA bool, isServer bool, host string, certSigner crypto.Signer, parent *x509.Certificate) (*CertKeyPair, error) {
	privateKey, privBytes, err := newPrivKey()
	if err != nil {
		return nil, err
	}

	template, err := newCertTemplate()
	if err != nil {
		return nil, err
	}

	tenYearsFromNow := time.Now().Add(time.Hour * 24 * 365 * 10)
	if isCA {
		template.NotAfter = tenYearsFromNow
		template.IsCA = true
		template.KeyUsage |= x509.KeyUsageCertSign | x509.KeyUsageCRLSign
		template.ExtKeyUsage = []x509.ExtKeyUsage{
			x509.ExtKeyUsageClientAuth,
			x509.ExtKeyUsageServerAuth,
		}
		template.BasicConstraintsValid = true
	} else {
		template.ExtKeyUsage = []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth}
	}
	if isServer {
		template.NotAfter = tenYearsFromNow
		template.ExtKeyUsage = append(template.ExtKeyUsage, x509.ExtKeyUsageServerAuth)
		if ip := net.ParseIP(host); ip != nil {
			template.IPAddresses = append(template.IPAddresses, ip)
		} else {
			template.DNSNames = append(template.DNSNames, host)
		}
	}
	// If no parent cert, it's a self signed cert
	if parent == nil || certSigner == nil {
		parent = &template
		certSigner = privateKey
	}
	rawBytes, err := x509.CreateCertificate(rand.Reader, &template, parent, &privateKey.PublicKey, certSigner)
	if err != nil {
		return nil, err
	}
	pubKey := encodePEM("CERTIFICATE", rawBytes)
	/*生成文件*/
	certfile, err := os.Create(generateStoreDir("/configs/tmp/x509/") + "cert.pem")
	if err != nil {
		log.Fatalf("Create certfile failed ! error := %s", err.Error())
	}
	defer certfile.Close()
	/*将证书导出*/
	pem.Encode(certfile, &pem.Block{Type: "CERTIFICATE", Bytes: rawBytes})

	block, _ := pem.Decode(pubKey)
	if block == nil { // Never comes unless x509 or pem has bug
		return nil, errors.Errorf("%s: wrong PEM encoding", pubKey)
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}
	privKey := encodePEM("EC PRIVATE KEY", privBytes)
	/*生成文件*/
	keyfile, err := os.Create(generateStoreDir("/configs/tmp/x509/") + "key.pem")
	if err != nil {
		log.Fatalf("Create keyfile failed ! error := %s", err.Error())
	}
	defer keyfile.Close()
	/*将私钥数据导出*/
	pem.Encode(keyfile, &pem.Block{Type: "PRIVATE KEY", Bytes: privBytes})
	return &CertKeyPair{
		Key:     privKey,
		Cert:    pubKey,
		Signer:  privateKey,
		TLSCert: cert,
	}, nil
}

/*将证书转化成PEM格式的数据*/
func encodePEM(keyType string, data []byte) []byte {
	return pem.EncodeToMemory(&pem.Block{Type: keyType, Bytes: data})
}

func main() {
	newCertKeyPair(true, false, "", nil, nil)
}
