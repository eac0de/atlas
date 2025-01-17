import { useState } from "react";
import reactLogo from "./assets/react.svg";
import "./App.css";
import Auth from "./components/Auth/Auth";
import AuthProvider from "./AuthContext";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import PrivateRoute from "./PrivateRoute";

export default function App() {
  return (
    <div className="App">
      <Router>
        <AuthProvider>
          <Routes>
            <Route path="/login" element={<Auth />} />
            <Route element={<PrivateRoute />}>
              <Route path="/" element={<div>Dashboard</div>} />
            </Route>
          </Routes>
        </AuthProvider>
      </Router>
    </div>
  );
}
