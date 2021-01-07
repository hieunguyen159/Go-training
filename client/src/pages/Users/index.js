import React, { useState, useEffect } from "react";
import { getAllUsers } from "../../api/users";
import DataTable from "../../components/DataTable";

export default function Users() {
  const [fetchingLoading, setFetchingLoading] = useState(false);
  const [users, setUsers] = useState([]);
  const fetchUsers = async () => {
    setFetchingLoading(true);
    setUsers(await getAllUsers());
    setFetchingLoading(false);
  };
  useEffect(() => {
    fetchUsers();
  }, []);

  return (
    <div className="table-root">
      {users && (
        <DataTable
          fetchUsers={fetchUsers}
          loading={fetchingLoading}
          data={users}
        />
      )}
    </div>
  );
}
