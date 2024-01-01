import React from 'react';
import "./style/index.css"
import {Link, Route, Routes} from "react-router-dom";
import HomePage from "../components/homepage/homepage";
import Teachers from "../components/teacher/teachers";
import Navbar from "../components/navbar/navbar";
import Groups from "../components/groups/groups";

const App = () => {
    return (
            <div className="container">
                <Navbar />
                <Routes>
                    <Route path={"/"} element={<HomePage/>}/>
                    <Route path={"/all_teachers"} element={<Teachers/>}/>
                    <Route path={"/all_groups"} element={<Groups/>}/>
                    <Route path={"/all_courses"} element={<Teachers/>}/>
                </Routes>
                <div className="block-hidden"></div>
            </div>

    );

};
export default App
