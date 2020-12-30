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
export default {
  sendMail,
};
