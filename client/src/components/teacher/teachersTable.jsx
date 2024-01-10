import React from "react";
import Table from "../table/table";
import {Link} from "react-router-dom";

const TeachersTable = ({ teachers }) => {
    const teacherColumns = {
        full_name: {
            name: "Преподаватель",
            path: "full_name",
            component: (teacher) => (
                <Link
                    to={`${teacher.full_name}`}
                >
                    {teacher.full_name}
                </Link>
            )
        },
        department: {
            name: "Кафедра",
            path: "department"
        },
        post: {
            name: "Должность",
            path: "post"
        },
        terms_of_attraction: {
            name: "Условия привлечения",
            path: "terms_of_attraction"
        },

        lectures: {
            name: "Лекционные часы",
            path: "lectures"
        },
        practice: {
            name: "Часы практики",
            path: "practice"
        },
        labs: {
            name: "Лабораторные часы",
            path: "labs"
        },
        total: {
            name: "Всего",
            path: "total"
        },
    }

    return (
        <Table data={teachers} columns={teacherColumns} />
    )
}

export default TeachersTable;