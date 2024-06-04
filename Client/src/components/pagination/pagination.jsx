import React from "react";
import _ from "lodash";
import "./pagination.css";
import {checkPage} from "../../utils/checkPage";
// teachers, pageSize, func, current
// 1 c-2 c-1 current c+1 c+2 end
const Pagination = ({ itemsCount, pageSize, onPageChange, currentPage }) => {
    const pageCount = Math.ceil(itemsCount / pageSize);
    // const pageCount = Math.ceil(itemsCount / pageSize);
    if (pageCount === 1) return null;
    // const pages = _.range(1, pageCount + 1);
    const fixedPages = _.range(1, 8);
    const allPages = _.range(1, pageCount + 1);
    return (
        <nav>
            <ul className="pagination">
                {pageCount >= 7
                    ? allPages.map(page => checkPage(currentPage, page, pageCount) ? <li key={`page_` + page}><a className={currentPage === page ? "active" : ""} onClick={() => onPageChange(page)}>{page}</a></li> : null)
                    : allPages.map(page => <li key={`page_` + page}><a className={currentPage === page ? "active" : ""} onClick={() => onPageChange(page)}>{page}</a></li>)
                }
            </ul>
        </nav>
    );
};

export default Pagination;

// {pages.map(page => <li key={`page_` + page}><a className={currentPage === page ? "active" : ""} onClick={() => onPageChange(page)}>{page}</a></li>)}