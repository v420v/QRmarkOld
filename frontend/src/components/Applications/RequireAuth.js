import { useContext, useEffect, useState } from "react";
import { useLocation, Navigate, Outlet, useNavigate } from "react-router-dom";
import axios from "../../api/axios";
import React from "react";
import CurrentUserContext from "../../context/AuthProvider";
import SessionContext from "../../context/SessionProvider";

const RequireAuth = ({ allowedRoles }) => {
  const [isPending, setIsPending] = useState(true);
  const location = useLocation();
  const [CurrentUser, setCurrentUser] = useContext(CurrentUserContext);
  const [session, setSession] = useContext(SessionContext);

  useEffect(() => {
    setIsPending(true);
    const FetchCurrentUser = async () => {
      try {
        const response = await axios.get("/user/current",
          {
            headers: {'Authorization': `Bearer ${session}`,'Content-Type': 'text/plain; charset=utf-8'}
          }
        );
        setCurrentUser(response.data);
      } catch (err) {
        setSession(null);
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
