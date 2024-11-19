import React from 'react';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { format } from "date-fns";
import { ChevronRight } from "lucide-react";
import { useNavigate } from 'react-router-dom';
import { useListEmailQuery } from '@/services';

const EmailList = () => {
  const navigate = useNavigate();
  const { data: smtps, isLoading, isError } = useListEmailQuery(undefined, {
    refetchOnMountOrArgChange: true
  });

  return (
    <div className="p-8">
      <Card>
        <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-4">
          <CardTitle>Sent Emails</CardTitle>
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
                  <TableHead className="w-[50px]"></TableHead>
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
                      <ChevronRight
                        className="w-4 h-4 opacity-0 group-hover:opacity-100 transition-opacity text-muted-foreground"
                      />
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