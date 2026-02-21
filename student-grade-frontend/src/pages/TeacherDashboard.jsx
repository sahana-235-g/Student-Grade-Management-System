import { Link } from "react-router-dom";

export default function TeacherDashboard() {
  return (
    <div className="mx-auto max-w-3xl px-4 py-10 sm:px-6">
      <div className="mb-8">
        <h1 className="text-2xl font-bold text-slate-800 sm:text-3xl">
          Teacher Dashboard
        </h1>
        <p className="mt-1 text-slate-500">
          Assign and manage grades
        </p>
      </div>

      <Link
        to="/assign-grade"
        className="group relative block overflow-hidden rounded-2xl bg-gradient-to-br from-violet-500 to-violet-600 p-8 text-white shadow-lg transition duration-200 hover:from-violet-600 hover:to-violet-700 hover:shadow-xl"
      >
        <div className="relative z-10 flex items-center gap-5">
          <span className="flex h-16 w-16 shrink-0 items-center justify-center rounded-2xl bg-white/20 text-3xl backdrop-blur-sm">
            ✏️
          </span>
          <div>
            <h2 className="text-xl font-semibold">Assign Grade</h2>
            <p className="mt-1 text-sm text-white/90">
              Record a grade for a student in a course
            </p>
          </div>
        </div>
        <div className="absolute -right-6 -top-6 h-32 w-32 rounded-full bg-white/10" />
      </Link>
    </div>
  );
}
