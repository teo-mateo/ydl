var React = require('react');
var numeral = require('numeral');

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

import PropTypes from 'prop-types';
import Pagination from './Pagination';
import Api from '../utils/Api';

class List extends React.Component{
    constructor(){
        super()
        this.state = {
            allfiles:[],
            files: [], 
            loading:true,
            paging: {
                page: 1,
                pageSize:5
            }
        }

        this.gotoPage = this.gotoPage.bind(this);

        console.log(this.state);
    }

    loadSongs(user){
        Api.GetList(user)
            .then(songs => {
                songs.forEach(function(s){ s.isSelected=false;});
                this.setState({ 
                    allfiles: songs,
                    files: songs, 
                    loading: false })
        });
    }

    componentDidMount(){
        this.loadSongs(this.props.selectedUser)
    }

    componentDidUpdate(){
        console.log('List.componentDidUpdate()');
        console.log(this.state);
    }

    componentWillReceiveProps(nextProps){

        if (this.props.selectedUser === nextProps.selectedUser)
            return;

        //parent will show loader
        this.setState({
            allfiles: [], 
            files: [], 
            loading: true
        });

        this.loadSongs(nextProps.selectedUser);      
    }

    onRowSelection(rows){
        if (rows === "all"){
            var newState = {
                files: this.state.files.map(f => {
                    f.isSelected = true;
                    return f;
                })
            };         
            this.setState(newState);   

            var selectedFiles = this.state.files.map(f => ({
                id: f.ID,
                size: f.FileSize
            }));
            this.props.onSelect(selectedFiles);

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

            var selectedFiles = [];
            this.props.onSelect(selectedFiles);
        } else {

            //rows has the indexes of the selected rows
            var selectedFiles = rows.map((x) => ({ 
                id: this.state.files[x].ID,
                size: this.state.files[x].FileSize
            }));


            this.props.onSelect(selectedFiles);

            this.setState(() =>{
                var newState = {
                    files: this.state.files
                        .map(f => {
                            var shouldBeSelected = selectedFiles.filter(x => x.id == f.ID).length > 0; 
                            f.isSelected = shouldBeSelected;
                            return f;
                        })

                };
                return newState;
            });            
        }
    }


    btnSkip_Click(){
        this.gotoPage(this.state.paging.page+1);  
    }

    gotoPage(page){
        var pageSize = this.state.paging.pageSize;
        var viewFiles = [];
        for (
            var i = (pageSize * (page-1))   ; 
            i < pageSize * page && i < this.state.allfiles.length; 
            i++){
            viewFiles.push(this.state.allfiles[i])
        }

        this.setState(() => ({
            files: viewFiles,
            paging: {
                pageSize: this.state.paging.pageSize,
                page: page
            }
        }));           
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
                    {/* <div>
                        <Pagination 
                            pageSize={this.state.paging.pageSize}
                            itemsCount={this.state.allfiles.length}
                            onPageChange={this.gotoPage}/>
                    </div> */}
                    <Table selectable={true} multiSelectable={true} 
                        style={{tableLayout: 'auto', padding: '0 !important'}} 
                        onRowSelection={this.onRowSelection.bind(this)}>
                        <TableHeader>
                            <TableRow style={myPaddingStyle}>
                                <TableHeaderColumn></TableHeaderColumn>
                                <TableHeaderColumn></TableHeaderColumn>
                                <TableHeaderColumn></TableHeaderColumn>
                                <TableHeaderColumn></TableHeaderColumn>
                            </TableRow>
                        </TableHeader>
                        <TableBody
                            deselectOnClickaway={false}
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
                                            <TableRowColumn style={{width: "50px"}}>
                                                {(song.YTKey.Valid) 
                                                    ? (
                                                        <a href={"https://www.youtube.com/watch?v="+song.YTKey.String} 
                                                            target="_blank">
                                                            <img 
                                                                src={"https://img.youtube.com/vi/"+song.YTKey.String+"/0.jpg"} 
                                                                className="videothumb" alt="Play on YouTube"
                                                                /> 
                                                        </a>
                                                    )
                                                    : (
                                                        <img src="https://img.youtube.com/" className="videothumb" />
                                                    )}

                                            </TableRowColumn>
                                            <TableRowColumn style={{width: "50px"}}>{song.Who.String}</TableRowColumn>
                                            <TableRowColumn>
                                                
                                                <div><a href={dlurl+song.ID}>{song.FileName}</a></div>
                                                <div style={{
                                                        fontSize: 10,
                                                        color: "gray"
                                                    }}>
                                                    <div style={{marginTop: 5}}>
                                                        {numeral(song.FileSize / (1024*1024)).format('0.0') + "MB"} 
                                                    </div>
                                                </div>
                                            </TableRowColumn>
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