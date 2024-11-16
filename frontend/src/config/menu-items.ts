import { Calendar, Home, Inbox, Search, Settings } from "lucide-react"

const MenuItems = [
  {
    title: "Home",
    url: "/",
    icon: Home,
  },
  {
    title: "Templates",
    url: "/template",
    icon: Inbox,
  },
  {
    title: "SMTP",
    url: "/smtp",
    icon: Calendar,
  },
  {
    title: "Emails",
    url: "/email",
    icon: Search,
  },
  {
    title: "Api Keys",
    url: "/api-key",
    icon: Settings,
  },
]

export default MenuItems;