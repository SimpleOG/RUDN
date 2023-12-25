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
                </table>
            </div>
        </>
    )
};

export default Table;
