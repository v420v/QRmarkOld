import { useContext, useEffect, useState } from "react";
import { useLocation, Navigate, Outlet } from "react-router-dom";
import axios from "../../api/axios";
import React from "react";
import CurrentUserContext from "../../context/AuthProvider";

const RequireAuth = ({ allowedRoles }) => {
  const [isPending, setIsPending] = useState(true);
  const location = useLocation();
  const [CurrentUser, setCurrentUser] = useContext(CurrentUserContext);

  useEffect(() => {
    setIsPending(true);
    const FetchCurrentUser = async () => {
      try {
        const response = await axios.get("/user/current",
          {
            headers: {
              'Content-Type': 'text/plain; charset=utf-8',
            },
            withCredentials: true
          }
        );
        setCurrentUser(response.data);
      } catch (err) {
        setCurrentUser(null);
      } finally {
        setIsPending(false);
      }
    };
    FetchCurrentUser();
  }, []);

  if (isPending) {
    return <div>読み込み中...</div>;
  }

  if (CurrentUser) {
    if (allowedRoles.length !== 0 && !allowedRoles.includes(CurrentUser.role)) {
      return (
        <Navigate to="/about" state={{ from: location }} replace />
      );
    }
    return <Outlet />;
  }

  return (
    <Navigate to="/about" state={{ from: location }} replace />
  );
};

export default RequireAuth;
