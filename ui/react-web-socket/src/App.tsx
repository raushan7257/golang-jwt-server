import {useEffect, useRef} from "react";




function App() {
  const ws = useRef<WebSocket | null>(null);

  useEffect(() => {
    ws.current = new WebSocket("ws://localhost:8080/ws");

    ws.current.onopen = () => {
      console.log("WebSocket connected");
    };

    ws.current.onmessage = (event: MessageEvent) => {
      console.log("Message from server:", event.data);
      alert(event.data);
    };

    ws.current.onclose = () => {
      console.log("WebSocket disconnected");
    };

    return () => {
      ws.current?.close();
    };
  }, []);

  const handleStart = () => {
    if (ws.current?.readyState === WebSocket.OPEN) {
      ws.current.send("start");
      
    } else {
      console.error("WebSocket not connected");
    }
  };

  const handleStop = () => {
    if (ws.current?.readyState === WebSocket.OPEN) {
      ws.current.send("stop");
    } else {
      console.error("WebSocket not connected");
    }
  };

  return (
    <>
      <h1 className="text-center" style={{textAlign: "center"}}>WebSocket Testing</h1>
      <div style={{display: "flex", justifyContent: "center", gap: "5rem"}}>
        <button  onClick={handleStart}  style={{
    margin: "10px",
    padding: "10px 20px",
    backgroundColor: "#4CAF50",
    color: "white",
    border: "none",
    borderRadius: "8px",
    cursor: "pointer",
    fontSize: "16px",
    fontWeight: "bold",
    boxShadow: "0px 4px 6px rgba(0,0,0,0.1)",
    transition: "all 0.3s ease",
  }}
  onMouseOver={(e) => {
    (e.currentTarget as HTMLButtonElement).style.backgroundColor = "#45a049";
  }}
  onMouseOut={(e) => {
    (e.currentTarget as HTMLButtonElement).style.backgroundColor = "#4CAF50";
  }}
  
  
  >
          Start
        </button>
        <button onClick={handleStop}
 style={{
    margin: "10px",
    padding: "10px 20px",
    backgroundColor: "#ffffff",
    color: "red",
    border: "none",
    borderRadius: "8px",
    cursor: "pointer",
    fontSize: "16px",
    fontWeight: "bold",
    boxShadow: "0px 4px 6px rgba(0,0,0,0.1)",
    transition: "all 0.3s ease",
  }}        
        
        
        >
          Stop
        </button>

        
      </div>
    </>
  );
}

export default App;
