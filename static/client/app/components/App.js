//var React = require('react')

import React from 'react';
import RaisedButton from 'material-ui/RaisedButton';
import injectTapEventPlugin from 'react-tap-event-plugin';

import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import getMuiTheme from 'material-ui/styles/getMuiTheme';
import {blueGrey900} from 'material-ui/styles/colors';

import List from './List'
import ChooseUser from './ChooseUser'

// Needed for onTouchTap
// http://stackoverflow.com/a/34015469/988941
injectTapEventPlugin();

// This replaces the textColor value on the palette
// and then update the keys for each component that depends on it.
// More on Colors: http://www.material-ui.com/#/customization/colors
const muiTheme = getMuiTheme({
  palette: {
    textColor: blueGrey900,
  },
  appBar: {
    height: 50,
  },
});

class App extends React.Component{
  constructor(){
    super()
    this.state = {
      selectedUser: ""
    }
  }

  onSelect(user){
    console.log("App knows new selection: " + user);
    this.setState({
      selectedUser: user
    });
  }

  render(){
    return (
      <MuiThemeProvider muiTheme={muiTheme}>
        <div>
          <span><h1>Get your fix here.</h1></span>
          <ChooseUser onSelect={this.onSelect.bind(this)}/>
          <List selectedUser={this.state.selectedUser} />
        </div>
      </MuiThemeProvider>
    )
  }
}

module.exports = App;