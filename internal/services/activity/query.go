package activity



const baseActivityListQuery = `
SELECT 
	a.id,
	a.user_id,
	a.project_id,
	p.project_name AS project_name,
	
	a.activity_type,
	a.activity_title,
	a.activity_body,
	a.amount,
	TO_CHAR(a.activity_date, 'YYYY-MM-DD') AS activity_date,
	
	a.event_category_id,
	CASE WHEN ec.event_category_child IS NULL THEN ec.event_category_parent
	ELSE CONCAT(ec.event_category_parent, ' (', ec.event_category_child, ')') END AS event_category_name,
	
	a.cost_category_id,
	CASE WHEN cc.cost_category_child IS NULL THEN cc.cost_category_parent 
	ELSE CONCAT(cc.cost_category_parent, ' (', cc.cost_category_child, ')') END AS cost_category_name,
	
	a.vendor_id,
	v.vendor_name AS vendor_name,
	
	a.photo_url,
	a.photo_thumbnail_url,
	
	a.flag_is_completed,
	a.flag_is_deleted,
	
	TO_CHAR(a.created_at, 'YYYY-MM-DD HH24:MI:SS') AS created_at,
	TO_CHAR(a.updated_at, 'YYYY-MM-DD HH24:MI:SS') AS updated_at
FROM user_project_activity a
LEFT JOIN master_project_list p ON a.project_id = p.id
LEFT JOIN master_event_category ec ON a.event_category_id = ec.id
LEFT JOIN master_cost_category cc ON a.cost_category_id = cc.id
LEFT JOIN master_vendor_list v ON a.vendor_id = v.id
WHERE 1=1
`


func BuildBaseActivityListQueryWithDeletedFilter(isDeleted bool) string {
	if isDeleted {
		filterOptions := "AND a.flag_is_deleted = true ORDER BY a.created_at DESC, a.activity_date DESC"
		return baseActivityListQuery + filterOptions
	} else {
		filterOptions := "AND a.flag_is_deleted = false ORDER BY a.created_at DESC, a.activity_date DESC"
		return baseActivityListQuery + filterOptions
	}
}


const baseActivityInsertExpense = `
INSERT INTO USER_PROJECT_ACTIVITY (
	USER_ID,
	PROJECT_ID,
	ACTIVITY_TYPE,
	ACTIVITY_TITLE,
	ACTIVITY_BODY,
	AMOUNT,
	ACTIVITY_DATE,
	COST_CATEGORY_ID,
	VENDOR_ID,
	PHOTO_URL
) VALUES (
	(SELECT UUID FROM USER_AUTH WHERE EMAIL = $1 LIMIT 1),
	(SELECT ID FROM MASTER_PROJECT_LIST WHERE PROJECT_NAME = $2 AND FLAG_IS_DELETED = FALSE LIMIT 1),
	'expense',
	$3,
	$4,
	$5,
	$6,
	(SELECT ID FROM MASTER_COST_CATEGORY WHERE CONCAT(COST_CATEGORY_PARENT, ' (', COST_CATEGORY_CHILD, ')') = $7 AND FLAG_IS_DELETED = FALSE LIMIT 1),
	(SELECT ID FROM MASTER_VENDOR_LIST WHERE VENDOR_NAME = $8 AND FLAG_IS_DELETED = FALSE LIMIT 1),
	$9
)
`

const baseActivityListByIdQuery = `
SELECT 
	a.id,
	a.user_id,
	a.project_id,
	p.project_name AS project_name,
	
	a.activity_type,
	a.activity_title,
	a.activity_body,
	a.amount,
	TO_CHAR(a.activity_date, 'YYYY-MM-DD') AS activity_date,
	
	a.event_category_id,
	CASE WHEN ec.event_category_child IS NULL THEN ec.event_category_parent
	ELSE CONCAT(ec.event_category_parent, ' (', ec.event_category_child, ')') END AS event_category_name,
	
	a.cost_category_id,
	CASE WHEN cc.cost_category_child IS NULL THEN cc.cost_category_parent 
	ELSE CONCAT(cc.cost_category_parent, ' (', cc.cost_category_child, ')') END AS cost_category_name,
	
	a.vendor_id,
	v.vendor_name AS vendor_name,
	
	a.photo_url,
	a.photo_thumbnail_url,
	
	a.flag_is_completed,
	a.flag_is_deleted,
	
	TO_CHAR(a.created_at, 'YYYY-MM-DD HH24:MI:SS') AS created_at,
	TO_CHAR(a.updated_at, 'YYYY-MM-DD HH24:MI:SS') AS updated_at
FROM user_project_activity a
LEFT JOIN master_project_list p ON a.project_id = p.id
LEFT JOIN master_event_category ec ON a.event_category_id = ec.id
LEFT JOIN master_cost_category cc ON a.cost_category_id = cc.id
LEFT JOIN master_vendor_list v ON a.vendor_id = v.id
WHERE 1=1 AND a.id = $1
ORDER BY a.activity_date DESC, a.created_at DESC;
`

const baseActivityTypeQuery = `
SELECT activity_type
FROM user_project_activity
WHERE id = $1;
`


const baseActivityByIdExpenseQuery = `
SELECT 
	a.id,
	p.project_name AS project_name,
	v.vendor_name AS vendor_name,
	cc.cost_category_parent AS cost_category_name,
	a.activity_title AS expense_title,
	a.activity_body AS expense_desc,
	a.amount,
	TO_CHAR(a.activity_date, 'YYYY-MM-DD') AS expense_date,
	a.photo_url
FROM user_project_activity a
LEFT JOIN master_project_list p ON a.project_id = p.id
LEFT JOIN master_cost_category cc ON a.cost_category_id = cc.id
LEFT JOIN master_vendor_list v ON a.vendor_id = v.id
WHERE a.flag_is_deleted = false AND a.id = $1 AND a.activity_type = 'expense'
ORDER BY a.activity_date DESC, a.created_at DESC;
`


const baseActivityUpdateExpense = `
UPDATE user_project_activity
SET
	project_id = (
		SELECT ID FROM MASTER_PROJECT_LIST 
		WHERE PROJECT_NAME = $2 
		AND FLAG_IS_DELETED = FALSE 
		LIMIT 1
	),
	activity_title = $3,
	activity_body = $4,
	amount = $5,
	activity_date = $6,
	cost_category_id = (
		SELECT ID FROM MASTER_COST_CATEGORY 
		WHERE 
			CASE WHEN COST_CATEGORY_CHILD IS NULL THEN COST_CATEGORY_PARENT = $7
			ELSE CONCAT(COST_CATEGORY_PARENT, ' (', COST_CATEGORY_CHILD, ')') = $7 END
			 AND FLAG_IS_DELETED = FALSE 
		LIMIT 1
	),
	vendor_id = (
		SELECT ID FROM MASTER_VENDOR_LIST 
		WHERE VENDOR_NAME = $8 
		AND FLAG_IS_DELETED = FALSE 
		LIMIT 1
	),
	photo_url = $9,
	updated_at = NOW()
WHERE id = $10 
AND activity_type = 'expense' 
AND user_id = (
	SELECT UUID FROM USER_AUTH 
	WHERE EMAIL = $1 
	LIMIT 1
);
`

const baseActivityDeleteSoft = `
UPDATE user_project_activity
SET flag_is_deleted = true, updated_at = NOW()
WHERE id = $1;
`
const baseActivityDeleteErase = `
DELETE FROM user_project_activity
WHERE id = $1;
`

const baseActivityInsertNewActivity = `
INSERT INTO USER_PROJECT_ACTIVITY (
	USER_ID,
	PROJECT_ID,
	ACTIVITY_TYPE,
	ACTIVITY_TITLE,
	ACTIVITY_BODY,
	AMOUNT,
	ACTIVITY_DATE,
	EVENT_CATEGORY_ID,
	COST_CATEGORY_ID,
	VENDOR_ID,
	PHOTO_URL
) VALUES (
	(SELECT UUID FROM USER_AUTH WHERE EMAIL = $1 LIMIT 1),
	(SELECT ID FROM MASTER_PROJECT_LIST WHERE PROJECT_NAME = $2 AND FLAG_IS_DELETED = FALSE LIMIT 1),
	$3,
	$4,
	$5,
	$6,
	$7,
	(SELECT ID FROM MASTER_EVENT_CATEGORY WHERE CONCAT(EVENT_CATEGORY_PARENT, ' (', EVENT_CATEGORY_CHILD, ')') = $8 AND FLAG_IS_DELETED = FALSE LIMIT 1),
	(SELECT ID FROM MASTER_COST_CATEGORY WHERE CONCAT(COST_CATEGORY_PARENT, ' (', COST_CATEGORY_CHILD, ')') = $9 AND FLAG_IS_DELETED = FALSE LIMIT 1),
	(SELECT ID FROM MASTER_VENDOR_LIST WHERE VENDOR_NAME = $10 AND FLAG_IS_DELETED = FALSE LIMIT 1),
	$11
)
`

const baseActivityUpdate = `
UPDATE user_project_activity
SET
	user_id = (
		SELECT UUID FROM USER_AUTH
		WHERE EMAIL = $1
		LIMIT 1
	),
	project_id = (
		SELECT ID FROM MASTER_PROJECT_LIST
		WHERE PROJECT_NAME = $2
		AND FLAG_IS_DELETED = FALSE
		LIMIT 1
	),
	activity_type = $3,
	activity_title = $4,
	activity_body = $5,
	amount = $6,
	activity_date = $7,
	event_category_id = (
		SELECT ID FROM MASTER_EVENT_CATEGORY
		WHERE CONCAT(EVENT_CATEGORY_PARENT, ' (', EVENT_CATEGORY_CHILD, ')') = $8
		AND FLAG_IS_DELETED = FALSE
		LIMIT 1
	),
	cost_category_id = (
		SELECT ID FROM MASTER_COST_CATEGORY
		WHERE CONCAT(COST_CATEGORY_PARENT, ' (', COST_CATEGORY_CHILD, ')') = $9
		AND FLAG_IS_DELETED = FALSE
		LIMIT 1
	),
	vendor_id = (
		SELECT ID FROM MASTER_VENDOR_LIST
		WHERE VENDOR_NAME = $10
		AND FLAG_IS_DELETED = FALSE
		LIMIT 1
	),
	photo_url = $11,
	updated_at = NOW()
WHERE id = $12
AND user_id = (
	SELECT UUID FROM USER_AUTH
	WHERE EMAIL = $1
	LIMIT 1
);
`
