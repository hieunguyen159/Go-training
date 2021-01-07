import React, { useState } from "react";
import { Redirect, Route } from "react-router-dom";

function PrivateRoute({ component: Component, ...rest }) {
  const [isAccess] = useState(localStorage.getItem("token"));
  return (
    <Route
      {...rest}
      render={(props) =>
        isAccess ? (
          <Component {...props} />
        ) : (
          <Redirect to={{ pathname: "/login" }} />
        )
      }
    />
  );
}

export default PrivateRoute;
