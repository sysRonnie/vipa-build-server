package note

const queryNoteList = `
SELECT
	n.id,
	u.email,
	n.project_id,
	p.project_name,
	CONCAT(c.customer_name_first, ' ', c.customer_name_last) AS customer_name,
	n.note_body,
	n.note_photo_url,
	n.flag_is_deleted,
	n.created_at,
	n.updated_at
FROM user_project_note n
JOIN user_auth u
	ON u.uuid = n.user_id
LEFT JOIN master_project_list p
	ON p.id = n.project_id
LEFT JOIN master_customer_list c 
	ON c.id = p.customer_id
WHERE
	u.email = $1
	AND n.flag_is_deleted = $2
ORDER BY
	n.updated_at DESC;
`

const queryNoteByID = `
SELECT
	n.id,
	u.email,
	n.project_id,
	p.project_name,
	CONCAT(c.customer_name_first, ' ', c.customer_name_last) AS customer_name,
	n.note_body,
	n.note_photo_url,
	n.flag_is_deleted,
	n.created_at,
	n.updated_at
FROM user_project_note n
JOIN user_auth u
	ON u.uuid = n.user_id
LEFT JOIN master_project_list p
	ON p.id = n.project_id
LEFT JOIN master_customer_list c
	ON c.id = p.customer_id
WHERE
	u.email = $1
	AND n.id = $2
LIMIT 1;
`

const queryInsertNote = `
INSERT INTO user_project_note (
	user_id,
	project_id,
	note_body,
	note_photo_url
)
SELECT
	u.uuid,
	p.id,
	$2,
	$3
FROM user_auth u
LEFT JOIN master_project_list p
	ON p.project_name = $4
WHERE
	u.email = $1;
`

const queryUpdateNote = `
UPDATE user_project_note n
SET
	project_id = p.id,
	note_body = $3,
	note_photo_url = $4,
	flag_is_deleted = FALSE,
	updated_at = NOW()
FROM master_project_list p
JOIN user_auth u
	ON u.email = $1
WHERE
	n.id = $2
	AND n.user_id = u.uuid
	AND (
		p.project_name = $5
		OR $5 = ''
	);
`

const queryRemoveNote = `
UPDATE user_project_note n
SET
	flag_is_deleted = TRUE,
	updated_at = NOW()
FROM user_auth u
WHERE
	n.id = $2
	AND n.user_id = u.uuid
	AND u.email = $1;
`

const queryEraseNote = `
DELETE FROM user_project_note n
USING user_auth u
WHERE
	n.id = $2
	AND n.user_id = u.uuid
	AND u.email = $1;
`