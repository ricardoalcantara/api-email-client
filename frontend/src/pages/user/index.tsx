import { useState } from "react";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import * as z from "zod";
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
import { Button } from "@/components/ui/button";
import { Alert, AlertDescription } from "@/components/ui/alert";
import { AlertCircle, CheckCircle } from "lucide-react";
import { getError } from "@/lib/error";
import { usePatchPasswordMutation } from "@/services";

const formSchema = z.object({
  current_password: z.string().min(1, "Current password is required"),
  new_password: z
    .string()
    .min(6, "Password must be at least 8 characters"),
  confirm_password: z.string().min(1, "Please confirm your password"),
}).refine((data) => data.new_password === data.confirm_password, {
  message: "Passwords do not match",
  path: ["confirm_password"],
});

const UserHome = () => {
  const [updatePassword, { isLoading: isUpdating }] = usePatchPasswordMutation();
  const [errorMsg, setErrorMsg] = useState("");
  const [successMsg, setSuccessMsg] = useState("");

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      current_password: "",
      new_password: "",
      confirm_password: "",
    },
  });

  async function onSubmit(values: z.infer<typeof formSchema>) {
    try {
      setErrorMsg("");
      setSuccessMsg("");
      await updatePassword(values).unwrap();
      setSuccessMsg("Password successfully updated");
      form.reset();
    } catch (err) {
      setErrorMsg(getError(err));
    }
  }

  return (
    <div className="p-8">
      <Card>
        <CardHeader>
          <CardTitle>Change Password</CardTitle>
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

          {successMsg && (
            <Alert className="mb-6 border-green-500 bg-green-50 text-green-700">
              <AlertDescription className="flex items-center gap-2">
                <CheckCircle className="h-4 w-4" />
                {successMsg}
              </AlertDescription>
            </Alert>
          )}

          <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-6">
              <FormField
                control={form.control}
                name="current_password"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Current Password</FormLabel>
                    <FormControl>
                      <Input
                        type="password"
                        placeholder="Enter your current password"
                        {...field}
                        disabled={isUpdating}
                      />
                    </FormControl>
                    <FormDescription>
                      Enter your current password to verify your identity
                    </FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />

              <FormField
                control={form.control}
                name="new_password"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>New Password</FormLabel>
                    <FormControl>
                      <Input
                        type="password"
                        placeholder="Enter your new password"
                        {...field}
                        disabled={isUpdating}
                      />
                    </FormControl>
                    <FormDescription>
                      Password must be at least 8 characters and include uppercase, lowercase, and numbers
                    </FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />

              <FormField
                control={form.control}
                name="confirm_password"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Confirm New Password</FormLabel>
                    <FormControl>
                      <Input
                        type="password"
                        placeholder="Confirm your new password"
                        {...field}
                        disabled={isUpdating}
                      />
                    </FormControl>
                    <FormDescription>
                      Re-enter your new password to confirm
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
                  disabled={isUpdating}
                >
                  Reset
                </Button>
                <Button type="submit" disabled={isUpdating}>
                  Update Password
                </Button>
              </div>
            </form>
          </Form>
        </CardContent>
      </Card>
    </div>
  );
};

export default UserHome;