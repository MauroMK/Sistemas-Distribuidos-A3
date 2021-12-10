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
  Select,
  FormControl,
  MenuItem
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

const Inputs = () => {
  const [products, setProducts] = useState([]);
  const [productSelectedVal, setProductSelectedVal] = useState([]);
  const [productSelected, setProductSelected] = useState([]);
  const [productSelectedKey, setProductSelectedKey] = useState([]);

  const [showInputs, setShowInputs] = useState(false);

  useEffect(() => {
    fetchProducts();
  }, []);

  const fetchProducts = () => {
    return new Promise((resolve, reject) => {
      axios.get("http://localhost:6060/product").then(({ data }) => {
        setProducts(data);
        resolve();
      });
    });
  };

  const handleChangeProductSelect = (event) => {
    let event_value = event.target.value.split("_");
    setProductSelectedVal(event.target.value);
    setProductSelected(event_value[1]);
    setProductSelectedKey(event_value[0]);
    setShowInputs(true);
  };

  return (
    
    <Box>
      <Typography variant="h4" textAlign="center">
        Select Product
      </Typography>
      <div style={{display: 'flex', flexDirection: 'column', justifyContent: 'center', width: 500}}>
      <FormControl style={{display: 'flex', width: '100%' }}>
        <Select
          labelId="select-products-label"
          id="select-products"
          value={productSelectedVal}
          onChange={handleChangeProductSelect}
          style={{display: 'flex', width: '100%', justifyContent: 'center', alignItems: 'center'}}
        >
        <MenuItem disabled value="">
          <em>Select Product</em>
        </MenuItem>
        {products.map((product) => (
          <MenuItem
            key={product.id}
            value={product.id + "_" + product.name}
          >
            {product.name}
          </MenuItem>
        ))}
        </Select>
      </FormControl>
      </div>
      {(() => {
        if (showInputs) {
          return <InputsHandler productName={productSelected} productId={productSelectedKey} />;
        }
      })()}
    </Box>
  );
};

const InputsHandler = (props) => {
  const [rows, setRows] = useState([]);
  const [rowSelected, setRowSelected] = useState(null);

  const [material, setMaterial] = useState("");
  const [quantity, setQuantity] = useState(0);

  const [materials, setMaterials] = useState([]);

  const [disableMaterial, setDisableMaterial] = useState(false)
  const { productId } = props;

  useEffect(() => {
    fetchInputs();
    fetchMaterials();
  }, [productId]);

  const handleChangeMaterialSelect = (event) => {
    setMaterial(event.target.value);
  };

  const fetchInputs = () => {
    return new Promise((resolve, reject) => {
      axios.get(`http://localhost:6060/raw_material/by_product/${productId}`).then(({ data }) => {
        setRows(data);
        resolve();
      });
    });
  };

  const fetchMaterials = () => {
    return new Promise((resolve, reject) => {
      axios.get(`http://localhost:6060/raw_material`).then(({ data }) => {
        setMaterials(data);
        resolve();
      });
    });
  };

  const saveInputHandler = () => {
    if (rowSelected.id === undefined) {
      axios
        .post(`http://localhost:6060/input`, {
          product: Number(productId),
          material: Number(material),
          quantity: Number(quantity),
        })
        .then(() => {
          fetchInputs();
        })
        .finally(() => {
          setRowSelected(null);
        });
    } else {
      axios
        .put(`http://localhost:6060/input`, {
          product: Number(productId),
          material: Number(material),
          quantity: Number(quantity),
        })
        .then(() => {
          fetchInputs();
        })
        .finally(() => {
          setRowSelected(null);
        });
    }
  };

  const deleteRowHandler = (row) => {
    console.log(Number(productId));
    console.log(Number(row.id));
    axios.delete(`http://localhost:6060/input`, { 
          data:
            {
              product: Number(productId),
              material: Number(row.id)
            }
        }).then(() => {
      fetchInputs();
    });
  };

  return (
      <div>
      <Typography variant="h4" textAlign="center">
        Insumos do produto {props.productName}
      </Typography>
      <Box
        sx={{ display: "flex" }}
        flexDirection="row"
        justifyContent="space-between"
      >
        <Typography variant="h5" textAlign="center">
          Lista de Insumos
        </Typography>
        <Button onClick={() => setRowSelected({})}>Adicionar Insumo</Button>
      </Box>
      <TableContainer>
        <Table sx={{ minWidth: 650 }} size="small" aria-label="a dense table">
          <TableHead>
            <TableRow>
              <TableCell>ID</TableCell>
              <TableCell>Name</TableCell>
              <TableCell>Necessary amount</TableCell>
              <TableCell>Delete</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {rows.map((row) => (
              <TableRow
                key={row.id}
                sx={{ "&:last-child td, &:last-child th": { border: 0 } }}
                onClick={() => {
                  setMaterial(row.id);
                  setQuantity(row.quantity);
                  setRowSelected(row);
                  setDisableMaterial(true);
                }}
              >
                <TableCell>{row.id}</TableCell>
                <TableCell>{row.name}</TableCell>
                <TableCell>{row.quantity}</TableCell>
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
        onClose={() => {
          setRowSelected(null);
          setDisableMaterial(false);
        }}
        aria-labelledby="modal-modal-title"
        aria-describedby="modal-modal-description"
      >
        <Box style={style}>
          <Typography id="modal-modal-title" variant="h6" component="h2">
            Cadastrar/Editar Insumo
          </Typography>
          <Box>
            <Box width="100%" paddingY={2}>
              <Select
                disabled={disableMaterial}
                labelId="select-products-label"
                id="select-products"
                value={material}
                onChange={handleChangeMaterialSelect}
              >
              {materials.map((material) => (
                <MenuItem
                  key={material.id}
                  value={material.id}
                >
                  {material.name}
                </MenuItem>
              ))}
              </Select>
            </Box>
            <Box width="100%" paddingY={2}>
              <TextField
                label="Quantity"
                variant="outlined"
                fullWidth
                value={quantity}
                onChange={(e) => setQuantity(e.target.value)}
              />
            </Box>
            <Button
              variant="contained"
              fullWidth
              onClick={() => saveInputHandler()}
            >
              Salvar
            </Button>
          </Box>
        </Box>
      </Modal>
      </div>
  );
}
export default Inputs;
