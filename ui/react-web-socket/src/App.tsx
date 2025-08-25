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
      <h1 style={{textAlign: "center"}}>WebSocket Testing</h1>
      <div style={{display: "flex", justifyContent: "center", gap: "5rem"}}>
        <button onClick={handleStart} style={{margin: "10px"}}>
          Start
        </button>
        <button onClick={handleStop} style={{margin: "10px"}}>
          Stop
        </button>
      </div>
    </>
  );
}

export default App;
