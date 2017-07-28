var React = require('react');

import {
    Table,
    TableBody,
    TableHeader,
    TableHeaderColumn,
    TableRow,
    TableRowColumn
} from 'material-ui/Table'
import RaisedButton from 'material-ui/RaisedButton';
import IconButton from 'material-ui/IconButton';
import FileCloudDownload from 'material-ui/svg-icons/file/cloud-download';
import LinearProgress from 'material-ui/LinearProgress';




import PropTypes from 'prop-types'

var Api = require('../utils/Api')


class List extends React.Component{
    constructor(){
        super()
        this.state = {
            files: [], 
            loading:true
        }

        console.log(this.state);

        //this.load = this.load.bind(this);
    }

    componentDidMount(){
        Api.GetList(this.props.selectedUser)
            .then(songs => {
                songs.forEach(function(s){ s.isSelected=false;});
                this.setState({ files: songs, loading: false })
        });
    }

    componentDidUpdate(){
        console.log('List.componentDidUpdate()');
        console.log(this.state);
    }

    componentWillReceiveProps(nextProps){

        if (this.props.selectedUser === nextProps.selectedUser)
            return;

        //parent will show loader
        this.setState(() => {
            return { files: [], loading: true}
        });

        console.log("List.componentWillReceiveProps()")
        Api.GetList(nextProps.selectedUser)
            .then(songs => {
                songs.forEach(function(s){ s.isSelected=false;});
                this.setState({ files: songs, loading: false })
        });        
    }

    onRowSelection(rows){

        // this.setState(() =>{
        //     return {
        //         files: this.state.files
        //     }
        // })

        if (rows === "all"){
            //change state: select all
            this.setState(() => {
                // var newState = {
                //     files: this.state.files.map(f => {
                //         f.isSelected = true;
                //         return f;
                //     })
                // };
                // return newState;
            });
            this.props.onSelect(this.state.files.map(f => f.ID ));
        } else if (rows === "none" || rows.length == 0){
            //change state: select none
            this.setState(() =>{
                var newState = {
                    files: this.state.files.map(f => {
                        f.isSelected = false;
                        return f;
                    })
                };
                return newState;
            });
            this.props.onSelect([]);
        } else {
            //rows has the indexes of the selected rows
            var rows = rows.map((x) =>{ return this.state.files[x].ID;});

            this.setState(() =>{
                var newState = {
                    files: this.state.files
                        .map(f => {
                            f.isSelected = (rows.indexOf(f.ID) >= 0)
                            return f;
                        })

                };
                return newState;
            });
            this.props.onSelect(rows);
        }
    }

    render(){

        let myPaddingStyle = {
            paddingTop: 0,
            paddingBottom: 0,
        }

        var loaderStyle = {
            display: (this.state.loading ? 'block' : 'none')
        }

        var dlurl = 'http://' + window.location.host + '/download?id='
        if (this.state.files != null && this.state.files.length > 0){
            return (
                <div>
                    <LinearProgress mode="indeterminate" style={loaderStyle}/>
                    <Table selectable={true} multiSelectable={true}
                        style={{tableLayout: 'auto', padding: '0 !important'}} 
                        onRowSelection={this.onRowSelection.bind(this)}>
                        <TableHeader>
                            <TableRow style={myPaddingStyle}>
                                <TableHeaderColumn></TableHeaderColumn>
                                <TableHeaderColumn></TableHeaderColumn>
                                <TableHeaderColumn></TableHeaderColumn>
                            </TableRow>
                        </TableHeader>
                        <TableBody 
                            stripedRows={false}
                            showRowHover={true}
                            displayRowCheckbox={true}>
                            {
                                this.state.files.map(function(song, index){
                                    return (
                                        <TableRow 
                                            selectable={true} 
                                            style={myPaddingStyle} 
                                            key={song.ID}
                                            selected={song.isSelected}>
                                            <TableRowColumn>{song.ID}</TableRowColumn>
                                            <TableRowColumn>{song.Who.String}</TableRowColumn>
                                            <TableRowColumn>{song.File.String}</TableRowColumn>
                                            <TableRowColumn>
                                                <a href={dlurl+song.ID}>
                                                    <IconButton>
                                                        <FileCloudDownload />
                                                    </IconButton>
                                                </a>
                                            </TableRowColumn>
                                        </TableRow>                                    
                                    );
                                })
                            }
                        </TableBody>
                    </Table>                    
                </div>
            )
        } else {
            return (
                <div>
                    <LinearProgress mode="indeterminate" style={loaderStyle}/>
                    <p>Loading.</p>
                </div>
                
            )
        }
    }
}

List.propTypes = {
    selectedUser: PropTypes.string,
    onSelect: PropTypes.func.isRequired,
    selectedSongs: PropTypes.array.isRequired
}

module.exports = List;