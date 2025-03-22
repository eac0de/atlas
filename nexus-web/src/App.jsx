import { React } from "react";
import Auth from "./components/Auth";
import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";
import Layout from "./components/Layout";
import HomePage from "./components/HomePage";
import styled from "styled-components";
import ProjectsPage from "./components/ProjectsPage";
import TasksPage from "./components/TasksPage";

const Main = styled.main`
  height: 100vh;
  width: 100vw;
  display: flex;
  flex-direction: column;
  background-color: #e9e9eb;
  color: #43506c;
`;

export default function App() {
  return (
    <Main>
      <BrowserRouter>
        <Routes>
          <Route path="/login" element={<Auth />} />
          <Route element={<Layout />}>
            <Route path="/" element={<HomePage />} />
            <Route path="/tasks" element={<TasksPage />} />
            <Route path="/projects" element={<ProjectsPage />} />
            <Route path="/haha" element={<div>haha</div>} />
            <Route path="*" element={<Navigate to="/" replace />} />
          </Route>
        </Routes>
      </BrowserRouter>
    </Main>
  );
}
