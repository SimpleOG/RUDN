// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: group_hours.sql

package db

import (
	"context"
)

const create_group_hours_discipline = `-- name: Create_group_hours_discipline :one
INSERT INTO "group_hours_discipline"(
         group_name,
         discpline_name,
         amount_id)

VALUES ($1, $2, $3)
RETURNING discpline_name, group_name, amount_id
`

type Create_group_hours_disciplineParams struct {
	GroupName     string `json:"group_name"`
	DiscplineName string `json:"discpline_name"`
	AmountID      int32  `json:"amount_id"`
}

func (q *Queries) Create_group_hours_discipline(ctx context.Context, arg Create_group_hours_disciplineParams) (GroupHoursDiscipline, error) {
	row := q.db.QueryRow(ctx, create_group_hours_discipline, arg.GroupName, arg.DiscplineName, arg.AmountID)
	var i GroupHoursDiscipline
	err := row.Scan(&i.DiscplineName, &i.GroupName, &i.AmountID)
	return i, err
}

const get_group_hours_discipline = `-- name: Get_group_hours_discipline :one
SELECT discpline_name, group_name, amount_id
FROM group_hours_discipline
WHERE group_name = $1 and discpline_name=$2
`

type Get_group_hours_disciplineParams struct {
	GroupName     string `json:"group_name"`
	DiscplineName string `json:"discpline_name"`
}

func (q *Queries) Get_group_hours_discipline(ctx context.Context, arg Get_group_hours_disciplineParams) (GroupHoursDiscipline, error) {
	row := q.db.QueryRow(ctx, get_group_hours_discipline, arg.GroupName, arg.DiscplineName)
	var i GroupHoursDiscipline
	err := row.Scan(&i.DiscplineName, &i.GroupName, &i.AmountID)
	return i, err
}
