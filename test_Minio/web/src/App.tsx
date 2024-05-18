import "./App.css";
import axios from "axios";
import React from "react";

function App() {
  const [url, setUrl] = React.useState<string>("");

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (file) {
      const formData = new FormData();
      formData.append("file", file);
      axios.post("/api/upload", formData).then((res) => {
        setUrl(res.data);
      });
    }
  };

  return (
    <>
      {url && (
        <img
          className="w-40 rounded"
          key={`http://199.241.138.79:9000/testminio/${url}`}
          src={`http://199.241.138.79:9000/testminio/${url}`}
        />
      )}
      <input type="file" onChange={handleFileChange} />
      {/* <Button variant="outline">Button</Button> */}
    </>
  );
}

export default App;
