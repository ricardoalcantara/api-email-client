import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { usePostSmtpMutation } from "@/services";
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { getError } from "@/lib/error";
import { Alert, AlertDescription } from "@/components/ui/alert";
import { AlertCircle } from "lucide-react";
import SmtpForm, { formSchema } from "../components/form";
import * as z from "zod";

const SmtpCreate = () => {
  const navigate = useNavigate();
  const [errorMsg, setErrorMsg] = useState("");
  const [createSmtp, { isLoading }] = usePostSmtpMutation();

  async function onSubmit(values: z.infer<typeof formSchema>) {
    try {
      setErrorMsg("");
      await createSmtp(values).unwrap();
      navigate('/smtp');
    } catch (err) {
      setErrorMsg(getError(err));
    }
  }

  return (
    <div className="p-8">
      <Card>
        <CardHeader>
          <CardTitle>Add SMTP Configuration</CardTitle>
        </CardHeader>
        <CardContent>
          {errorMsg && (
            <Alert variant="destructive" className="mb-6">
              <AlertDescription className="flex items-center gap-2">
                <AlertCircle className="h-4 w-4" />
                {errorMsg}
              </AlertDescription>
            </Alert>
          )}
          <SmtpForm onSubmit={onSubmit} isLoading={isLoading} />
        </CardContent>
      </Card>
    </div>
  );
};

export default SmtpCreate;