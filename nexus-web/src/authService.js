export async function logout() {
  try {
    const response = await fetch("http://localhost:8080/api/auth/token/", { method: "DELETE", credentials: "include" });
    if (!response.ok) {
      console.warn(`Ошибка сети: ${response.status} ${response.statusText}`);
    }
    localStorage.removeItem("nexus_access_token");
    window.location.reload();
  } catch (error) {
    console.error("Ошибка:", error.message);
  }
}
