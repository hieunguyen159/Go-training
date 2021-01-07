import React, { useState } from "react";
import "./index.css";
import { login, Register } from "../../api/auth";
import { useSnackbar } from "notistack";
import { useHistory } from "react-router-dom";
import Select from "@material-ui/core/Select";
import InputLabel from "@material-ui/core/InputLabel";
import { Redirect } from "react-router-dom";
import Button from "@material-ui/core/Button";
export default function Login() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [loading, setLoading] = useState(false);
  const [register, setRegister] = useState(false);
  const [roles, setRoles] = useState([]);
  const [isAuth] = useState(localStorage.getItem("token"));
  console.log("roles", roles);
  let history = useHistory();
  const { enqueueSnackbar } = useSnackbar();

  const handleSubmit = async () => {
    setLoading(true);
    const res = await login(email, password);
    if (res.status === 200) {
      setLoading(false);
      history.push("/home");
      enqueueSnackbar("Login successfully !", {
        variant: "success",
      });
    } else {
      enqueueSnackbar(res.response.data.message, {
        variant: "error",
      });
      setLoading(false);
    }
  };

  const handleChange = (event) => {
    const rolesArr = [...roles];
    rolesArr.push(event.target.value);
    setRoles(rolesArr);
  };

  const handleRegister = async () => {
    setLoading(true);
    const res = await Register(email, password, roles);
    if (res.status === 201) {
      setLoading(false);
      setRegister(false);
      enqueueSnackbar("Register successfully !", {
        variant: "success",
      });
    } else {
      enqueueSnackbar(res.response.data.message, {
        variant: "error",
      });
      setLoading(false);
    }
  };
  if (isAuth) {
    return <Redirect to="/home" />;
  }
  return (
    <div className="root">
      <div className="form__area">
        <div className="form__image--container"></div>
        {!register ? (
          <div className="form__input--wrapper">
            <div className="form__input--container">
              <h1 style={{ fontWeight: "bold" }}>Member Login</h1>
              <div className="form__input">
                <i className="fas fa-envelope" />
                <input
                  type="text"
                  name="email"
                  className="form__input--item"
                  placeholder="Email"
                  onChange={(e) => setEmail(e.target.value)}
                />
              </div>
              {email.trim().length === 0 && (
                <span style={{ color: "red", fontSize: 12 }}>Enter email</span>
              )}
              <div className="form__input">
                <i className="fas fa-unlock-alt" />
                <input
                  type="password"
                  name="password"
                  className="form__input--item"
                  placeholder="Password"
                  onChange={(e) => setPassword(e.target.value)}
                />
              </div>
              {password.trim().length === 0 && (
                <span style={{ color: "red", fontSize: 12 }}>
                  Enter password
                </span>
              )}
              <Button
                onClick={handleSubmit}
                disabled={loading}
                className="login--button"
                variant="contained"
                color="primary"
              >
                LOGIN
              </Button>
              <p className="forgot__link">
                Forgot
                <button className="forgot-button" href="#">
                  Username / Password ?
                </button>
              </p>
            </div>
            <div
              className="create__container"
              onClick={() => setRegister(true)}
            >
              <p> Create a new account</p>&nbsp;&nbsp;&nbsp;
              <i className="fas fa-arrow-right" />
            </div>
          </div>
        ) : (
          <div className="form__input--wrapper">
            <div className="form__input--container">
              <h1 style={{ fontWeight: "bold" }}>Member Register</h1>
              <div className="form__input">
                <i className="fas fa-envelope" />
                <input
                  type="text"
                  name="email"
                  className="form__input--item"
                  placeholder="Email"
                  onChange={(e) => setEmail(e.target.value)}
                />
              </div>
              {email.trim().length === 0 && (
                <span style={{ color: "red", fontSize: 12 }}>Enter email</span>
              )}
              <div className="form__input">
                <i className="fas fa-unlock-alt" />
                <input
                  type="password"
                  name="password"
                  className="form__input--item"
                  placeholder="Password"
                  onChange={(e) => setPassword(e.target.value)}
                />
              </div>
              {password.trim().length === 0 && (
                <span style={{ color: "red", fontSize: 12 }}>
                  Enter password
                </span>
              )}
              <div style={{ margin: "10px 0" }}>
                <InputLabel htmlFor="age-native-simple">Role</InputLabel>
                <Select
                  native
                  value={roles}
                  onChange={handleChange}
                  inputProps={{
                    name: "age",
                    id: "age-native-simple",
                  }}
                >
                  <option aria-label="None" value="" />
                  <option value="ADMIN">Admin</option>
                  <option value="USER">User</option>
                </Select>
              </div>
              <Button
                onClick={handleRegister}
                disabled={loading}
                className="login--button"
                variant="contained"
                color="primary"
              >
                REGISTER
              </Button>
            </div>
            <div
              className="create__container"
              onClick={() => setRegister(false)}
            >
              <i className="fas fa-arrow-left" />
              &nbsp;&nbsp;&nbsp;
              <p> Login</p>
            </div>
          </div>
        )}
      </div>
    </div>
  );
}
