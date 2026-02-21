import { useState, useEffect } from "react";
import { Link, useNavigate } from "react-router-dom";
import api from "../utils/api";

export default function CreateCourse() {
  const navigate = useNavigate();
  useEffect(() => {
    if (!localStorage.getItem("token")) navigate("/");
  }, [navigate]);
  const [name, setName] = useState("");
  const [code, setCode] = useState("");
  const [msg, setMsg] = useState("");
  const [loading, setLoading] = useState(false);

  const submit = async (e) => {
    e.preventDefault();
    setMsg("");
    setLoading(true);
    try {
      await api.post("/api/courses", { name, code });
      setMsg("Course created successfully!");
      setName("");
      setCode("");
    } catch (err) {
      setMsg(err.response?.data?.error || "Failed to create course");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="mx-auto max-w-md px-4 py-10 sm:px-6">
      <Link to="/admin" className="link-back mb-6 inline-flex">
        <span>←</span> Back to Dashboard
      </Link>

      <div className="card">
        <h1 className="text-2xl font-bold text-slate-800">Create Course</h1>
        <p className="mt-1 text-sm text-slate-500">
          Add a new course to the system
        </p>

        <form onSubmit={submit} className="mt-6 flex flex-col gap-4">
          <div>
            <label className="mb-1.5 block text-sm font-medium text-slate-700">
              Course Name
            </label>
            <input
              className="input-field"
              placeholder="e.g. Introduction to Programming"
              value={name}
              onChange={(e) => setName(e.target.value)}
              required
            />
          </div>
          <div>
            <label className="mb-1.5 block text-sm font-medium text-slate-700">
              Course Code
            </label>
            <input
              className="input-field"
              placeholder="e.g. CS101"
              value={code}
              onChange={(e) => setCode(e.target.value)}
              required
            />
          </div>

          <button
            type="submit"
            className="btn mt-2 w-full rounded-xl bg-emerald-600 py-3 font-semibold text-white shadow-md transition hover:bg-emerald-700 hover:shadow-lg disabled:opacity-50"
            disabled={loading}
          >
            {loading ? "Creating..." : "Create Course"}
          </button>
        </form>

        {msg && (
          <div
            className={`mt-4 rounded-xl px-4 py-3 text-sm ${
              msg.includes("Failed")
                ? "bg-red-50 text-red-600 border border-red-100"
                : "bg-emerald-50 text-emerald-700 border border-emerald-100"
            }`}
          >
            {msg}
          </div>
        )}
      </div>
    </div>
  );
}
