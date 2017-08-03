import React from 'react'
import IconButton from 'material-ui/IconButton';
import FlatButton from 'material-ui/FlatButton';
import FileCloudDownload from 'material-ui/svg-icons/file/cloud-download';
import ActionsDelete from 'material-ui/svg-icons/action/delete';
import numeral from 'numeral';
import Dialog from 'material-ui/Dialog';
import Api from '../utils/Api';

import PropTypes from 'prop-types';

class MultiActions extends React.Component{
    constructor(){
        super()

        this.state = {
            dialogOpen: false
        };

        this.handleClose = this.handleClose.bind(this);
        this.handleDelete = this.handleDelete.bind(this);
        this.closeAndRefresh = this.closeAndRefresh.bind(this);
    }

    handleClose(){
        this.setState({
            dialogOpen: false
        });
    }

    closeAndRefresh(){
        this.handleClose();
        window.location.reload()
    }

    handleDelete(){
        Api.MultiDelete(this.props.selectedSongs.map(s => s.id))
            .then(this.closeAndRefresh);
    }

    render(){
        
        var fullSize = 0;
        
        var hasSelection = false;
        if(this.props.selectedSongs.length > 0){
            hasSelection = true;
            fullSize = this.props.selectedSongs.reduce((s,v) => {
                return s + v.size;
            }, 0);            
        }


        const dialogActions=[
            <FlatButton label="Cancel" primary={true} onTouchTap={this.handleClose} />,
            <FlatButton label="OK" primary={true}  onTouchTap={this.handleDelete} />
        ];

        return (
            <div style={{
                fontSize: 12,
                color: "gray"
                }}>
                    <span>
                        <span>{numeral(fullSize / (1024*1024)).format('0.0') + "MB"} </span>&nbsp;|&nbsp;
                        <span>{ this.props.selectedSongs.length} songs selected</span>
                    </span>                
                <IconButton 
                    disabled={hasSelection ? false : true}
                    onClick={() => {
                        var multiDlUrl = "http://" + window.location.host + "/multidownload?ids="+this.props.selectedSongs.map(x => x.id).join(",");
                        window.open(multiDlUrl, "_blank");
                        }}>
                        <FileCloudDownload />
                </IconButton>                             
                <IconButton 
                    disabled={hasSelection ? false : true}
                    onClick={() => {
                        this.setState({
                            dialogOpen: true
                        });
                    }}>
                    <ActionsDelete />
                </IconButton> 
                <Dialog open={this.state.dialogOpen}
                    title={"Confirm deletion of " + this.props.selectedSongs.length + " files."}
                    actions={dialogActions}
                    modal={true} >
                </Dialog>                                        
            </div>    
        )
    }
}

// selected songs must be an array of {id, size} objects
MultiActions.propTypes = {
    selectedSongs: PropTypes.arrayOf(PropTypes.shape({
        id: PropTypes.number.isRequired,
        size: PropTypes.number.isRequired
    })).isRequired
}

module.exports = MultiActions;