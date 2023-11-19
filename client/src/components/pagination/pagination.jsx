import React from "react";
import _ from "lodash";
import "./pagination.css";

export const Pagination = ({ itemsCount, pageSize, onPageChange, currentPage }) => {
    const pageCount = Math.ceil(itemsCount / pageSize);
    if (pageCount === 1) return null;
    const pages = _.range(1, pageCount + 1);
    return (
        <nav>
            <ul className="pagination">
                {pages.map(page => <li key={`page_` + page}><a className={currentPage === page ? "active" : ""} onClick={() => onPageChange(page)}>{page}</a></li>)}
            </ul>
        </nav>
    );
};