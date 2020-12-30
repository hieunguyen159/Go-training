import axios from "axios";

const sendMail = async (email, receiver) => {
  try {
    const res = await axios.post("http://localhost:8080/mail/send-all", {
      email,
      receiver,
    });
    console.log(res.data);
    return res.data;
  } catch (e) {
    console.log("error", e);
  }
};
const getAllEmails = async () => {
  try {
    const res = await axios.get("http://localhost:8080/emails");
    console.log(res.data);
    return res.data;
  } catch (e) {
    console.log("error", e);
  }
};

const toggleAlert = async (id) => {
  try {
    const res = await axios.put("http://localhost:8080/emails/" + id);
    console.log(res.data);
    return res.data;
  } catch (e) {
    console.log("error", e);
  }
};
export default {
  sendMail,
  getAllEmails,
  toggleAlert,
};
