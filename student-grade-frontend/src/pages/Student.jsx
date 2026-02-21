import { useEffect } from "react";
import { useNavigate, Link } from "react-router-dom";
import StudentDashboard from "./StudentDashboard";

export default function Student() {
  const navigate = useNavigate();

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (!token) {
      navigate("/");
      return;
    }
    try {
      const payload = JSON.parse(atob(token.split(".")[1]));
      if (payload.role !== "student") {
        localStorage.removeItem("token");
        navigate("/");
      }
    } catch {
      navigate("/");
    }
  }, [navigate]);

  return (
    <div className="min-h-screen bg-slate-50/80">
      <nav className="sticky top-0 z-10 border-b border-slate-200/80 bg-white/90 backdrop-blur-md shadow-sm">
        <div className="mx-auto flex h-14 max-w-5xl items-center justify-between px-4 sm:px-6">
          <div className="flex items-center gap-2">
            <span className="flex h-8 w-8 items-center justify-center rounded-lg bg-sky-600 text-sm font-bold text-white shadow">
              S
            </span>
            <span className="font-semibold text-slate-800">Student Portal</span>
          </div>
          <Link
            to="/"
            onClick={() => localStorage.removeItem("token")}
            className="rounded-lg px-3 py-2 text-sm font-medium text-slate-600 transition hover:bg-slate-100 hover:text-slate-900"
          >
            Logout
          </Link>
        </div>
      </nav>
      <StudentDashboard />
    </div>
  );
}
