import React, {useEffect, useState} from "react";
import Pagination from "../pagination/pagination";
import {paginate} from "../../utils/paginate";
import {getFields} from "../../utils/getFields";
import TeachersTable from "./teachersTable";
import "./teachers.css"

const Teachers = () => {

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
            const data = await fetch("http://localhost:8080/teachers");
            const response = await data.json()
            console.log(response)
            setTeachers(response);
            response.info
            teacherColumns = getFields(response);
            console.log(teacherColumns);        }
        fetchData();
    }, [])


    const count = teachers.length;
    const teacherCrop = paginate(teachers, currentPage, pageSize);

    return (
        teachers.length > 0
            ?
            <div className="temp2">
                <TeachersTable teachers={teacherCrop}/>
                <Pagination
                    itemsCount={count}
                    pageSize={pageSize}
                    currentPage={currentPage}
                    onPageChange={handlePageChange}
                />
            </div>
            :
            "Loading..."
    );
};

export default Teachers;
