import React from "react";
import mailsService from "../../api/mails";

export default function ToggleRemind(props) {
  const id = props.match.params.id;
  const cancelAlert = async () => {
    await mailsService.toggleAlert(id);
  };
  return <button onClick={cancelAlert}>Turn Off</button>;
}
