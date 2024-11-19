import React from 'react';
import { Button } from "@/components/ui/button";
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Plus } from "lucide-react";
import { Link, useNavigate } from 'react-router-dom';
import { useListSmtpQuery } from '@/services';
import { SmtpDto } from '@/services/dto';

const SmtpList = () => {
  const navigate = useNavigate();
  const { data: smtps, isLoading, isError } = useListSmtpQuery(undefined, {
    refetchOnMountOrArgChange: true
  });

  const handleRowClick = (item: SmtpDto) => {
    navigate(`/smtp/${item.slug}`);
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
                  <TableHead className="text-center">Status</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {smtps?.list?.map((item) => (
                  <TableRow
                    key={item.id}
                    onClick={() => handleRowClick(item)}
                    className="cursor-pointer hover:bg-muted/50 transition-colors"
                  >
                    <TableCell className="font-medium">{item.name}</TableCell>
                    <TableCell>{item.server}</TableCell>
                    <TableCell>{item.port}</TableCell>
                    <TableCell>{item.email}</TableCell>
                    <TableCell>{item.user}</TableCell>
                    <TableCell className="text-center">
                      {item.default && (
                        <Badge variant="default">Default</Badge>
                      )}
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