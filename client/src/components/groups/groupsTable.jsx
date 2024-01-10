import React from "react";
import Table from "../table/table";

const GroupsTable = ({ groups }) => {
    const groupsColumns = {
        name: {
            name: "Название группы",
            path: "name"
        },
        code: {
            name: "Код",
            path: "code"
        },
        programname: {
            name: "Название программы",
            path: "name_of_the_program"
        },

    }

    return (
        <Table data={groups} columns={groupsColumns} />
    )
}

export default GroupsTable;