import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import ExitToAppIcon from "@material-ui/icons/ExitToApp";
import Typography from "@material-ui/core/Typography";
import Button from "@material-ui/core/Button";
import Container from "@material-ui/core/Container";
import "./layouts.css";
import {
  BrowserRouter as Router,
  Switch,
  Route,
  useRouteMatch,
  Redirect,
  useHistory,
  Link,
} from "react-router-dom";
import Home from "../pages/Home/Home";
import Users from "../pages/Users";

const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
  },
  menuButton: {
    marginRight: theme.spacing(2),
  },
  title: {
    flexGrow: 1,
  },
  navbarEle: {
    margin: "0 100px",
    display: "flex",
    justifyContent: "space-between",
  },
  container: {
    marginTop: 64,
    padding: "20px 0",
  },
}));
export default function Layouts() {
  const classes = useStyles();
  const match = useRouteMatch();
  const history = useHistory();
  const logOut = () => {
    setTimeout(() => {
      localStorage.clear();
      history.push("/login");
    }, 2000);
  };
  return (
    <div>
      <Router>
        <AppBar color="inherit" position="fixed">
          <Toolbar>
            <Typography variant="h5" className={classes.title}>
              Dashboard
            </Typography>
            <div className={classes.navbarEle}>
              <Link to="emails" target="_top">
                <Button color="inherit">Emails</Button>
              </Link>
              <Link to="users" target="_top">
                <Button color="inherit">Users</Button>
              </Link>
            </div>
            <Button
              onClick={logOut}
              variant="contained"
              color="secondary"
              endIcon={<ExitToAppIcon />}
            >
              Logout
            </Button>
          </Toolbar>
        </AppBar>
        <Container className={classes.container}>
          <Switch>
            <Redirect exact from="/home" to={`${match.url}/emails`} />
            <Route exact path="/home/emails" component={Home} />
            <Route exact path="/home/users" component={Users} />
          </Switch>
        </Container>
      </Router>
    </div>
  );
}
