import React from 'react';

function App() {
  const handleClick = () => {
    fetch('http://localhost:3000/api')
    .then(response => response.json())
    .then(data => console.log(data))
    .catch(error => console.error('Error:', error));
  };

  return (
    <div className="App">
      <button onClick={handleClick}>Call Golang Code</button>
    </div>
  );
}

export default App;