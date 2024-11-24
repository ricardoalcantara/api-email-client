import { useState } from "react";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import * as z from "zod";
import { useNavigate } from "react-router-dom";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Textarea } from "@/components/ui/textarea";
import { Alert, AlertDescription } from "@/components/ui/alert";
import { AlertCircle } from "lucide-react";
import { useListSmtpQuery, useListTemplateQuery, useSendEmailMutation } from "@/services";
import { getError } from "@/lib/error";

const formSchema = z.object({
  template_slug: z.string().min(1, "Template is required"),
  smtp_slug: z.string().min(1, "SMTP configuration is required"),
  to: z.string().email("Invalid email address"),
  subject: z.string().optional(),
  data: z.string().refine(
    (val) => {
      try {
        if (val === "") return true;
        JSON.parse(val);
        return true;
      } catch {
        return false;
      }
    },
    { message: "Invalid JSON format" }
  ),
});

const SendEmailForm = () => {
  const navigate = useNavigate();
  const [errorMsg, setErrorMsg] = useState("");
  const { data: templates, isLoading: templatesLoading } = useListTemplateQuery();
  const { data: smtps, isLoading: smtpLoading } = useListSmtpQuery();
  const [sendEmail, { isLoading: sendingEmail }] = useSendEmailMutation();

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      template_slug: "",
      smtp_slug: "",
      to: "",
      subject: "",
      data: "",
    },
  });

  async function onSubmit(values: z.infer<typeof formSchema>) {
    try {
      setErrorMsg("");
      const formattedValues = {
        ...values,
        data: values.data ? JSON.parse(values.data) : {},
      };
      await sendEmail(formattedValues).unwrap();
      navigate("/email"); // Adjust navigation path as needed
    } catch (err) {
      setErrorMsg(getError(err));
    }
  }

  const isLoading = templatesLoading || smtpLoading || sendingEmail;

  return (
    <div className="p-8">
      <Card>
        <CardHeader>
          <CardTitle>Send Email</CardTitle>
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

          <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-6">
              <div className="grid gap-6 md:grid-cols-2">
                <FormField
                  control={form.control}
                  name="template_slug"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Template</FormLabel>
                      <Select
                        disabled={isLoading}
                        onValueChange={field.onChange}
                        value={field.value}
                      >
                        <FormControl>
                          <SelectTrigger>
                            <SelectValue placeholder="Select a template" />
                          </SelectTrigger>
                        </FormControl>
                        <SelectContent>
                          {templates?.list?.map((template) => (
                            <SelectItem
                              key={template.slug}
                              value={template.slug}
                            >
                              {template.name}
                            </SelectItem>
                          ))}
                        </SelectContent>
                      </Select>
                      <FormDescription>
                        Select the email template to use
                      </FormDescription>
                      <FormMessage />
                    </FormItem>
                  )}
                />

                <FormField
                  control={form.control}
                  name="smtp_slug"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>SMTP Configuration</FormLabel>
                      <Select
                        disabled={isLoading}
                        onValueChange={field.onChange}
                        value={field.value}
                      >
                        <FormControl>
                          <SelectTrigger>
                            <SelectValue placeholder="Select SMTP config" />
                          </SelectTrigger>
                        </FormControl>
                        <SelectContent>
                          {smtps?.list?.map((smtp) => (
                            <SelectItem
                              key={smtp.slug}
                              value={smtp.slug}
                            >
                              {smtp.name}
                            </SelectItem>
                          ))}
                        </SelectContent>
                      </Select>
                      <FormDescription>
                        Choose the SMTP configuration to use
                      </FormDescription>
                      <FormMessage />
                    </FormItem>
                  )}
                />
              </div>

              <div className="grid gap-6 md:grid-cols-2">
                <FormField
                  control={form.control}
                  name="to"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>To</FormLabel>
                      <FormControl>
                        <Input
                          placeholder="recipient@example.com"
                          {...field}
                          disabled={isLoading}
                        />
                      </FormControl>
                      <FormDescription>
                        Email address of the recipient
                      </FormDescription>
                      <FormMessage />
                    </FormItem>
                  )}
                />

                <FormField
                  control={form.control}
                  name="subject"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Subject (Optional)</FormLabel>
                      <FormControl>
                        <Input
                          placeholder="Email subject"
                          {...field}
                          disabled={isLoading}
                        />
                      </FormControl>
                      <FormDescription>
                        Override template subject if needed
                      </FormDescription>
                      <FormMessage />
                    </FormItem>
                  )}
                />
              </div>

              <FormField
                control={form.control}
                name="data"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Template Data (JSON)</FormLabel>
                    <FormControl>
                      <Textarea
                        placeholder='{"name": "John", "code": "123"}'
                        className="font-mono"
                        {...field}
                        disabled={isLoading}
                      />
                    </FormControl>
                    <FormDescription>
                      Enter the template variables as JSON
                    </FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />

              <div className="flex justify-end space-x-4">
                <Button
                  variant="outline"
                  type="button"
                  onClick={() => form.reset()}
                  disabled={isLoading}
                >
                  Reset
                </Button>
                <Button type="submit" disabled={isLoading}>
                  Send Email
                </Button>
              </div>
            </form>
          </Form>
        </CardContent>
      </Card>
    </div>
  );
};

export default SendEmailForm;