import {
  Button,
  IconButton,
  Modal,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  TextField,
  Typography,
} from "@mui/material";
import DeleteIcon from "@mui/icons-material/Delete";

import { Box } from "@mui/system";
import axios from "axios";
import { useEffect, useState } from "react";

const style = {
  position: "absolute",
  top: "50%",
  left: "50%",
  transform: "translate(-50%, -50%)",
  width: 800,
  backgroundColor: "#FFF",
  border: "2px solid #000",
  boxShadow: 24,
  padding: "20px",
};

const Products = () => {
  const [rows, setRows] = useState([]);
  const [rowSelected, setRowSelected] = useState(null);

  const [name, setName] = useState("");
  const [value, setValue] = useState("");

  useEffect(() => {
    fetchProducts();
  }, []);

  const fetchProducts = () => {
    return new Promise((resolve, reject) => {
      axios.get("http://localhost:6060/product").then(({ data }) => {
        setRows(data);
        resolve();
      });
    });
  };

  const saveProductHandler = () => {
    if (rowSelected.id === undefined) {
      axios
        .post(`http://localhost:6060/product`, {
          name,
          value: Number(value),
        })
        .then(() => {
          fetchProducts();
        })
        .finally(() => {
          setRowSelected(null);
        });
    } else {
      axios
        .put(`http://localhost:6060/product/${rowSelected.id}`, {
          name,
          value: Number(value),
        })
        .then(() => {
          fetchProducts();
        })
        .finally(() => {
          setRowSelected(null);
        });
    }
  };

  const deleteRowHandler = (row) => {
    axios.delete(`http://localhost:6060/product/${row.id}`).then(() => {
      fetchProducts();
    });
  };

  return (
    <Box>
      <Typography variant="h4" textAlign="center">
        Products
      </Typography>
      <Box
        sx={{ display: "flex" }}
        flexDirection="row"
        justifyContent="space-between"
      >
        <Typography variant="h5" textAlign="center">
          Products List
        </Typography>
        <Button onClick={() => setRowSelected({})}>Add Product</Button>
      </Box>
      <TableContainer>
        <Table sx={{ minWidth: 650 }} size="small" aria-label="a dense table">
          <TableHead>
            <TableRow>
              <TableCell>ID</TableCell>
              <TableCell>Name</TableCell>
              <TableCell>Value</TableCell>
              <TableCell>Delete</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {rows.map((row) => (
              <TableRow
                key={row.id}
                sx={{ "&:last-child td, &:last-child th": { border: 0 } }}
                onClick={() => {
                  setName(row.name);
                  setValue(row.value);
                  setRowSelected(row);
                }}
              >
                <TableCell>{row.id}</TableCell>
                <TableCell>{row.name}</TableCell>
                <TableCell>R$ {row.value}</TableCell>
                <TableCell>
                  <IconButton>
                    <DeleteIcon
                      onClick={(e) => {
                        e.stopPropagation();
                        deleteRowHandler(row);
                      }}
                    />
                  </IconButton>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
      <Modal
        open={rowSelected !== null}
        onClose={() => setRowSelected(null)}
        aria-labelledby="modal-modal-title"
        aria-describedby="modal-modal-description"
      >
        <Box style={style}>
          <Typography id="modal-modal-title" variant="h6" component="h2">
            Create/Edit Product
          </Typography>
          <Box>
            <Box width="100%" paddingY={2}>
              <TextField
                label="Nome"
                variant="outlined"
                fullWidth
                value={name}
                onChange={(e) => setName(e.target.value)}
              />
            </Box>
            <Box width="100%" paddingY={2}>
              <TextField
                label="Valor"
                variant="outlined"
                fullWidth
                value={value}
                onChange={(e) => setValue(e.target.value)}
              />
            </Box>
            <Button
              variant="contained"
              fullWidth
              onClick={() => saveProductHandler()}
            >
              Save
            </Button>
          </Box>
        </Box>
      </Modal>
    </Box>
  );
};

export default Products;
