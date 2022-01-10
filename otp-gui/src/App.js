import logo from './logo.svg';
import './App.css';
import Grid from './components/Grid.js'
import OptionsBar from './components/OptionsBar';

function App() {
  return (
    <div className="App">
      <OptionsBar />
      <Grid />
    </div>
  );
}

export default App;
