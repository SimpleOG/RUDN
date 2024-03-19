// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: information_about_PPS.sql

package db

import (
	"context"
)

const create_information_about_PPS = `-- name: Create_information_about_PPS :one

INSERT INTO "information_about_PPS" ("department",
                                     "post",
                                     "terms_of_attraction",
                                     "full_name",
                                     "a_special_feature")
VALUES ($1, $2, $3, $4, $5)
RETURNING id, department, post, terms_of_attraction, full_name, a_special_feature
`

type Create_information_about_PPSParams struct {
	Department        string `json:"department"`
	Post              string `json:"post"`
	TermsOfAttraction string `json:"terms_of_attraction"`
	FullName          string `json:"full_name"`
	ASpecialFeature   string `json:"a_special_feature"`
}

func (q *Queries) Create_information_about_PPS(ctx context.Context, arg Create_information_about_PPSParams) (InformationAboutPP, error) {
	row := q.db.QueryRow(ctx, create_information_about_PPS,
		arg.Department,
		arg.Post,
		arg.TermsOfAttraction,
		arg.FullName,
		arg.ASpecialFeature,
	)
	var i InformationAboutPP
	err := row.Scan(
		&i.ID,
		&i.Department,
		&i.Post,
		&i.TermsOfAttraction,
		&i.FullName,
		&i.ASpecialFeature,
	)
	return i, err
}

const get_information_about_PPS = `-- name: Get_information_about_PPS :many

select distinct full_name ,department,post,terms_of_attraction from "information_about_PPS"
`

type Get_information_about_PPSRow struct {
	FullName          string `json:"full_name"`
	Department        string `json:"department"`
	Post              string `json:"post"`
	TermsOfAttraction string `json:"terms_of_attraction"`
}

func (q *Queries) Get_information_about_PPS(ctx context.Context) ([]Get_information_about_PPSRow, error) {
	rows, err := q.db.Query(ctx, get_information_about_PPS)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Get_information_about_PPSRow{}
	for rows.Next() {
		var i Get_information_about_PPSRow
		if err := rows.Scan(
			&i.FullName,
			&i.Department,
			&i.Post,
			&i.TermsOfAttraction,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}