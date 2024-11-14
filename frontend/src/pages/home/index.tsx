import { Button } from "@/components/ui/button"

import { Link } from "react-router-dom";

export function Home() {
  return (
    <>
      <h1 className="text-3xl font-bold underline">Hello world!</h1>
      <Button asChild variant="outline">
        <Link to="/login">Login</Link>
      </Button>
    </>
  );
}
