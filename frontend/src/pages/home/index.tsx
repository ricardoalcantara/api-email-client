import { Button } from "@/components/ui/button"
import { useTheme } from "@/hooks/useTheme";

import { Link } from "react-router-dom";

export function Home() {
  const [theme, setTheme] = useTheme();
  const nextThema = theme === "dark" ? "light" : "dark";

  return (
    <>
      <h1 className="text-3xl font-bold underline">Hello world!</h1>
      <Button asChild variant="outline">
        <Link to="/login">Login</Link>
      </Button>

      <Button variant="default" onClick={() => setTheme(nextThema)}>
        Switch Theme to {nextThema}
      </Button>
    </>
  );
}
