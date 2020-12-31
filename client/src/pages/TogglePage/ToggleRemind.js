import React, { useState } from "react";
import mailsService from "../../api/mails";
import IconButton from "@material-ui/core/IconButton";
import NotificationsIcon from "@material-ui/icons/Notifications";
import CircularProgress from "@material-ui/core/CircularProgress";
export default function ToggleRemind(props) {
  const [loading, setLoading] = useState(false);
  const id = props.match.params.id;
  const cancelAlert = async () => {
    setLoading(true);
    await mailsService.toggleAlert(id);
    setLoading(false);
  };
  return (
    <div
      style={{
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
      }}
    >
      <IconButton color="secondary" onClick={cancelAlert}>
        Turn Off
        <NotificationsIcon />
        {loading && <CircularProgress />}
      </IconButton>
    </div>
  );
}
