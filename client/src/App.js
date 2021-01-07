import React from "react";
import PrivateRoute from "./guards/PrivateRoute";
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Redirect,
} from "react-router-dom";
import { SnackbarProvider } from "notistack";
import ToggleRemind from "./pages/TogglePage/ToggleRemind";

import Login from "./pages/Login";
import Layouts from "./layouts";
export default function App() {
  return (
    <Router>
      <SnackbarProvider maxSnack={3}>
        <Switch>
          <Redirect exact from="/" to="/login" />
          <Route exact path="/login" component={Login} />
          <PrivateRoute path="/home" component={Layouts} />
          <Route exact path="/alert/:id" component={ToggleRemind} />
          <Route path="*" component={() => <div>404 Not Found</div>} />
        </Switch>
      </SnackbarProvider>
    </Router>
  );
}
