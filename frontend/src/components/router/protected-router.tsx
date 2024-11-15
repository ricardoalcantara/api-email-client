import { useLocalStorage } from "@/hooks/useLocalStorage";
import { Navigate, Outlet } from "react-router-dom";

export default function ProtectedRouter() {
  const access_token = localStorage.getItem("access_token");
  return access_token ? <Outlet /> : <Navigate to="/login" replace />;
}