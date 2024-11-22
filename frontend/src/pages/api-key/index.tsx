import React, { useEffect } from 'react';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Trash2 } from "lucide-react";
import { format } from "date-fns";
import { useDeleteApiKeyMutation, useListApiKeyQuery } from '@/services';
import DialogAddApiKeys from './dialog-add-api-keys';

const ApiKeyList = () => {
  const [isOpen, setIsOpen] = React.useState(false);
  const [deleteApiKey, { isLoading: isDeleting }] = useDeleteApiKeyMutation();

  const { data: apiKeys, isLoading, isError, refetch } = useListApiKeyQuery(undefined, {
    refetchOnMountOrArgChange: true
  });

  useEffect(() => {
    if (!isOpen) {
      refetch().catch(error => {
        console.error("Failed to refetch API keys:", error);
      });
    }
  }, [isOpen, refetch]);

  const handleDelete = async (id: number) => {
    try {
      await deleteApiKey(id).unwrap();
    } catch (err) {
      console.error("Failed to delete API key:", err);
    }

    try {
      await refetch().unwrap();
    } catch (err) {
      console.error("Failed to refetch API keys:", err);
    }
  };

  // Helper function to check if a date is in the past
  const isExpired = (date: string | null | undefined): boolean => {
    if (!date) return false;
    return new Date(date) < new Date();
  };

  // Format IP whitelist for display
  const formatIpWhitelist = (ips: string | null | undefined): string => {
    if (!ips) return "No IP restrictions";
    const parts = (ips).split(",");
    if (parts.length === 1) return parts[0];
    return `${parts[0]} +${parts.length - 1} more`;
  };

  return (
    <div className="p-8">
      <Card>
        <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-4">
          <CardTitle>API Keys</CardTitle>
          <DialogAddApiKeys open={isOpen} onOpenChange={setIsOpen} />
        </CardHeader>
        <CardContent>
          <div className="rounded-md border">
            <Table>
              <TableHeader>
                <TableRow>
                  <TableHead className="w-[100px]">ID</TableHead>
                  <TableHead>Name</TableHead>
                  <TableHead>Last Used</TableHead>
                  <TableHead>IP Whitelist</TableHead>
                  <TableHead>Expires At</TableHead>
                  <TableHead className="w-[100px]">Actions</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {apiKeys?.list?.map((item) => (
                  <TableRow key={item.id}>
                    <TableCell className="font-medium">#{item.id}</TableCell>
                    <TableCell>{item.name}</TableCell>
                    <TableCell className="text-muted-foreground">
                      {!!item.last_used ? format(new Date(item.last_used), "MMM d, yyyy HH:mm") : "-"}
                    </TableCell>
                    <TableCell>{formatIpWhitelist(item.ip_whitelist)}</TableCell>
                    <TableCell>
                      <Badge variant={isExpired(item.expires_at) ? "destructive" : "default"}>
                        {!!item.expires_at ? format(new Date(item.expires_at), "MMM d, yyyy") : "-"}
                      </Badge>
                    </TableCell>
                    <TableCell>
                      <Button
                        variant="ghost"
                        size="icon"
                        onClick={() => handleDelete(item.id)}
                        className="hover:text-destructive"
                        disabled={isDeleting || isLoading}
                      >
                        <Trash2 className="h-4 w-4" />
                      </Button>
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </div>
        </CardContent>
      </Card>
    </div>
  );
};

export default ApiKeyList;