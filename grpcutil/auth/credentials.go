package auth

import (
	"crypto/x509"
	"errors"
	"fmt"
	"net/url"

	"encoding/base64"
	"encoding/json"
	"encoding/pem"

	cleanhttp "github.com/hashicorp/go-cleanhttp"
	"google.golang.org/api/googleapi"
)

const (
	serviceAccountPublicKeyURLTemplate = "https://iam.videocoin.net/service_accounts/v1/metadata/x509/%s?alt=json"
)

// PublicKey returns a public key from a Google PEM key file (type TYPE_X509_PEM_FILE).
func PublicKey(pemString string) (interface{}, error) {
	// Attempt to base64 decode
	pemBytes := []byte(pemString)
	if b64decoded, err := base64.StdEncoding.DecodeString(pemString); err == nil {
		pemBytes = b64decoded
	}

	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, errors.New("Unable to find pem block in key")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}

	return cert.PublicKey, nil
}

// ServiceAccountPublicKey returns the service account public key for a given key id.
func ServiceAccountPublicKey(pubKeyURLTemplate string, serviceAccount string, keyID string) (interface{}, error) {
	keyURL := fmt.Sprintf(pubKeyURLTemplate, url.PathEscape(serviceAccount))
	res, err := cleanhttp.DefaultClient().Get(keyURL)
	if err != nil {
		return nil, err
	}
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}

	jwks := map[string]interface{}{}
	if err := json.NewDecoder(res.Body).Decode(&jwks); err != nil {
		return nil, fmt.Errorf("unable to decode JSON response: %v", err)
	}
	kRaw, ok := jwks[keyID]
	if !ok {
		return nil, fmt.Errorf("service account %q key %q not found at GET %q", keyID, serviceAccount, keyURL)
	}

	kStr, ok := kRaw.(string)
	if !ok {
		return nil, fmt.Errorf("unexpected error - decoded JSON key value %v is not string", kRaw)
	}
	return PublicKey(kStr)
}
