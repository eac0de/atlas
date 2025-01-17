import React from "react";
import { Navigate, Outlet } from "react-router-dom";

export default function PrivateRoute() {
  const token = localStorage.getItem("nexus_access_token");
  if (!token) return <Navigate to="/login" />;
  return <Outlet />;
}
