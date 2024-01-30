import React, {useState} from "react";
import "./teacherPage.css";
import Courses from "../courses/courses";
import Groups from "../groups/groups";
import {Modal} from "../shared/ui/modal/modal";
import TeacherForm from "../teacherForm/teacherForm";

const TeacherPage = ({name}) => {
    const [select, setSelect] = useState("course");

    const [isOpened, setIsOpened] = useState(false);

    const [format, setFormat] = useState(null);
    const obj = {
        "Шифр": "TheCodeOfTheOOPRUDN",
        "Код направления": "direction_code",
        "Наименование программы": "name_of_the_program",
        "Блок": "block",
        "Компонента": "component",
        "№ в РУП": "n_v_RUP",
        "доп.инфо": "dop_info",
        "Наименование дисциплины или вида учебной работы": "name_of_the_discipline_or_type_of_academic_work",
        "Семестр ; Модуль": "SemesterOrModule",
        "Недель в семестре (модуле)": "weeks_per_semester_module",
        "Вид учебной работы": "type_of_educational_work",
        "Вид ПА или ГИА": "type_of_PA_or_GIA",
        "Курс. работы": "course_works",
        "Курс. проекты": "course_projects",
        "Уч. пр. (ЗЕ по РУП)": "course_Uch_ave_ZE_on_RUP",
        "Пр. пр. (ЗЕ по РУП)": "pr_ZE_on_RUP",
        "НИР (ЗЕ по РУП)": "NIR_ZE_by_RUP",
        "Код": "string",
        "Номер группы": "string",
        "Подгрупп": "of_groups",
        "Групп": "subgroups",
        "Всего": "total_people",
        "РФ": "RF",
        "ИН": "foreign",
        "Норматив": "standard",
        "Рассчетных": "calculated",
        "ПК": "PK",
        "Кафедра/департамент": "department",
        "должность": "post",
        "условия привлечения ": "terms_of_attraction",
        "Фамилия И.О.  преподавателя": "full_name",
        "Особый признак": "a_special_feature",
        "Лекции":"lectures",
        "Практика / Семинары": "practice_or_Seminars",
        "Лаб. работы / Клинические занятия": "Lab_works_or_Clinical_classes",
        "Текущий контроль": "current_control",
        "Промежуточная аттестация (ПА) по БРС": "interim_certification_PO_for_BRS",
        "Оформление результатов ПА": "registration_of_PA_results",
        "Текущие консультации по дисциплине": "ongoing_consultations_on_the_discipline",
        "Курсовые работы": "course_works",
        "Курсовые проекты": "course_projects",
        "Учебная практика": "educational_practice",
        "Произв., педагогическая и преддипломная практики": "proc_pedagogical_and_pre_graduate_practices",
        "НИР":"NIR",
        "Практики (в т.ч. НИР) цифровых магистратур": "practices_including_research_of_digital_magistracies",
        "Рецензирование рефератов аспирантов": "reviewing_the_abstracts_of_graduate_students",
        "Кандидатский экзамен": "candidates_exam",
        "Научное руководство": "scientific_guidance",
        "Руководство ВКР или НКР, в том числе Организация и сопровождение Первичной аккредитации МИ": "the_leadership_of_the_WRC_or_the_NKR",
        "Рецензирование ВКР": "review_of_the_WRC",
        "ГЭК ": "GEK",
        "ИТОГО": "total",

    }

    const initialState = Object.fromEntries(
        Object.keys(obj).map((key) => [key, false])
    )
    const [data, setData] = useState(
        initialState
    )
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
        console.log(e.target.name);
        setFormat(e.target.name);
    }

    const handleChange = (target) => {
        console.log("сработало handle change", target)
        setData((prevState) => ({
            ...prevState,
            [target.name]: target.value
        }));
    }

    function saveFile(url, filename) {
        console.log("working");
        const a = document.createElement("a");
        a.href = url;
        a.download = filename || "file-name";
        document.body.appendChild(a);
        a.click();
        document.body.removeChild(a);
    }

    const handleSubmit = async (e) => {
        e.preventDefault();
        const pay = Object.values(Object.keys(data).filter((key) => Boolean(data[key])));
        console.log(pay)
        const payload = []
        for (let i = 0; i < pay.length; i++) {
            payload.push(obj[pay[i]])
        }
        console.log(payload)

        const response = await fetch(`http://localhost:8080/getWordFile/${name}`,
            {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json;charset=utf-8'
                },
                body: JSON.stringify({array: payload})
            },
        )

        const responseBlob = await response.blob();
        const blobUrl = URL.createObjectURL(responseBlob);
        saveFile(blobUrl, `${name}.doc`);
    }

    return (
        <div className="teacher_container">
            <Modal
                className={isOpened ? "opened" : null}
                onClose={onClose}
            >
                <TeacherForm data={data} onChange={handleChange} onSubmit={handleSubmit}
                             onClick={handleClickSubmitButton}/>
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
                {select === "course" ? <Courses name={name}/> : <Groups name={name}/>}
            </div>
        </div>
    )
}

export default TeacherPage;