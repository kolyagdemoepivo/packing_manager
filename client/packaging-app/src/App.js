import React, { useState } from "react";
import './App.css';

function App() {

  const [totalItems, setTotalItems] = useState("");
  const [packageSizes, setPackageSizes] = useState([""]);
  const [calRes, setCalRes] = useState({})
  const [errorMsg, setErrorMsg] = useState("")
  const apiUrl = process.env.REACT_APP_API_URL;

  const handleTotalItemsChange = (e) => {
    const value = e.target.value;
    if (/^\d*$/.test(value)) {
      setTotalItems(value);
    }
  };


  const handlePackageSizeChange = (index, value) => {
    if (/^\d*$/.test(value)) {
      const newPackageSizes = [...packageSizes];
      newPackageSizes[index] = value;
      setPackageSizes(newPackageSizes);
    }
  };


  const addPackageSize = () => {
    setPackageSizes([...packageSizes, ""]);
  };


  const removePackageSize = (index) => {
    const newPackageSizes = packageSizes.filter((_, i) => i !== index);
    setPackageSizes(newPackageSizes);
  };


  const handleCalculate = () => {
    fetch(`${apiUrl}/calculate?totalItems=${totalItems}&packages=${packageSizes.join(',')}`).then(res => {
      return res.json()
    }).then(json => {
      if (json.error) {
        setCalRes({})
        setErrorMsg(json.error)
      } else {
        setErrorMsg("")
        setCalRes({...json})
      }
      
    }).catch(e => {
      console.log(e)
    })
  };

  return (
    <div className="App">
      <div style={{ width: "1024px", margin: "0 auto", padding: "20px", border: "1px solid #ccc" }}>

        <div style={{ textAlign: "right", marginBottom: "20px" }}>
          <button onClick={addPackageSize} style={{ padding: "10px 20px", cursor: "pointer" }}>
            Add Package
          </button>
        </div>


        <div style={{ marginBottom: "20px", width: '100%' }}>
          <label>Total Items: </label>
          <input
            type="text"
            value={totalItems}
            onChange={handleTotalItemsChange}
            style={{ padding: "5px" }}
          />
        </div>

        <div style={{alignItems: 'center', width: '100%', flexDirection: 'row', justifyContent: 'center'}}>
          {packageSizes.map((size, index) => (
            <div key={index} style={{ marginBottom: "10px", width: '100%', display: 'flex', alignItems: 'center', justifyContent: 'center' }}>
              <div style={{width: '200px'}}>
                <label>Package Size: </label>
              </div>
              <div style={{width: '100px'}}>
                <input
                  type="text"
                  value={size}
                  onChange={(e) => handlePackageSizeChange(index, e.target.value)}
                  style={{ padding: "5px", width: "100px" }}
                />
              </div>
              {packageSizes.length === 1 && (<div style={{width: '94px'}}/>)}
              {packageSizes.length > 1 && (
                <button
                  onClick={() => removePackageSize(index)}
                  style={{ marginLeft: "30px", padding: "5px 10px", cursor: "pointer" }}
                >
                  Delete
                </button>
              )}
              
            </div>
          ))}
        </div>
        <div style={{ marginBottom: "20px", width: '100%' }}>
          {Object.keys(calRes).map((val, index) => {
            return <p key={index}>A package of size {val} will have a quantity of {calRes[val]}</p>
          })}
        </div>
        <div style={{ marginBottom: "20px", width: '100%' }}>
          <p>{errorMsg}</p>
        </div>
        <div style={{ textAlign: "right", marginTop: "20px" }}>
          <button
            onClick={handleCalculate}
            style={{ padding: "10px 20px", cursor: "pointer" }}
          >
            Calculate
          </button>
        </div>
      </div>
    </div>
  );
}

export default App;
