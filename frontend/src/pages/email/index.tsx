import React from 'react';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { format } from "date-fns";
import { Mail, RefreshCcw } from "lucide-react";
import { Link, useNavigate } from 'react-router-dom';
import { useListEmailQuery, useResendEmailMutation } from '@/services';
import { Button } from '@/components/ui/button';

const EmailList = () => {
  const navigate = useNavigate();
  const [resendEmail, { isLoading: isSending }] = useResendEmailMutation();

  const { data: smtps, isLoading, isError, refetch } = useListEmailQuery(undefined, {
    refetchOnMountOrArgChange: true
  });

  const handleResend = async (id: number) => {
    try {
      await resendEmail(id).unwrap();
    } catch (err) {
      console.error("Failed to resend email:", err);
    }
  };

  return (
    <div className="p-8">
      <Card>
        <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-4">
          <CardTitle>Sent Emails</CardTitle>
          <div className="flex gap-2">
            <Button variant='outline' onClick={refetch} disabled={isLoading || isSending}>
              <RefreshCcw className="h-4 w-4" />
            </Button>
            <Button asChild>
              <Link to="/email/send">
                <Mail className="mr-2 h-4 w-4" />
                Send Email
              </Link>
            </Button>
          </div>
        </CardHeader>
        <CardContent>
          <div className="rounded-md border">
            <Table>
              <TableHeader>
                <TableRow>
                  <TableHead className="w-[100px]">ID</TableHead>
                  <TableHead>SMTP</TableHead>
                  <TableHead>From</TableHead>
                  <TableHead>To</TableHead>
                  <TableHead>Subject</TableHead>
                  <TableHead className="text-right">Sent At</TableHead>
                  <TableHead className="w-[100px]">Actions</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {smtps?.list?.map((item) => (
                  <TableRow key={item.id} className="cursor-pointer group hover:bg-muted/50 transition-colors"                  >
                    <TableCell className="font-medium">#{item.id}</TableCell>
                    <TableCell>{item.smtp_name}</TableCell>
                    <TableCell className="font-mono text-sm">{item.from}</TableCell>
                    <TableCell className="font-mono text-sm">{item.to}</TableCell>
                    <TableCell className="max-w-[300px] truncate">
                      {item.subject}
                    </TableCell>
                    <TableCell className="text-right text-muted-foreground">
                      {item.sent_at && format(new Date(item.sent_at), "MMM d, yyyy HH:mm:ss")}
                    </TableCell>
                    <TableCell>
                      <Button
                        variant="ghost"
                        size="icon"
                        onClick={() => handleResend(item.id)}
                        className="hover:text-destructive"
                        disabled={isSending || isLoading}
                      >
                        <Mail className="h-4 w-4" />
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

export default EmailList;