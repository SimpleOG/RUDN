import React from "react";
import "./style.css"
import {Link} from "react-router-dom";

const HomePage = () => {
    return (
        <div className="container">
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
                            <Link to={"/all_teachers"}>
                                Преподаватели
                            </Link>
                        </li>
                        <li>
                            <Link to={"/all_groups"}>
                                Группы
                            </Link>
                        </li>
                        <li>
                            <Link to={"/all_courses"}>
                                Курсы
                            </Link>
                        </li>
                    </ul>
                </div>
            </nav>

            <div className="content">
                <div className="content__secondary">
                    <h4 className="text__h4">Работа с базой данных кафедры</h4>
                </div>
                <div className="content__primary">
                    <a href="https://www.rudn.ru/" target="_blank"><h1> RUDN UNIVERSITY </h1></a>
                </div>
            </div>
        </div>


        // <div className="bgimage">
        //     <div className="menu">
        //
        //         <div className="leftmenu">
        //             <a href="http://localhost:8081/home"><h4 id="firstlist"> Главная страница </h4></a>
        //         </div>
        //         <div className="rightmenu">
        //             <ul>
        //                 <a href="http://localhost:8080/all_teachers">
        //                     <li>Преподаватели</li>
        //                 </a>
        //                 <a href="http://localhost:8080/all_groups">
        //                     <li>Группы</li>
        //                 </a>
        //                 <a href="http://localhost:8080/all_courses">
        //                     <li>Курсы</li>
        //                 </a>
        //             </ul>
        //         </div>
        //
        //     </div>
        //
        //     <div className="text">
        //         <h4 className="text__h4">Работа с базой данных кафедры</h4>
        //         <a href="https://www.rudn.ru/" target="_blank"><h1> RUDN UNIVERSITY </h1></a>
        //     </div>
        //
        // </div>
    )
}
export default HomePage;