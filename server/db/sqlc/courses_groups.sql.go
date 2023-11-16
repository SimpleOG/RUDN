// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: courses_groups.sql

package db

import (
	"context"
)

const createGroupsCourse = `-- name: CreateGroupsCourse :one
INSERT INTO courses_groups (
    course_name,
    groups_id

) VALUES (
             $1, $2
         ) RETURNING course_name, groups_id
`

type CreateGroupsCourseParams struct {
	CourseName string `json:"course_name"`
	GroupsID   int32  `json:"groups_id"`
}

func (q *Queries) CreateGroupsCourse(ctx context.Context, arg CreateGroupsCourseParams) (CoursesGroup, error) {
	row := q.db.QueryRow(ctx, createGroupsCourse, arg.CourseName, arg.GroupsID)
	var i CoursesGroup
	err := row.Scan(&i.CourseName, &i.GroupsID)
	return i, err
}

const getGroupsCourse = `-- name: GetGroupsCourse :one
SELECT course_name, groups_id FROM courses_groups
WHERE course_name =$1 and groups_id=$2  LIMIT 1
`

type GetGroupsCourseParams struct {
	CourseName string `json:"course_name"`
	GroupsID   int32  `json:"groups_id"`
}

func (q *Queries) GetGroupsCourse(ctx context.Context, arg GetGroupsCourseParams) (CoursesGroup, error) {
	row := q.db.QueryRow(ctx, getGroupsCourse, arg.CourseName, arg.GroupsID)
	var i CoursesGroup
	err := row.Scan(&i.CourseName, &i.GroupsID)
	return i, err
}

const listAllCourseGroups = `-- name: ListAllCourseGroups :many
SELECT course_name, groups_id, c.id, c.name, lecture_hours, laboratories_hours, practise_hours, g.id, code, number, g.name FROM courses_groups join courses c on courses_groups.course_name=$1 and
                               c.name = courses_groups.course_name
                             join groups as g  on g.id = courses_groups.groups_id

ORDER BY course_name
`

type ListAllCourseGroupsRow struct {
	CourseName        string `json:"course_name"`
	GroupsID          int32  `json:"groups_id"`
	ID                int32  `json:"id"`
	Name              string `json:"name"`
	LectureHours      int32  `json:"lecture_hours"`
	LaboratoriesHours int32  `json:"laboratories_hours"`
	PractiseHours     int32  `json:"practise_hours"`
	ID_2              int32  `json:"id_2"`
	Code              string `json:"code"`
	Number            int32  `json:"number"`
	Name_2            string `json:"name_2"`
}

func (q *Queries) ListAllCourseGroups(ctx context.Context, courseName string) ([]ListAllCourseGroupsRow, error) {
	rows, err := q.db.Query(ctx, listAllCourseGroups, courseName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListAllCourseGroupsRow{}
	for rows.Next() {
		var i ListAllCourseGroupsRow
		if err := rows.Scan(
			&i.CourseName,
			&i.GroupsID,
			&i.ID,
			&i.Name,
			&i.LectureHours,
			&i.LaboratoriesHours,
			&i.PractiseHours,
			&i.ID_2,
			&i.Code,
			&i.Number,
			&i.Name_2,
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

const listAllGroupCourses = `-- name: ListAllGroupCourses :many
SELECT course_name, groups_id, g.id, code, number, g.name, c.id, c.name, lecture_hours, laboratories_hours, practise_hours FROM courses_groups join groups as g on groups_id=$1 and
                                                 g.id=courses_groups.groups_id
    join courses c on c.name = courses_groups.course_name

ORDER BY course_name
`

type ListAllGroupCoursesRow struct {
	CourseName        string `json:"course_name"`
	GroupsID          int32  `json:"groups_id"`
	ID                int32  `json:"id"`
	Code              string `json:"code"`
	Number            int32  `json:"number"`
	Name              string `json:"name"`
	ID_2              int32  `json:"id_2"`
	Name_2            string `json:"name_2"`
	LectureHours      int32  `json:"lecture_hours"`
	LaboratoriesHours int32  `json:"laboratories_hours"`
	PractiseHours     int32  `json:"practise_hours"`
}

func (q *Queries) ListAllGroupCourses(ctx context.Context, groupsID int32) ([]ListAllGroupCoursesRow, error) {
	rows, err := q.db.Query(ctx, listAllGroupCourses, groupsID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListAllGroupCoursesRow{}
	for rows.Next() {
		var i ListAllGroupCoursesRow
		if err := rows.Scan(
			&i.CourseName,
			&i.GroupsID,
			&i.ID,
			&i.Code,
			&i.Number,
			&i.Name,
			&i.ID_2,
			&i.Name_2,
			&i.LectureHours,
			&i.LaboratoriesHours,
			&i.PractiseHours,
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