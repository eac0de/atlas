import { React, useState } from "react";
import styled from "styled-components";
import { logout } from "../authService";
import iconPath from "../assets/new_atlas.svg";
import accountIconPath from "../assets/me.jpg";

const Main = styled.div`
  display: flex;
  flex-direction: row;
  color: #43506c;
  align-items: center;
  justify-content: space-between;
  font-size: 22px;
  background-color: #e9e9eb;
  box-shadow: 0 0 5px #43506c;
  z-index: 3;
  width: 100%;
  height: 50px;
  flex-shrink: 0;
`;

const Logo = styled.img`
  width: 2.5%;
  margin-right: 9%;
`;

const IconHeaderButton = styled.button`
  background: none;
  border: none;
  font-size: 28px;
  color: #43506c;
  cursor: pointer;
`;

const AccountMenu = styled.button`
  display: flex;
  flex-direction: row;
  background: none;
  border: none;
  color: #43506c;
  background-color: #e9e9eb;
  margin: 0 10px 0 10px;
  align-items: center;
  cursor: pointer;
`;

const AccountIconMenu = styled.img`
  width: 50px;
  height: 50px;
  border-radius: 50%;
`;

const AccountNameAndEmailMenu = styled.div`
  background: none;
  border: none;
  font-size: 18px;
  color: #43506c;
  margin: 10px;
  white-space: nowrap;
`;

const AccountNameMenu = styled.div`
  background: none;
  border: none;
  font-size: 18px;
  color: #43506c;
  margin: 10px;
  white-space: nowrap;
  cursor: pointer;
`;

const AccountEmailMenu = styled.div`
  background: none;
  border: none;
  font-size: 14px;
  color: #43506c;
  margin: 10px;
  cursor: pointer;
`;

const Workspace = styled.button`
  background: none;
  border: none;
  display: flex;
  font-size: 18px;
  color: #43506c;
  align-items: center;
  justify-content: center;
  cursor: pointer;
`;

const HeaderRight = styled.div`
  display: flex;
  flex-direction: row;
  height: 100%;
  align-items: center;
  justify-content: center;
  margin-right: 10px;
  gap: 10px;
`;

const ProfileMenu = styled.ul`
  display: flex;
  flex-direction: column;
  position: absolute;
  top: 40px;
  right: 0px;
  background-color: #e9e9eb;
  box-shadow: 0 2px 5px #43506c;
  padding: 10px;
  z-index: 2;
  list-style: none;
`;
const ProfileMenuItem = styled.button`
  display: flex;
  background: none;
  border: none;
  cursor: pointer;
  color: #43506c;
  font-size: 16px;
  padding: 7px;
`;

const WorkspaceDropdown = styled.ul`
  display: flex;
  flex-direction: column;
  position: absolute;
  top: 40px;
  right: 0px;
  width: 180px;
  height: 300px;
  background-color: #e9e9eb;
  box-shadow: 0 2px 5px #43506c;
  padding: 10px;
  overflow-y: auto;
  z-index: 2;
  list-style: none;

  scrollbar-width: none;
  &::-webkit-scrollbar {
    display: none;
  }
`;

const WorkspaceDropdownItem = styled.li`
  display: flex;
  justify-content: center;
  align-items: center;
  background: none;
  box-shadow: 0 1px 2px #43506c;
  color: #43506c;
  font-size: 16px;
  padding: 7px;
  height: 20px;
  cursor: pointer;
`;

export default function Header() {
  const [isProfileOpen, setIsProfileOpen] = useState(false);
  const [isWorkspaceOpen, setIsWorkspaceOpen] = useState(false);

  return (
    <Main>
      <Logo src={iconPath} />
      <div>NEXUS | ATLAS</div>

      <HeaderRight>
        <div style={{ position: "relative", height: "100%", display: "flex" }}>
          <Workspace onClick={() => setIsWorkspaceOpen(!isWorkspaceOpen)}>
            MY_WORKSPACE<span className="material-symbols-outlined">expand_more</span>
          </Workspace>
          {isWorkspaceOpen && (
            <WorkspaceDropdown>
              <WorkspaceDropdownItem onClick={() => alert("change workspace")}>OTHER_WORKSPACE</WorkspaceDropdownItem>
              <WorkspaceDropdownItem onClick={() => alert("change workspace")}>MORE_WORKSPACE</WorkspaceDropdownItem>
              <WorkspaceDropdownItem onClick={() => alert("change workspace")}>HOME</WorkspaceDropdownItem>
              <WorkspaceDropdownItem className="material-symbols-outlined" style={{ marginTop: "20px" }} onClick={() => alert("add workspace")}>
                add
              </WorkspaceDropdownItem>
            </WorkspaceDropdown>
          )}
        </div>
        <div style={{ position: "relative", height: "100%", display: "flex" }}>
          <IconHeaderButton onClick={() => setIsProfileOpen(!isProfileOpen)} className="material-symbols-outlined">
            account_circle
          </IconHeaderButton>
          {isProfileOpen && (
            <ProfileMenu>
              <AccountMenu onClick={() => alert("click")}>
                <AccountIconMenu src={accountIconPath} />
                <AccountNameAndEmailMenu>
                  <AccountNameMenu>George Beetle</AccountNameMenu>
                  <AccountEmailMenu>gbeetle@gmail.com</AccountEmailMenu>
                </AccountNameAndEmailMenu>
              </AccountMenu>
              <hr style={{ width: "100%" }} />
              <ProfileMenuItem onClick={() => alert("Settings")}>Settings</ProfileMenuItem>
              <ProfileMenuItem onClick={logout}>Sign out</ProfileMenuItem>
            </ProfileMenu>
          )}
        </div>
      </HeaderRight>
    </Main>
  );
}
