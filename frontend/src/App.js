import { useState } from "react";
import { Box, Button, ButtonGroup, Container, Typography } from "@mui/material";
import Products from "./pages/Products";
import RawMaterials from "./pages/RawMaterial";
import Inputs from "./pages/Inputs";
import Reports from "./pages/Reports";

function App() {
  const [currentPage, setCurrentPage] = useState(null);

  const handleChangePage = (page) => {
    setCurrentPage(page);
  };

  return (
    <div style={{display: 'flex', flexDirection: 'column', width: '100%', height: '100%'}}>
      <div style={{display: 'flex', margin: 20, marginLeft: 40, marginRight: 40}}>
        <ButtonGroup
          style={{display:'flex', width:'100%'}}
          variant="contained"
          aria-label="outlined primary button group"
        >
          <Button style={{ display:'flex', flex:1 }} onClick={() => handleChangePage("products")}>Products</Button>
          <Button style={{ display:'flex', flex:1 }} onClick={() => handleChangePage("raw_material")}>Raw Material</Button>
          <Button style={{ display:'flex', flex:1 }} onClick={() => handleChangePage("inputs")}>Inputs</Button>
          <Button style={{ display:'flex', flex:1 }} onClick={() => handleChangePage("report_all")}>Total Report</Button>
          <Button style={{ display:'flex', flex:1 }} onClick={() => handleChangePage("report_valuable")}>Most valuable report</Button>
        </ButtonGroup>
        </div>

      <div style={{display: 'flex', justifyContent: 'center', alignItems: 'center', width: '100%', height: '100%'}}>
      {(() => {
        switch((currentPage)) {
          case 'products':
            return <Products/>
            break;
          case 'raw_material':
            return <RawMaterials/>
            break;
          case 'inputs':
            return <Inputs/>
            break;
          case 'report_all':
            return <Reports reportType={"all"} />
            break;
          case 'report_valuable':
            return <Reports reportType={"valuable"} />
            break;
        }
        })()}
      </div>
    </div>
  );
}

export default App;
