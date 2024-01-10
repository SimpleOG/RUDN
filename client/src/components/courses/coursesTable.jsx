import React from "react";
import Table from "../table/table";

const CoursesTable = ({ courses }) => {
    const coursesColumns = {
        name: {
            name: "Название курса",
            path: "name_of_the_discipline_or_type_of_academic_work"
        },
        typeofwork: {
            name: "Тип академической работы",
            path: "type_of_educational_work"
        },
        total: {
            name: "Всего часов ",
            path: "total"
        },

    }

    return (
        <Table data={courses} columns={coursesColumns} />
    )
}

export default CoursesTable;