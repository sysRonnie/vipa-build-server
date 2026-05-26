package note



type NoteRow struct {
	ID int `json:"id"`

	UserID string `json:"user_id"`

	ProjectID *int `json:"project_id,omitempty"`
	ProjectName *string `json:"project_name,omitempty"`
	CustomerName *string `json:"customer_name,omitempty"`

	NoteBody string `json:"note_body"`

	PhotoURL *string `json:"photo_url,omitempty"`


	FlagIsDeleted bool `json:"flag_is_deleted"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type NoteRowList struct {
	Notes []NoteRow `json:"notes"`
}

type NoteDeleteRequest struct {
	ID int `json:"id"`
}