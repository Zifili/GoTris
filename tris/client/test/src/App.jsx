import { useState, useEffect } from "react";
import WebSocket from "react-websocket";

function Square({ value, onSquareClick }) {
  return (
    <button className="square" onClick={onSquareClick}>
      {" "}
      {value}{" "}
    </button>
  );
}

function Reset({ resetOnClick }) {
  return (
    <button className="reset" onClick={resetOnClick}>
      Quit
    </button>
  );
}

export default function Board() {
  const [xIsNext, setxIsNext] = useState(true);
  const [squares, setSquares] = useState(Array(9).fill(null));
  const api = "http://localhost:8080";
  //const winner = calculateWinner(squares);
  const [winner, setWinner] = useState(null);
  useEffect(() => {
    fetch(api+'/game-state')
      .then(response => response.json())
      .then(data => {
        setxIsNext(data.xIsNext);
        setSquares(data.squares);
        setWinner(data.winner);
      });
  }, []);
  function handleData(data) {
    const gameState = JSON.parse(data);
    setxIsNext(gameState.xIsNext);
    setSquares(gameState.squares);
    setWinner(gameState.winner);
  }
  let status;
  if (winner) {
    status = "Winner: " + winner;
  } else {
    status = "Next player: " + (xIsNext ? "X" : "O");
  }
  function handleClick(i) {
    if (squares[i] || winner) {
      return;
    }
    
    const currentSquares = squares.slice();
    if (xIsNext) {
      currentSquares[i] = "X";
    } else {
      currentSquares[i] = "O";
    }
    fetch(api+'/make-move', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ squares: currentSquares }),
    })
      .then(response => response.json())
      .then(data => {
        setxIsNext(!xIsNext);
        setSquares(data.squares);
        setWinner(data.winner);
      });
    }
  function reset() {
    // Effettua una richiesta al backend per resettare il gioco
    fetch(api+'/reset-game', {
      method: 'POST',
    })
      .then(response => response.json())
      .then(data => {
        setxIsNext(data.xIsNext);
        setSquares(data.squares);
        setWinner(data.winner);
      });
  }

  return (
    <>
      <WebSocket
        url="ws://127.0.0.1:8080/ws"
        onMessage={handleData}
        reconnect={true}
      />
      <div className="status">{status}</div>
      <div className="board-row">
        <Square value={squares[0]} onSquareClick={() => handleClick(0)} />
        <Square value={squares[1]} onSquareClick={() => handleClick(1)} />
        <Square value={squares[2]} onSquareClick={() => handleClick(2)} />
      </div>
      <div className="board-row">
        <Square value={squares[3]} onSquareClick={() => handleClick(3)} />
        <Square value={squares[4]} onSquareClick={() => handleClick(4)} />
        <Square value={squares[5]} onSquareClick={() => handleClick(5)} />
      </div>
      <div className="board-row">
        <Square value={squares[6]} onSquareClick={() => handleClick(6)} />
        <Square value={squares[7]} onSquareClick={() => handleClick(7)} />
        <Square value={squares[8]} onSquareClick={() => handleClick(8)} />
      </div>
      <br></br>
      <Reset resetOnClick={reset} />
    </>
  );
}

