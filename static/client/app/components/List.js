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
import PropTypes from 'prop-types'

var Api = require('../utils/Api')


function SingleRow(props){
    var dlurl = 'http://' + window.location.host + '/download?id='
    return (
        <TableRow selectable={true}>
            <TableRowColumn>{props.song.ID}</TableRowColumn>
            <TableRowColumn>{props.song.Who.String}</TableRowColumn>
            <TableRowColumn>{props.song.File.String}</TableRowColumn>
            <TableRowColumn>
                <a href={dlurl+props.song.ID}>
                    <IconButton>
                        <FileCloudDownload />
                    </IconButton>
                </a>
            </TableRowColumn>
        </TableRow>
    );
}

class List extends React.Component{
    constructor(){
        super()
        this.state = {
            files: [], 
            settings: {
                fixedHeader: true,
                fixedFooter: true,
                stripedRows: true,
                showRowHover: true,
                selectable: false,
                multiSelectable: false,
                enableSelectAll: false,
                deselectOnClickaway: true,
                showCheckboxes: false
            }
        }

        console.log(this.state);

        //this.load = this.load.bind(this);
    }

    componentDidMount(){
        Api.GetList(this.props.selectedUser)
            .then(songs => {
                this.setState({ files: songs })
        });
    }

    componentDidUpdate(){
        console.log('List.componentDidUpdate()');
        console.log(this.state);
    }

    componentWillReceiveProps(nextProps){
        console.log("List.componentWillReceiveProps()")
        Api.GetList(nextProps.selectedUser)
            .then(songs => {
                this.setState({ files: songs })
        });        
    }

    render(){

        var settings = this.state.settings;
        console.log(settings)
        return (
            <Table 
                height={settings.height}
                fixedHeader={settings.fixedHeader}
                fixedFooter={settings.fixedFooter}
                selectable={settings.selectable}
                multiSelectable={settings.multiSelectable}            
                style={{tableLayout: 'auto'}} >
                <TableHeader
                    displaySelectAll={settings.showCheckboxes}
                    adjustForCheckbox={settings.showCheckboxes}
                    enableSelectAll={settings.enableSelectAll}>
                    <TableRow>
                        <TableHeaderColumn></TableHeaderColumn>
                        <TableHeaderColumn></TableHeaderColumn>
                        <TableHeaderColumn></TableHeaderColumn>
                    </TableRow>
                </TableHeader>
                <TableBody
                    displayRowCheckbox={settings.showCheckboxes}
                    deselectOnClickaway={settings.deselectOnClickaway}
                    showRowHover={settings.showRowHover}
                    stripedRows={settings.stripedRows}>

                    {
                        !this.state.files
                        ? "NO DATA"
                        : this.state.files.map(function(song, index){
                            return <SingleRow song={song} key={index} />
                        })
                        
                    }
                </TableBody>
            </Table>
        )
    }
}

List.propTypes = {
    selectedUser: PropTypes.string
}

module.exports = List;