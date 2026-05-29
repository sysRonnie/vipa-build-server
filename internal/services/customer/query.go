package customer

const baseCustomerSelect = `
	SELECT
		id,
		customer_name_first,
		customer_name_last,
		COALESCE(customer_phone, '') AS customer_phone,
		COALESCE(customer_email, '') AS customer_email,
		COALESCE(customer_address_street, '') AS customer_address_street,
		COALESCE(customer_address_unit, '') AS customer_address_unit,
		COALESCE(customer_address_city, '') AS customer_address_city,
		COALESCE(customer_address_state, '') AS customer_address_state,
		COALESCE(customer_address_zip, '') AS customer_address_zip,
		COALESCE(comment, '') AS comment,
		flag_is_deleted,
		TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS') AS created_at,
		TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS') AS updated_at
	FROM master_customer_list
	WHERE 1=1
`

func buildCustomerListQuery(recycled bool) string {
	query := baseCustomerSelect
	query += ` AND flag_is_deleted = $1`
	query += ` ORDER BY created_at DESC`
	return query
}

func buildCustomerDetailQuery() string {
	query := baseCustomerSelect
	query += ` AND id = $1`
	return query
}

const baseCustomerUpdate = `
		UPDATE master_customer_list
		SET
			customer_name_first = $1,
			customer_name_last = $2,
			customer_phone = $3,
			customer_email = $4,
			comment = $5,
			customer_address_street = $6,
			customer_address_unit = $7,
			customer_address_city = $8,
			customer_address_state = $9,
			customer_address_zip = $10,
			flag_is_deleted = false,
			updated_at = NOW()
		WHERE 1=1
		AND customer_name_first = $1
	AND customer_name_last = $2
`

const baseCustomerSoftDelete = `
		UPDATE master_customer_list
		SET flag_is_deleted = true, updated_at = NOW()
		WHERE id = $1 AND flag_is_deleted = false
`

const baseCustomerInsert = `
		INSERT INTO master_customer_list (
			customer_name_first,
			customer_name_last,
			customer_phone,
			customer_email,
			customer_address_street,
			customer_address_unit,
			customer_address_city,
			customer_address_state,
			customer_address_zip,
			comment
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
`

const baseCustomerExistsQuery = `
		SELECT EXISTS (
			SELECT 1
			FROM master_customer_list
			WHERE customer_name_first = $1
			AND customer_name_last = $2
			AND flag_is_deleted = false
		)
`
const baseCustomerExistsRecycledQuery = `
		SELECT EXISTS (
			SELECT 1
			FROM master_customer_list
			WHERE customer_name_first = $1
			AND customer_name_last = $2
			AND flag_is_deleted = true
		)
`

const baseCustomerNamesQuery = `
		SELECT CONCAT(customer_name_first, ' ', customer_name_last) AS full_name
		FROM master_customer_list
		WHERE flag_is_deleted = false
		ORDER BY created_at DESC
`


const baseCustomerNamesLatestQuery = `
		SELECT CONCAT(customer_name_first, ' ', customer_name_last) AS full_name
		FROM master_customer_list
		WHERE flag_is_deleted = false
		ORDER BY created_at DESC
		LIMIT 1
`
