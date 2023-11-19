import React, { useEffect, useState } from "react";
import Table from "../table/table";
import { paginate } from "../../utils/paginate";
import { Pagination } from "../pagination/pagination";

const Teachers = () => {

    console.log("rendered");
    const [teachers, setTeachers] = useState([]);
    const [currentPage, setCurrentPage] = useState(1);

    const pageSize = 10;

    const handlePageChange = (pageIndex) => {
        setCurrentPage(pageIndex);
    };


    useEffect( () => {
        const fetchData = async () => {
            const data = await fetch("http://localhost:8080/all_teachers");
            const response = await data.json()
            setTeachers(response);
        }

        fetchData();
    }, [])


    const count = teachers.length;
    const teacherCrop = paginate(teachers, currentPage, pageSize);

    return (
        teachers.length > 0
            ?
            <>
                <Table data={teacherCrop}/>
                <Pagination
                    itemsCount={count}
                    pageSize={pageSize}
                    currentPage={currentPage}
                    onPageChange={handlePageChange}
                />
            </>
             :
            "Loading..."
    );
};

export default Teachers;
