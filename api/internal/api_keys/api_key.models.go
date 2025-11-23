package apikeys

type APIKey struct {
	Name          string   `json:"name"`
	Slug          string   `json:"slug"`
	Prefix        string   `json:"prefix"`
	PublicID      string   `json:"public_id"`
	Scope         []string `json:"scope"`
	EncryptionIv  string   `json:"encryption_iv"`
	EncryptedText string   `json:"encrypted_text"`
	Algorithm     string   `json:"algorithm"`
	Rotations     int32    `json:"rotations"`
	LastRotatedAt string   `json:"last_rotated_at"`
}
