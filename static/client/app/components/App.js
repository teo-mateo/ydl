//var React = require('react')

import React from 'react';
import RaisedButton from 'material-ui/RaisedButton';
import injectTapEventPlugin from 'react-tap-event-plugin';

import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import getMuiTheme from 'material-ui/styles/getMuiTheme';

import AppBar from 'material-ui/AppBar';


import Header from './Header';
import MultiActions from './MultiActions';

import Main from './Main';

var ReactRouter = require('react-router-dom');
var Router = ReactRouter.HashRouter;
var Route = ReactRouter.Route;
var Switch = ReactRouter.Switch;

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
      selectedSongs: [],
      searchTerm: ""
    }

    this.onSearch = this.onSearch.bind(this);
    this.renderMain = this.renderMain.bind(this);
  }

  renderMain({match}){
    var searchTerm = this.state.searchTerm;
    return (
      <Main selectedUser={match.params.selectedUser || ""} searchTerm={searchTerm} />
    )
  }

  onSearch(term){
    this.setState({
      searchTerm: term
    });
  }

  render(){
    return (
      <MuiThemeProvider>
        <div>
          <AppBar title={<Header onSearch={this.onSearch} />}/>

          <Router>
            <Switch>
              <Route path='/' exact render={this.renderMain} />
              <Route path='/users/:selectedUser?' render={this.renderMain} />
            </Switch>
          </Router>
        </div>
      </MuiThemeProvider>
    );
  }
}

class Root extends React.Component{
  constructor(){
    super()
  }

  render(){
    return (
      <div>ROOT </div>
    )
  }
}

module.exports = App; 