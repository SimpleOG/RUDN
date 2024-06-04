import React from "react";
import { render } from "react-dom";
import App from "./App/App";
import {BrowserRouter} from "react-router-dom";

render(
    <BrowserRouter>
        <App />
    </BrowserRouter>,

    document.getElementById("root")
)




// const tbody = document.getElementById("tbody");
// const clearButton = document.getElementById("clear");
//
//
// async function fetchTeachers(URL) {
//     const data = await fetch(URL);
//     return data
// }
//
//
// function createNode(teacher) {
//     const tableRow = document.createElement("tr");
//
//     const validKey = {
//         id: true,
//         full_name: true,
//         department: true
//     }
//
//     for (let key in validKey) {
//         const tableData = document.createElement("td");
//         tableData.innerText = teacher[key];
//         tableRow.append(tableData);
//     }
//
//     return tableRow;
// }
//
//
// fetchTeachers("https://jsonplaceholder.typicode.com/users")
//     .then(async (data) => {
//         const teachers = await data.json();
//         renderTeachers(teachers);
//     })
//
// function clearChild(node) {
//     node.innerHTML = "";
// }
//
// clearButton.onclick = () => {
//     clearChild(tbody);
// }
//
//
// function renderTeachers(teacherSlice) {
//     for (let i = 0; i < teachers.length; i++) {
//         tbody.append(createNode(teachers[i]));
//     }
// }









