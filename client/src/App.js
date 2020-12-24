import React, { useState, useEffect } from "react";
import axios from "axios";
import { connect, sendMsg } from "./socket";
export default function App() {
  const [email, setEmail] = useState("");
  const [receiver1, setReceiver1] = useState("");
  const [receiver2, setReceiver2] = useState("");
  const [receiver3, setReceiver3] = useState("");
  const [loading, setLoading] = useState(false);

  useEffect(() => connect(), []);
  const hello = () => {
    sendMsg("Hello from Client");
  };
  const sendMail = async () => {
    try {
      var receiver = [];
      receiver.push(receiver1);
      receiver.push(receiver2);
      receiver.push(receiver3);
      console.log("receiver", receiver);

      setLoading(true);
      const res = await axios.post("http://localhost:8080/mail/send-all", {
        email,
        receiver,
      });
      setLoading(false);
      console.log(res);
      return res.data;
    } catch (e) {
      setLoading(false);
      console.log(e);
    }
  };

  return (
    <div style={{ textAlign: "center" }}>
      <div>
        <h1>Send Email</h1>
        <div>
          <p>
            <label>From: </label>
          </p>
          <input
            style={{ padding: 10 }}
            onChange={(e) => setEmail(e.target.value)}
            type="email"
            name="email"
            value={email}
            placeholder="Enter email..."
          />
        </div>
        <div
          style={{
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
          }}
        >
          <p>
            <label>To: </label>
          </p>
          <input
            style={{ padding: 10 }}
            onChange={(e) => setReceiver1(e.target.value)}
            type="email"
            name="receiver"
            value={receiver1}
            placeholder="Enter receiver..."
          />
          <input
            style={{ padding: 10 }}
            onChange={(e) => setReceiver2(e.target.value)}
            type="email"
            name="receiver"
            value={receiver2}
            placeholder="Enter receiver..."
          />
          <input
            style={{ padding: 10 }}
            onChange={(e) => setReceiver3(e.target.value)}
            type="email"
            name="receiver"
            value={receiver3}
            placeholder="Enter receiver..."
          />
        </div>
        <button style={{ marginTop: 20 }} onClick={sendMail}>
          Send mail
        </button>
        {loading && "Loading..."}

        <button onClick={hello}>Say hello to backend server</button>
      </div>
    </div>
  );
}
