import logo from './logo.svg';
import './App.css';
import Grid from './components/Grid.js'
import OptionsDrawer from './components/OptionsDrawer';

function App() {
  return (
    <div className="App">
      <OptionsDrawer />
      <Grid />
    </div>
  );
}

export default App;
