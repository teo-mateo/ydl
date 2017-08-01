var React = require('react');
var ChooseUser = require('./ChooseUser');
var MultiActions = require('./MultiActions');
var List = require('./List');

class Main extends React.Component{
    constructor(){
        super()
        this.state = {
            selectedSongs: []
        }
    }

    onSelectSongs(songs){
        console.log("App knows new song selection: " + songs)
        this.setState({
            selectedSongs: songs
        });
    }

    componentWillReceiveProps(nextProps){
        if (this.props.selectedUser !== nextProps.selectedUser){
            this.setState({
                selectedSongs: []
            });
        }
    }

    render(){
        return (
        <div>
            <div className="flex-container">
                <div className="flex-item">
                    <ChooseUser selectedUser={this.props.selectedUser}/>
                </div>
                <div className="flex-item">    
                    <MultiActions  selectedSongs={this.state.selectedSongs} />
                </div> 
            </div> 
            <div>
                <List 
                    onSelect={this.onSelectSongs.bind(this)} 
                    selectedUser={this.props.selectedUser} selectedSongs={[]} />
            </div>          

        </div>
        );
    }
}

module.exports = Main;