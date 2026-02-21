import { useState, useEffect } from "react";
import { Link, useNavigate } from "react-router-dom";
import api from "../utils/api";

const roleStyles = {
  admin: "bg-rose-100 text-rose-800",
  teacher: "bg-violet-100 text-violet-800",
  student: "bg-sky-100 text-sky-800",
};

export default function ListUsers() {
  const navigate = useNavigate();
  useEffect(() => {
    if (!localStorage.getItem("token")) navigate("/");
  }, [navigate]);
  const [users, setUsers] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {
    const fetchUsers = async () => {
      try {
        const res = await api.get("/api/users");
        setUsers(Array.isArray(res.data) ? res.data : []);
      } catch (err) {
        setError(err.response?.data?.error || "Failed to fetch users");
      } finally {
        setLoading(false);
      }
    };
    fetchUsers();
  }, []);

  return (
    <div className="mx-auto max-w-4xl px-4 py-10 sm:px-6">
      <Link to="/admin" className="link-back mb-6 inline-flex">
        <span>←</span> Back to Dashboard
      </Link>

      <div className="mb-6">
        <h1 className="text-2xl font-bold text-slate-800">All Users</h1>
        <p className="mt-1 text-slate-500">
          Copy Student IDs for enrollment. Admin-only.
        </p>
      </div>

      {loading ? (
        <div className="card py-12 text-center">
          <div className="inline-block h-8 w-8 animate-spin rounded-full border-2 border-indigo-500 border-t-transparent" />
          <p className="mt-4 text-slate-500">Loading users...</p>
        </div>
      ) : error ? (
        <div className="card border border-red-100 bg-red-50/50">
          <p className="text-red-600">{error}</p>
        </div>
      ) : (
        <div className="card overflow-hidden p-0">
          <div className="overflow-x-auto">
            <table className="w-full min-w-[500px]">
              <thead>
                <tr className="border-b border-slate-200 bg-slate-50/80">
                  <th className="px-4 py-3 text-left text-xs font-semibold uppercase tracking-wide text-slate-500">
                    ID
                  </th>
                  <th className="px-4 py-3 text-left text-xs font-semibold uppercase tracking-wide text-slate-500">
                    Name
                  </th>
                  <th className="px-4 py-3 text-left text-xs font-semibold uppercase tracking-wide text-slate-500">
                    Email
                  </th>
                  <th className="px-4 py-3 text-left text-xs font-semibold uppercase tracking-wide text-slate-500">
                    Role
                  </th>
                </tr>
              </thead>
              <tbody className="divide-y divide-slate-100">
                {users.map((u) => (
                  <tr
                    key={u.id || u._id}
                    className="transition hover:bg-slate-50/50"
                  >
                    <td className="px-4 py-3 font-mono text-xs text-slate-600 break-all">
                      {(typeof u.id === "string" ? u.id : u.id?.$oid) ||
                        u._id ||
                        "—"}
                    </td>
                    <td className="px-4 py-3 font-medium text-slate-800">
                      {u.name}
                    </td>
                    <td className="px-4 py-3 text-slate-600">{u.email}</td>
                    <td className="px-4 py-3">
                      <span
                        className={`inline-flex rounded-lg px-2.5 py-1 text-xs font-medium ${
                          roleStyles[u.role] || "bg-slate-100 text-slate-700"
                        }`}
                      >
                        {u.role}
                      </span>
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
          {users.length === 0 && (
            <div className="py-12 text-center text-slate-500">
              No users found.
            </div>
          )}
        </div>
      )}
    </div>
  );
}
