import React, {useState} from "react";
import "./teacherPage.css";
import Courses from "../courses/courses";
import Groups from "../groups/groups";
import {Modal} from "../shared/ui/modal/modal";
import TeacherForm from "../teacherForm/teacherForm";

const TeacherPage = ({ name }) => {
    const [select, setSelect] = useState("course");

    const [isOpened, setIsOpened] = useState(false);

    const [format, setFormat] = useState(null);

    const [data, setData] = useState({
        ["sosi hui mrazota ebanaya mat ebal"]: false,
        d: false,
        e: false,
        f: false,
        g: false,
        h: false,
        i: false,
        j: false,
        k: false,
        l: false,
        m: false,
        n: false,
        o: false,
        p: false,
        q: false,
        r: false,
        s: false,
        t: false,
        u: false,
        v: false,
        w: false,
        x: false,
        y: false,
        z: false,
        aa: false,
        ab: false,
        ac: false,
        ad: false,
        ae: false,
        af: false,
        ag: false,
        ah: false,
        ai: false,
        aj: false,
        ak: false,
        al: false,
        am: false,
        an: false,
        ao: false,
        ap: false,
        aq: false,
        ar: false,
        as: false,
        at: false,
        au: false,
        av: false,
        aw: false,
        ax: false,
        ay: false,
        az: false,
        ba: false,
        bb: false,
        bc: false,
        bd: false,
        be: false
    })

    const selectCourse = () => {
        setSelect("course");
    }

    const selectGroup = () => {
        setSelect("group");
    }

    const onOpen = () => {
        setIsOpened(true);
    }

    const onClose = () => {
        setIsOpened(false);
    }

    const handleClickSubmitButton = (e) => {
        setFormat(e.target.name);
    }

    const handleChange = (target) => {
        setData((prevState) => ({
            ...prevState,
            [target.name]: target.value
        }));
    }

    const handleSubmit = async (e) => {
        e.preventDefault();
        console.log(data);
        const payload = Object.keys(data).filter((key) => Boolean(data[key]));
        console.log(payload);
        const response = await fetch(`http://localhost:8080/getWordFile/${name}`,
            {
                mode: 'no-cors',
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json;charset=utf-8'
                },
                body: JSON.stringify(payload)
            },
        )
        console.log(response);
    }

    return (
        <div className="teacher_container">
            <Modal
                className={isOpened ? "opened" : null}
                onClose={onClose}
            >
                <TeacherForm data={data} onChange={handleChange} onSubmit={handleSubmit} onClick={handleClickSubmitButton}/>
            </Modal>
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
                    <button
                        className="select_button download_button"
                        onClick={onOpen}
                    >
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