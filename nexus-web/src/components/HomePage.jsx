import { React, useState } from "react";
import styled from "styled-components";

const Main = styled.div`
  display: flex;
  flex-direction: row;
  width: 100%;
  height: 100%;
  padding: 1%;
`;

const TasksBlock = styled.div`
  display: flex;
  flex-direction: column;
  max-height: 40%;
  background-color: #e9e9eb;
  box-shadow: 0 0 5px #43506c;
`;

const TasksBlockTitle = styled.div`
  display: flex;
  flex-direction: column;
  box-shadow: 0 3px 10px #43506c52;
  z-index: 1;
`;

const TasksBlockTitleUp = styled.div`
  display: flex;
  width: 100%;
  font-size: 18px;
  justify-content: space-between;
  padding: 15px;
  align-items: center;
`;
const TasksBlockTitleDown = styled.div`
  display: flex;
  font-weight: 600;
  font-size: 16px;
  width: 100%;
  padding-bottom: 5px;
  gap: 10px;
`;

const TasksBlockTitleDownItem = styled.div`
  display: flex;
  align-items: center;
  justify-content: center;
`;

const FilterButton = styled.button`
  background: none;
  border: none;
  font-size: 22px;
  color: inherit;
  cursor: pointer;
  width: fit-content;
`;
const Tasks = styled.ul`
  width: 100%;
  height: 100%;
  background-color: #e9e9eb;
  list-style: none;
  overflow-y: auto;
  padding: 0;
  margin: 0;

  scrollbar-width: none;
  &::-webkit-scrollbar {
    display: none;
  }
`;

const Task = styled.div`
  display: flex;
  font-size: 14px;
  width: 100%;
  padding: 20px 0;
  box-shadow: 0 3px 10px #43506c52;
  gap: 10px;
`;

const TaskItem = styled.div`
  display: flex;
  align-items: center;
  justify-content: center;
`;

const TaskTitle = styled.div`
  display: flex;
  flex-direction: column;
  width: 70%;
  font-size: 14px;
  background-color: #e9e9eb;
`;

export default function HomePage() {
  let my_tasks = [
    {
      id: 6344867,
      project: "❗️Баги и оптимизация (общий проект команды)",
      status: "Решена",
      progress: "100%",
      title: "Адаптация функционала преобразования заявки в обращение после перехода на новый API",
      deadline: "05.03(-1d)",
    },
    {
      id: 2141565,
      project: "❗️Баги и оптимизация (общий проект команды)",
      status: "Новая",
      progress: "50%",
      title: "Переименовать название столбца (Направление - Взимается).",
      deadline: "06.03(today)",
    },
    {
      id: 453395,
      project: "Разработка",
      status: "ТЗ готово",
      progress: "80%",
      title: "Преобразование заявок в обращения",
      deadline: "05.03(-1d)",
    },
    {
      id: 497396,
      project: "Новая разработка",
      status: "Системный анализ",
      progress: "0%",
      title: "Изменить поведение фильтра «Просроченные» в ЖЗ",
      deadline: "03.03(-4d)",
    },
    {
      id: 497395,
      project: "Новая разработка",
      status: "Готово к тестированию",
      progress: "70%",
      title: "Интеграция с обращениями ГИС — Загрузка/выгрузка + файлы",
      deadline: "21.12(256d)",
    },
  ];
  return (
    <Main>
      <TasksBlock>
        <TasksBlockTitle>
          <TasksBlockTitleUp>
            <div>My Tasks</div>
            <FilterButton className="material-symbols-outlined">filter_alt</FilterButton>
          </TasksBlockTitleUp>
          <TasksBlockTitleDown>
            <TasksBlockTitleDownItem style={{ width: "100px" }}>#</TasksBlockTitleDownItem>
            <TasksBlockTitleDownItem style={{ width: "150px", justifyContent: "flex-start" }}>Project</TasksBlockTitleDownItem>
            <TasksBlockTitleDownItem style={{ width: "100px" }}>Status</TasksBlockTitleDownItem>
            <TasksBlockTitleDownItem style={{ width: "120px" }}>Progress</TasksBlockTitleDownItem>
            <TasksBlockTitleDownItem style={{ width: "400px", justifyContent: "flex-start" }}>Title</TasksBlockTitleDownItem>
            <TasksBlockTitleDownItem style={{ width: "100px", justifyContent: "flex-start" }}>Deadline</TasksBlockTitleDownItem>
          </TasksBlockTitleDown>
        </TasksBlockTitle>
        <Tasks>
          {my_tasks.map((task) => (
            <Task>
              <TaskItem style={{ width: "100px" }}>{task.id}</TaskItem>
              <TaskItem style={{ width: "150px", justifyContent: "flex-start" }}>
                <span style={{ whiteSpace: "nowrap", overflow: "hidden", textOverflow: "ellipsis" }}>{task.project}</span>
              </TaskItem>
              <TaskItem style={{ width: "100px" }}>
                <span style={{ whiteSpace: "nowrap", overflow: "hidden", textOverflow: "ellipsis" }}>{task.status}</span>
              </TaskItem>
              <TaskItem style={{ width: "120px" }}>{task.progress}</TaskItem>
              <TaskItem style={{ display: "-webkit-box", width: "400px", justifyContent: "flex-start", WebkitBoxOrient: "vertical", WebkitLineClamp: 2, overflow: "hidden" }}>{task.title}</TaskItem>
              <TaskItem style={{ width: "100px", justifyContent: "flex-start" }}>{task.deadline}</TaskItem>
            </Task>
          ))}
        </Tasks>
        <div style={{ height: "10px" }}></div>
      </TasksBlock>
    </Main>
  );
}
