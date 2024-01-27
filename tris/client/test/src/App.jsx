import {useState} from 'react';
import { connect, sendMsg } from "../api/api.js";

function Square({value, onSquareClick}){
  return <button className="square" onClick={onSquareClick}> {value} </button>;
}

function Reset({resetOnClick}){
  return <button className='reset'onClick={resetOnClick}>Quit</button>;
}

export default function Board() {
  const [xIsNext,setXIsNext] = useState(true);
  const [squares, setSquares] = useState(Array(9).fill(null));
  const winner = calculateWinner(squares);
  let status;
  if(winner){
    status = "Winner: " + winner;
  } else {
    status = "Next player: " + (xIsNext ? "X" : "O");
  }
  function handleClick(i){
    if (squares[i] || winner){
      return;
    }
    const currentSquares = squares.slice();
    if(xIsNext){
      currentSquares[i] = "X";
    } else {
      currentSquares[i] = "O";
    }
    setSquares(currentSquares);
    setXIsNext(!xIsNext);
  }
  function reset(){
    const reset = Array(9).fill(null);
    setSquares(reset);
    setXIsNext(true);
  }

  return (
  <>
    <div className="status">{status}</div>
    <div className="board-row">
      <Square value={squares[0]} onSquareClick={() => handleClick(0)}/>
      <Square value={squares[1]} onSquareClick={() => handleClick(1)}/>
      <Square value={squares[2]} onSquareClick={() => handleClick(2)}/>
    </div>
    <div className="board-row">
      <Square value={squares[3]} onSquareClick={() => handleClick(3)}/>
      <Square value={squares[4]} onSquareClick={() => handleClick(4)}/>
      <Square value={squares[5]} onSquareClick={() => handleClick(5)}/>
    </div>
    <div className="board-row">
      <Square value={squares[6]} onSquareClick={() => handleClick(6)}/>
      <Square value={squares[7]} onSquareClick={() => handleClick(7)}/>
      <Square value={squares[8]} onSquareClick={() => handleClick(8)}/>
    </div>
    <br></br>
    <Reset resetOnClick={reset}/>
  </>
  );
}

function calculateWinner(squares) {
  const lines = [
    [0, 1, 2],
    [3, 4, 5],
    [6, 7, 8],
    [0, 3, 6],
    [1, 4, 7],
    [2, 5, 8],
    [0, 4, 8],
    [2, 4, 6]
  ];
  for (let i = 0; i < lines.length; i++) {
    const [a, b, c] = lines[i];
    if (squares[a] && squares[a] === squares[b] && squares[a] === squares[c]) {
      return squares[a];
    }
  }
  return null;
}