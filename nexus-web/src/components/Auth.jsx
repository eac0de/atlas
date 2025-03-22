import { useNavigate } from "react-router-dom";
import styled from "styled-components";
import { React, useState, useEffect } from "react";
import iconPath from "../assets/new_atlas.png";

const emailRegex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;

const MainAuth = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 100%;
  height: 100%;
`;

const Title = styled.div`
  margin: 15px 0px 70px 0px;
  font-size: 32px;
`;

const Input = styled.input`
  margin-top: 0px;
  font-size: 18px;
  background: none;
  border: none;
  width: 500px;
  text-align: center;
  padding: 15px;
  color: inherit;
  outline: none;
  border-bottom: 2px solid #ccc;
  align-items: center;
  justify-content: center;
`;

const ButtonIn = styled.button`
  margin-top: 100px;
  background: none;
  border: none;
  font-size: 38px;
  color: inherit;
  cursor: pointer;
  width: fit-content;
`;
const Logo = styled.img`
  width: 3%;
  margin-top: 12%;
`;
export default function Auth() {
  const navigate = useNavigate();

  const token = localStorage.getItem("nexus_access_token");

  useEffect(() => {
    if (token) {
      navigate("/");
    }
  }, [token]);

  const [email, setEmail] = useState("");
  const [isValidEmail, setIsValidEmail] = useState(false);
  const [emailCodeId, setEmailCodeId] = useState("");

  const [code, setCode] = useState("");

  const handleInputChange = (event) => {
    if (emailCodeId === "") {
      setEmail(event.target.value);
      setIsValidEmail(emailRegex.test(event.target.value));
    } else {
      setCode(event.target.value);
    }
  };

  const handleSubmit = () => {
    if (emailCodeId === "") {
      generateCode();
    } else {
      verifyCode();
    }
  };

  async function generateCode() {
    const data = {
      email: email,
    };
    try {
      const response = await fetch("http://localhost:8080/api/auth/code/generate/", { body: JSON.stringify(data), method: "POST", headers: { "Content-Type": "application/json" } });
      if (!response.ok) {
        throw new Error(`Ошибка сети: ${response.status} ${response.statusText}`);
      }
      const result = await response.json();
      setEmailCodeId(result.email_code_id);
    } catch (error) {
      console.error("Ошибка:", error.message);
    }
  }
  async function verifyCode() {
    if (code.length !== 4 || isNaN(parseInt(code))) {
      return;
    }
    const data = {
      email_code_id: emailCodeId,
      code: parseInt(code),
    };
    try {
      const response = await fetch("http://localhost:8080/api/auth/code/verify/", {
        body: JSON.stringify(data),
        method: "POST",
        headers: { "Content-Type": "application/json" },
        credentials: "include",
      });
      if (!response.ok) {
        if (response.status === 410) {
          setEmailCodeId("");
          return;
        }
        throw new Error(`Ошибка сети: ${response.status} ${response.statusText}`);
      }
      const result = await response.json();
      localStorage.setItem("nexus_access_token", result.access_token);
      navigate("/");
      return null;
    } catch (error) {
      console.error("Ошибка:", error.message);
    }
  }

  async function onEnter(event) {
    if (event.code == "Enter") {
      handleSubmit();
    }
  }

  return (
    <MainAuth>
      <Logo src={iconPath} />
      <Title>Welcome to NEXUS</Title>
      {!emailCodeId ? (
        <Input type={"email"} placeholder="your@email.com" value={email} onChange={handleInputChange} onKeyDown={onEnter} />
      ) : (
        <Input type="text" placeholder={`code from ${email}`} value={code} onChange={handleInputChange} maxLength={4} onKeyDown={onEnter} />
      )}
      <div>
        {emailCodeId ? (
          <ButtonIn className="material-symbols-outlined" style={{ paddingRight: "200px" }} onClick={() => setEmailCodeId("")}>
            chevron_left
          </ButtonIn>
        ) : null}
        <ButtonIn className="material-symbols-outlined" style={{ opacity: isValidEmail ? "100%" : "20%" }} disabled={!isValidEmail} onClick={handleSubmit}>
          {!emailCodeId ? "chevron_right" : "check"}
        </ButtonIn>
      </div>
    </MainAuth>
  );
}
