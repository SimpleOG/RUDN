import Pagination from "../pagination/pagination";
import React, {useEffect, useState} from "react";

import {paginate} from "../../utils/paginate";
import GroupsTable from "./groupsTable";

const Groups = ({ name }) => {

    const [groups, setgroups] = useState([]);
    const [currentPage, setCurrentPage] = useState(1);

    const pageSize = 10;

    const handlePageChange = (pageIndex) => {
        setCurrentPage(pageIndex);
    };


    useEffect(() => {
        const fetchData = async () => {
            const data = await fetch(`http://localhost:8080/groups/${name}`);
            const response = await data.json()
            console.log(response);
            setgroups(response);
        }

        fetchData();
    }, [])


    const count = groups.length;
    const groupsCrop = paginate(groups, currentPage, pageSize);

    return (
        (
            groups.length > 0
                ?
                <>
                    <GroupsTable groups={groupsCrop}/>
                    <Pagination
                        itemsCount={count}
                        pageSize={pageSize}
                        currentPage={currentPage}
                        onPageChange={handlePageChange}
                    />
                </>
                :
                "Loading..."
        )
    )
}

export default Groups;