import axios from "axios";

export const login = async (email, password) => {
  try {
    const res = await axios.post("http://localhost:8080/auth/login", {
      email,
      password,
    });
    localStorage.setItem("token", JSON.stringify(res.data.token));
    return res;
  } catch (e) {
    console.log({ e });
    return e;
  }
};

export const Register = async (email, password, roles) => {
  try {
    const res = await axios.post("http://localhost:8080/auth/register", {
      email,
      password,
      roles,
    });
    console.log(res);
    return res;
  } catch (e) {
    console.log(e);
    return e;
  }
};
