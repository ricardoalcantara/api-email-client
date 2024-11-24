import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import {
  Mail,
  FileText,
  Server,
  Key
} from "lucide-react";
import { useGetDashboardQuery } from "@/services";

const StatCard = ({
  title,
  value,
  icon: Icon,
  description
}: {
  title: string;
  value: number;
  icon: React.ElementType;
  description: string;
}) => (
  <Card>
    <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
      <CardTitle className="text-sm font-medium">
        {title}
      </CardTitle>
      <Icon className="h-4 w-4 text-muted-foreground" />
    </CardHeader>
    <CardContent>
      <div className="text-2xl font-bold">{value}</div>
      <p className="text-xs text-muted-foreground">
        {description}
      </p>
    </CardContent>
  </Card>
);

const Dashboard = () => {
  const { data: stats, isLoading } = useGetDashboardQuery();

  if (isLoading) {
    return (
      <div className="p-8">
        <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
          {[...Array(4)].map((_, i) => (
            <Card key={i} className="animate-pulse">
              <CardHeader className="space-y-0 pb-2">
                <div className="h-4 w-24 bg-muted rounded" />
              </CardHeader>
              <CardContent>
                <div className="h-7 w-16 bg-muted rounded" />
                <div className="h-3 w-32 bg-muted rounded mt-2" />
              </CardContent>
            </Card>
          ))}
        </div>
      </div>
    );
  }

  const metrics = [
    {
      title: "Total Templates",
      value: stats?.templates ?? 0,
      icon: FileText,
      description: "Email templates in the system"
    },
    {
      title: "Emails Sent",
      value: stats?.emails ?? 0,
      icon: Mail,
      description: "Total emails delivered"
    },
    {
      title: "SMTP Configs",
      value: stats?.smtps ?? 0,
      icon: Server,
      description: "Active SMTP configurations"
    },
    {
      title: "API Keys",
      value: stats?.api_keys ?? 0,
      icon: Key,
      description: "Generated API keys"
    }
  ];

  return (
    <div className="p-8">
      <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
        {metrics.map((metric) => (
          <StatCard
            key={metric.title}
            title={metric.title}
            value={metric.value}
            icon={metric.icon}
            description={metric.description}
          />
        ))}
      </div>
    </div>
  );
};

export default Dashboard;