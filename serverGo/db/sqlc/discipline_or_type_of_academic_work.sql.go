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
              "block"                                          ,
              "component"                                      ,
              "n_v_rup"                                        ,
              "dop_info"                                       ,
              "name_of_the_discipline_or_type_of_academic_work")
VALUES ($1, $2, $3, $4,$5)
RETURNING id, block, component, n_v_rup, dop_info, name_of_the_discipline_or_type_of_academic_work
`

type Create_Discipline_or_type_of_academic_workParams struct {
	Block                                   string `json:"block"`
	Component                               string `json:"component"`
	NVRup                                   string `json:"n_v_rup"`
	DopInfo                                 string `json:"dop_info"`
	NameOfTheDisciplineOrTypeOfAcademicWork string `json:"name_of_the_discipline_or_type_of_academic_work"`
}

func (q *Queries) Create_Discipline_or_type_of_academic_work(ctx context.Context, arg Create_Discipline_or_type_of_academic_workParams) (DisciplineOrTypeOfAcademicWork, error) {
	row := q.db.QueryRow(ctx, create_Discipline_or_type_of_academic_work,
		arg.Block,
		arg.Component,
		arg.NVRup,
		arg.DopInfo,
		arg.NameOfTheDisciplineOrTypeOfAcademicWork,
	)
	var i DisciplineOrTypeOfAcademicWork
	err := row.Scan(
		&i.ID,
		&i.Block,
		&i.Component,
		&i.NVRup,
		&i.DopInfo,
		&i.NameOfTheDisciplineOrTypeOfAcademicWork,
	)
	return i, err
}

const get_Discipline_or_type_of_academic_work = `-- name: Get_Discipline_or_type_of_academic_work :one

SELECT id, block, component, n_v_rup, dop_info, name_of_the_discipline_or_type_of_academic_work
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
		&i.NVRup,
		&i.DopInfo,
		&i.NameOfTheDisciplineOrTypeOfAcademicWork,
	)
	return i, err
}

const list_All_Teacher_Disciplines = `-- name: List_All_Teacher_Disciplines :many

SELECT  type_of_educational_work,name_of_the_discipline_or_type_of_academic_work,total,group_name from discipline_or_type_of_academic_work d
join together t on d.id = t.discipline_id
join k_w kw on t.k_w_id = kw.id
join the_amount_of_teaching_work_of_the_teaching_staff taotwotts on t.amount_id = taotwotts.id
join "information_about_PPS" iaP on iaP.id = t.teacher_id and iap.full_name=$1
join the_contingent_of_students tcos on t.group_id = tcos.id
join semester s on s.id = t.semestr_id and s.semester_type=$2
`

type List_All_Teacher_DisciplinesParams struct {
	FullName     string `json:"full_name"`
	SemesterType string `json:"semester_type"`
}

type List_All_Teacher_DisciplinesRow struct {
	TypeOfEducationalWork                   string  `json:"type_of_educational_work"`
	NameOfTheDisciplineOrTypeOfAcademicWork string  `json:"name_of_the_discipline_or_type_of_academic_work"`
	Total                                   float64 `json:"total"`
	GroupName                               string  `json:"group_name"`
}

func (q *Queries) List_All_Teacher_Disciplines(ctx context.Context, arg List_All_Teacher_DisciplinesParams) ([]List_All_Teacher_DisciplinesRow, error) {
	rows, err := q.db.Query(ctx, list_All_Teacher_Disciplines, arg.FullName, arg.SemesterType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []List_All_Teacher_DisciplinesRow{}
	for rows.Next() {
		var i List_All_Teacher_DisciplinesRow
		if err := rows.Scan(
			&i.TypeOfEducationalWork,
			&i.NameOfTheDisciplineOrTypeOfAcademicWork,
			&i.Total,
			&i.GroupName,
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