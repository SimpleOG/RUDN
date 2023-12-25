import React from "react";

const TableBody = ({ data, columns }) => {
    console.log(data);
    console.log(columns);
    return (
        <tbody id="tbody">
        {data.map((item) => (
            <tr>
                {Object.keys(columns).map((column) => (
                    <td>
                        {item[columns[column].path]}

                    </td>

                ))}
            </tr>
        ))}
        </tbody>
    )
}

export default TableBody;