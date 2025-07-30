import { useState } from 'react';
import api from '../api/client';
import { TextField, Button, Box } from '@mui/material';

export default function UrlForm({ onSuccess }: { onSuccess: () => void }) {
  const [url, setUrl] = useState('');

  const handleSubmit = async () => {
    try {
      await api.post('/urls', { url });
      setUrl('');
      onSuccess(); // trigger reload in parent
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <Box display="flex" gap={2} mb={2}>
      <TextField
        fullWidth
        label="Enter website URL"
        value={url}
        onChange={(e) => setUrl(e.target.value)}
      />
      <Button variant="contained" onClick={handleSubmit}>
        Add
      </Button>
    </Box>
  );
}
