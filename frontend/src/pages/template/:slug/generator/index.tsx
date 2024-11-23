import React, { useEffect } from 'react';
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
import { useGenerateTemplateMutation, useGetTemplateQuery, usePatchTemplateMutation } from '@/services';
import { RequestTemplateGeneratorDto } from '@/services/dto';

// Zod schemas
const buttonSchema = z.object({
  Color: z.string(),
  TextColor: z.string(),
  Text: z.string(),
  Link: z.string().url().optional(),
});

const actionSchema = z.object({
  Instructions: z.string(),
  Button: buttonSchema,
  InviteCode: z.string(),
});

const tableRowSchema = z.object({
  Key: z.string(),
  Value: z.string(),
});

const productSchema = z.object({
  Name: z.string().min(1, "Product name is required"),
  Link: z.string().url().optional(),
  Logo: z.string().url().optional(),
  Copyright: z.string(),
  TroubleText: z.string(),
});

const configSchema = z.object({
  TextDirection: z.enum(["ltr", "rtl"]),
  Product: productSchema,
  DisableCSSInlining: z.boolean(),
});

const emailSchema = z.object({
  Name: z.string().min(1, "Template name is required"),
  Title: z.string(),
  Intros: z.array(z.string()),
  Dictionary: z.array(z.object({
    Key: z.string(),
    Value: z.string(),
  })),
  Actions: z.array(actionSchema),
  Outros: z.array(z.string()),
  Greeting: z.string(),
  Signature: z.string(),
  FreeMarkdown: z.string(),
});

const formSchema = z.object({
  theme: z.enum(["Default", "Flat"]).optional(),
  config: configSchema,
  email: emailSchema,
});

// Schemas and types remain the same until defaultValues
const defaultValues: RequestTemplateGeneratorDto = {
  theme: "Default",
  config: {
    TextDirection: "ltr",
    Product: {
      Name: "",
      Link: "",
      Logo: "",
      Copyright: "",
      TroubleText: ""
    },
    DisableCSSInlining: false
  },
  email: {
    Name: "",
    Title: "",
    Intros: [""],
    Dictionary: [],
    Actions: [],
    Outros: [],
    Greeting: "",
    Signature: "",
    FreeMarkdown: ""
  }
};

const EmailTemplateGenerator: React.FC = () => {
  const { slug } = useParams<{ slug: string }>();
  const { data: template, isLoading: isLoadingTemplate } = useGetTemplateQuery(slug!);
  const [updateTemplate, { isLoading }] = usePatchTemplateMutation();
  const [generateTemplate, { isLoading: isLoadingGen }] = useGenerateTemplateMutation();
  const [generated, setGenerated] = React.useState(false);
  const navigate = useNavigate();

  const form = useForm<RequestTemplateGeneratorDto>({
    resolver: zodResolver(formSchema),
    defaultValues,
  });

  const { formState: { errors } } = form;

  useEffect(() => {
    if (errors) {
      console.log(errors);
    }
  }, [errors]);

  const [generatedOutput, setGeneratedOutput] = React.useState({
    template_html: '',
    template_text: ''
  });

  const onSubmit = async (data: RequestTemplateGeneratorDto) => {
    try {
      const response = await generateTemplate(data).unwrap();
      setGeneratedOutput(response);
      setGenerated(true);
    } catch (err) {
      console.error(err);
    }
  };

  const onSave = async () => {
    await updateTemplate({ slug: slug!, template: generatedOutput }).unwrap();
    navigate(`/template/${slug}`);
  };

  const addListItem = (fieldName: 'Intros' | 'Outros' | 'Dictionary' | 'Actions', defaultValue: any) => {
    const currentValue = form.getValues(`email.${fieldName}`);
    form.setValue(`email.${fieldName}`, [...currentValue, defaultValue]);
  };

  const removeListItem = (fieldName: 'Intros' | 'Outros' | 'Dictionary' | 'Actions', index: number) => {
    const currentValue = form.getValues(`email.${fieldName}`) as Array<any>;
    form.setValue(
      `email.${fieldName}`,
      currentValue.filter((_: unknown, i: number) => i !== index)
    );
  };

  return (
    <div className="w-full max-w-6xl mx-auto space-y-4 p-4">
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4">
          {/* Configuration Card */}
          <Card>
            <CardHeader>
              <CardTitle>Generate Template For {template?.name} - {template?.subject}</CardTitle>
            </CardHeader>
            <CardContent className="space-y-6">
              <div className="grid grid-cols-2 gap-4">
                <FormField
                  control={form.control}
                  name="theme"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Theme</FormLabel>
                      <Select onValueChange={field.onChange} defaultValue={field.value}>
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
                  name="config.TextDirection"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Text Direction</FormLabel>
                      <Select onValueChange={field.onChange} defaultValue={field.value}>
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
                  name="config.Product.Name"
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

                <FormField
                  control={form.control}
                  name="config.Product.Link"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Product Link</FormLabel>
                      <FormControl>
                        <Input {...field} type="url" />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />

                <FormField
                  control={form.control}
                  name="config.Product.Logo"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Product Logo URL</FormLabel>
                      <FormControl>
                        <Input {...field} type="url" />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />

                <FormField
                  control={form.control}
                  name="config.Product.Copyright"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Copyright Text</FormLabel>
                      <FormControl>
                        <Input {...field} />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />

                <FormField
                  control={form.control}
                  name="config.Product.TroubleText"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Trouble Text</FormLabel>
                      <FormControl>
                        <Input {...field} />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />

                <FormField
                  control={form.control}
                  name="config.DisableCSSInlining"
                  render={({ field }) => (
                    <FormItem className="flex items-center gap-2">
                      <FormControl>
                        <Switch checked={field.value} onCheckedChange={field.onChange} />
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
          <Card>
            <CardHeader>
              <CardTitle>Email Content</CardTitle>
            </CardHeader>
            <CardContent className="space-y-6">
              {/* Basic Info */}
              <div className="grid grid-cols-2 gap-4">
                <FormField
                  control={form.control}
                  name="email.Name"
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
                  name="email.Title"
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
                    onClick={() => addListItem('Intros', '')}
                  >
                    <Plus className="h-4 w-4 mr-2" /> Add Intro
                  </Button>
                </div>
                {form.watch('email.Intros').map((_: unknown, index: number) => (
                  <FormField
                    key={index}
                    control={form.control}
                    name={`email.Intros.${index}`}
                    render={({ field }) => (
                      <FormItem className="flex gap-2 mt-2">
                        <FormControl>
                          <Input {...field} />
                        </FormControl>
                        <Button
                          type="button"
                          variant="outline"
                          size="icon"
                          onClick={() => removeListItem('Intros', index)}
                        >
                          <Minus className="h-4 w-4" />
                        </Button>
                        <FormMessage />
                      </FormItem>
                    )}
                  />
                ))}
              </div>

              {/* Dictionary */}
              <div>
                <div className="flex justify-between items-center mb-2">
                  <FormLabel>Dictionary</FormLabel>
                  <Button
                    type="button"
                    variant="outline"
                    onClick={() => addListItem('Dictionary', { Key: '', Value: '' })}
                  >
                    <Plus className="h-4 w-4 mr-2" /> Add Dictionary Item
                  </Button>
                </div>
                {form.watch('email.Dictionary').map((_: unknown, index: number) => (
                  <div key={index} className="flex gap-2 mt-2">
                    <FormField
                      control={form.control}
                      name={`email.Dictionary.${index}.Key`}
                      render={({ field }) => (
                        <FormItem className="flex-1">
                          <FormControl>
                            <Input {...field} placeholder="Key" />
                          </FormControl>
                          <FormMessage />
                        </FormItem>
                      )}
                    />
                    <FormField
                      control={form.control}
                      name={`email.Dictionary.${index}.Value`}
                      render={({ field }) => (
                        <FormItem className="flex-1">
                          <FormControl>
                            <Input {...field} placeholder="Value" />
                          </FormControl>
                          <FormMessage />
                        </FormItem>
                      )}
                    />
                    <Button
                      type="button"
                      variant="outline"
                      size="icon"
                      onClick={() => removeListItem('Dictionary', index)}
                    >
                      <Minus className="h-4 w-4" />
                    </Button>
                  </div>
                ))}
              </div>

              {/* Actions */}
              <div className="space-y-4">
                <div className="flex justify-between items-center">
                  <FormLabel>Actions</FormLabel>
                  <Button
                    type="button"
                    variant="outline"
                    onClick={() => addListItem('Actions', {
                      Instructions: '',
                      Button: {
                        Color: '#000000',
                        TextColor: '#ffffff',
                        Text: '',
                        Link: ''
                      },
                      InviteCode: ''
                    })}
                  >
                    <Plus className="h-4 w-4 mr-2" /> Add Action
                  </Button>
                </div>
                {form.watch('email.Actions').map((_: unknown, index: number) => (
                  <Card key={index}>
                    <CardContent className="space-y-4 pt-4">
                      <div className="flex justify-between items-center">
                        <FormLabel>Action {index + 1}</FormLabel>
                        <Button
                          type="button"
                          variant="outline"
                          size="icon"
                          onClick={() => removeListItem('Actions', index)}
                        >
                          <Minus className="h-4 w-4" />
                        </Button>
                      </div>

                      <FormField
                        control={form.control}
                        name={`email.Actions.${index}.Instructions`}
                        render={({ field }) => (
                          <FormItem>
                            <FormLabel>Instructions</FormLabel>
                            <FormControl>
                              <Input {...field} placeholder="Action instructions" />
                            </FormControl>
                            <FormMessage />
                          </FormItem>
                        )}
                      />

                      <div className="grid grid-cols-2 gap-4">
                        <FormField
                          control={form.control}
                          name={`email.Actions.${index}.Button.Color`}
                          render={({ field }) => (
                            <FormItem>
                              <FormLabel>Button Color</FormLabel>
                              <FormControl>
                                <Input {...field} type="color" />
                              </FormControl>
                              <FormMessage />
                            </FormItem>
                          )}
                        />

                        <FormField
                          control={form.control}
                          name={`email.Actions.${index}.Button.TextColor`}
                          render={({ field }) => (
                            <FormItem>
                              <FormLabel>Text Color</FormLabel>
                              <FormControl>
                                <Input {...field} type="color" />
                              </FormControl>
                              <FormMessage />
                            </FormItem>
                          )}
                        />
                      </div>

                      <FormField
                        control={form.control}
                        name={`email.Actions.${index}.Button.Text`}
                        render={({ field }) => (
                          <FormItem>
                            <FormLabel>Button Text</FormLabel>
                            <FormControl>
                              <Input {...field} placeholder="Click me" />
                            </FormControl>
                            <FormMessage />
                          </FormItem>
                        )}
                      />

                      <FormField
                        control={form.control}
                        name={`email.Actions.${index}.Button.Link`}
                        render={({ field }) => (
                          <FormItem>
                            <FormLabel>Button Link</FormLabel>
                            <FormControl>
                              <Input {...field} type="url" placeholder="https://" />
                            </FormControl>
                            <FormMessage />
                          </FormItem>
                        )}
                      />

                      <FormField
                        control={form.control}
                        name={`email.Actions.${index}.InviteCode`}
                        render={({ field }) => (
                          <FormItem>
                            <FormLabel>Invite Code</FormLabel>
                            <FormControl>
                              <Input {...field} placeholder="Optional invite code" />
                            </FormControl>
                            <FormMessage />
                          </FormItem>
                        )}
                      />
                    </CardContent>
                  </Card>
                ))}
              </div>

              {/* Outros */}
              <div>
                <div className="flex justify-between items-center mb-2">
                  <FormLabel>Outros</FormLabel>
                  <Button
                    type="button"
                    variant="outline"
                    onClick={() => addListItem('Outros', '')}
                  >
                    <Plus className="h-4 w-4 mr-2" /> Add Outro
                  </Button>
                </div>
                {form.watch('email.Outros').map((_: unknown, index: number) => (
                  <FormField
                    key={index}
                    control={form.control}
                    name={`email.Outros.${index}`}
                    render={({ field }) => (
                      <FormItem className="flex gap-2 mt-2">
                        <FormControl>
                          <Input {...field} placeholder="Outro text" />
                        </FormControl>
                        <Button
                          type="button"
                          variant="outline"
                          size="icon"
                          onClick={() => removeListItem('Outros', index)}
                        >
                          <Minus className="h-4 w-4" />
                        </Button>
                        <FormMessage />
                      </FormItem>
                    )}
                  />
                ))}
              </div>

              {/* Additional Content */}
              <div className="space-y-4">
                <FormField
                  control={form.control}
                  name="email.Greeting"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Greeting</FormLabel>
                      <FormControl>
                        <Input {...field} placeholder="Hello" />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />

                <FormField
                  control={form.control}
                  name="email.Signature"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Signature</FormLabel>
                      <FormControl>
                        <Input {...field} placeholder="Best regards" />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />

                <FormField
                  control={form.control}
                  name="email.FreeMarkdown"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Free Markdown</FormLabel>
                      <FormControl>
                        <Textarea
                          {...field}
                          placeholder="Free markdown content that replaces all content other than header and footer"
                          className="min-h-[100px]"
                        />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />
              </div>

              {/* Submit Button */}
              <div className="flex justify-end gap-2">
                <Button type="submit" disabled={isLoadingGen}>
                  {isLoadingGen ? 'Generating...' : 'Generate Email'}
                </Button>
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
            <Button onClick={onSave} disabled={isLoading || !generated}>
              {isLoading ? 'Saving...' : 'Save Template'}
            </Button>
          </div>
        </CardContent>
      </Card>
    </div>
  );
};

export default EmailTemplateGenerator;