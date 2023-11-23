import React from "react";
import Table from "../table/table";

const GroupsTable = ({ groups }) => {
    const groupsColumns = {
        id: {
            name: "ID",
            path: "id"
        },
        code: {
            name: "Код",
            path: "code"
        },
        number: {
            name: "Номер",
            path: "number"
        },
        name: {
            name: "Название",
            path: "name"
        }
    }

    return (
        <Table data={groups} columns={groupsColumns} />
    )
}

export default GroupsTable;