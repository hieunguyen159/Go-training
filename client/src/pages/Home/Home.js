import React, { useState, useEffect } from "react";
import { makeStyles } from "@material-ui/core/styles";
import LinearProgress from "@material-ui/core/LinearProgress";
import Button from "@material-ui/core/Button";
import Fab from "@material-ui/core/Fab";
import AddIcon from "@material-ui/icons/Add";
import "./home.css";
import sendMailApi from "../../api/mails";

const useStyles = makeStyles((theme) => ({
  root: {
    width: "30%",
    "& > * + *": {
      marginTop: 10,
    },
    padding: 30,
  },
}));

var socket = new WebSocket("ws://localhost:8080/ws");
export default function App() {
  const [email, setEmail] = useState("");
  const [receiver, setReceiver] = useState([""]);
  const [loading, setLoading] = useState(false);

  const [emailList, setEmailList] = useState([]);

  const [mess, setMess] = useState([]);
  console.log("rec", receiver);

  const classes = useStyles();

  // socket connection
  useEffect(() => {
    console.log("Attempting Connection...");
    socket.onopen = () => {
      console.log("Successfully Connected");
    };
    socket.onmessage = (msg) => {
      console.log("on message: ", msg.data);
      let data = JSON.parse(msg.data);
      setMess(data);
    };
    socket.onerror = (error) => {
      console.log("Socket Error: ", error);
    };
    return () =>
      (socket.onclose = (event) => {
        console.log("Socket Closed Connection: ", event);
      });
  }, []);

  // handle real-time sending mails
  useEffect(() => {
    let listClone = [...emailList];
    for (let i = 0; i < listClone.length; i++) {
      if (listClone[i].id.trim() === mess[i].id.trim()) {
        let changeItem = {
          ...listClone[i],
          status: "done",
        };
        listClone[i] = changeItem;
      }
      setEmailList(listClone);
    }
  }, [mess]);

  // send mails API calls
  const sendMailToUsers = async () => {
    try {
      setEmailList([]);
      setLoading(true);
      const fetchEmailList = async () => {
        const res = await sendMailApi.sendMail(email, receiver);
        setEmailList(res);
        setLoading(false);
      };
      fetchEmailList();
    } catch (e) {
      setLoading(false);
      console.log(e);
    }
  };

  // handle text change array
  const onChangeText = (e, i) => {
    const receiversValues = [...receiver];
    receiversValues[i] = e.target.value;
    setReceiver(receiversValues);
  };

  const onFieldAdded = () => {
    const inputField = [...receiver];
    inputField.push("");
    setReceiver(inputField);
  };

  return (
    <div className="App">
      <Fab
        className="fab-button"
        color="primary"
        aria-label="add"
        onClick={onFieldAdded}
      >
        <AddIcon />
      </Fab>
      <div className="from--area">
        <h1>Send Email</h1>
        <p style={{ textAlign: "left" }}>Enter your email:</p>

        <div style={{ display: "flex", alignItems: "center", margin: 10 }}>
          <input
            style={{ padding: 10 }}
            onChange={(e) => setEmail(e.target.value)}
            type="email"
            name="email"
            value={email}
            placeholder="Enter email..."
          />
          <Button
            style={{ marginLeft: 20 }}
            variant="contained"
            color="primary"
            onClick={sendMailToUsers}
            disabled={loading}
          >
            Send mail
          </Button>
        </div>
        {loading && (
          <div className={classes.root}>
            <LinearProgress color="secondary" />
          </div>
        )}

        <div className="email-list-container">
          {emailList &&
            emailList.length > 0 &&
            emailList.map((e, i) => (
              <p
                key={i}
                style={{ color: e.status === "done" ? "green" : "red" }}
              >
                Email: {e.email} - Status: {e.status}
              </p>
            ))}
        </div>
      </div>
      <div className="to--area">
        <p>
          <label>To: </label>
        </p>
        {receiver.map((input, i) => (
          <input
            key={i}
            disabled={loading}
            style={{ padding: 10, margin: "10px 0" }}
            onChange={(e) => onChangeText(e, i)}
            type="email"
            name={i}
            value={receiver[i]}
            placeholder="Enter receiver..."
          />
        ))}
      </div>
    </div>
  );
}
