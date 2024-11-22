import { useEffect, useState } from 'react';
import { Button } from "@/components/ui/button";
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogTrigger } from "@/components/ui/dialog";
import { Plus } from "lucide-react";
import { usePostApiKeyMutation } from '@/services';
import { getError } from '@/lib/error';
import AlertError from '@/components/alert-error';
import FormApiKey, { formSchema } from './form-api-key';
import { ApiKeyDto } from '@/services/dto';
import ApiKeyResultView from './api-key-result-view';
import * as z from "zod"

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