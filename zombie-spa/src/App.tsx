import React, { FunctionComponent } from 'react';
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Redirect,
} from 'react-router-dom';
import IllnessesScreen from './components/IllnessesScreen';
import HospitalsScreen from './components/HospitalsScreen';
import SeverityLevelScreen from './components/SeverityLevelScreen';

const App: FunctionComponent = () => {
  return (
    <Router>
      <Switch>
        <Route exact path="/">
          <Redirect to="/illnesses" />
        </Route>
        <Route exact path="/illnesses" component={IllnessesScreen} />
        <Route exact path="/illnesses/:illnessId/severity-level/">
          <SeverityLevelScreen />
        </Route>
        <Route exact path="/hospitals" component={HospitalsScreen} />
      </Switch>
    </Router>
  );
};

export default App;
