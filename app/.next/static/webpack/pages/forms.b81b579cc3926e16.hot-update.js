"use strict";
/*
 * ATTENTION: An "eval-source-map" devtool has been used.
 * This devtool is neither made for production nor for readable output files.
 * It uses "eval()" calls to create a separate source file with attached SourceMaps in the browser devtools.
 * If you are trying to read the output file, select a different devtool (https://webpack.js.org/configuration/devtool/)
 * or disable the default devtool with "devtool: false".
 * If you are looking for production-ready output files, see mode: "production" (https://webpack.js.org/configuration/mode/).
 */
self["webpackHotUpdate_N_E"]("pages/forms",{

/***/ "./pages/forms.tsx":
/*!*************************!*\
  !*** ./pages/forms.tsx ***!
  \*************************/
/***/ (function(module, __webpack_exports__, __webpack_require__) {

eval(__webpack_require__.ts("__webpack_require__.r(__webpack_exports__);\n/* harmony import */ var react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! react/jsx-dev-runtime */ \"./node_modules/react/jsx-dev-runtime.js\");\n/* harmony import */ var react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__);\n/* harmony import */ var _jsonforms_react__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! @jsonforms/react */ \"./node_modules/@jsonforms/react/lib/index.js\");\n/* harmony import */ var _jsonforms_react__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(_jsonforms_react__WEBPACK_IMPORTED_MODULE_1__);\n/* harmony import */ var _jsonforms_material_renderers__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! @jsonforms/material-renderers */ \"./node_modules/@jsonforms/material-renderers/lib/index.js\");\n/* harmony import */ var _jsonforms_material_renderers__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(_jsonforms_material_renderers__WEBPACK_IMPORTED_MODULE_2__);\nvar _this = undefined;\n\n\n\nvar schema = {\n    \"type\": \"array\",\n    \"items\": {\n        \"type\": \"object\",\n        \"properties\": {\n            \"follower\": {\n                \"type\": \"object\",\n                \"properties\": {\n                    \"accountId\": {\n                        \"type\": \"string\"\n                    },\n                    \"userLink\": {\n                        \"type\": \"string\"\n                    }\n                },\n                \"required\": [\n                    \"accountId\",\n                    \"userLink\"\n                ]\n            }\n        },\n        \"required\": [\n            \"follower\"\n        ]\n    }\n};\nvar Forms = function() {\n    return /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"div\", {\n        children: /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_jsonforms_react__WEBPACK_IMPORTED_MODULE_1__.JsonForms, {\n            schema: schema,\n            // uischema={uischema}\n            data: [],\n            renderers: _jsonforms_material_renderers__WEBPACK_IMPORTED_MODULE_2__.materialRenderers,\n            cells: _jsonforms_material_renderers__WEBPACK_IMPORTED_MODULE_2__.materialCells\n        }, void 0, false, {\n            fileName: \"/Users/nsesti/basin-node/app/pages/forms.tsx\",\n            lineNumber: 31,\n            columnNumber: 9\n        }, _this)\n    }, void 0, false, {\n        fileName: \"/Users/nsesti/basin-node/app/pages/forms.tsx\",\n        lineNumber: 30,\n        columnNumber: 5\n    }, _this);\n};\n_c = Forms;\n/* harmony default export */ __webpack_exports__[\"default\"] = (Forms);\nvar _c;\n$RefreshReg$(_c, \"Forms\");\n\n\n;\n    // Wrapped in an IIFE to avoid polluting the global scope\n    ;\n    (function () {\n        var _a, _b;\n        // Legacy CSS implementations will `eval` browser code in a Node.js context\n        // to extract CSS. For backwards compatibility, we need to check we're in a\n        // browser context before continuing.\n        if (typeof self !== 'undefined' &&\n            // AMP / No-JS mode does not inject these helpers:\n            '$RefreshHelpers$' in self) {\n            // @ts-ignore __webpack_module__ is global\n            var currentExports = module.exports;\n            // @ts-ignore __webpack_module__ is global\n            var prevExports = (_b = (_a = module.hot.data) === null || _a === void 0 ? void 0 : _a.prevExports) !== null && _b !== void 0 ? _b : null;\n            // This cannot happen in MainTemplate because the exports mismatch between\n            // templating and execution.\n            self.$RefreshHelpers$.registerExportsForReactRefresh(currentExports, module.id);\n            // A module can be accepted automatically based on its exports, e.g. when\n            // it is a Refresh Boundary.\n            if (self.$RefreshHelpers$.isReactRefreshBoundary(currentExports)) {\n                // Save the previous exports on update so we can compare the boundary\n                // signatures.\n                module.hot.dispose(function (data) {\n                    data.prevExports = currentExports;\n                });\n                // Unconditionally accept an update to this module, we'll check if it's\n                // still a Refresh Boundary later.\n                // @ts-ignore importMeta is replaced in the loader\n                module.hot.accept();\n                // This field is set when the previous version of this module was a\n                // Refresh Boundary, letting us know we need to check for invalidation or\n                // enqueue an update.\n                if (prevExports !== null) {\n                    // A boundary can become ineligible if its exports are incompatible\n                    // with the previous exports.\n                    //\n                    // For example, if you add/remove/change exports, we'll want to\n                    // re-execute the importing modules, and force those components to\n                    // re-render. Similarly, if you convert a class component to a\n                    // function, we want to invalidate the boundary.\n                    if (self.$RefreshHelpers$.shouldInvalidateReactRefreshBoundary(prevExports, currentExports)) {\n                        module.hot.invalidate();\n                    }\n                    else {\n                        self.$RefreshHelpers$.scheduleUpdate();\n                    }\n                }\n            }\n            else {\n                // Since we just executed the code for the module, it's possible that the\n                // new exports made it ineligible for being a boundary.\n                // We only care about the case when we were _previously_ a boundary,\n                // because we already accepted this update (accidental side effect).\n                var isNoLongerABoundary = prevExports !== null;\n                if (isNoLongerABoundary) {\n                    module.hot.invalidate();\n                }\n            }\n        }\n    })();\n//# sourceURL=[module]\n//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiLi9wYWdlcy9mb3Jtcy50c3guanMiLCJtYXBwaW5ncyI6Ijs7Ozs7OztBQUFBOztBQUUwQztBQUNvQztBQUU5RSxJQUFNRyxNQUFNLEdBQUc7SUFDWCxNQUFNLEVBQUUsT0FBTztJQUNmLE9BQU8sRUFBRTtRQUNMLE1BQU0sRUFBRSxRQUFRO1FBQ2hCLFlBQVksRUFBRTtZQUNWLFVBQVUsRUFBRTtnQkFDUixNQUFNLEVBQUUsUUFBUTtnQkFDaEIsWUFBWSxFQUFFO29CQUNWLFdBQVcsRUFBRTt3QkFDVCxNQUFNLEVBQUUsUUFBUTtxQkFDbkI7b0JBQ0QsVUFBVSxFQUFFO3dCQUNSLE1BQU0sRUFBRyxRQUFRO3FCQUNwQjtpQkFDSjtnQkFDRCxVQUFVLEVBQUU7b0JBQUMsV0FBVztvQkFBRSxVQUFVO2lCQUFDO2FBQ3hDO1NBQ0o7UUFDRCxVQUFVLEVBQUU7WUFBQyxVQUFVO1NBQUM7S0FDM0I7Q0FDSjtBQUVELElBQU1DLEtBQUssR0FBYSxXQUFNO0lBQzVCLHFCQUNFLDhEQUFDQyxLQUFHO2tCQUNBLDRFQUFDTCx1REFBUztZQUNORyxNQUFNLEVBQUVBLE1BQU07WUFDZCxzQkFBc0I7WUFDdEJHLElBQUksRUFBRSxFQUFFO1lBQ1JDLFNBQVMsRUFBRU4sNEVBQWlCO1lBQzVCTyxLQUFLLEVBQUVOLHdFQUFhOzs7OztpQkFFdEI7Ozs7O2FBQ0EsQ0FDUDtDQUNGO0FBYktFLEtBQUFBLEtBQUs7QUFlWCwrREFBZUEsS0FBSyIsInNvdXJjZXMiOlsid2VicGFjazovL19OX0UvLi9wYWdlcy9mb3Jtcy50c3g/ZTJhOSJdLCJzb3VyY2VzQ29udGVudCI6WyJpbXBvcnQgdHlwZSB7IE5leHRQYWdlIH0gZnJvbSAnbmV4dCdcbmltcG9ydCBzdHlsZXMgZnJvbSAnLi4vc3R5bGVzL0hvbWUubW9kdWxlLmNzcydcbmltcG9ydCB7SnNvbkZvcm1zfSBmcm9tIFwiQGpzb25mb3Jtcy9yZWFjdFwiXG5pbXBvcnQge21hdGVyaWFsUmVuZGVyZXJzLCBtYXRlcmlhbENlbGxzfSBmcm9tIFwiQGpzb25mb3Jtcy9tYXRlcmlhbC1yZW5kZXJlcnNcIlxuXG5jb25zdCBzY2hlbWEgPSB7XG4gICAgXCJ0eXBlXCI6IFwiYXJyYXlcIixcbiAgICBcIml0ZW1zXCI6IHtcbiAgICAgICAgXCJ0eXBlXCI6IFwib2JqZWN0XCIsXG4gICAgICAgIFwicHJvcGVydGllc1wiOiB7XG4gICAgICAgICAgICBcImZvbGxvd2VyXCI6IHtcbiAgICAgICAgICAgICAgICBcInR5cGVcIjogXCJvYmplY3RcIixcbiAgICAgICAgICAgICAgICBcInByb3BlcnRpZXNcIjoge1xuICAgICAgICAgICAgICAgICAgICBcImFjY291bnRJZFwiOiB7XG4gICAgICAgICAgICAgICAgICAgICAgICBcInR5cGVcIjogXCJzdHJpbmdcIlxuICAgICAgICAgICAgICAgICAgICB9LFxuICAgICAgICAgICAgICAgICAgICBcInVzZXJMaW5rXCI6IHtcbiAgICAgICAgICAgICAgICAgICAgICAgIFwidHlwZVwiIDogXCJzdHJpbmdcIlxuICAgICAgICAgICAgICAgICAgICB9XG4gICAgICAgICAgICAgICAgfSxcbiAgICAgICAgICAgICAgICBcInJlcXVpcmVkXCI6IFtcImFjY291bnRJZFwiLCBcInVzZXJMaW5rXCJdXG4gICAgICAgICAgICB9XG4gICAgICAgIH0sXG4gICAgICAgIFwicmVxdWlyZWRcIjogW1wiZm9sbG93ZXJcIl1cbiAgICB9XG59O1xuXG5jb25zdCBGb3JtczogTmV4dFBhZ2UgPSAoKSA9PiB7XG4gIHJldHVybiAoXG4gICAgPGRpdj5cbiAgICAgICAgPEpzb25Gb3Jtc1xuICAgICAgICAgICAgc2NoZW1hPXtzY2hlbWF9XG4gICAgICAgICAgICAvLyB1aXNjaGVtYT17dWlzY2hlbWF9XG4gICAgICAgICAgICBkYXRhPXtbXX1cbiAgICAgICAgICAgIHJlbmRlcmVycz17bWF0ZXJpYWxSZW5kZXJlcnN9XG4gICAgICAgICAgICBjZWxscz17bWF0ZXJpYWxDZWxsc31cbiAgICAgICAgICAgIC8vIG9uQ2hhbmdlPXsoeyBlcnJvcnMsIGRhdGEgfSkgPT4gc2V0RGF0YShkYXRhKX1cbiAgICAgICAgLz5cbiAgICA8L2Rpdj5cbiAgKVxufVxuXG5leHBvcnQgZGVmYXVsdCBGb3Jtc1xuIl0sIm5hbWVzIjpbIkpzb25Gb3JtcyIsIm1hdGVyaWFsUmVuZGVyZXJzIiwibWF0ZXJpYWxDZWxscyIsInNjaGVtYSIsIkZvcm1zIiwiZGl2IiwiZGF0YSIsInJlbmRlcmVycyIsImNlbGxzIl0sInNvdXJjZVJvb3QiOiIifQ==\n//# sourceURL=webpack-internal:///./pages/forms.tsx\n"));

/***/ })

});