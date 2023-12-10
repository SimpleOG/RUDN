import React, {useEffect, useState} from "react";
import Table from "../table/table";
import Pagination from "../pagination/pagination";
import {paginate} from "../../utils/paginate";
import {getFields} from "../../utils/getFields";
import TeachersTable from "./teachersTable";

const Teachers = () => {

    console.log("rendered");
    const [teachers, setTeachers] = useState([]);
    const [currentPage, setCurrentPage] = useState(1);
    // const [teacherColumn, setTeacherColumn] = useState({});

    const pageSize = 10;
    let teacherColumns;

    const handlePageChange = (pageIndex) => {
        setCurrentPage(pageIndex);
    };


    useEffect(() => {
        const fetchData = async () => {
            console.log("хуй");
            const data = await fetch("http://localhost:8080/teachers");
            const response = await data.json()
            setTeachers(response);
            response.info
            teacherColumns = getFields(response);
            console.log(teacherColumns);


        }

        fetchData();
    }, [])


    const count = teachers.length;
    const teacherCrop = paginate(teachers, currentPage, pageSize);

    return (
        teachers.length > 0
            ?
            <>
                <TeachersTable teachers={teacherCrop}/>
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
