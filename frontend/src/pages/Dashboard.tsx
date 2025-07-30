import { useEffect, useState } from 'react';
import api from '../api/client';
import UrlForm from '../components/UrlForm';
import {
  Table, TableHead, TableRow, TableCell, TableBody, Paper,
  Container, Typography, CircularProgress, TableContainer
} from '@mui/material';
import { useNavigate } from 'react-router-dom';

export default function Dashboard() {
  const [urls, setUrls] = useState<any[]>([]);
  const [loading, setLoading] = useState(true);
  const navigate = useNavigate();

  const fetchUrls = async () => {
    try {
      const res = await api.get('/urls');
      setUrls(res.data);
    } catch (err) {
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchUrls();
    const interval = setInterval(fetchUrls, 5000); // polling every 5s
    return () => clearInterval(interval);
  }, []);

  return (
    <Container>
      <Typography variant="h4" gutterBottom>
        URL Dashboard
      </Typography>

      <UrlForm onSuccess={fetchUrls} />

      {loading ? (
        <CircularProgress />
      ) : (
        <TableContainer component={Paper}>
          <Table>
            <TableHead>
              <TableRow>
                <TableCell>URL</TableCell>
                <TableCell>Status</TableCell>
                <TableCell>Title</TableCell>
                <TableCell>HTML Version</TableCell>
                <TableCell>Internal</TableCell>
                <TableCell>External</TableCell>
                <TableCell>Broken</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {urls.map((row) => (
                <TableRow
                  key={row.id}
                  hover
                  onClick={() => navigate(`/urls/${row.id}`)}
                  style={{ cursor: 'pointer' }}
                >
                  <TableCell>{row.url}</TableCell>
                  <TableCell>{row.status}</TableCell>
                  <TableCell>{row.title}</TableCell>
                  <TableCell>{row.html_version}</TableCell>
                  <TableCell>{row.internal_links}</TableCell>
                  <TableCell>{row.external_links}</TableCell>
                  <TableCell>{row.broken_links_count}</TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      )}
    </Container>
  );
}
