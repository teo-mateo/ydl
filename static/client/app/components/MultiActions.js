import React from 'react'
import IconButton from 'material-ui/IconButton';
import FileCloudDownload from 'material-ui/svg-icons/file/cloud-download';
import ActionsDelete from 'material-ui/svg-icons/action/delete';
import numeral from 'numeral';

import PropTypes from 'prop-types';

class MultiActions extends React.Component{
    constructor(){
        super()
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
                    <IconButton disabled={hasSelection ? false : true}>
                        <ActionsDelete />
                    </IconButton>                                         
            </div>    
        )
    }
}

MultiActions.propTypes = {
    selectedSongs: PropTypes.array.isRequired
}

module.exports = MultiActions;