import {
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Typography,
} from "@mui/material";

import { Box } from "@mui/system";
import axios from "axios";
import { useEffect, useState } from "react";

const Reports = (props) => {
  const [rows, setRows] = useState([]);

  const { reportType } = props;

  useEffect(() => {
    fetchReport();
  }, [reportType]);

  const fetchReport = () => {
    return new Promise((resolve, reject) => {
      axios.get(`http://localhost:6060/report/${reportType}`).then(({ data }) => {
        setRows(data);
        resolve();
      });
    });
  };

  return (
    <Box>
      { reportType === "all" 
      ?
      <Typography variant="h4" textAlign="center">
        Total report
      </Typography>
      :
      <Typography variant="h4" textAlign="center">
        More valuable report
      </Typography>
      }
      <Box
        sx={{ display: "flex" }}
        flexDirection="row"
        justifyContent="space-between"
      >
      </Box>
      <TableContainer>
        <Table sx={{ minWidth: 650 }} size="small" aria-label="a dense table">
          <TableHead>
            <TableRow>
              <TableCell>ID</TableCell>
              <TableCell>Name</TableCell>
              <TableCell>Possible amount to produce</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {rows.map((row) => (
              <TableRow
                key={row.id}
                sx={{ "&:last-child td, &:last-child th": { border: 0 } }}
              >
                <TableCell>{row.ID}</TableCell>
                <TableCell>{row.Name}</TableCell>
                <TableCell>{row.Quantity}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </Box>
  );
};

export default Reports;
