import React from "react";
import Table from "../table/table";

const TeachersTable = ({ teachers }) => {
    const teacherColumns = {
        id: {
            name: "ID",
            path: "id"
        },
        full_name: {
            name: "Имя",
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
        conditions: {
            name: "Наём",
            path: "conditions"
        }
    }

    return (
        <Table data={teachers} columns={teacherColumns} />
    )
}

export default TeachersTable;