import Pagination from "../pagination/pagination";
import React, {useEffect, useState} from "react";

import {getFields} from "../../utils/getFields";
import {paginate} from "../../utils/paginate";
import GroupsTable from "./groupsTable";

const Groups = () => {

    const [groups, setGroups] = useState([]);
    const [currentPage, setCurrentPage] = useState(1);
    // const [teacherColumn, setTeacherColumn] = useState({});

    const pageSize = 5;
    let groupsColumns;

    const handlePageChange = (pageIndex) => {
        setCurrentPage(pageIndex);
    };


    useEffect(() => {
        const fetchData = async () => {
            const data = await fetch("http://localhost:8080/all_groups");
            const response = await data.json()
            setGroups(response);

            groupsColumns = getFields(response[0]);

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