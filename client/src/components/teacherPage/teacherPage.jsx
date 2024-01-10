import React, {useState} from "react";
import "./teacherPage.css";
import Courses from "../courses/courses";
import Groups from "../groups/groups";

const TeacherPage = ({ name }) => {
    const [select, setSelect] = useState("course");

    const selectCourse = () => {
        setSelect("course");
    }

    const selectGroup = () => {
        setSelect("group");
    }

    return (
        <div className="teacher_container">
            <div className="side_information">
                <div className="teacher_profile">
                    <h1>{name}</h1>
                </div>

                <div className="toggle_buttons">
                    <button
                        className={`select_button ${select === "course" ? "active" : null}`}
                        onClick={selectCourse}
                    >
                        Курсы
                    </button>
                    <button
                        className={`select_button ${select === "group" ? "active" : null}`}
                        onClick={selectGroup}
                    >
                        Группы
                    </button>
                </div>
                <div className="save_button">
                    <button className="select_button download_button">
                        Скачать информацию
                    </button>
                </div>
            </div>

            <div className="extra_information">
                {select === "course" ? <Courses name={name} /> : <Groups name={name} /> }
            </div>
        </div>
    )
}

export default TeacherPage;