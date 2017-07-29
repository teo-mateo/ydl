import React from 'react'
import IconButton from 'material-ui/IconButton';
import FileCloudDownload from 'material-ui/svg-icons/file/cloud-download';
import ActionsDelete from 'material-ui/svg-icons/action/delete';

import PropTypes from 'prop-types';

class MultiActions extends React.Component{
    constructor(){
        super()
    }

    render(){
        var hasSelection = false;
        if(this.props.selectedSongs.length > 0){
            hasSelection = true;
        }

        return (
            <div>
                <span>{ this.props.selectedSongs.length} songs selected</span>
                <IconButton 
                    disabled={hasSelection ? false : true}
                    onClick={() => {
                        var multiDlUrl = "http://" + window.location.host + "/multidownload?ids="+this.props.selectedSongs.join(",");
                        window.open(multiDlUrl, "_blank");
                        }}>
                        <FileCloudDownload />
                    </IconButton>                             
                    <IconButton disabled={hasSelection ? false : true}>
                        <ActionsDelete />
                    </IconButton>
                <div>
                    
                </div>                                           
            </div>    
        )
    }
}

MultiActions.propTypes = {
    selectedSongs: PropTypes.array.isRequired
}

module.exports = MultiActions;