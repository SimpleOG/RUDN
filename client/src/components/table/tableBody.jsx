import React from "react";

const TableBody = ({ data, columns }) => {
    function renderData(item, column) {
        if (columns[column].component) {
            const component = columns[column].component;
            return component(item);
        }
        return item[columns[column].path];
    }
    return (
        <tbody id="tbody">
        {data.map((item) => (
            <tr>
                {Object.keys(columns).map((column) => (
                    <td>
                        {renderData(item, column)}
                        {/*{item[columns[column].path]}*/}
                    </td>

                ))}
            </tr>
        ))}
        </tbody>
    )
}

export default TableBody;