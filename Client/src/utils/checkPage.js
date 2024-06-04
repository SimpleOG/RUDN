export function checkPage(currentPage, page, lastPage) {
    if (page === 1 || page === lastPage || page === currentPage || page - 1 === currentPage || page - 2 === currentPage || page + 1 === currentPage || page + 2 === currentPage) {
        return true;
    }
    return false;
}

//                             ? <li key={`page_` + page}><a className={currentPage === page ? "active" : ""} onClick={() => onPageChange(page)}>{page}</a></li>)