import { useLocalStorage } from "@/hooks/useLocalStorage";
import { useEffect, useState } from "react";
import { Navigate, Outlet } from "react-router-dom";

export default function ProtectedRouter() {
  const token = localStorage.getItem("access_token");
  const [isValid, setIsValid] = useState(isJwtValid(token));

  useEffect(() => {
    const timer = setInterval(() => {
      const validate = isJwtValid(token)
      if (isValid !== validate) {
        setIsValid(validate)
      }
    }, 5000);
    return () => clearInterval(timer);
  }, [token]);

  return isValid ? <Outlet /> : <Navigate to="/login" replace />;
}

const isJwtValid = (token: string | null): boolean => {
  try {
    if (!token) return false;

    const payload = JSON.parse(atob(token.split('.')[1]));
    const expTime = payload.exp * 1000; // Convert to milliseconds
    const exp = 5 * 1000; // 5 seconds
    return Date.now() < (expTime - exp);
  } catch (error) {
    return false;
  }
};
