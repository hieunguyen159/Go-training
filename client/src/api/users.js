import axios from "axios";

export const getAllUsers = async () => {
  const token = JSON.parse(localStorage.getItem("token"));
  try {
    const res = await axios.get("http://localhost:8080/users", {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    console.log("user", res.data);
    return res.data;
  } catch (e) {
    console.log("error", e);
  }
};

export const toggleUser = async (id, status) => {
  const token = JSON.parse(localStorage.getItem("token"));
  try {
    const res = await axios.put(
      "http://localhost:8080/users/active/" + id,
      { status },
      {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      }
    );
    console.log(res);
    return res;
  } catch (e) {
    console.log("error", { e });
    return e;
  }
};

export const setRolesUser = async (id, roles) => {
  const token = JSON.parse(localStorage.getItem("token"));
  try {
    const res = await axios.put(
      "http://localhost:8080/users/roles/" + id,
      { roles },
      {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      }
    );
    console.log(res.data);
    return res;
  } catch (e) {
    console.log("error", e);
    return e;
  }
};
