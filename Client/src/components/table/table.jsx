import React from "react";
import "./style.css"
import TableHeader from "./tableHeader";
import TableBody from "./tableBody";

const Table = ({ data, columns }) => {
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
