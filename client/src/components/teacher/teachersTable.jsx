import React from "react";
import Table from "../table/table";

const TeachersTable = ({ teachers }) => {
    const teacherColumns = {
        full_name: {
            name: "Преподаватель",
            path: "full_name"
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
    }

    return (
        <Table data={teachers} columns={teacherColumns} />
    )
}

export default TeachersTable;