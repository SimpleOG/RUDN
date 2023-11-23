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
        {/*{data.map((teacher) => {*/}
        {/*    return (*/}
        {/*        <tr>*/}
        {/*            {Object.keys(validKey).map((key) => {*/}
        {/*                if (key === "full_name") {*/}
        {/*                    return (*/}
        {/*                        <td><a href="">{teacher[key]}</a></td>*/}
        {/*                    )*/}
        {/*                }*/}
        {/*                return (*/}
        {/*                    <td>{teacher[key]}</td>*/}
        {/*                )*/}
        {/*            })}*/}
        {/*        </tr>*/}
        {/*    )*/}
        {/*})}*/}
        </tbody>
    )
}

export default TableBody;