import { useNavigate } from "react-router-dom";
import "./Auth.css";
import { React, useState, useEffect } from "react";

const emailRegex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;

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

  return (
    <div className="main">
      <div className="title">NEXUS</div>
      <input type="email" className={"input " + (!emailCodeId ? "active" : "")} placeholder="your@email.com" value={email} onChange={handleInputChange} />
      {emailCodeId ? <input type="text" className={"input " + (emailCodeId ? "active" : "")} placeholder="code" value={code} onChange={handleInputChange} maxLength={4} /> : null}
      <div>
        {emailCodeId ? (
          <button style={{ paddingRight: "150px" }} className="material-symbols-outlined button-in" onClick={() => setEmailCodeId("")}>
            chevron_left
          </button>
        ) : null}
        <button style={{ opacity: isValidEmail ? "100%" : "20%" }} disabled={!isValidEmail} className="material-symbols-outlined button-in" onClick={handleSubmit}>
          chevron_right
        </button>
      </div>
    </div>
  );
}
