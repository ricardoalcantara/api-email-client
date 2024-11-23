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
import { Switch } from "@/components/ui/switch";
import { SmtpDto } from "@/services/dto";

export const formSchema = z.object({
  name: z.string().min(2, {
    message: "Name must be at least 2 characters.",
  }),
  slug: z.string().min(2, {  // Add this
    message: "Slug must be at least 2 characters.",
  }),
  server: z.string().min(2, {
    message: "Server address is required.",
  }),
  port: z.coerce.number().min(1).max(65535, {
    message: "Port must be between 1 and 65535.",
  }),
  email: z.string().email({
    message: "Please enter a valid email address.",
  }),
  user: z.string().min(2, {
    message: "Username is required.",
  }),
  password: z.string().min(2, {
    message: "Password is required.",
  }),
  default: z.boolean().default(false),
});

export const formSchemaEdit = z.object({
  name: z.string().min(2, {
    message: "Name must be at least 2 characters.",
  }),
  slug: z.string().min(2, {  // Add this
    message: "Slug must be at least 2 characters.",
  }),
  server: z.string().min(2, {
    message: "Server address is required.",
  }),
  port: z.coerce.number().min(1).max(65535, {
    message: "Port must be between 1 and 65535.",
  }),
  email: z.string().email({
    message: "Please enter a valid email address.",
  }),
  user: z.string().min(2, {
    message: "Username is required.",
  }),
  password: z.string().default(""),
  default: z.boolean().default(false),
});

export interface SmtpFormProps {
  onSubmit: (values: z.infer<typeof formSchema>) => void;
  isLoading: boolean;
  defaultValues?: SmtpDto;
}

const SmtpForm = ({ onSubmit, isLoading, defaultValues }: SmtpFormProps) => {
  const form = useForm({
    resolver: zodResolver(defaultValues ? formSchemaEdit : formSchema),
    defaultValues: defaultValues ?? {
      name: "",
      slug: "",
      server: "",
      port: 587,
      email: "",
      user: "",
      password: "",
      default: false,
    },
  });

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-6">
        <div className="grid gap-6 md:grid-cols-2">
          <FormField
            control={form.control}
            name="name"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Configuration Name</FormLabel>
                <FormControl>
                  <Input placeholder="Primary SMTP" {...field} />
                </FormControl>
                <FormDescription>
                  A friendly name for this SMTP configuration.
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
                  <Input placeholder="primary-smtp" {...field} />
                </FormControl>
                <FormDescription>
                  A unique identifier for this SMTP configuration.
                </FormDescription>
                <FormMessage />
              </FormItem>
            )}
          />
        </div>

        <div className="grid gap-6 md:grid-cols-2">
          <FormField
            control={form.control}
            name="server"
            render={({ field }) => (
              <FormItem>
                <FormLabel>SMTP Server</FormLabel>
                <FormControl>
                  <Input placeholder="smtp.example.com" {...field} />
                </FormControl>
                <FormDescription>
                  The SMTP server address.
                </FormDescription>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="port"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Port</FormLabel>
                <FormControl>
                  <Input
                    type="number"
                    placeholder="587"
                    min={1}
                    max={65535}
                    {...field}
                  />
                </FormControl>
                <FormDescription>
                  Common ports: 25, 465, 587, 2525
                </FormDescription>
                <FormMessage />
              </FormItem>
            )}
          />
        </div>

        <div className="grid gap-6 md:grid-cols-2">
          <FormField
            control={form.control}
            name="user"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Username</FormLabel>
                <FormControl>
                  <Input placeholder="SMTP username" {...field} />
                </FormControl>
                <FormDescription>
                  The SMTP authentication username.
                </FormDescription>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="password"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Password</FormLabel>
                <FormControl>
                  <Input
                    type="password"
                    placeholder="Enter SMTP password"
                    {...field}
                  />
                </FormControl>
                <FormDescription>
                  The SMTP authentication password.
                </FormDescription>
                <FormMessage />
              </FormItem>
            )}
          />
        </div>

        <div className="grid gap-6 md:grid-cols-2">
          <FormField
            control={form.control}
            name="email"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Email Address</FormLabel>
                <FormControl>
                  <Input
                    type="email"
                    placeholder="noreply@example.com"
                    {...field}
                  />
                </FormControl>
                <FormDescription>
                  The email address used for sending.
                </FormDescription>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="default"
            render={({ field }) => (
              <FormItem className="flex flex-row items-center justify-between rounded-lg border p-4">
                <div className="space-y-0.5">
                  <FormLabel className="text-base">
                    Set as Default
                  </FormLabel>
                  <FormDescription>
                    Make this the default SMTP configuration.
                  </FormDescription>
                </div>
                <FormControl>
                  <Switch
                    checked={field.value}
                    onCheckedChange={field.onChange}
                  />
                </FormControl>
              </FormItem>
            )}
          />
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
          <Button type="submit" disabled={isLoading}>
            Save Configuration
          </Button>
        </div>
      </form>
    </Form>
  );
};

export default SmtpForm;