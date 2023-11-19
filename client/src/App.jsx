import React from 'react';
import {Link, Route, Routes} from "react-router-dom";
import HomePage from "./components/homepage/homepage";
import Teachers from "./components/teacher/teachers";

const App = () => {
    return (
        <div className="container">
            <Routes>
                <Route path={"/"} element={<HomePage/>}/>
                <Route path={"/all_teachers"} element={<Teachers/>}/>
            </Routes>
        </div>
    );

};
export default App
