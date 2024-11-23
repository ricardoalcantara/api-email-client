import React from 'react';
import { Button } from "@/components/ui/button";
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Plus, Trash2 } from "lucide-react";
import { Link, useNavigate } from 'react-router-dom';
import { useDeleteSmtpMutation, useListSmtpQuery } from '@/services';
import { SmtpDto } from '@/services/dto';

const SmtpList = () => {
  const navigate = useNavigate();
  const [deleteSmtp, { isLoading: isDeleting }] = useDeleteSmtpMutation();
  const { data: smtps, isLoading, isError, refetch } = useListSmtpQuery(undefined, {
    refetchOnMountOrArgChange: true
  });

  const handleRowClick = (item: SmtpDto) => {
    navigate(`/smtp/${item.slug}`);
  };

  const handleDelete = async (slug: string) => {
    try {
      await deleteSmtp(slug).unwrap();
    } catch (err) {
      console.error("Failed to delete API key:", err);
    }

    try {
      await refetch().unwrap();
    } catch (err) {
      console.error("Failed to refetch API keys:", err);
    }
  };

  return (
    <div className="p-8">
      <Card>
        <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-4">
          <CardTitle>SMTP Configurations</CardTitle>
          <Button asChild>
            <Link to="/smtp/new">
              <Plus className="mr-2 h-4 w-4" />
              Add SMTP
            </Link>
          </Button>
        </CardHeader>
        <CardContent>
          <div className="rounded-md border">
            <Table>
              <TableHeader>
                <TableRow>
                  <TableHead>Name</TableHead>
                  <TableHead>Server</TableHead>
                  <TableHead>Port</TableHead>
                  <TableHead>Email</TableHead>
                  <TableHead>User</TableHead>
                  <TableHead className="text-center">Is Default</TableHead>
                  <TableHead className="w-[100px]">Actions</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {smtps?.list?.map((item) => (
                  <TableRow
                    key={item.id}
                    className="cursor-pointer hover:bg-muted/50 transition-colors"
                  >
                    <TableCell onClick={() => handleRowClick(item)} className="font-medium">{item.name}</TableCell>
                    <TableCell onClick={() => handleRowClick(item)}>{item.server}</TableCell>
                    <TableCell onClick={() => handleRowClick(item)}>{item.port}</TableCell>
                    <TableCell onClick={() => handleRowClick(item)}>{item.email}</TableCell>
                    <TableCell onClick={() => handleRowClick(item)}>{item.user}</TableCell>
                    <TableCell className="text-center">
                      {item.default && (
                        <Badge variant="default">Default</Badge>
                      )}
                    </TableCell>
                    <TableCell>
                      <Button
                        variant="ghost"
                        size="icon"
                        onClick={() => handleDelete(item.slug)}
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

export default SmtpList;