import { useEffect, useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import api from "../utils/api";

function GradeBadge({ grade }) {
  const n = Number(grade);
  const isHigh = !Number.isNaN(n) && n >= 80;
  const isMid = !Number.isNaN(n) && n >= 60 && n < 80;
  const color = isHigh
    ? "bg-emerald-100 text-emerald-800"
    : isMid
      ? "bg-amber-100 text-amber-800"
      : "bg-slate-100 text-slate-700";
  return (
    <span
      className={`inline-flex items-center rounded-lg px-3 py-1 text-sm font-semibold ${color}`}
    >
      {grade}
    </span>
  );
}

export default function ViewGrades() {
  const navigate = useNavigate();
  const [grades, setGrades] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (!token) {
      navigate("/");
      return;
    }

    const fetchGrades = async () => {
      try {
        const payload = JSON.parse(atob(token.split(".")[1]));
        const studentId = payload.user_id;

        const res = await api.get(`/api/grades/${studentId}`);
        setGrades(Array.isArray(res.data) ? res.data : []);
      } catch (err) {
        setError(err.response?.data?.error || "Failed to fetch grades");
        setGrades([]);
      } finally {
        setLoading(false);
      }
    };

    fetchGrades();
  }, [navigate]);

  if (loading) {
    return (
      <div className="mx-auto max-w-2xl px-4 py-16 text-center">
        <div className="inline-block h-8 w-8 animate-spin rounded-full border-2 border-indigo-500 border-t-transparent" />
        <p className="mt-4 text-slate-500">Loading grades...</p>
      </div>
    );
  }

  if (error) {
    return (
      <div className="mx-auto max-w-md px-4 py-10">
        <div className="card rounded-2xl border border-red-100 bg-red-50/50">
          <p className="text-red-600">{error}</p>
          <Link to="/student" className="link-back mt-4 inline-flex">
            ← Back to Dashboard
          </Link>
        </div>
      </div>
    );
  }

  return (
    <div className="mx-auto max-w-2xl px-4 py-10 sm:px-6">
      <Link to="/student" className="link-back mb-6 inline-flex">
        <span>←</span> Back to Dashboard
      </Link>

      <div className="mb-6">
        <h1 className="text-2xl font-bold text-slate-800">My Grades</h1>
        <p className="mt-1 text-slate-500">
          Your grades for enrolled courses
        </p>
      </div>

      {grades.length === 0 ? (
        <div className="card text-center py-12">
          <div className="mx-auto flex h-16 w-16 items-center justify-center rounded-2xl bg-slate-100 text-3xl">
            📋
          </div>
          <p className="mt-4 font-medium text-slate-700">No grades yet</p>
          <p className="mt-1 text-sm text-slate-500">
            Grades will appear here once your teacher assigns them.
          </p>
        </div>
      ) : (
        <ul className="space-y-3">
          {grades.map((g) => (
            <li
              key={g.id || g._id || g.student_id + g.course_id}
              className="card flex flex-wrap items-center justify-between gap-4 border border-slate-100 transition hover:border-slate-200 hover:shadow-md"
            >
              <div className="min-w-0">
                <p className="text-xs font-medium uppercase tracking-wide text-slate-400">
                  Course ID
                </p>
                <p className="font-mono text-sm text-slate-700 break-all">
                  {g.course_id || "N/A"}
                </p>
              </div>
              <div className="flex items-center gap-2">
                <GradeBadge grade={g.grade} />
              </div>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}
