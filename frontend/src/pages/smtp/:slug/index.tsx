import * as z from "zod";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { useGetSmtpQuery, usePostSmtpMutation, usePutSmtpMutation } from "@/services";
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { getError } from "@/lib/error";
import { Alert, AlertDescription } from "@/components/ui/alert";
import { AlertCircle } from "lucide-react";
import SmtpForm, { formSchema } from "../components/form";
import { useParams } from 'react-router-dom';

const SmtpEdit = () => {
  const { slug } = useParams<{ slug: string }>();
  const { data: smtp, isLoading: isLoadingSmtp } = useGetSmtpQuery(slug!);

  const navigate = useNavigate();

  const [errorMsg, setErrorMsg] = useState("");
  const [updateSmtp, { isLoading }] = usePutSmtpMutation();

  async function onSubmit(values: z.infer<typeof formSchema>) {
    try {
      setErrorMsg("")
      await updateSmtp({ slug: slug!, smtp: values as any }).unwrap();
      navigate(`/smtp`);
    } catch (err) {
      setErrorMsg(getError(err));
    }
  }

  return (
    <div className="p-8">
      <Card>
        <CardHeader>
          <CardTitle>Update Smtp: {slug}</CardTitle>
        </CardHeader>
        <CardContent>
          {errorMsg && (
            <Alert variant="destructive" className="">
              <AlertDescription className="flex items-center gap-2">
                <AlertCircle className="h-4 w-4" />
                {errorMsg}
              </AlertDescription>
            </Alert>
          )}
          {smtp && (<SmtpForm onSubmit={onSubmit} isLoading={isLoading || isLoadingSmtp} defaultValues={smtp} />)}
        </CardContent>
      </Card>
    </div>
  );
};

export default SmtpEdit;
