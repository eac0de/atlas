import { React } from "react";
import styled from "styled-components";
import { Link, useLocation } from "react-router-dom";
const Main = styled.aside`
  display: flex;
  flex-direction: column;
  width: 55px;
  height: 100%;
  flex-shrink: 0;
  background-color: #43506c;
`;

const Item = styled(Link)`
  text-decoration: none;
  display: flex;
  align-items: center;
  justify-content: center;
  height: 7%;
  color: #e9e9eb;
  background-color: ${({ $isActive }) => ($isActive ? "#e9e9eb" : "#43506c")};
  color: ${({ $isActive }) => ($isActive ? "#43506c" : "#e9e9eb")};
  cursor: pointer;
`;

export default function Menu() {
  const loc = useLocation();
  return (
    <Main>
      <div style={{ height: "10%" }}></div>
      <Item to={"/"} className="material-symbols-outlined" $isActive={loc.pathname === "/"}>
        home
      </Item>
      <Item to="/tasks" className="material-symbols-outlined" $isActive={loc.pathname === "/tasks"}>
        task
      </Item>
      <Item to="/projects" className="material-symbols-outlined" $isActive={loc.pathname === "/projects"}>
        emoji_food_beverage
      </Item>
      <Item to="/haha" className="material-symbols-outlined" $isActive={loc.pathname === "/haha"}>
        browse_activity
      </Item>
    </Main>
  );
}
