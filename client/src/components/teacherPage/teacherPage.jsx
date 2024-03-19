import React, {useState} from "react";
import "./teacherPage.css";
import Courses from "../courses/courses";
import {Modal} from "../shared/ui/modal/modal";
import TeacherForm from "../teacherForm/teacherForm";

const TeacherPage = ({name}) => {
    const [select, setSelect] = useState("Autumn");

    const [isOpened, setIsOpened] = useState(false);

    const [format, setFormat] = useState(null);
    const obj = {
        "Шифр": "the_code_of_the_oop_rudn",
        "Код направления": "direction_code",
        "Наименование программы": "name_of_the_program",
        "Блок": "block",
        "Компонента": "component",
        "№ в РУП": "n_v_rup",
        "доп.инфо": "dop_info",
        "Наименование дисциплины или вида учебной работы": "name_of_the_discipline_or_type_of_academic_work",
        "Семестр ; Модуль": "semester_or_module",
        "Недель в семестре (модуле)": "weeks_per_semester_module",
        "Вид учебной работы": "type_of_educational_work",
        "Вид ПА или ГИА": "type_of_pa_or_gia",
        "Курс. работы": "kw_course_works",
        "Курс. проекты": "kw_course_projects",
        "Уч. пр. (ЗЕ по РУП)": "course_uch_ave_ze_on_rup",
        "Пр. пр. (ЗЕ по РУП)": "pr_ze_on_rup",
        "НИР (ЗЕ по РУП)": "nir_ze_by_rup",
        "Код": "code",
        "Номер группы": "group_number",
        "Подгрупп": "of_groups",
        "Групп": "subgroups",
        "Всего": "total_people",
        "РФ": "rf",
        "ИН": "foreign",
        "Норматив": "standard",
        "Рассчетных": "calculated",
        "ПК": "pk",
        "Кафедра/департамент": "department",
        "должность": "post",
        "условия привлечения ": "terms_of_attraction",
        "Фамилия И.О.  преподавателя": "full_name",
        "Особый признак": "a_special_feature",
        "Лекции":"lectures",
        "Практика / Семинары": "practice_or_seminars",
        "Лаб. работы / Клинические занятия": "lab_works_or_clinical_classes",
        "Текущий контроль": "current_control",
        "Промежуточная аттестация (ПА) по БРС": "interim_certification_po_for_brs",
        "Оформление результатов ПА": "registration_of_pa_results",
        "Текущие консультации по дисциплине": "ongoing_consultations_on_the_discipline",
        "Курсовые работы": "course_works",
        "Курсовые проекты": "course_projects",
        "Учебная практика": "educational_practice",
        "Произв., педагогическая и преддипломная практики": "proc_pedagogical_and_pre_graduate_practices",
        "НИР":"nir",
        "Практики (в т.ч. НИР) цифровых магистратур": "practices_including_research_of_digital_magistracies",
        "Рецензирование рефератов аспирантов": "reviewing_the_abstracts_of_graduate_students",
        "Кандидатский экзамен": "candidates_exam",
        "Научное руководство": "scientific_guidance",
        "Руководство ВКР или НКР, в том числе Организация и сопровождение Первичной аккредитации МИ": "the_leadership_of_the_wrc_or_the_nkr",
        "Рецензирование ВКР": "review_of_the_wrc",
        "ГЭК ": "gek",
        "ИТОГО": "total",

    }

    const initialState = Object.fromEntries(
        Object.keys(obj).map((key) => [key, false])
    )
    const [data, setData] = useState(
        initialState
    )
    const selectAutumn = () => {
        setSelect("Осенний");
    }

    const selectSpring = () => {
        setSelect("Весенний");
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
        //идёшь по пути который я тебе присылаю
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
                        className={`select_button ${select === "Осенний" ? "active" : null}`}
                        onClick={selectAutumn}
                    >
                        Осенний
                    </button>
                    <button
                        className={`select_button ${select === "Весенний" ? "active" : null}`}
                        onClick={selectSpring}
                    >
                        Весенний
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
                { <Courses name={name} type={select}/> }
            </div>
        </div>
    )
}

export default TeacherPage;