package apikeys

import (
	"crypto/aes"
	"errors"
	"fmt"
	"log"

	"github.com/gosimple/slug"
	"github.com/qryne/api/lib"
	"github.com/qryne/api/utility"
)

type APIKeyServices struct {
	APIKeyRepo IAPIKeyRepository
}

func (service *APIKeyServices) GenerateAPIKey(name, prefix, setup_id string, scope []string) (APIKey, error) {
	if len(name) == 0 {
		return APIKey{}, errors.New("API key name cannot be empty")
	}

	slugified_name := slug.Make(name)

	key, err := utility.GetENVString("SETUP_API_SECRET")

	if err != nil {
		return APIKey{}, err
	}

	iv, err := utility.GenerateIV()
	if err != nil {
		return APIKey{}, err
	}

	plain_text, err := utility.RandomString(16)
	if err != nil {
		return APIKey{}, err
	}

	aes_cbc := lib.AESCBC{}

	cipherText := fmt.Sprintf("%v", aes_cbc.Ase256Encode(plain_text, key, string(iv), aes.BlockSize))

	public_id, err := utility.RandomString(8)

	record, err := service.APIKeyRepo.CreateAPIKey(name, slugified_name, prefix, public_id, string(iv), cipherText, "AES256", setup_id, scope)
	if err != nil {
		return APIKey{}, err
	}

	log.Fatal(record)

	return APIKey{
		Name:          record.Name,
		Slug:          record.Slug,
		Prefix:        record.Prefix,
		PublicID:      record.PublicID,
		Scope:         record.Scope,
		EncryptionIv:  record.EncryptionIv,
		EncryptedText: record.EncryptedText,
		Algorithm:     record.Algorithm,
		Rotations:     record.Rotations,
		LastRotatedAt: record.LastRotatedAt.Time.String(),
	}, nil
}
