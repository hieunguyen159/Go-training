import React from "react";
import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom";
import ToggleRemind from "./pages/TogglePage/ToggleRemind";
import Home from "./pages/Home/Home";
export default function App() {
  return (
    <Router>
      <Switch>
        <Route exact path="/" component={Home} />
        <Route exact path="/:id" component={ToggleRemind} />
      </Switch>
    </Router>
  );
}
