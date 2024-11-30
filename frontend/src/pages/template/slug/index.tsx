import * as z from "zod";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { useGetTemplateQuery, usePostTemplateMutation, usePutTemplateMutation } from "@/services";
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { getError } from "@/lib/error";
import { Alert, AlertDescription } from "@/components/ui/alert";
import { AlertCircle } from "lucide-react";
import TemplateForm, { formSchema } from "../components/form";
import { useParams } from 'react-router-dom';

const TemplateEdit = () => {
  const { slug } = useParams<{ slug: string }>();
  const { data: template, isLoading: isLoadingTemplate } = useGetTemplateQuery(slug!, {
    refetchOnMountOrArgChange: true,
  });

  const navigate = useNavigate();

  const [errorMsg, setErrorMsg] = useState("");
  const [updateTemplate, { isLoading }] = usePutTemplateMutation();

  async function onSubmit(values: z.infer<typeof formSchema>) {
    try {
      setErrorMsg("")
      await updateTemplate({ slug: slug!, template: values as any }).unwrap();
      navigate(`/template`);
    } catch (err) {
      setErrorMsg(getError(err));
    }
  }

  return (
    <div className="p-8">
      <Card>
        <CardHeader>
          <CardTitle>Update Template: {slug}</CardTitle>
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
          {template && (<TemplateForm onSubmit={onSubmit} isLoading={isLoading || isLoadingTemplate} defaultValues={template} slug={slug} />)}
        </CardContent>
      </Card>
    </div>
  );
};

export default TemplateEdit;
