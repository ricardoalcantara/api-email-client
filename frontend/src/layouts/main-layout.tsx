// src/layouts/DashboardLayout.tsx
import { Link, Outlet, useLocation } from "react-router-dom"
import { cn } from "@/lib/utils"
import {
  LayoutDashboard,
  Mail,
  FileCode,
  Server,
  Key
} from "lucide-react"

export default function DashboardLayout() {
  const location = useLocation()

  const menuItems = [
    { path: "/", label: "Dashboard", icon: LayoutDashboard },
    { path: "/templates", label: "Templates", icon: FileCode },
    { path: "/smtp", label: "SMTP", icon: Server },
    { path: "/emails", label: "Emails", icon: Mail },
    { path: "/api-keys", label: "API Keys", icon: Key },
  ]

  return (
    <div className="min-h-screen flex dark:bg-gray-950">
      {/* Sidebar */}
      <aside className="w-64 border-r dark:border-gray-800">
        <nav className="space-y-1 px-2">
          {menuItems.map((item) => {
            const Icon = item.icon
            return (
              <Link
                key={item.path}
                to={item.path}
                className={cn(
                  "flex items-center gap-3 px-3 py-2 rounded-md transition-colors",
                  "hover:bg-gray-100 dark:hover:bg-gray-800",
                  location.pathname === item.path && "bg-gray-100 dark:bg-gray-800"
                )}
              >
                <Icon className="h-5 w-5" />
                <span>{item.label}</span>
              </Link>
            )
          })}
        </nav>
      </aside>

      {/* Main Content */}
      <main className="flex-1">
        <Outlet />
      </main>
    </div>
  )
}