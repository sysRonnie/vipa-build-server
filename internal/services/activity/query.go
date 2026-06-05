package activity



const baseActivityListQuery = `
SELECT 
	a.id,
	a.user_id,
	a.project_id,
	p.project_name AS project_name,

	c.id AS customer_id,
	CONCAT(c.customer_name_first, ' ', c.customer_name_last) AS customer_name,



	
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

	a.income_category_id,
	CASE
		WHEN ic.income_category_child IS NULL THEN ic.income_category_parent
		ELSE CONCAT(ic.income_category_parent, ' (', ic.income_category_child, ')')
	END AS income_category_name,
	
	a.photo_url,
	a.photo_thumbnail_url,
	
	a.flag_is_completed,
	a.flag_is_deleted,
	
	TO_CHAR(a.created_at, 'YYYY-MM-DD HH24:MI:SS') AS created_at,
	TO_CHAR(a.updated_at, 'YYYY-MM-DD HH24:MI:SS') AS updated_at
FROM user_project_activity a
LEFT JOIN master_project_list p ON a.project_id = p.id
LEFT JOIN master_customer_list c on p.customer_id = c.id
LEFT JOIN master_event_category ec ON a.event_category_id = ec.id
LEFT JOIN master_cost_category cc ON a.cost_category_id = cc.id
LEFT JOIN master_vendor_list v ON a.vendor_id = v.id
LEFT JOIN master_income_category ic ON a.income_category_id = ic.id
WHERE 1=1
`


func BuildBaseActivityListQueryWithDeletedFilter(isDeleted bool) string {
	if isDeleted {
		filterOptions := "AND a.flag_is_deleted = true AND p.flag_is_deleted = false ORDER BY a.created_at DESC, a.activity_date DESC"
		return baseActivityListQuery + filterOptions
	} else {
		filterOptions := "AND a.flag_is_deleted = false AND p.flag_is_deleted = false ORDER BY a.created_at DESC, a.activity_date DESC"
		return baseActivityListQuery + filterOptions
	}
}

const baseActivityListQueryNotes = `
SELECT 
	a.id,
	a.user_id,
	a.project_id,
	p.project_name AS project_name,

	c.id AS customer_id,
	CONCAT(c.customer_name_first, ' ', c.customer_name_last) AS customer_name,
	
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

	a.income_category_id,
	CASE
		WHEN ic.income_category_child IS NULL THEN ic.income_category_parent
		ELSE CONCAT(ic.income_category_parent, ' (', ic.income_category_child, ')')
	END AS income_category_name,
	
	a.photo_url,
	a.photo_thumbnail_url,
	
	a.flag_is_completed,
	a.flag_is_deleted,
	
	TO_CHAR(a.created_at, 'YYYY-MM-DD HH24:MI:SS') AS created_at,
	TO_CHAR(a.updated_at, 'YYYY-MM-DD HH24:MI:SS') AS updated_at
FROM user_project_activity a
LEFT JOIN master_project_list p ON a.project_id = p.id
LEFT JOIN master_customer_list c on p.customer_id = c.id
LEFT JOIN master_event_category ec ON a.event_category_id = ec.id
LEFT JOIN master_cost_category cc ON a.cost_category_id = cc.id
LEFT JOIN master_vendor_list v ON a.vendor_id = v.id
LEFT JOIN master_income_category ic ON a.income_category_id = ic.id
WHERE 1=1
AND a.activity_type = 'note'
`

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
	INCOME_CATEGORY_ID,
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
	(SELECT ID FROM MASTER_INCOME_CATEGORY WHERE CONCAT(INCOME_CATEGORY_PARENT, ' (', INCOME_CATEGORY_CHILD, ')') = $9 AND FLAG_IS_DELETED = FALSE LIMIT 1),
	$10
)
`

const baseActivityListByProjectIdQuery = `
SELECT 
	a.id,
	a.user_id,
	a.project_id,
	p.project_name AS project_name,
	c.id AS customer_id,
	CONCAT(c.customer_name_first, ' ', c.customer_name_last) AS customer_name,
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
	a.income_category_id,
	CASE
		WHEN ic.income_category_child IS NULL THEN ic.income_category_parent
		ELSE CONCAT(ic.income_category_parent, ' (', ic.income_category_child, ')')
	END AS income_category_name,
	a.photo_url,
	a.photo_thumbnail_url,
	a.flag_is_completed,
	a.flag_is_deleted,
	TO_CHAR(a.created_at, 'YYYY-MM-DD HH24:MI:SS') AS created_at,
	TO_CHAR(a.updated_at, 'YYYY-MM-DD HH24:MI:SS') AS updated_at
FROM user_project_activity a
LEFT JOIN master_project_list p ON a.project_id = p.id
LEFT JOIN master_customer_list c on p.customer_id = c.id
LEFT JOIN master_event_category ec ON a.event_category_id = ec.id
LEFT JOIN master_cost_category cc ON a.cost_category_id = cc.id
LEFT JOIN master_vendor_list v ON a.vendor_id = v.id
LEFT JOIN master_income_category ic ON a.income_category_id = ic.id
WHERE p.id = $1
AND a.flag_is_deleted = false
ORDER BY a.activity_date DESC, a.created_at DESC;
`
const baseActivityListByIdQuery = `
SELECT 
	a.id,
	a.user_id,
	a.project_id,
	p.project_name AS project_name,

	c.id AS customer_id,
	CONCAT(c.customer_name_first, ' ', c.customer_name_last) AS customer_name,

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
	a.income_category_id,
	CASE
		WHEN ic.income_category_child IS NULL THEN ic.income_category_parent
		ELSE CONCAT(ic.income_category_parent, ' (', ic.income_category_child, ')')
	END AS income_category_name,
	
	a.photo_url,
	a.photo_thumbnail_url,
	
	a.flag_is_completed,
	a.flag_is_deleted,
	
	TO_CHAR(a.created_at, 'YYYY-MM-DD HH24:MI:SS') AS created_at,
	TO_CHAR(a.updated_at, 'YYYY-MM-DD HH24:MI:SS') AS updated_at
FROM user_project_activity a
LEFT JOIN master_project_list p ON a.project_id = p.id
LEFT JOIN master_customer_list c on p.customer_id = c.id
LEFT JOIN master_event_category ec ON a.event_category_id = ec.id
LEFT JOIN master_cost_category cc ON a.cost_category_id = cc.id
LEFT JOIN master_vendor_list v ON a.vendor_id = v.id
LEFT JOIN master_income_category ic ON a.income_category_id = ic.id
WHERE 1=1 
AND a.id = $1
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
LEFT JOIN master_customer_list c on p.customer_id = c.id
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
INSERT INTO user_project_activity (
	user_id,
	project_id,
	activity_type,
	activity_title,
	activity_body,
	activity_date,
	amount,
	event_category_id,
	cost_category_id,
	vendor_id,
	income_category_id,
	photo_url
) VALUES (
	(SELECT uuid FROM user_auth WHERE email = $1 LIMIT 1),
	(SELECT id FROM master_project_list WHERE project_name = $2 AND flag_is_deleted = false LIMIT 1),
	$3,
	$4,
	$5,
	COALESCE($6::date, NOW()::date),
	$7,
	(SELECT id FROM master_event_category WHERE event_category_parent = $8 AND flag_is_deleted = false LIMIT 1),
	(SELECT id FROM master_cost_category WHERE CONCAT(cost_category_parent, ' (', cost_category_child, ')') = $9 AND flag_is_deleted = false LIMIT 1),
	(SELECT id FROM master_vendor_list WHERE vendor_name = $10 AND flag_is_deleted = false LIMIT 1),
	(SELECT id FROM master_income_category WHERE income_category_parent = $11 AND flag_is_deleted = false LIMIT 1),
	$12
)
`

const baseActivityUpdate = `
UPDATE user_project_activity
SET
	project_id = (
		SELECT id
		FROM master_project_list
		WHERE project_name = $2
		AND flag_is_deleted = FALSE
		LIMIT 1
	),
	activity_type = $3,
	activity_title = $4,
	activity_body = $5,
	amount = $6,
	activity_date = $7,
	event_category_id = (
		SELECT id
		FROM master_event_category
		WHERE
			CASE
				WHEN event_category_child IS NULL THEN event_category_parent = $8
				ELSE CONCAT(event_category_parent, ' (', event_category_child, ')') = $8
			END
		AND flag_is_deleted = FALSE
		LIMIT 1
	),
	cost_category_id = (
		SELECT id
		FROM master_cost_category
		WHERE
			CASE
				WHEN cost_category_child IS NULL THEN cost_category_parent = $9
				ELSE CONCAT(cost_category_parent, ' (', cost_category_child, ')') = $9
			END
		AND flag_is_deleted = FALSE
		LIMIT 1
	),
	vendor_id = (
		SELECT id
		FROM master_vendor_list
		WHERE vendor_name = $10
		AND flag_is_deleted = FALSE
		LIMIT 1
	),
	income_category_id = (
		SELECT id
		FROM master_income_category
		WHERE
			CASE
				WHEN income_category_child IS NULL THEN income_category_parent = $11
				ELSE CONCAT(income_category_parent, ' (', income_category_child, ')') = $11
			END
		AND flag_is_deleted = FALSE
		LIMIT 1
	),
	photo_url = $12,
	flag_is_completed = $13,
	updated_at = NOW()
WHERE id = $14
AND user_id = (
	SELECT uuid
	FROM user_auth
	WHERE email = $1
	LIMIT 1
)
RETURNING
	id,
	user_id,
	project_id,
	(
		SELECT project_name
		FROM master_project_list
		WHERE id = user_project_activity.project_id
	) AS project_name,
	activity_type,
	activity_title,
	activity_body,
	amount,
	TO_CHAR(activity_date, 'YYYY-MM-DD') AS activity_date,
	event_category_id,
	(
		SELECT
			CASE
				WHEN event_category_child IS NULL THEN event_category_parent
				ELSE CONCAT(event_category_parent, ' (', event_category_child, ')')
			END
		FROM master_event_category
		WHERE id = user_project_activity.event_category_id
	) AS event_category_name,
	cost_category_id,
	(
		SELECT
			CASE
				WHEN cost_category_child IS NULL THEN cost_category_parent
				ELSE CONCAT(cost_category_parent, ' (', cost_category_child, ')')
			END
		FROM master_cost_category
		WHERE id = user_project_activity.cost_category_id
	) AS cost_category_name,
	vendor_id,
	(
		SELECT vendor_name
		FROM master_vendor_list
		WHERE id = user_project_activity.vendor_id
	) AS vendor_name,
	income_category_id,
	(
		SELECT
			CASE
				WHEN income_category_child IS NULL THEN income_category_parent
				ELSE CONCAT(income_category_parent, ' (', income_category_child, ')')
			END
		FROM master_income_category
		WHERE id = user_project_activity.income_category_id
	) AS income_category_name,
	photo_url,
	photo_thumbnail_url,
	flag_is_completed,
	flag_is_deleted,
	TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS') AS created_at,
	TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS') AS updated_at;
`

var baseProjectListNamesQuery = `
SELECT DISTINCT 
	CONCAT(A.PROJECT_NAME, ' (', B.CUSTOMER_NAME_FIRST, ' ', B.CUSTOMER_NAME_LAST, ')') AS PROJECT_NAME
FROM MASTER_PROJECT_LIST A
LEFT JOIN MASTER_CUSTOMER_LIST B ON A.CUSTOMER_ID = B.ID
WHERE A.FLAG_IS_DELETED = FALSE
ORDER BY CONCAT(A.PROJECT_NAME, ' (', B.CUSTOMER_NAME_FIRST, ' ', B.CUSTOMER_NAME_LAST, ')')
`

var baseVendorListNamesQuery = `
SELECT DISTINCT 
	A.VENDOR_NAME
FROM MASTER_VENDOR_LIST A
WHERE A.FLAG_IS_DELETED = FALSE
ORDER BY A.VENDOR_NAME
`

var baseCostListNamesQuery = `
SELECT 
	CASE WHEN A.COST_CATEGORY_CHILD IS NULL OR A.COST_CATEGORY_CHILD = '' THEN A.COST_CATEGORY_PARENT 
	ELSE CONCAT(A.COST_CATEGORY_PARENT, ' (', A.COST_CATEGORY_CHILD, ')') END AS COST_CATEGORY_NAME
FROM MASTER_COST_CATEGORY A
WHERE A.FLAG_IS_DELETED = FALSE
ORDER BY COST_CATEGORY_NAME
`

var baseEventListNamesQuery = `
SELECT 
	CASE WHEN 
		A.EVENT_CATEGORY_CHILD IS NULL THEN A.EVENT_CATEGORY_PARENT 
		ELSE A.EVENT_CATEGORY_PARENT || ' (' || A.EVENT_CATEGORY_CHILD || ')' 
		END AS EVENT_CATEGORY_NAME
FROM MASTER_EVENT_CATEGORY A
WHERE A.FLAG_IS_DELETED = FALSE
ORDER BY A.EVENT_CATEGORY_PARENT
`

var baseIncomeListQuery = `
SELECT DISTINCT
	CASE WHEN income_category_child IS NULL THEN income_category_parent
	ELSE CONCAT(income_category_parent, ' (', income_category_child, ')') END AS income_category_name
FROM master_income_category
WHERE flag_is_deleted = false
ORDER BY income_category_name
`