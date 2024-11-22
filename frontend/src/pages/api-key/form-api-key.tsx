import React, { useState } from 'react';
import { useForm } from "react-hook-form"
import { zodResolver } from "@hookform/resolvers/zod"
import * as z from "zod"
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogTrigger } from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Trash2, Plus, AlertCircle } from "lucide-react";
import { format, addDays } from "date-fns";
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form"
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select"
import { usePostApiKeyMutation } from '@/services';
import { getError } from '@/lib/error';
import { Alert, AlertDescription } from '@/components/ui/alert';
import AlertError from '@/components/alert-error';

const EXPIRATION_OPTIONS = {
  '30_DAYS': '30 Days',
  '60_DAYS': '60 Days',
  '90_DAYS': '90 Days',
  'CUSTOM': 'Custom',
  'NO_EXPIRE': 'No Expiration'
} as const;

const formSchema = z.object({
  name: z.string().min(1, "Name is required"),
  expirationType: z.enum(['30_DAYS', '60_DAYS', '90_DAYS', 'CUSTOM', 'NO_EXPIRE']),
  expires_at: z.string().optional(),
  ip_whitelist: z.string().default(""),
})

export interface FormApiKeyProps {
  onSubmit: (values: z.infer<typeof formSchema>) => void;
  disabled?: boolean;
}

const FormApiKey = ({ onSubmit, disabled }: FormApiKeyProps) => {

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      name: "",
      expirationType: '30_DAYS',
      expires_at: format(addDays(new Date(), 30), "yyyy-MM-dd"),
      ip_whitelist: "",
    },
  })

  const expirationType = form.watch('expirationType');

  // Handle expiration type change
  const handleExpirationTypeChange = (value: string) => {
    form.setValue('expirationType', value as keyof typeof EXPIRATION_OPTIONS);

    if (value === 'NO_EXPIRE') {
      form.setValue('expires_at', '');
    } else if (value !== 'CUSTOM') {
      const days = parseInt(value.split('_')[0]);
      form.setValue('expires_at', format(addDays(new Date(), days), "yyyy-MM-dd"));
    }
  };

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="grid gap-4 py-4">
        <FormField
          control={form.control}
          name="name"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Name</FormLabel>
              <FormControl>
                <Input placeholder="API Key Name" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="expirationType"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Expiration Type</FormLabel>
              <Select
                onValueChange={handleExpirationTypeChange}
                defaultValue={field.value}
              >
                <FormControl>
                  <SelectTrigger>
                    <SelectValue placeholder="Select expiration type" />
                  </SelectTrigger>
                </FormControl>
                <SelectContent>
                  {Object.entries(EXPIRATION_OPTIONS).map(([value, label]) => (
                    <SelectItem key={value} value={value}>
                      {label}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>
              <FormMessage />
            </FormItem>
          )}
        />

        {expirationType === 'CUSTOM' && (
          <FormField
            control={form.control}
            name="expires_at"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Custom Expiration Date</FormLabel>
                <FormControl>
                  <Input
                    type="date"
                    {...field}
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
        )}

        <FormField
          control={form.control}
          name="ip_whitelist"
          render={({ field }) => (
            <FormItem>
              <FormLabel>IP Whitelist</FormLabel>
              <FormControl>
                <Input
                  placeholder="Enter IP addresses (comma-separated)"
                  {...field}
                />
              </FormControl>
              <FormDescription>
                Enter multiple IP addresses separated by commas (e.g., 192.168.1.1,192.168.1.2)
              </FormDescription>
              <FormMessage />
            </FormItem>
          )}
        />

        <Button type="submit" className="mt-2" disabled={disabled}>
          Create Key
        </Button>
      </form>
    </Form>
  )
}

export default FormApiKey;