// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: educational_program.sql

package db

import (
	"context"
)

const create_EducationalProgram = `-- name: Create_EducationalProgram :one
INSERT INTO "educational_program" ("the_code_of_the_OOP_RUDN",
                                   "direction_code",
                                   "name_of_the_program")

VALUES ($1, $2, $3)
RETURNING id, "the_code_of_the_OOP_RUDN", direction_code, name_of_the_program
`

type Create_EducationalProgramParams struct {
	TheCodeOfTheOOPRUDN string `json:"the_code_of_the_OOP_RUDN"`
	DirectionCode       string `json:"direction_code"`
	NameOfTheProgram    string `json:"name_of_the_program"`
}

func (q *Queries) Create_EducationalProgram(ctx context.Context, arg Create_EducationalProgramParams) (EducationalProgram, error) {
	row := q.db.QueryRow(ctx, create_EducationalProgram, arg.TheCodeOfTheOOPRUDN, arg.DirectionCode, arg.NameOfTheProgram)
	var i EducationalProgram
	err := row.Scan(
		&i.ID,
		&i.TheCodeOfTheOOPRUDN,
		&i.DirectionCode,
		&i.NameOfTheProgram,
	)
	return i, err
}

const get_EducationalProgram = `-- name: Get_EducationalProgram :one
SELECT id, "the_code_of_the_OOP_RUDN", direction_code, name_of_the_program
FROM educational_program
WHERE "id" = $1
LIMIT 1
`

func (q *Queries) Get_EducationalProgram(ctx context.Context, id int32) (EducationalProgram, error) {
	row := q.db.QueryRow(ctx, get_EducationalProgram, id)
	var i EducationalProgram
	err := row.Scan(
		&i.ID,
		&i.TheCodeOfTheOOPRUDN,
		&i.DirectionCode,
		&i.NameOfTheProgram,
	)
	return i, err
}
