import React from "react";
import "./App.css";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import LRUcache from "./LRUcache";

function App() {
  return (
    <div className="App">
      <h1>React App is ready to start works</h1>
      <LRUcache />
    </div>
  );
}

export default App;
