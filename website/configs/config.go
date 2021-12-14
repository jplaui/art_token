package main

import "os"

type Config struct {
	StaticAssetsDir   string
	TlsCertPath       string
	TlsKeyPath        string
	HttpAddr          string
	HttpPort          string
	HttpsAddr         string
	HttpsPort         string
	HttpSessionSecret string
}

var (
	config *Config
)

func init() {

	StaticAssetsDir := os.Getenv("STATIC_ASSETS_DIR")
	if StaticAssetsDir == "" {
		StaticAssetsDir = "./frontend/build"
	}

	TlsCertPath := os.Getenv("TLS_CERT_PATH")
	if TlsCertPath == "" {
		TlsCertPath = "./configs/ssl/test-cert.pem"
	}

	TlsKeyPath := os.Getenv("TLS_KEY_PATH")
	if TlsKeyPath == "" {
		TlsKeyPath = "./configs/ssl/test-privkey.pem"
	}

	HttpAddr := os.Getenv("HTTP_ADDR")
	if HttpAddr == "" {
		HttpAddr = "0.0.0.0"
	}

	HttpPort := os.Getenv("HTTP_PORT")
	if HttpPort == "" {
		HttpPort = "6080"
	}

	HttpsAddr := os.Getenv("HTTPS_ADDR")
	if HttpsAddr == "" {
		HttpsAddr = "0.0.0.0"
	}

	HttpsPort := os.Getenv("HTTPS_PORT")
	if HttpsPort == "" {
		HttpsPort = "6443"
	}

	HttpSessionSecret := os.Getenv("HTTP_SESSION_SECRET")
	if HttpSessionSecret == "" {
		HttpSessionSecret = "41ed725e56ee7fcc43da77f14d6e0ed3c2e570378e73af98241dbf84ea8fb882"
	}

	config = &Config{
		StaticAssetsDir:   StaticAssetsDir,
		TlsCertPath:       TlsCertPath,
		TlsKeyPath:        TlsKeyPath,
		HttpAddr:          HttpAddr,
		HttpPort:          HttpPort,
		HttpsAddr:         HttpsAddr,
		HttpsPort:         HttpsPort,
		HttpSessionSecret: HttpSessionSecret,
	}
}

func GetConfig() *Config {
	return config
}
