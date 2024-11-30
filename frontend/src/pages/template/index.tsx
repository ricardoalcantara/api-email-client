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
import { Copy, Plus, Trash2 } from "lucide-react";
import { Link, useNavigate } from "react-router-dom";
import { useCloneTemplateMutation, useDeleteTemplateMutation, useListTemplateQuery } from "@/services";
import { TemplateDto } from "@/services/dto";

const TemplateList = () => {
  const navigate = useNavigate();
  const [deleteTemplate, { isLoading: isDeleting }] = useDeleteTemplateMutation();
  const [cloneTemplate, { isLoading: isCloning }] = useCloneTemplateMutation();
  const { data: templates, isLoading, isError, refetch } = useListTemplateQuery(undefined, {
    refetchOnMountOrArgChange: true
  });

  const handleRowClick = (item: TemplateDto) => {
    navigate(`/template/${item.slug}`);
  };

  const handleDelete = async (slug: string) => {
    try {
      await deleteTemplate(slug).unwrap();
    } catch (err) {
      console.error("Failed to delete API key:", err);
    }

    try {
      await refetch().unwrap();
    } catch (err) {
      console.error("Failed to refetch API keys:", err);
    }
  };

  const handleClone = async (slug: string) => {
    let result: TemplateDto | undefined;
    try {
      result = await cloneTemplate(slug).unwrap();
    } catch (err) {
      console.error("Failed to clone template:", err);
    }

    try {
      await refetch().unwrap();
    } catch (err) {
      console.error("Failed to refetch API keys:", err);
    }

    if (result) {
      navigate(`/template/${result.slug}`);
    }
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
                  <TableHead className="w-[100px]">Actions</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {templates?.list?.map((item) => (
                  <TableRow
                    key={item.id}
                    className="cursor-pointer hover:bg-muted/50 transition-colors"
                  >
                    <TableCell onClick={() => handleRowClick(item)} className="font-medium">{item.id}</TableCell>
                    <TableCell onClick={() => handleRowClick(item)}>{item.name}</TableCell>
                    <TableCell onClick={() => handleRowClick(item)}>{item.subject}</TableCell>
                    <TableCell onClick={() => handleRowClick(item)}>{item.slug}</TableCell>
                    <TableCell className="flex gap-2">
                      <Button
                        variant="ghost"
                        size="icon"
                        onClick={(e) => {
                          e.stopPropagation();
                          handleClone(item.slug);
                        }}
                        className="hover:text-primary"
                        disabled={isCloning || isLoading}
                      >
                        <Copy className="h-4 w-4" />
                      </Button>
                      <Button
                        variant="ghost"
                        size="icon"
                        onClick={(e) => {
                          e.stopPropagation();
                          handleDelete(item.slug);
                        }}
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
