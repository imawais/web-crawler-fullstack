import { useParams } from 'react-router-dom';
import { useEffect, useState } from 'react';
import api from '../api/client';
import {
  Container, Typography, CircularProgress, Table, TableRow, TableCell, TableHead, TableBody, Paper, Grid,
} from '@mui/material';
import {
  PieChart, Pie, Cell, Tooltip, Legend,
} from 'recharts';

export default function Details() {
  const { id } = useParams();
  const [urlData, setUrlData] = useState<any>(null);
  const [brokenLinks, setBrokenLinks] = useState<any[]>([]);
  const [loading, setLoading] = useState(true);

  const fetchDetails = async () => {
    try {
      const res = await api.get(`/urls`);
      const selected = res.data.find((u: any) => u.id === Number(id));
      setUrlData(selected);

      const brokenRes = await api.get(`/broken-links/${id}`);
      setBrokenLinks(brokenRes.data);
    } catch (e) {
      console.error(e);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchDetails();
  }, [id]);

  const COLORS = ['#0088FE', '#00C49F'];

  const pieData = [
    { name: 'Internal Links', value: urlData?.internal_links || 0 },
    { name: 'External Links', value: urlData?.external_links || 0 },
  ];

  if (loading) return <CircularProgress />;

  return (
    <Container>
      <Typography variant="h4" gutterBottom>
        Details: {urlData?.url}
      </Typography>

      <Typography variant="subtitle1">
        Status: {urlData.status} | Title: {urlData.title}
      </Typography>

      <Grid container spacing={4} mt={2}>
        <Grid item xs={12} md={6}>
          <Typography variant="h6">Link Distribution</Typography>
          <PieChart width={300} height={300}>
            <Pie
              data={pieData}
              dataKey="value"
              nameKey="name"
              outerRadius={100}
              label
            >
              {pieData.map((_, index) => (
                <Cell key={index} fill={COLORS[index % COLORS.length]} />
              ))}
            </Pie>
            <Tooltip />
            <Legend />
          </PieChart>
        </Grid>

        <Grid item xs={12} md={6}>
          <Typography variant="h6">Broken Links</Typography>
          <Paper>
            <Table>
              <TableHead>
                <TableRow>
                  <TableCell>Link</TableCell>
                  <TableCell>Status Code</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {brokenLinks.map((b, idx) => (
                  <TableRow key={idx}>
                    <TableCell>{b.link}</TableCell>
                    <TableCell>{b.status_code}</TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </Paper>
        </Grid>
      </Grid>
    </Container>
  );
}
