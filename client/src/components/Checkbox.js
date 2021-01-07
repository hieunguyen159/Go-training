import React, { useState } from "react";

export default function Checkbox({ value, onSelected }) {
  const [checked, setChecked] = useState(false);
  return (
    <p>
      <input
        type="checkbox"
        value={value}
        checked={checked}
        onChange={() => {
          let c = !checked;
          setChecked(c);
          onSelected({ value: value, checked: c });
        }}
      />
      {value}
    </p>
  );
}
