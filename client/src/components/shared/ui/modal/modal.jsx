import React from "react";
import "./modal.css";

export const Modal = ({children, onClose, className}) => {

    const handleClose = (e) => {
        if (e.target !== e.currentTarget) {
            e.stopPropagation();
            return;
        }
        onClose();
    }

    return (
        <div className={`modal ${className}`}>
            <div className="overlay" onClick={handleClose}>
                <div className="contentModal">
                    {children}
                </div>
            </div>
        </div>
    )
}