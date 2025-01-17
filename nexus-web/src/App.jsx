import { React } from "react";
import "./App.css";
import Auth from "./components/Auth/Auth";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import PrivateRoute from "./PrivateRoute";
import { logout } from "./authService";

export default function App() {
  return (
    <div className="App">
      <BrowserRouter>
        <Routes>
          <Route path="/login" element={<Auth />} />
          <Route element={<PrivateRoute />}>
            <Route
              path="/"
              element={
                <div>
                  <button onClick={logout}>Logout</button>
                </div>
              }
            />
          </Route>
        </Routes>
      </BrowserRouter>
    </div>
  );
}
