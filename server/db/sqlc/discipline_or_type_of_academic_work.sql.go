// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: discipline_or_type_of_academic_work.sql

package db

import (
	"context"
)

const create_Discipline_or_type_of_academic_work = `-- name: Create_Discipline_or_type_of_academic_work :one
INSERT INTO "discipline_or_type_of_academic_work" (
               "block",
               "component",
               "n_v_RUP",
               "name_of_the_discipline_or_type_of_academic_work",
               "dop_info")
VALUES ($1, $2, $3, $4,$5)
RETURNING id, block, component, "n_v_RUP", dop_info, name_of_the_discipline_or_type_of_academic_work
`

type Create_Discipline_or_type_of_academic_workParams struct {
	Block                                   string `json:"block"`
	Component                               string `json:"component"`
	NVRUP                                   string `json:"n_v_RUP"`
	NameOfTheDisciplineOrTypeOfAcademicWork string `json:"name_of_the_discipline_or_type_of_academic_work"`
	DopInfo                                 string `json:"dop_info"`
}

func (q *Queries) Create_Discipline_or_type_of_academic_work(ctx context.Context, arg Create_Discipline_or_type_of_academic_workParams) (DisciplineOrTypeOfAcademicWork, error) {
	row := q.db.QueryRow(ctx, create_Discipline_or_type_of_academic_work,
		arg.Block,
		arg.Component,
		arg.NVRUP,
		arg.NameOfTheDisciplineOrTypeOfAcademicWork,
		arg.DopInfo,
	)
	var i DisciplineOrTypeOfAcademicWork
	err := row.Scan(
		&i.ID,
		&i.Block,
		&i.Component,
		&i.NVRUP,
		&i.DopInfo,
		&i.NameOfTheDisciplineOrTypeOfAcademicWork,
	)
	return i, err
}

const get_Discipline_or_type_of_academic_work = `-- name: Get_Discipline_or_type_of_academic_work :one
SELECT id, block, component, "n_v_RUP", dop_info, name_of_the_discipline_or_type_of_academic_work
FROM "discipline_or_type_of_academic_work"
WHERE "id" = $1
LIMIT 1
`

func (q *Queries) Get_Discipline_or_type_of_academic_work(ctx context.Context, id int32) (DisciplineOrTypeOfAcademicWork, error) {
	row := q.db.QueryRow(ctx, get_Discipline_or_type_of_academic_work, id)
	var i DisciplineOrTypeOfAcademicWork
	err := row.Scan(
		&i.ID,
		&i.Block,
		&i.Component,
		&i.NVRUP,
		&i.DopInfo,
		&i.NameOfTheDisciplineOrTypeOfAcademicWork,
	)
	return i, err
}
