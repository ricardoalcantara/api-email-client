import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import * as z from "zod";
import { Button } from "@/components/ui/button";
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
import { Input } from "@/components/ui/input";
import { Textarea } from "@/components/ui/textarea";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { usePostTemplateMutation } from "@/services";
import { useState } from "react";
import { useNavigate } from "react-router-dom";

const formSchema = z.object({
  name: z.string().min(2, {
    message: "Name must be at least 2 characters.",
  }),
  slug: z.string().min(2, {
    message: "Slug must be at least 2 characters.",
  }),
  subject: z.string().min(2, {
    message: "Subject must be at least 2 characters.",
  }),
  json_schema: z
    .string()
    .min(2, {
      message: "JSON Schema is required.",
    })
    .refine(
      (val) => {
        try {
          JSON.parse(val);
          return true;
        } catch {
          return false;
        }
      },
      {
        message: "Invalid JSON Schema format.",
      }
    ),
  template_html: z.string().min(2, {
    message: "HTML Template is required.",
  }),
  template_text: z.string().min(2, {
    message: "Text Template is required.",
  }),
});

const TemplateCreate = () => {
  const navigate = useNavigate();
  const form = useForm({
    resolver: zodResolver(formSchema),
    defaultValues: {
      name: "Welcome Email",
      slug: "welcome-email",
      json_schema:
        '{"type": "object", "properties": {"name": {"type": "string"}, "email": {"type": "string"}}}',
      subject: "Welcome to Our Service!",
      template_html:
        "<html><body><h1>Welcome, {{name}}!</h1><p>We're glad to have you.</p></body></html>",
      template_text: "Welcome, {{name}}!\nWe're glad to have you.",
    },
  });

  const [errorMsg, setErrorMsg] = useState("");
  const [createTemplate, { isLoading, error }] = usePostTemplateMutation();

  async function onSubmit(values: z.infer<typeof formSchema>) {
    try {
      const response = await createTemplate(values as any).unwrap();
      // navigate(`/templates/${response.slug}`);
      navigate(`/template`);
    } catch (err) {
      setErrorMsg("Error");
    }
  }

  return (
    <div className="p-8">
      <Card>
        <CardHeader>
          <CardTitle>Create Template</CardTitle>
        </CardHeader>
        <CardContent>
          <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-6">
              <div className="grid gap-6 md:grid-cols-2">
                <FormField
                  control={form.control}
                  name="name"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Template Name</FormLabel>
                      <FormControl>
                        <Input
                          placeholder="Enter template name..."
                          {...field}
                        />
                      </FormControl>
                      <FormDescription>
                        The name of your email template.
                      </FormDescription>
                      <FormMessage />
                    </FormItem>
                  )}
                />
                <FormField
                  control={form.control}
                  name="slug"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Slug</FormLabel>
                      <FormControl>
                        <Input placeholder="template-name" {...field} />
                      </FormControl>
                      <FormDescription>
                        A unique identifier for this template
                      </FormDescription>
                      <FormMessage />
                    </FormItem>
                  )}
                />
              </div>

              <FormField
                control={form.control}
                name="subject"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Subject</FormLabel>
                    <FormControl>
                      <Input placeholder="Enter email subject..." {...field} />
                    </FormControl>
                    <FormDescription>
                      The subject line of the email.
                    </FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />

              <FormField
                control={form.control}
                name="json_schema"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>JSON Schema</FormLabel>
                    <FormControl>
                      <Textarea
                        placeholder="Enter JSON schema..."
                        className="font-mono"
                        rows={4}
                        {...field}
                      />
                    </FormControl>
                    <FormDescription>
                      The JSON schema defining the template variables.
                    </FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />

              <div className="space-y-4">
                <FormLabel>Template Content</FormLabel>
                <Tabs defaultValue="html" className="w-full">
                  <TabsList className="grid w-full grid-cols-2">
                    <TabsTrigger value="html">HTML Template</TabsTrigger>
                    <TabsTrigger value="text">Text Template</TabsTrigger>
                  </TabsList>
                  <TabsContent value="html" className="mt-4">
                    <FormField
                      control={form.control}
                      name="template_html"
                      render={({ field }) => (
                        <FormItem>
                          <FormControl>
                            <Textarea
                              placeholder="Enter HTML template..."
                              className="font-mono min-h-[300px]"
                              {...field}
                            />
                          </FormControl>
                          <FormDescription>
                            The HTML version of your email template.
                          </FormDescription>
                          <FormMessage />
                        </FormItem>
                      )}
                    />
                  </TabsContent>
                  <TabsContent value="text" className="mt-4">
                    <FormField
                      control={form.control}
                      name="template_text"
                      render={({ field }) => (
                        <FormItem>
                          <FormControl>
                            <Textarea
                              placeholder="Enter text template..."
                              className="min-h-[300px]"
                              {...field}
                            />
                          </FormControl>
                          <FormDescription>
                            The plain text version of your email template.
                          </FormDescription>
                          <FormMessage />
                        </FormItem>
                      )}
                    />
                  </TabsContent>
                </Tabs>
              </div>

              <div className="flex justify-end space-x-4">
                <Button
                  variant="outline"
                  type="button"
                  onClick={() => form.reset()}
                  disabled={isLoading}
                >
                  Reset
                </Button>
                <Button type="submit">Create Template</Button>
              </div>
            </form>
          </Form>
        </CardContent>
      </Card>
    </div>
  );
};

export default TemplateCreate;
