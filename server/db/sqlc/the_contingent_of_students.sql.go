// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: the_contingent_of_students.sql

package db

import (
	"context"
)

const create_the_contingent_of_students = `-- name: Create_the_contingent_of_students :one
INSERT INTO "the_contingent_of_students" (
              "code"         ,
              "group_number" ,
              "of_groups"    ,
              "subgroups"    ,
              "total_people" ,
              "RF"           ,
              "foreign"      ,
              "standard"     ,
              "calculated" ,
              "PK")
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING id, code, group_number, of_groups, subgroups, total_people, "RF", "foreign", standard, calculated, "PK"
`

type Create_the_contingent_of_studentsParams struct {
	Code        string `json:"code"`
	GroupNumber string `json:"group_number"`
	OfGroups    string `json:"of_groups"`
	Subgroups   string `json:"subgroups"`
	TotalPeople string `json:"total_people"`
	RF          string `json:"RF"`
	Foreign     string `json:"foreign"`
	Standard    string `json:"standard"`
	Calculated  string `json:"calculated"`
	PK          string `json:"PK"`
}

func (q *Queries) Create_the_contingent_of_students(ctx context.Context, arg Create_the_contingent_of_studentsParams) (TheContingentOfStudent, error) {
	row := q.db.QueryRow(ctx, create_the_contingent_of_students,
		arg.Code,
		arg.GroupNumber,
		arg.OfGroups,
		arg.Subgroups,
		arg.TotalPeople,
		arg.RF,
		arg.Foreign,
		arg.Standard,
		arg.Calculated,
		arg.PK,
	)
	var i TheContingentOfStudent
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.GroupNumber,
		&i.OfGroups,
		&i.Subgroups,
		&i.TotalPeople,
		&i.RF,
		&i.Foreign,
		&i.Standard,
		&i.Calculated,
		&i.PK,
	)
	return i, err
}

const get_the_contingent_of_students = `-- name: Get_the_contingent_of_students :one
SELECT id, code, group_number, of_groups, subgroups, total_people, "RF", "foreign", standard, calculated, "PK"
FROM the_contingent_of_students
WHERE "group_number" = $1
LIMIT 1
`

func (q *Queries) Get_the_contingent_of_students(ctx context.Context, groupNumber string) (TheContingentOfStudent, error) {
	row := q.db.QueryRow(ctx, get_the_contingent_of_students, groupNumber)
	var i TheContingentOfStudent
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.GroupNumber,
		&i.OfGroups,
		&i.Subgroups,
		&i.TotalPeople,
		&i.RF,
		&i.Foreign,
		&i.Standard,
		&i.Calculated,
		&i.PK,
	)
	return i, err
}