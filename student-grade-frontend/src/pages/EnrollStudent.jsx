import { useState, useEffect } from "react";
import { Link, useNavigate } from "react-router-dom";
import api from "../utils/api";

export default function EnrollStudent() {
  const navigate = useNavigate();
  useEffect(() => {
    if (!localStorage.getItem("token")) navigate("/");
  }, [navigate]);
  const [studentId, setStudentId] = useState("");
  const [courseId, setCourseId] = useState("");
  const [courses, setCourses] = useState([]);
  const [msg, setMsg] = useState("");
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    const fetchCourses = async () => {
      try {
        const res = await api.get("/api/courses");
        setCourses(Array.isArray(res.data) ? res.data : []);
      } catch {
        setCourses([]);
      }
    };
    fetchCourses();
  }, []);

  const submit = async (e) => {
    e.preventDefault();
    setMsg("");
    setLoading(true);
    try {
      await api.post(`/api/courses/${courseId}/enroll`, {
        student_id: studentId,
      });
      setMsg("Student enrolled successfully!");
      setStudentId("");
      setCourseId("");
    } catch (err) {
      setMsg(err.response?.data?.error || "Failed to enroll");
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
        <h1 className="text-2xl font-bold text-slate-800">Enroll Student</h1>
        <p className="mt-1 text-sm text-slate-500">
          Add a student to a course. Get Student IDs from List Users.
        </p>

        <form onSubmit={submit} className="mt-6 flex flex-col gap-4">
          <div>
            <label className="mb-1.5 block text-sm font-medium text-slate-700">
              Student ID
            </label>
            <input
              className="input-field font-mono text-sm"
              placeholder="MongoDB ObjectID"
              value={studentId}
              onChange={(e) => setStudentId(e.target.value)}
              required
            />
          </div>
          <div>
            <label className="mb-1.5 block text-sm font-medium text-slate-700">
              Course
            </label>
            <select
              className="input-field appearance-none bg-[length:1.25rem] bg-[right_0.75rem_center] bg-no-repeat pr-10"
              style={{
                backgroundImage: `url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' fill='none' viewBox='0 0 20 20'%3E%3Cpath stroke='%236b7280' stroke-linecap='round' stroke-linejoin='round' stroke-width='1.5' d='m6 8 4 4 4-4'/%3E%3C/svg%3E")`,
              }}
              value={courseId}
              onChange={(e) => setCourseId(e.target.value)}
              required
            >
              <option value="">Select course</option>
              {courses.map((c) => (
                <option key={c.id || c._id} value={c.id || c._id}>
                  {c.name} ({c.code || c.id})
                </option>
              ))}
            </select>
          </div>

          <button
            type="submit"
            className="btn mt-2 w-full rounded-xl bg-sky-600 py-3 font-semibold text-white shadow-md transition hover:bg-sky-700 hover:shadow-lg disabled:opacity-50"
            disabled={loading}
          >
            {loading ? "Enrolling..." : "Enroll Student"}
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
