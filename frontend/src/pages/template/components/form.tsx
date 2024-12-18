import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import * as z from "zod";
import { Button } from "@/components/ui/button";
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
import { TemplateDto } from "@/services/dto";
import { Link } from "react-router-dom";
import { useState } from "react";
import { Switch } from "@/components/ui/switch";

export const formSchema = z.object({
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
  template_html: z.string(),
  template_text: z.string(),
});

export interface TemplateFormProps {
  onSubmit: (values: z.infer<typeof formSchema>, generated: boolean) => void;
  isLoading: boolean;
  defaultValues?: TemplateDto;
  slug?: string
}

const TemplateForm = ({ onSubmit, isLoading, defaultValues, slug }: TemplateFormProps) => {
  const [generated, setGenerated] = useState(false)
  const form = useForm({
    resolver: zodResolver(formSchema),
    defaultValues: defaultValues ?? {
      name: "",
      slug: "",
      json_schema: "",
      subject: "",
      template_html: "",
      template_text: "",
    },
  });

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit((value) => onSubmit(value, generated))} className="space-y-6">
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
          <FormLabel className="flex gap-2 items-center">
            <span>
              Template Content -
            </span>
            {slug ? (

              <Button asChild variant='outline'>
                <Link to={`/template/${slug}/generator`}>
                  Regenerate
                </Link>
              </Button>
            ) : (
              <FormItem className="flex items-center gap-2">
                <FormLabel style={{ marginTop: 0 }}>Raw</FormLabel>
                <FormControl>
                  <Switch
                    checked={generated}
                    onCheckedChange={setGenerated}
                  />
                </FormControl>
                <FormLabel style={{ marginTop: 0 }}>Generated</FormLabel>
              </FormItem>
            )}
          </FormLabel>
          {!generated && (
            <Tabs defaultValue="html" className="w-full">
              <TabsList className="grid w-full grid-cols-3">
                <TabsTrigger value="html">HTML Template</TabsTrigger>
                <TabsTrigger value="text">Text Template</TabsTrigger>
                <TabsTrigger value="preview">Preview</TabsTrigger>
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
              <TabsContent value="preview" className="mt-4">
                <div
                  className="border p-4 rounded bg-white min-h-[200px]"
                  dangerouslySetInnerHTML={{ __html: form.getValues("template_html") }}
                />
              </TabsContent>
            </Tabs>
          )}
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
          <Button type="submit" disabled={isLoading}>Save</Button>
        </div>
      </form>
    </Form>)
}

export default TemplateForm;