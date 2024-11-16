import React from 'react';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogTrigger } from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Trash2, Plus, X } from "lucide-react";
import { format, addMonths } from "date-fns";

const ApiKeyList = () => {
  const [isOpen, setIsOpen] = React.useState(false);
  const [newKey, setNewKey] = React.useState({
    name: '',
    expiresAt: format(addMonths(new Date(), 1), "yyyy-MM-dd"),
    ipWhitelist: ['']
  });

  // Sample data
  const [data] = React.useState([
    {
      id: 1,
      name: "Production API Key",
      lastUsed: "2024-11-16T10:30:00Z",
      ipWhitelist: ["192.168.1.1", "10.0.0.1", "172.16.0.1"],
      expiresAt: "2025-12-31T23:59:59Z"
    },
    {
      id: 2,
      name: "Development API Key",
      lastUsed: "2024-11-15T15:45:00Z",
      ipWhitelist: ["192.168.1.2"],
      expiresAt: "2024-12-31T23:59:59Z"
    },
    {
      id: 3,
      name: "Testing API Key",
      lastUsed: "2024-11-14T09:20:00Z",
      ipWhitelist: [],
      expiresAt: "2024-11-30T23:59:59Z"
    }
  ]);

  const handleDelete = (id) => {
    console.log("Delete API key:", id);
  };

  const handleCreateKey = () => {
    console.log("Create new API key:", newKey);
    // Reset form and close modal
    setNewKey({
      name: '',
      expiresAt: format(addMonths(new Date(), 1), "yyyy-MM-dd"),
      ipWhitelist: ['']
    });
    setIsOpen(false);
  };

  const addIpInput = () => {
    setNewKey(prev => ({
      ...prev,
      ipWhitelist: [...prev.ipWhitelist, '']
    }));
  };

  const removeIpInput = (index) => {
    setNewKey(prev => ({
      ...prev,
      ipWhitelist: prev.ipWhitelist.filter((_, i) => i !== index)
    }));
  };

  const updateIpInput = (index, value) => {
    setNewKey(prev => ({
      ...prev,
      ipWhitelist: prev.ipWhitelist.map((ip, i) => i === index ? value : ip)
    }));
  };

  // Helper function to check if a date is in the past
  const isExpired = (date) => {
    return new Date(date) < new Date();
  };

  // Format IP whitelist for display
  const formatIpWhitelist = (ips) => {
    if (ips.length === 0) return "No IP restrictions";
    if (ips.length === 1) return ips[0];
    return `${ips[0]} +${ips.length - 1} more`;
  };

  return (
    <div className="p-8">
      <Card>
        <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-4">
          <CardTitle>API Keys</CardTitle>
          <Dialog open={isOpen} onOpenChange={setIsOpen}>
            <DialogTrigger asChild>
              <Button>
                <Plus className="mr-2 h-4 w-4" />
                Create API Key
              </Button>
            </DialogTrigger>
            <DialogContent className="sm:max-w-[425px]">
              <DialogHeader>
                <DialogTitle>Create New API Key</DialogTitle>
              </DialogHeader>
              <div className="grid gap-4 py-4">
                <div className="grid gap-2">
                  <Label htmlFor="name">Name</Label>
                  <Input
                    id="name"
                    placeholder="API Key Name"
                    value={newKey.name}
                    onChange={(e) => setNewKey(prev => ({ ...prev, name: e.target.value }))}
                  />
                </div>
                <div className="grid gap-2">
                  <Label htmlFor="expires">Expiration Date</Label>
                  <Input
                    id="expires"
                    type="date"
                    value={newKey.expiresAt}
                    onChange={(e) => setNewKey(prev => ({ ...prev, expiresAt: e.target.value }))}
                  />
                </div>
                <div className="grid gap-2">
                  <Label>IP Whitelist</Label>
                  {newKey.ipWhitelist.map((ip, index) => (
                    <div key={index} className="flex gap-2">
                      <Input
                        placeholder="Enter IP address"
                        value={ip}
                        onChange={(e) => updateIpInput(index, e.target.value)}
                      />
                      <Button
                        variant="ghost"
                        size="icon"
                        type="button"
                        onClick={() => removeIpInput(index)}
                        disabled={newKey.ipWhitelist.length === 1}
                      >
                        <X className="h-4 w-4" />
                      </Button>
                    </div>
                  ))}
                  <Button
                    type="button"
                    variant="outline"
                    onClick={addIpInput}
                    className="mt-2"
                  >
                    Add IP Address
                  </Button>
                </div>
                <Button onClick={handleCreateKey} className="mt-2">
                  Create Key
                </Button>
              </div>
            </DialogContent>
          </Dialog>
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
                {data.map((item) => (
                  <TableRow key={item.id}>
                    <TableCell className="font-medium">#{item.id}</TableCell>
                    <TableCell>{item.name}</TableCell>
                    <TableCell className="text-muted-foreground">
                      {format(new Date(item.lastUsed), "MMM d, yyyy HH:mm")}
                    </TableCell>
                    <TableCell>{formatIpWhitelist(item.ipWhitelist)}</TableCell>
                    <TableCell>
                      <Badge variant={isExpired(item.expiresAt) ? "destructive" : "default"}>
                        {format(new Date(item.expiresAt), "MMM d, yyyy")}
                      </Badge>
                    </TableCell>
                    <TableCell>
                      <Button
                        variant="ghost"
                        size="icon"
                        onClick={() => handleDelete(item.id)}
                        className="hover:text-destructive"
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