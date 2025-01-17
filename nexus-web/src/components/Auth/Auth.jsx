import "./Auth.css";
import { useState } from "react";

const emailRegex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;

export default function Auth() {
  const [email, setEmail] = useState("");
  const [isValidEmail, setIsValidEmail] = useState(false);

  const handleInputChange = (event) => {
    setEmail(event.target.value);
    setIsValidEmail(emailRegex.test(event.target.value));
  };

  const handleSubmit = () => {
    if (isValidEmail) {
      alert("Email: " + email);
    } else {
      alert("Invalid email");
    }
  };

  return (
    <div className="main">
      <div className="title">NEXUS</div>
      <input type="email" className="input" placeholder="your@email.com" value={email} onChange={handleInputChange} />
      <button style={{ opacity: isValidEmail ? "100%" : "20%" }} className="material-symbols-outlined button-in" onClick={handleSubmit}>
        chevron_right
      </button>
    </div>
  );
}
