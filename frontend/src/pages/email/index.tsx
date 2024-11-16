import React from 'react';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { format } from "date-fns";
import { ChevronRight } from "lucide-react";

const EmailList = () => {
  // Sample data
  const [data] = React.useState([
    {
      id: 1,
      smtpName: "Primary SMTP",
      from: "noreply@example.com",
      to: "user1@example.com",
      subject: "Welcome to Our Service!",
      sentAt: "2024-11-16T10:30:00Z"
    },
    {
      id: 2,
      smtpName: "Primary SMTP",
      from: "noreply@example.com",
      to: "user2@example.com",
      subject: "Your Monthly Report",
      sentAt: "2024-11-16T10:28:00Z"
    },
    {
      id: 3,
      smtpName: "Secondary SMTP",
      from: "support@example.com",
      to: "user3@example.com",
      subject: "Password Reset Request",
      sentAt: "2024-11-16T10:25:00Z"
    }
  ]);

  const handleRowClick = (item: any) => {
    // Navigate to detail view
    console.log("Navigate to email details:", item);
  };

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
                {data.map((item) => (
                  <TableRow
                    key={item.id}
                    onClick={() => handleRowClick(item)}
                    className="cursor-pointer group hover:bg-muted/50 transition-colors"
                  >
                    <TableCell className="font-medium">#{item.id}</TableCell>
                    <TableCell>{item.smtpName}</TableCell>
                    <TableCell className="font-mono text-sm">{item.from}</TableCell>
                    <TableCell className="font-mono text-sm">{item.to}</TableCell>
                    <TableCell className="max-w-[300px] truncate">
                      {item.subject}
                    </TableCell>
                    <TableCell className="text-right text-muted-foreground">
                      {format(new Date(item.sentAt), "MMM d, yyyy HH:mm:ss")}
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