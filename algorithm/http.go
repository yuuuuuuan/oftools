package algorithm

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"time"
)

func generateSelfSignedCert() (cert []byte, key []byte, err error) {
	// 生成私钥
	priv, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate private key: %w", err)
	}

	// 创建自签名证书模板
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization: []string{"SelfSigned"},
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(365 * 24 * time.Hour), // 有效期1年

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	// 使用自签名方式创建证书
	certBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create certificate: %w", err)
	}

	// 编码证书和私钥为 PEM 格式
	certPem := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certBytes})
	keyBytes, err := x509.MarshalECPrivateKey(priv)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal private key: %w", err)
	}
	keyPem := pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: keyBytes})

	return certPem, keyPem, nil
}

func Http(port string) error {
	// 动态生成自签名证书
	certPem, keyPem, err := generateSelfSignedCert()
	if err != nil {
		fmt.Printf("Error generating certificate: %v\n", err)
		return err
	}

	// 将证书和密钥写入临时文件
	certFile, err := os.CreateTemp("", "cert.pem")
	if err != nil {
		fmt.Printf("Error creating temp cert file: %v\n", err)
		return err
	}
	defer os.Remove(certFile.Name())
	certFile.Write(certPem)
	certFile.Close()

	keyFile, err := os.CreateTemp("", "key.pem")
	if err != nil {
		fmt.Printf("Error creating temp key file: %v\n", err)
		return err
	}
	defer os.Remove(keyFile.Name())
	keyFile.Write(keyPem)
	keyFile.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Secure HTTPS Request received!"))
	})

	fmt.Printf("Listening on https://localhost:%s", port)
	portstring := ":" + port
	err = http.ListenAndServeTLS(portstring, certFile.Name(), keyFile.Name(), nil)
	if err != nil {
		fmt.Printf("Error starting HTTPS server: %v\n", err)
	}
	return nil
}
