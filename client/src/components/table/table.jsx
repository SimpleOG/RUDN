import React from "react";
import "./style.css"
import TableHeader from "./tableHeader";
import TableBody from "./tableBody";

const Table = ({ data, columns }) => {
    // const validKey = {
    //     id: true,
    //     full_name: true,
    //     department: true,
    //     post: true,
    //     conditions: true
    // }
    return (
        <>
            <div className="table_container">
                <table className="table_dark">
                    <TableHeader columns={columns}/>
                    <TableBody data={data} columns={columns}/>
                    {/*<thead>*/}
                    {/*<tr>*/}
                    {/*    <th>ID</th>*/}
                    {/*    <th>Имя</th>*/}
                    {/*    <th>Кафедра</th>*/}
                    {/*    <th>Должность</th>*/}
                    {/*    <th>Условия содержания</th>*/}
                    {/*</tr>*/}
                    {/*</thead>*/}
                    {/*<tbody id="tbody">*/}
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
                    {/*</tbody>*/}
                </table>
            </div>
        </>
    )
};

export default Table;
