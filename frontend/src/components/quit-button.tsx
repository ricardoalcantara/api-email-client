// src/components/ThemeToggle.tsx
import { LogOut } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Link } from "react-router-dom";

export function QuitButton() {

  return (
    <Button
      variant="ghost"
      size="icon"
      className="h-8 w-8"
      asChild
    >
      <Link to="/login"><LogOut className="h-4 w-4" /></Link>
    </Button>
  );
}
