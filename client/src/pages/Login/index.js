import React, { useState } from "react";
import "./index.css";
import { login, Register } from "../../api/auth";
import { useSnackbar } from "notistack";
import { useHistory } from "react-router-dom";
import Select from "@material-ui/core/Select";
import InputLabel from "@material-ui/core/InputLabel";
export default function Login() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [loading, setLoading] = useState(false);
  const [register, setRegister] = useState(false);
  const [roles, setRoles] = useState([]);
  console.log("roles", roles);
  let history = useHistory();
  const { enqueueSnackbar } = useSnackbar();

  const handleSubmit = async () => {
    setLoading(true);
    const res = await login(email, password);
    setLoading(false);
    history.push("/home");
    if (res.status === 200) {
      enqueueSnackbar("Login successfully !");
    } else enqueueSnackbar(res.response.data.message);
  };

  const handleChange = (event) => {
    const rolesArr = [...roles];
    rolesArr.push(event.target.value);
    setRoles(rolesArr);
  };

  const handleRegister = async () => {
    setLoading(true);
    const res = await Register(email, password, roles);
    setLoading(false);
    setRegister(false);
    if (res.status === 201) {
      enqueueSnackbar("Register successfully !");
    } else enqueueSnackbar(res.response.data.message);
  };
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
                <p style={{ color: "red", fontSize: 12 }}>Enter email</p>
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
                <p style={{ color: "red", fontSize: 12 }}>Enter password</p>
              )}
              <button
                className="login--button"
                onClick={handleSubmit}
                disabled={loading}
              >
                LOGIN
              </button>
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
                <p style={{ color: "red", fontSize: 12 }}>Enter email</p>
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
                <p style={{ color: "red", fontSize: 12 }}>Enter password</p>
              )}
              <div
                style={{
                  margin: "10px 0",
                  //    display: "flex",
                  //    alignItems: "center",
                  //    justifyContent: "space-around",
                }}
              >
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
              <button
                className="login--button"
                onClick={handleRegister}
                disabled={loading}
              >
                REGISTER
              </button>
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
