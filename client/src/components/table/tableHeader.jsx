import React from "react";

const TableHeader = ({ columns }) => {

    return (
        <thead>
            <tr>
                {Object.keys(columns).map(column => (
                    <th>
                        {columns[column].name}
                    </th>
                ))}
            </tr>
        </thead>
    )
}

export default TableHeader;
// <thead>
// <tr>
//     <th>ID</th>
//     <th>Имя</th>
//     <th>Кафедра</th>
//     <th>Должность</th>
//     <th>Условия содержания</th>
// </tr>
// </thead>