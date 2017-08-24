package validator

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/ed25519"

	"github.com/stratumn/sdk/cs"
	"github.com/stratumn/sdk/store"

	log "github.com/Sirupsen/logrus"
	"github.com/xeipuuv/gojsonschema"
)

type schemaValidator struct {
	Type    string
	Options map[string]interface{}
	Schema  *gojsonschema.Schema
}

type signatureJSON struct {
	Type      string `json:"type"`
	PublicKey string `json:"publicKey"`
	Signature string `json:"signature"`
}

func newSchemaValidator(segmentType string, data []byte) (*schemaValidator, error) {
	schema, err := gojsonschema.NewSchema(gojsonschema.NewBytesLoader(data))

	if err != nil {
		return nil, err
	}

	return &schemaValidator{Type: segmentType, Schema: schema}, nil
}

func (sv schemaValidator) Filter(_ store.Reader, segment *cs.Segment) bool {
	// TODO: standardise action as string
	segmentAction, ok := segment.Link.Meta["action"].(string)
	if !ok {
		log.Debug("No action found in segment %v", segment)
		return false
	}

	if segmentAction != sv.Type {
		return false
	}

	return true
}

func (sv schemaValidator) Validate(_ store.Reader, segment *cs.Segment) error {
	segmentBytes, err := json.Marshal(segment.Link.State)
	if err != nil {
		return err
	}

	segmentData := gojsonschema.NewBytesLoader(segmentBytes)

	result, err := sv.Schema.Validate(segmentData)
	if err != nil {
		return err
	}

	if !result.Valid() {
		return fmt.Errorf("segment validation failed: %s", result.Errors())
	}

	if sv.Options != nil {
		v, ok := sv.Options["checkSignatures"].(bool)
		if !ok {
			return fmt.Errorf("`checkSignatures` is not boolean")
		}
		if v {
			signatures, ok := segment.Link.Meta["signatures"].([]signatureJSON)
			if !ok {
				return fmt.Errorf("signatures array is missing")
			}
			for _, signature := range signatures {
				var pubKeyBytes, signatureBytes []byte
				var err error

				if pubKeyBytes, err = decodeHex(signature.PublicKey); err != nil {
					return err
				}

				if signatureBytes, err = decodeHex(signature.Signature); err != nil {
					return err
				}

				switch signature.Type {
				case "EdDSA":
					if len(pubKeyBytes) != ed25519.PublicKeySize {
						return fmt.Errorf("public key size %s is incorrect", signature.PublicKey)
					}

					if len(signatureBytes) != ed25519.SignatureSize {
						return fmt.Errorf("signature size %s is incorrect", signature.PublicKey)
					}

					if !ed25519.Verify(pubKeyBytes, segmentBytes, signatureBytes) {
						return fmt.Errorf("signature verification failed")
					}
				// TODO: RSA, DSA, ECDSA
				default:
					return fmt.Errorf("Unknown/missing signature type %s", signature.Type)
				}

			}

		}
	}

	return nil
}

func decodeHex(input string) ([]byte, error) {
	src := []byte(input)

	dst := make([]byte, hex.DecodedLen(len(src)))
	_, err := hex.Decode(dst, src)
	if err != nil {
		return nil, err
	}

	return dst, nil
}
