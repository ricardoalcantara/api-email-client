import React from "react";
import { Button } from "@/components/ui/button";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Plus } from "lucide-react";
import { Link, useNavigate } from "react-router-dom";
import { useListTemplateQuery } from "@/services";
import { TemplateDto } from "@/services/dto";

const TemplateList = () => {
  const navigate = useNavigate();
  const { data: templates, isLoading, isError } = useListTemplateQuery(undefined, {
    refetchOnMountOrArgChange: true
  });

  const handleRowClick = (item: TemplateDto) => {
    navigate(`/template/${item.slug}`);
  };

  return (
    <div className="p-8">
      <Card>
        <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-4">
          <CardTitle>Templates</CardTitle>
          <Button asChild>
            <Link to="/template/new">
              <Plus className="mr-2 h-4 w-4" />
              Create New
            </Link>
          </Button>
        </CardHeader>
        <CardContent>
          <div className="rounded-md border">
            <Table>
              <TableHeader>
                <TableRow>
                  <TableHead className="w-24">ID</TableHead>
                  <TableHead>Name</TableHead>
                  <TableHead>Subject</TableHead>
                  <TableHead>Slug</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {templates?.list?.map((item) => (
                  <TableRow
                    key={item.id}
                    onClick={() => handleRowClick(item)}
                    className="cursor-pointer hover:bg-muted/50 transition-colors"
                  >
                    <TableCell className="font-medium">{item.id}</TableCell>
                    <TableCell>{item.name}</TableCell>
                    <TableCell>{item.subject}</TableCell>
                    <TableCell>{item.slug}</TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </div>

          <div className="flex items-center justify-end space-x-2 py-4">
            <Button variant="outline" size="sm" disabled={true}>
              Previous
            </Button>
            <div className="text-sm">Page 1 of 999</div>
            <Button variant="outline" size="sm" disabled={true}>
              Next
            </Button>
          </div>
        </CardContent>
      </Card>
    </div>
  );
};

export default TemplateList;
