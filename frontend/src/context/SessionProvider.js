import React from "react";

export const SessionContext = React.createContext();

export const SessionProvider = ({ children }) => {
  const [session, setSession] = React.useState(null)

  return (
    <SessionContext.Provider value={[ session, setSession ]}>
      {children}
    </SessionContext.Provider>
  )
}

export default SessionContext