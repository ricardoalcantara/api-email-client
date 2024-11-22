import React, { useEffect, useState } from 'react';
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
import FormApiKey from './form-api-key';
import { ApiKeyDto } from '@/services/dto';
import ApiKeyResultView from './api-key-result-view';

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

export interface DialogAddApiKeysProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
}

const DialogAddApiKeys = ({ open, onOpenChange }: DialogAddApiKeysProps) => {
  const [errorMsg, setErrorMsg] = useState("");
  const [createApiKey, { isLoading }] = usePostApiKeyMutation();
  const [apiKeyDto, setApiKeyDto] = useState<ApiKeyDto | null>(null);

  useEffect(() => {
    if (!open) {
      setErrorMsg("");
      setApiKeyDto(null);
    }
  }, [open]);

  async function onSubmit(values: z.infer<typeof formSchema>) {
    try {
      setErrorMsg("");
      const { expirationType, ...rest } = values;
      const formData = {
        ...rest,
        expires_at: expirationType === 'NO_EXPIRE' ? null : values.expires_at || null,
      };
      const result = await createApiKey(formData).unwrap();
      setApiKeyDto(result);
    } catch (err) {
      setErrorMsg(getError(err));
    }
  }

  const handleClose = () => {
    onOpenChange(false);
  };


  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogTrigger asChild>
        <Button>
          <Plus className="mr-2 h-4 w-4" />
          Create API Key
        </Button>
      </DialogTrigger>
      <DialogContent className='min-w-fit'>
        <DialogHeader>
          <DialogTitle>Create New API Key</DialogTitle>
        </DialogHeader>
        <AlertError error={errorMsg} />
        {apiKeyDto ? (
          <ApiKeyResultView
            apiKeyDto={apiKeyDto}
            onClose={handleClose}
          />
        ) : (
          <FormApiKey onSubmit={onSubmit} disabled={isLoading} />
        )}
      </DialogContent>
    </Dialog>
  )
}

export default DialogAddApiKeys;