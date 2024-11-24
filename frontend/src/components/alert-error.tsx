import { AlertCircle } from "lucide-react";
import { Alert, AlertDescription } from "./ui/alert";

const AlertError = ({ error }: { error: string | null | undefined }) => {
  if (!error) return null;
  return (
    <Alert variant="destructive">
      <AlertDescription className="flex items-center gap-2">
        <AlertCircle className="h-4 w-4" />
        {error}
      </AlertDescription>
    </Alert>)
}

export default AlertError;