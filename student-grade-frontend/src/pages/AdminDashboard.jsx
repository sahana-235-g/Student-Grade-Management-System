import { Link } from "react-router-dom";

const actions = [
  {
    to: "/create-course",
    label: "Create Course",
    description: "Add a new course to the system",
    icon: "📚",
    className: "from-emerald-500 to-emerald-600 hover:from-emerald-600 hover:to-emerald-700 shadow-emerald-500/25",
  },
  {
    to: "/enroll-student",
    label: "Enroll Student",
    description: "Enroll a student in a course",
    icon: "➕",
    className: "from-sky-500 to-sky-600 hover:from-sky-600 hover:to-sky-700 shadow-sky-500/25",
  },
  {
    to: "/list-users",
    label: "List Users",
    description: "View all users and copy Student IDs",
    icon: "👥",
    className: "from-slate-600 to-slate-700 hover:from-slate-700 hover:to-slate-800 shadow-slate-500/25",
  },
];

export default function AdminDashboard() {
  return (
    <div className="mx-auto max-w-3xl px-4 py-10 sm:px-6">
      <div className="mb-8">
        <h1 className="text-2xl font-bold text-slate-800 sm:text-3xl">
          Admin Dashboard
        </h1>
        <p className="mt-1 text-slate-500">
          Manage courses and enrollments
        </p>
      </div>

      <div className="grid gap-4 sm:grid-cols-1">
        {actions.map((action) => (
          <Link
            key={action.to}
            to={action.to}
            className={`group relative overflow-hidden rounded-2xl bg-gradient-to-br ${action.className} p-6 text-white shadow-lg transition duration-200 hover:shadow-xl`}
          >
            <div className="relative z-10 flex items-start gap-4">
              <span className="flex h-12 w-12 shrink-0 items-center justify-center rounded-xl bg-white/20 text-2xl backdrop-blur-sm">
                {action.icon}
              </span>
              <div>
                <h2 className="font-semibold text-lg">{action.label}</h2>
                <p className="mt-0.5 text-sm text-white/90">
                  {action.description}
                </p>
              </div>
            </div>
            <div className="absolute -right-4 -top-4 h-24 w-24 rounded-full bg-white/10" />
          </Link>
        ))}
      </div>
    </div>
  );
}
