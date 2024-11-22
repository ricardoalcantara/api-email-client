import React from 'react';
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import * as z from "zod";
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs';
import { Button } from '@/components/ui/button';
import { Switch } from '@/components/ui/switch';
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
import { Textarea } from "@/components/ui/textarea";
import { Plus, Minus } from 'lucide-react';
import { useNavigate, useParams } from 'react-router-dom';
import { useGetTemplateQuery, usePutTemplateMutation } from '@/services';

// Zod schema for form validation
const buttonSchema = z.object({
  color: z.string(),
  textColor: z.string(),
  text: z.string(),
  link: z.string().url().optional(),
});

const actionSchema = z.object({
  instructions: z.string(),
  button: buttonSchema,
  inviteCode: z.string(),
});

const dictionaryItemSchema = z.object({
  key: z.string(),
  value: z.string(),
});

const productSchema = z.object({
  name: z.string().min(1, "Product name is required"),
  link: z.string().url().optional(),
  logo: z.string().url().optional(),
  copyright: z.string(),
  troubleText: z.string(),
});

const formSchema = z.object({
  config: z.object({
    theme: z.enum(["Default", "Flat"]),
    textDirection: z.enum(["ltr", "rtl"]),
    product: productSchema,
    disableCSSInlining: z.boolean(),
  }),
  emailContent: z.object({
    name: z.string().min(1, "Template name is required"),
    title: z.string().min(1, "Email title is required"),
    intros: z.array(z.string()),
    dictionary: z.array(dictionaryItemSchema),
    actions: z.array(actionSchema),
    outros: z.array(z.string()),
    greeting: z.string(),
    signature: z.string(),
    freeMarkdown: z.string(),
  }),
});

type FormValues = z.infer<typeof formSchema>;

const defaultValues: FormValues = {
  config: {
    theme: "Default",
    textDirection: "ltr",
    product: {
      name: "",
      link: "",
      logo: "",
      copyright: "",
      troubleText: ""
    },
    disableCSSInlining: false
  },
  emailContent: {
    name: "",
    title: "",
    intros: [""],
    dictionary: [{ key: "", value: "" }],
    actions: [{
      instructions: "",
      button: {
        color: "#000000",
        textColor: "#ffffff",
        text: "",
        link: ""
      },
      inviteCode: ""
    }],
    outros: [""],
    greeting: "",
    signature: "",
    freeMarkdown: ""
  }
};

const EmailTemplateGenerator: React.FC = () => {
  const { slug } = useParams<{ slug: string }>();
  const { data: template, isLoading: isLoadingTemplate } = useGetTemplateQuery(slug!);
  const [updateTemplate, { isLoading }] = usePutTemplateMutation();
  const [genarated, setGenarated] = React.useState(false);

  const navigate = useNavigate();


  const form = useForm<FormValues>({
    resolver: zodResolver(formSchema),
    defaultValues,
  });

  const [generatedOutput, setGeneratedOutput] = React.useState({
    template_html: '',
    template_text: ''
  });

  const onSubmit = async (data: FormValues) => {
    console.log('Form submitted:', data);
    // Simulate generation
    setGeneratedOutput({
      template_html: '<div>Generated HTML based on form data</div>',
      template_text: 'Generated plain text based on form data'
    });
  };

  const onSave = async () => {
    await updateTemplate({ slug: slug!, template: generatedOutput }).unwrap();
    navigate(`/template/${slug}`);
  };

  const addListItem = (fieldName: 'intros' | 'outros' | 'dictionary' | 'actions', defaultValue: any) => {
    const currentValue = form.getValues(`emailContent.${fieldName}`);
    form.setValue(`emailContent.${fieldName}`, [...currentValue, defaultValue]);
  };

  const removeListItem = (fieldName: 'intros' | 'outros' | 'dictionary' | 'actions', index: number) => {
    const currentValue = form.getValues(`emailContent.${fieldName}`) as Array<any>;
    form.setValue(
      `emailContent.${fieldName}`,
      currentValue.filter((_, i) => i !== index)
    );
  };

  return (
    <div className="w-full max-w-6xl mx-auto space-y-4 p-4">
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)}>
          {/* Configuration Card */}
          <Card>
            <CardHeader>
              <CardTitle>Generate Template For {template?.name} - {template?.subject}</CardTitle>
            </CardHeader>
            <CardContent className="space-y-4">
              <div className="grid grid-cols-2 gap-4">
                <FormField
                  control={form.control}
                  name="config.theme"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Theme</FormLabel>
                      <Select
                        onValueChange={field.onChange}
                        defaultValue={field.value}
                      >
                        <FormControl>
                          <SelectTrigger>
                            <SelectValue placeholder="Select theme" />
                          </SelectTrigger>
                        </FormControl>
                        <SelectContent>
                          <SelectItem value="Default">Default</SelectItem>
                          <SelectItem value="Flat">Flat</SelectItem>
                        </SelectContent>
                      </Select>
                      <FormMessage />
                    </FormItem>
                  )}
                />

                <FormField
                  control={form.control}
                  name="config.textDirection"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Text Direction</FormLabel>
                      <Select
                        onValueChange={field.onChange}
                        defaultValue={field.value}
                      >
                        <FormControl>
                          <SelectTrigger>
                            <SelectValue placeholder="Select direction" />
                          </SelectTrigger>
                        </FormControl>
                        <SelectContent>
                          <SelectItem value="ltr">LTR</SelectItem>
                          <SelectItem value="rtl">RTL</SelectItem>
                        </SelectContent>
                      </Select>
                      <FormMessage />
                    </FormItem>
                  )}
                />
              </div>

              {/* Product Details */}
              <div className="space-y-4">
                <FormField
                  control={form.control}
                  name="config.product.name"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Product Name</FormLabel>
                      <FormControl>
                        <Input {...field} />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />

                {/* Add similar FormField components for other product fields */}
                {/* Link, Logo, Copyright, TroubleText */}

                <FormField
                  control={form.control}
                  name="config.disableCSSInlining"
                  render={({ field }) => (
                    <FormItem className="flex items-center gap-2">
                      <FormControl>
                        <Switch
                          checked={field.value}
                          onCheckedChange={field.onChange}
                        />
                      </FormControl>
                      <FormLabel>Disable CSS Inlining</FormLabel>
                      <FormMessage />
                    </FormItem>
                  )}
                />
              </div>
            </CardContent>
          </Card>

          {/* Email Content Card */}
          <Card className="mt-4">
            <CardHeader>
              <CardTitle>Email Content</CardTitle>
            </CardHeader>
            <CardContent className="space-y-4">
              {/* Basic Info */}
              <div className="grid grid-cols-2 gap-4">
                <FormField
                  control={form.control}
                  name="emailContent.name"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Template Name</FormLabel>
                      <FormControl>
                        <Input {...field} />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />

                <FormField
                  control={form.control}
                  name="emailContent.title"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Email Title</FormLabel>
                      <FormControl>
                        <Input {...field} />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />
              </div>

              {/* Intros */}
              <div>
                <div className="flex justify-between items-center mb-2">
                  <FormLabel>Intros</FormLabel>
                  <Button
                    type="button"
                    variant="outline"
                    onClick={() => addListItem('intros', '')}
                  >
                    <Plus className="h-4 w-4 mr-2" /> Add Intro
                  </Button>
                </div>
                {form.watch('emailContent.intros').map((_, index) => (
                  <FormField
                    key={index}
                    control={form.control}
                    name={`emailContent.intros.${index}`}
                    render={({ field }) => (
                      <FormItem className="flex gap-2 mt-2">
                        <FormControl>
                          <Input {...field} />
                        </FormControl>
                        <Button
                          type="button"
                          variant="outline"
                          size="icon"
                          onClick={() => removeListItem('intros', index)}
                        >
                          <Minus className="h-4 w-4" />
                        </Button>
                        <FormMessage />
                      </FormItem>
                    )}
                  />
                ))}
              </div>

              {/* Similar patterns for Dictionary, Actions, and Outros */}
              {/* ... */}

              {/* Submit Button */}
              <div className="flex justify-end gap-2">
                <Button type="submit">Generate Email</Button>
              </div>
            </CardContent>
          </Card>
        </form>
      </Form>

      {/* Generated Output Card */}
      <Card>
        <CardHeader>
          <CardTitle>Generated Output</CardTitle>
        </CardHeader>
        <CardContent className="space-y-4">
          <Tabs defaultValue="html">
            <TabsList>
              <TabsTrigger value="html">HTML</TabsTrigger>
              <TabsTrigger value="text">Text</TabsTrigger>
              <TabsTrigger value="preview">Preview</TabsTrigger>
            </TabsList>

            <TabsContent value="html">
              <Textarea
                value={generatedOutput.template_html}
                readOnly
                className="font-mono h-64"
              />
            </TabsContent>

            <TabsContent value="text">
              <Textarea
                value={generatedOutput.template_text}
                readOnly
                className="font-mono h-64"
              />
            </TabsContent>

            <TabsContent value="preview">
              <div
                className="border p-4 rounded bg-white min-h-[200px]"
                dangerouslySetInnerHTML={{ __html: generatedOutput.template_html }}
              />
            </TabsContent>
          </Tabs>

          <div className="flex justify-end gap-2">
            <Button onClick={onSave} disabled={isLoading || !genarated}>
              Save Template
            </Button>
          </div>
        </CardContent>
      </Card>
    </div>
  );
};

export default EmailTemplateGenerator;