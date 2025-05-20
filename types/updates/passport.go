package updates

type PassportData struct {
	Data        []*EncryptedPassportElement `json:"data"`
	Credentials *EncryptedCredentials       `json:"credentials"`
}

type EncryptedPassportElement struct {
	Type        string          `json:"type"`
	Data        string          `json:"data,omitempty"`
	PhoneNumber string          `json:"phone_number,omitempty"`
	Email       string          `json:"email,omitempty"`
	Files       []*PassportFile `json:"files,omitempty"`
	FrontSide   *PassportFile   `json:"front_side,omitempty"`
	ReverseSide *PassportFile   `json:"reverse_side,omitempty"`
	Selfie      *PassportFile   `json:"selfie,omitempty"`
	Translation []*PassportFile `json:"translation,omitempty"`
	Hash        string          `json:"hash,omitempty"`
}

type EncryptedCredentials struct {
	Data   string `json:"data"`
	Hash   string `json:"hash"`
	Secret string `json:"secret"`
}

type PassportFile struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	FileSize     int    `json:"file_size"`
	FileDate     int    `json:"file_date"`
}
