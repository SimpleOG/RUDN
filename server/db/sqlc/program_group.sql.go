// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: program_group.sql

package db

import (
	"context"
)

const create_Program_group = `-- name: Create_Program_group :one
INSERT INTO  "program_group"(
    name_of_the_program,
    group_name
    )

VALUES ($1, $2)
RETURNING name_of_the_program, group_name
`

type Create_Program_groupParams struct {
	NameOfTheProgram string `json:"name_of_the_program"`
	GroupName        string `json:"group_name"`
}

func (q *Queries) Create_Program_group(ctx context.Context, arg Create_Program_groupParams) (ProgramGroup, error) {
	row := q.db.QueryRow(ctx, create_Program_group, arg.NameOfTheProgram, arg.GroupName)
	var i ProgramGroup
	err := row.Scan(&i.NameOfTheProgram, &i.GroupName)
	return i, err
}

const get_Program_group = `-- name: Get_Program_group :one
SELECT name_of_the_program, group_name
FROM Program_group
WHERE name_of_the_program = $1 and group_name=$2
`

type Get_Program_groupParams struct {
	NameOfTheProgram string `json:"name_of_the_program"`
	GroupName        string `json:"group_name"`
}

func (q *Queries) Get_Program_group(ctx context.Context, arg Get_Program_groupParams) (ProgramGroup, error) {
	row := q.db.QueryRow(ctx, get_Program_group, arg.NameOfTheProgram, arg.GroupName)
	var i ProgramGroup
	err := row.Scan(&i.NameOfTheProgram, &i.GroupName)
	return i, err
}