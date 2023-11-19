/*
 * ATTENTION: The "eval" devtool has been used (maybe by default in mode: "development").
 * This devtool is neither made for production nor for readable output files.
 * It uses "eval()" calls to create a separate source file in the browser devtools.
 * If you are trying to read the output file, select a different devtool (https://webpack.js.org/configuration/devtool/)
 * or disable the default devtool with "devtool: false".
 * If you are looking for production-ready output files, see mode: "production" (https://webpack.js.org/configuration/mode/).
 */
/******/ (() => { // webpackBootstrap
/******/ 	var __webpack_modules__ = ({

/***/ "./src/index.js":
/*!**********************!*\
  !*** ./src/index.js ***!
  \**********************/
/***/ (() => {

eval("\r\nconst tbody = document.getElementById(\"tbody\");\r\nconst clearButton = document.getElementById(\"clear\");\r\n\r\n\r\nasync function fetchTeachers(URL) {\r\n    const data = await fetch(URL);\r\n    return data\r\n}\r\n\r\n\r\nfunction createNode(teacher) {\r\n    const tableRow = document.createElement(\"tr\");\r\n\r\n    const validKey = {\r\n        id: true,\r\n        full_name: true,\r\n        department: true\r\n    }\r\n\r\n    for (let key in validKey) {\r\n        const tableData = document.createElement(\"td\");\r\n        tableData.innerText = teacher[key];\r\n        tableRow.append(tableData);\r\n    }\r\n\r\n    return tableRow;\r\n}\r\n\r\n\r\nfetchTeachers(\"https://jsonplaceholder.typicode.com/users\")\r\n    .then(async (data) => {\r\n        const teachers = await data.json();\r\n        renderTeachers(teachers);\r\n    })\r\n\r\nfunction clearChild(node) {\r\n    node.innerHTML = \"\";\r\n}\r\n\r\nclearButton.onclick = () => {\r\n    clearChild(tbody);\r\n}\r\n\r\n\r\nfunction renderTeachers(teacherSlice) {\r\n    for (let i = 0; i < teachers.length; i++) {\r\n        tbody.append(createNode(teachers[i]));\r\n    }\r\n}\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\n\n//# sourceURL=webpack://dlyadolbaeba/./src/index.js?");

/***/ })

/******/ 	});
/************************************************************************/
/******/ 	
/******/ 	// startup
/******/ 	// Load entry module and return exports
/******/ 	// This entry module can't be inlined because the eval devtool is used.
/******/ 	var __webpack_exports__ = {};
/******/ 	__webpack_modules__["./src/index.js"]();
/******/ 	
/******/ })()
;