import { useState, useEffect } from "react";

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
  const [Turn, setTurn] = useState(true);
  const [squares, setSquares] = useState(Array(9).fill(null));
  const api = "http://localhost:8080";
  //const winner = calculateWinner(squares);
  const [winner, setWinner] = useState(null);
  useEffect(() => {
    fetch(api+'/game-state')
      .then(response => response.json())
      .then(data => {
        setTurn(data.Turn);
        setSquares(data.squares);
        setWinner(data.winner);
      });
  }, []);

  let status;
  if (winner) {
    status = "Winner: " + winner;
  } else {
    status = "Next player: " + (Turn ? "X" : "O");
  }
  function handleClick(i) {
    if (squares[i] || winner) {
      return;
    }
    if(!Turn){
      return;
    }
    const currentSquares = squares.slice();
    if (Turn) {
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
        setTurn(data.Turn);
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
        setTurn(data.Turn);
        setSquares(data.squares);
        setWinner(data.winner);
      });
  }

  return (
    <>
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

