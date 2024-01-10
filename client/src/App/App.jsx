import React from 'react';
import "./style/index.css"
import {Link, Route, Routes} from "react-router-dom";
import HomePage from "../components/homepage/homepage";
import Teachers from "../components/teacher/teachers";
import Navbar from "../components/navbar/navbar";
import Courses from "../components/courses/courses";

const App = () => {
    return (
            <div className="container">
                <Navbar />
                <Routes>
                    <Route path={"/"} element={<HomePage/>}/>
                    <Route path={"/teachers/:name?"} element={<Teachers/>}/>
                    <Route path={"/groups"} element={<Courses/>}/>
                    <Route path={"/courses"} element={<Teachers/>}/>
                </Routes>
                <div className="block-hidden"></div>
            </div>

    );

};
export default App
