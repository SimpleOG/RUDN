import React from "react";
import {Link} from "react-router-dom";
import "./navbar.css"

const Navbar = () => {
    return (
        <nav className="navbar">

            <div className="navbar__main">
                <ul>
                    <li>
                        <Link to={"/"}>
                            Главная страница
                        </Link>
                    </li>
                </ul>
            </div>

            <div className="navbar__secondary">
                <ul>
                    <li>
                        <Link to={"/teachers"}>
                            Преподаватели
                        </Link>
                    </li>
                    <li>
                        <Link to={"/groups"}>
                            Группы
                        </Link>
                    </li>
                    <li>
                        <Link to={"/courses"}>
                            Курсы
                        </Link>
                    </li>
                </ul>
            </div>
        </nav>
    )
};

export default Navbar;