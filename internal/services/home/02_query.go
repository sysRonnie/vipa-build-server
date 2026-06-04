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

		SUM(
			CASE
				WHEN activity_type = 'expense'
				THEN amount
				ELSE 0
			END
		) AS total_expenses,

		SUM(
			CASE
				WHEN activity_type = 'income'
				THEN amount
				ELSE 0
			END
		) AS total_income,

		COUNT(
			CASE
				WHEN activity_type = 'event'
				THEN 1
			END
		) AS total_events,

		COUNT(
			CASE
				WHEN activity_type = 'event'
				AND flag_is_completed = true
				THEN 1
			END
		) AS total_events_complete,

		COALESCE(ROUND(
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
			)*100,
			0	
		),0) AS total_percent,

		MAX(created_at) AS last_activity_at,
		CURRENT_DATE - MIN(activity_date)  as total_project_days

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
	COALESCE(B.total_events_complete, 0)::int AS total_events_complete,
	COALESCE(B.total_percent, 0)::int AS total_percent,
	COALESCE(B.total_project_days, 0) as total_project_days,
	COALESCE(B.last_activity_at, NOW()) AS last_activity_at
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
			&card.TotalEventsComplete,
			&card.TotalEventsPercent,
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

	return HomeDashboard{Projects: cards}, nil
}