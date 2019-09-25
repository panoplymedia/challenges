import React from 'react';
import './App.css';
import DataPage from './components/DataPage'
import Header from './components/Header'
import Sales from './components/Sales'


function App() {
  return (
    <div className="App">
        <Header/>
        <DataPage/>
        <Sales/>
    </div>
  );
}

export default App;
