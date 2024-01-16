import React, {useState} from "react";
import "./teacherForm.css";
import {CheckBoxForm} from "../shared/ui/form/checkBoxForm/checkBoxForm";

const TeacherForm = ({ data, onChange, onSubmit, onClick }) => {


    const cropAmount = Math.ceil(Object.keys(data).length / 5);
    let tempAr = [];
    for (let i = 0; i < cropAmount; i++) {
        tempAr.push(i);
    }

    const dataCrop = (data, cropSize, currIndex) => {
        const startIndex = currIndex * cropSize;
        return [...data].splice(startIndex, cropSize);
    }

    return (
        <form onSubmit={onSubmit}>
            <div className="checkBoxesContainer">
                {tempAr.map((index) => (
                    <div className="checkBoxItem">
                        {(dataCrop(Object.keys(data), 5, index)).map((key) => (<CheckBoxForm
                            name={key}
                            value={data[key]}
                            onChange={onChange}
                        />))}
                    </div>
                ))}
            </div>
            <div className="sendFormButtons">
                <button
                    className="select_button"
                    type="submit"
                    name="pdf"
                    onClick={onClick}
                >
                    Save pdf
                </button>
                <button
                    className="select_button"
                    type="submit"
                    name="doc"
                    onClick={onClick}
                >
                    Send doc
                </button>
            </div>
        </form>
    )
}

export default TeacherForm;