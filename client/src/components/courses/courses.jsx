import Pagination from "../pagination/pagination";
import React, {useEffect, useState} from "react";

import {paginate} from "../../utils/paginate";
import CoursesTable from "./coursesTable";

const Courses = ({ name ,type}) => {

    const [courses, setCourses] = useState([]);
    const [currentPage, setCurrentPage] = useState(1);

    const pageSize = 10;

    const handlePageChange = (pageIndex) => {
        setCurrentPage(pageIndex);
    };


    useEffect(() => {
        const fetchData = async () => {
            let s=`http://localhost:8080/course/${name}?sem_type=${type}`
            const data = await fetch(s);
            console.log(s)
            const response = await data.json()
            console.log(response);
            setCourses(response);
        }

        fetchData();
    }, [type])


    const count = courses.length;
    const coursesCrop = paginate(courses, currentPage, pageSize);

    return (
        (
            courses.length > 0
                ?
                <>
                    <CoursesTable courses={coursesCrop}/>
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

export default Courses;