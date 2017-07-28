//var React = require('react')

import React from 'react';
import RaisedButton from 'material-ui/RaisedButton';
import injectTapEventPlugin from 'react-tap-event-plugin';

import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import getMuiTheme from 'material-ui/styles/getMuiTheme';

import AppBar from 'material-ui/AppBar';


import Header from './Header';
import MultiActions from './MultiActions';

import List from './List';
import ChooseUser from './ChooseUser';
import TestTable from './tests/TestTable';

// Needed for onTouchTap
// http://stackoverflow.com/a/34015469/988941
injectTapEventPlugin();

// This replaces the textColor value on the palette
// and then update the keys for each component that depends on it.
// More on Colors: http://www.material-ui.com/#/customization/colors


class App extends React.Component{
  constructor(){
    super()
    this.state = {
      selectedUser: "", 
      selectedSongs: []
    }
  }

  onSelectUser(user){
    console.log("App knows new user selection: " + user);
    this.setState({
      selectedUser: user
    });
  }

  onSelectSongs(songs){
    console.log("App knows new song selection: " + songs)
    this.setState(() => {
      return {
        selectedSongs: songs, 
        selectedUser: this.state.selectedUser
      }
    });
  }

  // shouldComponentUpdate(nextProps, nextState){
  //   if (this.state.selectedUser !== nextState.selectedUser){
  //     return true;
  //   } else {
  //     return false;
  //   }
  // }

  render(){
    var a= 2;
    return (
      <MuiThemeProvider>
        <div>
          <AppBar title={<Header />}/>
  
          <div className="flex-container">
            <div className="flex-item">
              <ChooseUser onSelect={this.onSelectUser.bind(this)}/>
            </div>
            <div className="flex-item">    
              <MultiActions  selectedSongs={this.state.selectedSongs}/>
            </div>              
          </div>

          {/* <TestTable /> */}

           <List 
            selectedUser={this.state.selectedUser} 
            onSelect={this.onSelectSongs.bind(this)}
            selectedSongs={this.state.selectedSongs}
            /> 
        </div>
      </MuiThemeProvider>
    );
  }
}

module.exports = App;