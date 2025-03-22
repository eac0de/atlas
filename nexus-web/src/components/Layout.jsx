import { React } from "react";
import Header from "./Header";
import Menu from "./Menu";
import { Navigate, Outlet } from "react-router-dom";
import styled from "styled-components";

const Main = styled.main`
  display: flex;
  flex-direction: row;
  width: 100%;
  height: 100%;
  overflow: hidden;
`;

const AddTaskButton = styled.button`
  position: fixed; /* Фиксированное позиционирование относительно окна браузера */
  bottom: 30px; /* Отступ от нижнего края */
  right: 30px; /* Отступ от правого края */
  z-index: 1000; /* Достаточно высокий z-index, чтобы быть поверх других элементов */

  background-color: #43506c;
  color: #fff;
  border: none;
  border-radius: 50%; /* Круглая кнопка */
  width: 50px;
  height: 50px;
  font-size: 32px;

  display: flex;
  align-items: center;
  justify-content: center;

  cursor: pointer;
  box-shadow: 0px 2px 8px rgba(0, 0, 0, 0.2);
`;

export default function Layout() {
  const token = localStorage.getItem("nexus_access_token");
  if (!token) return <Navigate to="/login" />;
  return (
    <>
      <Header />
      <Main>
        <Menu />
        <Outlet />
        <AddTaskButton onClick={() => alert("Add task")}>+</AddTaskButton>
      </Main>
    </>
  );
}
