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
import { Trash2, Plus, AlertCircle, CheckCircle, Copy } from "lucide-react";
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

export interface ApiKeyResultViewProps {
  apiKeyDto: ApiKeyDto;
  onClose: () => void;
}

const ApiKeyResultView = ({ apiKeyDto, onClose }: ApiKeyResultViewProps) => {
  const [copied, setCopied] = useState(false);

  const handleCopy = async () => {
    await navigator.clipboard.writeText(apiKeyDto.key ?? "");
    setCopied(true);
    setTimeout(() => setCopied(false), 2000);
  };

  return (
    <div className="space-y-6">
      <Alert className="flex items-center gap-2">
        <div>
          <AlertCircle className="h-4 w-4" />
        </div>
        <AlertDescription>
          Make sure to copy your API key now. You won't be able to see it again!
        </AlertDescription>
      </Alert>

      <div className="space-y-4">
        <div>
          <Label className="text-muted-foreground">Name</Label>
          <p className="mt-1 font-medium">{apiKeyDto.name}</p>
        </div>

        <div>
          <Label className="text-muted-foreground">API Key</Label>
          <div className="mt-1 flex items-center gap-2">
            <code className="flex-1 rounded bg-muted px-2 py-1 font-mono">
              {apiKeyDto.key}
            </code>
            <Button
              variant="outline"
              size="icon"
              onClick={handleCopy}
              className="flex-shrink-0"
            >
              {copied ? (
                <CheckCircle className="h-4 w-4 text-green-500" />
              ) : (
                <Copy className="h-4 w-4" />
              )}
            </Button>
          </div>
        </div>

        {apiKeyDto.expires_at && (
          <div>
            <Label className="text-muted-foreground">Expires At</Label>
            <p className="mt-1 font-medium">{format(new Date(apiKeyDto.expires_at), "MMMM d, yyyy")}</p>
          </div>
        )}

        {apiKeyDto.ip_whitelist && (
          <div>
            <Label className="text-muted-foreground">IP Whitelist</Label>
            <p className="mt-1 font-medium">{apiKeyDto.ip_whitelist}</p>
          </div>
        )}
      </div>

      <div className="flex justify-end">
        <Button onClick={onClose}>
          Close
        </Button>
      </div>
    </div>
  );
};

export default ApiKeyResultView;