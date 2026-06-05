package home

import "database/sql"

const baseHomeRead = `
WITH CTE_PROJECT_KEYS AS (
	SELECT 
		A.id,
		A.project_name,
		CONCAT(B.customer_name_first, ' ', B.customer_name_last) AS customer_name
	FROM master_project_list A
	LEFT JOIN master_customer_list B ON A.customer_id = B.id
	WHERE A.flag_is_deleted = false
),

CTE_PROJECT_ACTIVITY AS (
	SELECT
		project_id,

		MIN(activity_date) AS project_start_date_raw,

		SUM(CASE WHEN activity_type = 'expense' THEN amount ELSE 0 END) AS total_expenses,
		SUM(CASE WHEN activity_type = 'income' THEN amount ELSE 0 END) AS total_income,

		SUM(
			CASE
				WHEN activity_type = 'income'
				AND flag_is_completed = false
				THEN amount
				ELSE 0
			END
		) AS outstanding_income,

		SUM(
			CASE
				WHEN activity_type = 'expense'
				AND flag_is_completed = false
				THEN amount
				ELSE 0
			END
		) AS unpaid_expenses,

		COUNT(
			CASE
				WHEN activity_type IN ('expense', 'income', 'event', 'note')
				THEN 1
			END
		) AS activity_count,

		COUNT(CASE WHEN activity_type = 'event' THEN 1 END) AS total_events,

		COUNT(
			CASE
				WHEN activity_type = 'event'
				AND flag_is_completed = true
				THEN 1
			END
		) AS total_events_complete,

		COALESCE(
			ROUND(
				(
					COUNT(
						CASE
							WHEN activity_type = 'event'
							AND flag_is_completed = true
							THEN 1
						END
					)::numeric
					/
					NULLIF(
						COUNT(
							CASE
								WHEN activity_type = 'event'
								THEN 1
							END
						),
						0
					)
				) * 100,
				0
			),
			0
		) AS total_percent,

		MAX(created_at) AS last_activity_at

	FROM user_project_activity
	WHERE flag_is_deleted = false
	GROUP BY project_id
)

SELECT 
	A.id AS project_id,
	A.project_name,
	A.customer_name,

	COALESCE(B.total_events, 0)::int AS total_events,

	ROUND(COALESCE(B.total_expenses, 0), 2)::int AS total_expenses,
	ROUND(COALESCE(B.total_income, 0), 2)::int AS total_income,
	ROUND(COALESCE(B.total_income, 0) - COALESCE(B.total_expenses, 0), 2)::int AS net_total,

	ROUND(COALESCE(B.outstanding_income, 0), 2)::int AS outstanding_income,
	ROUND(COALESCE(B.unpaid_expenses, 0), 2)::int AS unpaid_expenses,

	COALESCE(B.activity_count, 0)::int AS activity_count,

	COALESCE(B.total_events_complete, 0)::int AS total_events_complete,
	COALESCE(B.total_percent, 0)::int AS total_percent,

	COALESCE(TO_CHAR(B.project_start_date_raw, 'YYYY-MM-DD'), '') AS project_start_date,
	COALESCE(GREATEST(CURRENT_DATE - B.project_start_date_raw, 0), 0)::int AS total_project_days,
	COALESCE(TO_CHAR(B.last_activity_at, 'YYYY-MM-DD'), '') AS last_activity_at

FROM CTE_PROJECT_KEYS A
LEFT JOIN CTE_PROJECT_ACTIVITY B ON A.id = B.project_id
ORDER BY B.last_activity_at DESC NULLS LAST, A.project_name;
`

func (s *Store) ScanHomeProjectCards(rows *sql.Rows) (HomeDashboard, error) {
	cards := make([]HomeProjectCard, 0)

	for rows.Next() {
		var card HomeProjectCard

		err := rows.Scan(
			&card.ProjectID,
			&card.ProjectName,
			&card.CustomerName,
			&card.TotalEvents,
			&card.TotalExpenses,
			&card.TotalIncome,
			&card.NetTotal,
			&card.OutstandingIncome,
			&card.UnpaidExpenses,
			&card.ActivityCount,
			&card.TotalEventsComplete,
			&card.TotalEventsPercent,
			&card.ProjectStartDate,
			&card.TotalProjectDays,
			&card.LastActivityAt,
		)

		if err != nil {
			return HomeDashboard{}, err
		}

		cards = append(cards, card)
	}

	if err := rows.Err(); err != nil {
		return HomeDashboard{}, err
	}

	return HomeDashboard{
		Projects: cards,
	}, nil
}