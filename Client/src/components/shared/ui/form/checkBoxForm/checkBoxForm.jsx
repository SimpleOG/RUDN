import React from "react";
import "./checkBoxForm.css";

export const CheckBoxForm = ({ name, value, onChange }) => {
    const className = value ? "checked" : "";
    const handleChange = () => {
        onChange({ name: name, value: !value });
    };

    return (
        <div className="item">
            <p>{name}</p>
            <div className={`toggle-pill-color ${className}`}>
                <input
                    type="checkbox"
                    id={name}
                    name={name}
                    value={value}
                    onChange={handleChange}
                />
                <label htmlFor={name}></label>
            </div>
        </div>
    )
}