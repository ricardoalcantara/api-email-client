import * as z from "zod";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { usePostTemplateMutation } from "@/services";
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { getError } from "@/lib/error";
import { Alert, AlertDescription } from "@/components/ui/alert";
import { AlertCircle } from "lucide-react";
import TemplateForm, { formSchema } from "../components/form";

const TemplateCreate = () => {
  const navigate = useNavigate();

  const [errorMsg, setErrorMsg] = useState("");
  const [createTemplate, { isLoading }] = usePostTemplateMutation();

  async function onSubmit(values: z.infer<typeof formSchema>, generated: boolean) {
    try {
      setErrorMsg("")
      const response = await createTemplate(values as any).unwrap();
      if (generated) {
        navigate(`/template/${response.slug}/generator`);
      } else {
        navigate(`/template`);
      }
    } catch (err) {
      setErrorMsg(getError(err));
    }
  }

  return (
    <div className="p-8">
      <Card>
        <CardHeader>
          <CardTitle>Create Template</CardTitle>
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
          <TemplateForm onSubmit={onSubmit} isLoading={isLoading} />
        </CardContent>
      </Card>
    </div>
  );
};

export default TemplateCreate;
