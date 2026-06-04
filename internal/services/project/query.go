package project


var baseProjectListQuery = `
SELECT 
	A.ID,
	CONCAT(B.CUSTOMER_NAME_FIRST, ' ', B.CUSTOMER_NAME_LAST) AS CUSTOMER_NAME,
	A.PROJECT_NAME,
	COALESCE(TO_CHAR(A.PROJECT_START_DATE, 'YYYY-MM-DD'), '') AS PROJECT_START_DATE,
	COALESCE(TO_CHAR(A.PROJECT_END_DATE_EST, 'YYYY-MM-DD'), '') AS PROJECT_END_DATE_EST,
	COALESCE(TO_CHAR(A.PROJECT_END_DATE_ACTUAL, 'YYYY-MM-DD'), '') AS PROJECT_END_DATE_ACTUAL,
	ROUND(A.PROJECT_PRICE, 0)::int AS PROJECT_PRICE,
	ROUND(A.PROJECT_BUDGET, 0)::int AS PROJECT_BUDGET,
	A.PROJECT_NOTE,
	A.FLAG_IS_DELETED,
	A.CREATED_AT,
	A.UPDATED_AT
FROM MASTER_PROJECT_LIST A
LEFT JOIN MASTER_CUSTOMER_LIST B ON A.CUSTOMER_ID = B.ID
WHERE 1=1
`	

func buildProjectListQuery() string {
	return baseProjectListQuery + " AND A.FLAG_IS_DELETED = FALSE"
}

func buildProjectListRecycledQuery() string {
	return baseProjectListQuery + " AND A.FLAG_IS_DELETED = TRUE"
}
var baseProjectIncomeListByIDQuery = `
		WITH CTE_VENDOR_EXPENSE AS (
			SELECT project_id, activity_title, income_category_id, SUM(amount) AS total_amount
			FROM USER_PROJECT_ACTIVITY
			WHERE activity_type = 'income'
			AND project_id = $1
			GROUP BY project_id, activity_title, income_category_id
		)
		SELECT 
			A.activity_title as vendor_name, 
			B.income_category_parent,
			CASE WHEN B.income_category_child IS NULL THEN '' ELSE B.income_category_child END AS cost_category_child,
			A.total_amount
		FROM CTE_VENDOR_EXPENSE A
		LEFT JOIN MASTER_INCOME_CATEGORY B ON B.id = A.income_category_id
`

var baseProjectExpenseListByIDQuery = `
		WITH CTE_VENDOR_EXPENSE AS (
			SELECT project_id, vendor_id, cost_category_id, SUM(amount) AS total_amount
			FROM USER_PROJECT_ACTIVITY
			WHERE activity_type = 'expense'
			AND project_id = $1
			GROUP BY project_id, vendor_id, cost_category_id
		)
		SELECT 
			B.vendor_name,
			C.cost_category_parent,
			CASE WHEN C.cost_category_child IS NULL THEN '' ELSE C.cost_category_child END AS cost_category_child,
			A.total_amount
		FROM CTE_VENDOR_EXPENSE A
		LEFT JOIN MASTER_VENDOR_LIST B ON B.id = A.vendor_id
		LEFT JOIN MASTER_COST_CATEGORY C ON C.id = A.cost_category_id

`

var baseProjectListNamesQuery = `
SELECT DISTINCT 
	A.PROJECT_NAME
FROM MASTER_PROJECT_LIST A
WHERE A.FLAG_IS_DELETED = FALSE
ORDER BY A.PROJECT_NAME
`


var baseProjectInsert = `
INSERT INTO MASTER_PROJECT_LIST (
	CUSTOMER_ID,
	PROJECT_NAME,
	PROJECT_START_DATE,
	PROJECT_END_DATE_EST,
	PROJECT_END_DATE_ACTUAL,
	PROJECT_PRICE,
	PROJECT_BUDGET,
	PROJECT_NOTE
) VALUES ((select id from master_customer_list where CONCAT(customer_name_first, ' ', customer_name_last) = $1), $2, $3, $4, $5, $6, $7, $8)
`

var baseProjectByIDQuery = `
SELECT 
	A.ID,
	CONCAT(B.CUSTOMER_NAME_FIRST, ' ', B.CUSTOMER_NAME_LAST) AS CUSTOMER_NAME,
	A.PROJECT_NAME,
	COALESCE(TO_CHAR(A.PROJECT_START_DATE, 'YYYY-MM-DD'), '') AS PROJECT_START_DATE,
	COALESCE(TO_CHAR(A.PROJECT_END_DATE_EST, 'YYYY-MM-DD'), '') AS PROJECT_END_DATE_EST,
	COALESCE(TO_CHAR(A.PROJECT_END_DATE_ACTUAL, 'YYYY-MM-DD'), '') AS PROJECT_END_DATE_ACTUAL,
	A.PROJECT_PRICE,
	A.PROJECT_BUDGET,
	A.PROJECT_NOTE,
	A.FLAG_IS_DELETED,
	A.CREATED_AT,
	A.UPDATED_AT
FROM MASTER_PROJECT_LIST A
LEFT JOIN MASTER_CUSTOMER_LIST B ON A.CUSTOMER_ID = B.ID
WHERE A.ID = $1 
`


var baseProjectDelete = `
UPDATE MASTER_PROJECT_LIST
SET FLAG_IS_DELETED = TRUE, UPDATED_AT = NOW()
WHERE ID = $1
`

var baseProjectUpdate = `
UPDATE MASTER_PROJECT_LIST
SET
	CUSTOMER_ID = (select id from master_customer_list where CONCAT(customer_name_first, ' ', customer_name_last) = $1),
	PROJECT_NAME = $2,
	PROJECT_START_DATE = $3,
	PROJECT_END_DATE_EST = $4,
	PROJECT_END_DATE_ACTUAL = $5,
	PROJECT_PRICE = $6,
	PROJECT_BUDGET = $7,
	PROJECT_NOTE = $8,
	UPDATED_AT = NOW(),
	FLAG_IS_DELETED = false
WHERE ID = $9
`

const baseProjectNameLatestQuery = `
SELECT 
	CONCAT(A.PROJECT_NAME,' (',b.CUSTOMER_NAME_FIRST, ' ', b.CUSTOMER_NAME_LAST,')') as project_name
FROM MASTER_PROJECT_LIST A
LEFT JOIN MASTER_CUSTOMER_LIST B ON A.customer_id = b.id
WHERE A.FLAG_IS_DELETED = FALSE
ORDER BY A.CREATED_AT DESC
LIMIT 1
`